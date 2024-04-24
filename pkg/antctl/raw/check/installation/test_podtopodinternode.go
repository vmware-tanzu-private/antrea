// Copyright 2024 Antrea Authors.
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

package installation

import (
	"context"
	"fmt"

	"antrea.io/antrea/pkg/antctl/raw/check"
)

type PodToPodInterNodeConnectivityTest struct{}

func init() {
	RegisterTest("pod-to-pod-internode-connectivity", &PodToPodInterNodeConnectivityTest{})
}

func (t *PodToPodInterNodeConnectivityTest) Run(ctx context.Context, testContext *testContext) error {
	if testContext.echoOtherNodePod == nil {
		return fmt.Errorf("Skipping Inter-Node test because multiple Nodes are not available")
	}
	for _, clientPod := range testContext.clientPods {
		srcPod := testContext.namespace + "/" + clientPod.Name
		dstPod := testContext.namespace + "/" + testContext.echoOtherNodePod.Name
		for _, ipInfo := range testContext.echoOtherNodePod.Status.PodIPs {
			echoIP := ipInfo.IP
			testContext.Log("Validating from Pod %s to Pod %s...", srcPod, dstPod)
			_, _, err := check.ExecInPod(ctx, testContext.client, testContext.config, testContext.namespace, clientPod.Name, "", agnhostConnectCommand(echoIP+":80"))
			if err != nil {
				return fmt.Errorf("client Pod %s was not able to communicate with echo Pod %s (%s): %w", clientPod.Name, testContext.echoOtherNodePod.Name, echoIP, err)
			}
			testContext.Log("client Pod %s was able to communicate with echo Pod %s (%s)", clientPod.Name, testContext.echoOtherNodePod.Name, echoIP)
		}
	}
	return nil
}
