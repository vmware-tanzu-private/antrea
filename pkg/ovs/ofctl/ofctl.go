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

package ofctl

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type OfctlClient struct {
	bridge string
}

func NewClient(bridge string) *OfctlClient {
	return &OfctlClient{bridge}
}

func (c *OfctlClient) DumpFlows(args ...string) ([]string, error) {
	flowDump, err := c.RunOfctlCmd("dump-flows", args...)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(flowDump)))
	scanner.Split(bufio.ScanLines)
	// Skip the first line.
	scanner.Scan()
	flowList := []string{}
	for scanner.Scan() {
		flowList = append(flowList, scanner.Text())
	}

	return flowList, nil
}

func (c *OfctlClient) DumpGroups(args ...string) ([][]string, error) {
	groupsDump, err := c.RunOfctlCmd("dump-groups", args...)
	if err != nil {
		return nil, err
	}
	groupsDumpStr := strings.TrimSpace(string(groupsDump))

	scanner := bufio.NewScanner(strings.NewReader(groupsDumpStr))
	scanner.Split(bufio.ScanLines)
	// Skip the first line.
	scanner.Scan()
	rawGroupItems := []string{}
	for scanner.Scan() {
		rawGroupItems = append(rawGroupItems, scanner.Text())
	}

	var groupList [][]string
	for _, rawGroupItem := range rawGroupItems {
		rawGroupItem = strings.TrimSpace(rawGroupItem)
		elems := strings.Split(rawGroupItem, ",bucket=")
		groupList = append(groupList, elems)
	}
	return groupList, nil
}

func (c *OfctlClient) DumpTableFlows(table uint8) ([]string, error) {
	return c.DumpFlows(fmt.Sprintf("table=%d", table))
}

func (c *OfctlClient) RunOfctlCmd(cmd string, args ...string) ([]byte, error) {
	cmdStr := fmt.Sprintf("ovs-ofctl -O Openflow13 %s %s", cmd, c.bridge)
	cmdStr = cmdStr + " " + strings.Join(args, " ")
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
