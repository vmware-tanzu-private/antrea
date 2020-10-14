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

package store

import (
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	"github.com/vmware-tanzu/antrea/pkg/apis/controlplane"
	"github.com/vmware-tanzu/antrea/pkg/apiserver/storage"
	"github.com/vmware-tanzu/antrea/pkg/apiserver/storage/ram"
	"github.com/vmware-tanzu/antrea/pkg/controller/types"
)

// groupEvent implements storage.InternalEvent.
type groupEvent struct {
	// The current version of the stored Group.
	CurrGroup *types.Group
	// The previous version of the stored Group.
	PrevGroup *types.Group
	// The current version of the transferred Group, which will be used in Added events.
	CurrObject *controlplane.Group
	// The previous version of the transferred Group, which will be used in Deleted events.
	// Note that only metadata will be set in Deleted events for efficiency.
	PrevObject *controlplane.Group
	// The patch object of the message for transferring, which will be used in Modified events.
	PatchObject *controlplane.GroupPatch
	// The key of this Group.
	Key             string
	ResourceVersion uint64
}

// ToWatchEvent converts the groupEvent to *watch.Event based on the provided Selectors. It has the following features:
// 1. Added event will be generated if the Selectors was not interested in the object but is now.
// 2. Modified event will be generated if the Selectors was and is interested in the object.
// 3. Deleted event will be generated if the Selectors was interested in the object but is not now.
// 4. If nodeName is specified, only GroupMembers that hosted by the Node will be in the event.
func (event *groupEvent) ToWatchEvent(selectors *storage.Selectors, isInitEvent bool) *watch.Event {
	prevObjSelected, currObjSelected := isSelected(event.Key, event.PrevGroup, event.CurrGroup, selectors, isInitEvent)

	switch {
	case !currObjSelected && !prevObjSelected:
		// Watcher is not interested in that object.
		return nil
	case currObjSelected && !prevObjSelected:
		// Watcher was not interested in that object but is now, an added event will be generated.
		return &watch.Event{Type: watch.Added, Object: event.CurrObject}
	case currObjSelected && prevObjSelected:
		// Watcher was and is interested in that object, a modified event will be generated, unless there's no address change.
		if event.PatchObject == nil {
			return nil
		}
		return &watch.Event{Type: watch.Modified, Object: event.PatchObject}
	case !currObjSelected && prevObjSelected:
		// Watcher was interested in that object but is not interested now, a deleted event will be generated.
		return &watch.Event{Type: watch.Deleted, Object: event.PrevObject}
	}
	return nil
}

func (event *groupEvent) GetResourceVersion() uint64 {
	return event.ResourceVersion
}

var _ storage.GenEventFunc = genGroupEvent

// genGroupEvent generates InternalEvent from the given versions of an Group.
func genGroupEvent(key string, prevObj, currObj interface{}, rv uint64) (storage.InternalEvent, error) {
	if reflect.DeepEqual(prevObj, currObj) {
		return nil, nil
	}

	event := &groupEvent{Key: key, ResourceVersion: rv}

	if prevObj != nil {
		event.PrevGroup = prevObj.(*types.Group)
		event.PrevObject = new(controlplane.Group)
		ToGroupMsg(event.PrevGroup, event.PrevObject, false)
	}

	if currObj != nil {
		event.CurrGroup = currObj.(*types.Group)
		event.CurrObject = new(controlplane.Group)
		ToGroupMsg(event.CurrGroup, event.CurrObject, true)
	}

	// Calculate PatchObject in advance so that we don't need to do it for
	// each watcher when generating *event.Event.
	if event.PrevGroup != nil && event.CurrGroup != nil {
		var addedMembers, removedMembers []controlplane.GroupMember
		for memberHash, member := range event.CurrGroup.GroupMembers {
			if _, exists := event.PrevGroup.GroupMembers[memberHash]; !exists {
				addedMembers = append(addedMembers, *member)
			}
		}
		for memberHash, member := range event.PrevGroup.GroupMembers {
			if _, exists := event.CurrGroup.GroupMembers[memberHash]; !exists {
				removedMembers = append(removedMembers, *member)
			}
		}
		// PatchObject will not be generated when only span changes.
		if len(addedMembers)+len(removedMembers) > 0 {
			event.PatchObject = new(controlplane.GroupPatch)
			event.PatchObject.UID = event.CurrGroup.UID
			event.PatchObject.Name = event.CurrGroup.Name
			event.PatchObject.AddedGroupMembers = addedMembers
			event.PatchObject.RemovedGroupMembers = removedMembers
		}
	}

	return event, nil
}

// ToGroupMsg converts the stored Group to its message form.
// If includeBody is true, GroupMembers will be copied.
func ToGroupMsg(in *types.Group, out *controlplane.Group, includeBody bool) {
	out.Name = in.Name
	out.UID = in.UID
	if !includeBody {
		return
	}
	for _, member := range in.GroupMembers {
		out.GroupMembers = append(out.GroupMembers, *member)
	}
}

// GroupKeyFunc knows how to get the key of an Group.
func GroupKeyFunc(obj interface{}) (string, error) {
	group, ok := obj.(*types.Group)
	if !ok {
		return "", fmt.Errorf("object is not *types.Group: %v", obj)
	}
	return group.Name, nil
}

// NewGroupStore creates a store of Group.
func NewGroupStore() storage.Interface {
	indexers := cache.Indexers{
		cache.NamespaceIndex: func(obj interface{}) ([]string, error) {
			g, ok := obj.(*types.Group)
			if !ok {
				return []string{}, nil
			}
			return []string{g.Selector.Namespace}, nil
		},
	}
	return ram.NewStore(GroupKeyFunc, indexers, genGroupEvent, keyAndSpanSelectFunc, func() runtime.Object { return new(controlplane.Group) })
}
