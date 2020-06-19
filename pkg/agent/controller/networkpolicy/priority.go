// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkpolicy

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"k8s.io/klog"

	. "github.com/vmware-tanzu/antrea/pkg/agent/types"
)

const (
	PriorityBottomCNP     = uint16(100)
	InitialPriorityOffest = uint16(130)
	InitialPriorityZones  = 100
	DefaultTierStart      = uint16(13100)
)

// priorityAssigner is a struct that maintains the current boundaries of
// all ClusterNetworkPolicy category/priorities and rule priorities, and knows
// how to re-assgin priorities if certain section overflows.
type priorityAssigner struct {
	// priorityHash maintains the current mapping between a known CNP priority to OF priority.
	priorityHash map[Priority]uint16
	// priorityOffset stores the current size of a priority zone.
	// When tiering is introduced, each tier will keep its own priorityOffset.
	priorityOffset uint16
	// numPriorityZones stores the current number of numPriorityZones (within the default tier).
	// When tiering is introduced, each tier will keep its own numPriorityZones.
	numPriorityZones int32
}

func newPriorityAssinger() *priorityAssigner {
	pa := &priorityAssigner{
		priorityHash:     map[Priority]uint16{},
		priorityOffset:   InitialPriorityOffest,
		numPriorityZones: InitialPriorityZones,
	}
	return pa
}

// getPriorityZoneStart returns the priorityZone index and size for the given priority.
// It maps policyPriority [0.0-1.0) to 0, [1.0-2.0) to 1 and so on so forth.
// policyPriorities over 99.0 will be mapped to zone 99, without zone expansion for now.
func (pa *priorityAssigner) getPriorityZoneIndex(p Priority) int32 {
	floorPriority := int32(math.Floor(p.PolicyPriority))
	if floorPriority > pa.numPriorityZones-1 {
		floorPriority = pa.numPriorityZones - 1
	}
	return floorPriority
}

// getPriorityZoneStart returns the starting OF priority for the priorityZone for the input.
func (pa *priorityAssigner) getPriorityZoneStart(p Priority) uint16 {
	priorityIndex := pa.getPriorityZoneIndex(p)
	return DefaultTierStart - pa.priorityOffset*uint16(priorityIndex)
}

// getPriorityZoneSize returns the size of the priorityZone for the input.
func (pa *priorityAssigner) getPriorityZoneSize(p Priority) uint16 {
	zoneStart := pa.getPriorityZoneStart(p)
	if zoneStart-pa.priorityOffset < PriorityBottomCNP {
		return zoneStart - PriorityBottomCNP
	}
	return pa.priorityOffset
}

// sortPriorities sorts a list of priorities.
func (pa *priorityAssigner) sortPriorities(priorities []Priority) {
	sort.Slice(priorities, func(i, j int) bool {
		if priorities[i].PolicyPriority == priorities[j].PolicyPriority {
			return priorities[i].RulePriority < priorities[j].RulePriority
		}
		return priorities[i].PolicyPriority < priorities[j].PolicyPriority
	})
}

// getIndexSamePriorityZone returns a list of sorted priorities that needs to be present in the same
// priority zone of the input priority.
func (pa *priorityAssigner) getIndexSamePriorityZone(p Priority) []Priority {
	affected := []Priority{p}
	for k := range pa.priorityHash {
		if pa.getPriorityZoneStart(k) == pa.getPriorityZoneStart(p) {
			affected = append(affected, k)
		}
	}
	pa.sortPriorities(affected)
	return affected
}

// syncPriorityZone computes the new expected OF priorties for each priority in the same priority zone
// of the input priority, and returns installed priorities that need to be re-assigned if necessary.
func (pa *priorityAssigner) syncPriorityZone(p Priority) (*uint16, map[uint16]uint16, error) {

	// newPriority is the OF priority to be assigned for a new priority.
	// For priority Forget, newPriority returned should be nil.
	var newPriority uint16
	// priorityUpdates stores all the OF priority re-assignments to be performed by client
	priorityUpdates := map[uint16]uint16{}

	affected := pa.getIndexSamePriorityZone(p)
	if uint16(len(affected)) > pa.getPriorityZoneSize(p) {
		// TODO: Dynamically adjust priorityZone size to handle overflow
		return nil, priorityUpdates, fmt.Errorf("priorityZone for [%v %v) has overflowed",
			pa.getPriorityZoneIndex(p), pa.getPriorityZoneIndex(p)+1)
	}
	for offset, priority := range affected {
		computedPriority := pa.getPriorityZoneStart(p) - uint16(offset)
		oldOFPriority, updateExisting := pa.priorityHash[priority]
		if updateExisting && computedPriority != oldOFPriority {
			klog.V(2).Infof("Original priority %d needs to be reassigned %d now.", oldOFPriority, computedPriority)
			priorityUpdates[oldOFPriority] = computedPriority
		} else if !updateExisting {
			// A new Priority has been added to priorityHash
			newPriority = computedPriority
		}
		pa.priorityHash[priority] = computedPriority
	}
	return &newPriority, priorityUpdates, nil
}

// GetOFPriority retrieves the OFPriority for the input CompleteRule to be installed,
// and returns installed priorities that need to be re-assigned if necessary.
func (pa *priorityAssigner) GetOFPriority(p Priority) (*uint16, map[uint16]uint16, error) {
	ofPriority, exists := pa.priorityHash[p]
	if !exists {
		return pa.syncPriorityZone(p)
	}
	return &ofPriority, map[uint16]uint16{}, nil
}

// Forget removes the priority that currently corresponds to the input OFPriority from the priorityHash,
// and returns installed priorities that need to be re-assigned if necessary.
func (pa *priorityAssigner) Release(priorityStr string) error {
	priorityNum, err := strconv.ParseUint(priorityStr, 10, 16)
	if err != nil {
		// Cannot parse the priority str. Theoretically this should never happen.
		return err
	}
	for priorityKey, p := range pa.priorityHash {
		if uint16(priorityNum) == p {
			delete(pa.priorityHash, priorityKey)
			return nil
		}
	}
	klog.Infof("OF priority %s not stored in hash, skip forgetting priority.", priorityStr)
	return nil
}
