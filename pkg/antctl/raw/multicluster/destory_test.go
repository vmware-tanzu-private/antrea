// Copyright 2022 Antrea Authors
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

package multicluster

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	mcsv1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	mcscheme "antrea.io/antrea/pkg/antctl/raw/multicluster/scheme"
)

func TestDestroy(t *testing.T) {
	tests := []struct {
		name           string
		expectedOutput string
		namespace      string
	}{
		{
			name:           "destroy successfully",
			expectedOutput: "ClusterSet \"test-clusterset\" deleted in Namespace default\nClusterProperty \"cluster.clusterset.k8s.io\" deleted in Namespace default\nClusterProperty \"clusterset.k8s.io\" deleted in Namespace default\n",
			namespace:      "default",
		},
		{
			name:           "fail to destroy due to empty Namespace",
			expectedOutput: "Namespace must be specified",
			namespace:      "",
		},
	}

	cmd := NewDestroyCommand()
	buf := new(bytes.Buffer)
	cmd.SetOutput(buf)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.Flag("clusterset").Value.Set("test-clusterset")
	clusterSet := &mcsv1alpha1.ClusterSet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-clusterset",
		},
	}
	clusterProperty1 := &mcsv1alpha1.ClusterProperty{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "cluster.clusterset.k8s.io",
		},
	}
	clusterProperty2 := &mcsv1alpha1.ClusterProperty{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "clusterset.k8s.io",
		},
	}
	fakeClient := fake.NewClientBuilder().WithScheme(mcscheme.Scheme).WithObjects(clusterSet, clusterProperty1, clusterProperty2).Build()
	destroyOpts.K8sClient = fakeClient
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd.Flag("namespace").Value.Set(tt.namespace)
			err := cmd.Execute()
			if err != nil {
				assert.Equal(t, tt.expectedOutput, err.Error())
			} else {
				assert.Equal(t, tt.expectedOutput, buf.String())
			}
		})
	}
}
