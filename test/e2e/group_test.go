// Copyright 2021 Antrea Authors
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

package e2e

import (
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	crdv1alpha1 "antrea.io/antrea/pkg/apis/crd/v1alpha1"
	crdv1alpha3 "antrea.io/antrea/pkg/apis/crd/v1alpha3"
)

func testInvalidGroupIPBlockWithPodSelector(t *testing.T) {
	invalidErr := fmt.Errorf("group created with ipblock and podSelector")
	gName := "ipb-pod"
	pSel := &metav1.LabelSelector{MatchLabels: map[string]string{"pod": "x"}}
	cidr := "10.0.0.10/32"
	ipb := []crdv1alpha1.IPBlock{{CIDR: cidr}}
	g := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gName,
			Namespace: "x",
		},
		Spec: crdv1alpha3.GroupSpec{
			PodSelector: pSel,
			IPBlocks:    ipb,
		},
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g); err == nil {
		// Above creation of Group must fail as it is an invalid spec.
		failOnError(invalidErr, t)
	}
}

func testInvalidGroupIPBlockWithNSSelector(t *testing.T) {
	invalidErr := fmt.Errorf("group created with ipblock and namespaceSelector")
	gName := "ipb-ns"
	nSel := &metav1.LabelSelector{MatchLabels: map[string]string{"ns": "y"}}
	cidr := "10.0.0.10/32"
	ipb := []crdv1alpha1.IPBlock{{CIDR: cidr}}
	g := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gName,
			Namespace: "x",
		},
		Spec: crdv1alpha3.GroupSpec{
			NamespaceSelector: nSel,
			IPBlocks:          ipb,
		},
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g); err == nil {
		// Above creation of Group must fail as it is an invalid spec.
		failOnError(invalidErr, t)
	}
}

func testInvalidGroupServiceRefWithPodSelector(t *testing.T) {
	invalidErr := fmt.Errorf("group created with serviceReference and podSelector")
	gName := "svcref-pod-selector"
	pSel := &metav1.LabelSelector{MatchLabels: map[string]string{"pod": "x"}}
	svcRef := &crdv1alpha1.NamespacedName{
		Namespace: "y",
		Name:      "test-svc",
	}
	g := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gName,
			Namespace: "y",
		},
		Spec: crdv1alpha3.GroupSpec{
			PodSelector:      pSel,
			ServiceReference: svcRef,
		},
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g); err == nil {
		// Above creation of Group must fail as it is an invalid spec.
		failOnError(invalidErr, t)
	}
}

func testInvalidGroupServiceRefWithNSSelector(t *testing.T) {
	invalidErr := fmt.Errorf("group created with serviceReference and namespaceSelector")
	gName := "svcref-ns-selector"
	nSel := &metav1.LabelSelector{MatchLabels: map[string]string{"ns": "y"}}
	svcRef := &crdv1alpha1.NamespacedName{
		Namespace: "y",
		Name:      "test-svc",
	}
	g := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gName,
			Namespace: "y",
		},
		Spec: crdv1alpha3.GroupSpec{
			NamespaceSelector: nSel,
			ServiceReference:  svcRef,
		},
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g); err == nil {
		// Above creation of Group must fail as it is an invalid spec.
		failOnError(invalidErr, t)
	}
}

func testInvalidGroupServiceRefWithIPBlock(t *testing.T) {
	invalidErr := fmt.Errorf("group created with ipblock and namespaceSelector")
	gName := "ipb-svcref"
	cidr := "10.0.0.10/32"
	ipb := []crdv1alpha1.IPBlock{{CIDR: cidr}}
	svcRef := &crdv1alpha1.NamespacedName{
		Namespace: "y",
		Name:      "test-svc",
	}
	g := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gName,
			Namespace: "y",
		},
		Spec: crdv1alpha3.GroupSpec{
			ServiceReference: svcRef,
			IPBlocks:         ipb,
		},
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g); err == nil {
		// Above creation of Group must fail as it is an invalid spec.
		failOnError(invalidErr, t)
	}
}

var (
	testChildGroupName      = "test-child-grp"
	testChildGroupNamespace = "x"
)

func createChildGroupForTest(t *testing.T) {
	g := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testChildGroupName,
			Namespace: testChildGroupNamespace,
		},
		Spec: crdv1alpha3.GroupSpec{
			PodSelector: &metav1.LabelSelector{},
		},
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g); err != nil {
		failOnError(err, t)
	}
}

func cleanupChildGroupForTest(t *testing.T) {
	if err := k8sUtils.DeleteV1Alpha3Group(testChildGroupNamespace, testChildGroupName); err != nil {
		failOnError(err, t)
	}
}

func testInvalidGroupChildGroupWithPodSelector(t *testing.T) {
	invalidErr := fmt.Errorf("group created with childGroups and podSelector")
	gName := "child-group-pod-selector"
	pSel := &metav1.LabelSelector{MatchLabels: map[string]string{"pod": "x"}}
	g := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gName,
			Namespace: testChildGroupNamespace,
		},
		Spec: crdv1alpha3.GroupSpec{
			PodSelector: pSel,
			ChildGroups: []crdv1alpha3.ClusterGroupReference{crdv1alpha3.ClusterGroupReference(testChildGroupName)},
		},
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g); err == nil {
		// Above creation of Group must fail as it is an invalid spec.
		failOnError(invalidErr, t)
	}
}

func testInvalidGroupChildGroupWithServiceReference(t *testing.T) {
	invalidErr := fmt.Errorf("group created with childGroups and ServiceReference")
	gName := "child-group-svcref"
	svcRef := &crdv1alpha1.NamespacedName{
		Namespace: testChildGroupNamespace,
		Name:      "test-svc",
	}
	g := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gName,
			Namespace: testChildGroupNamespace,
		},
		Spec: crdv1alpha3.GroupSpec{
			ServiceReference: svcRef,
			ChildGroups:      []crdv1alpha3.ClusterGroupReference{crdv1alpha3.ClusterGroupReference(testChildGroupName)},
		},
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g); err == nil {
		// Above creation of Group must fail as it is an invalid spec.
		failOnError(invalidErr, t)
	}
}

func testInvalidGroupMaxNestedLevel(t *testing.T) {
	invalidErr := fmt.Errorf("group created with childGroup which has childGroups itself")
	gName1, gName2 := "g-nested-1", "g-nested-2"
	g1 := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{Namespace: testChildGroupNamespace, Name: gName1},
		Spec: crdv1alpha3.GroupSpec{
			ChildGroups: []crdv1alpha3.ClusterGroupReference{crdv1alpha3.ClusterGroupReference(testChildGroupName)},
		},
	}
	g2 := &crdv1alpha3.Group{
		ObjectMeta: metav1.ObjectMeta{Namespace: testChildGroupNamespace, Name: gName2},
		Spec: crdv1alpha3.GroupSpec{
			ChildGroups: []crdv1alpha3.ClusterGroupReference{crdv1alpha3.ClusterGroupReference(gName1)},
		},
	}
	// Try to create g-nested-1 first and then g-nested-2.
	// The creation of g-nested-2 should fail as it breaks the max nested level
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g1); err != nil {
		// Above creation of Group must succeed as it is a valid spec.
		failOnError(err, t)
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g2); err == nil {
		// Above creation of Group must fail as g-nested-2 cannot have g-nested-1 as childGroup.
		failOnError(invalidErr, t)
	}
	// cleanup g-nested-1
	if err := k8sUtils.DeleteV1Alpha3Group(testChildGroupNamespace, gName1); err != nil {
		failOnError(err, t)
	}
	// Try to create g-nested-2 first and then g-nested-1.
	// The creation of g-nested-1 should fail as it breaks the max nested level
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g2); err != nil {
		// Above creation of Group must succeed as it is a valid spec.
		failOnError(err, t)
	}
	if _, err := k8sUtils.CreateOrUpdateV1Alpha3Group(g1); err == nil {
		// Above creation of Group must fail as g-nested-2 cannot have g-nested-1 as childGroup.
		failOnError(invalidErr, t)
	}
	// cleanup cg-nested-2
	if err := k8sUtils.DeleteV1Alpha3Group(testChildGroupNamespace, gName2); err != nil {
		failOnError(err, t)
	}
}

func TestGroup(t *testing.T) {
	skipIfHasWindowsNodes(t)
	skipIfAntreaPolicyDisabled(t)

	data, err := setupTest(t)
	if err != nil {
		t.Fatalf("Error when setting up test: %v", err)
	}
	defer teardownTest(t, data)
	initialize(t, data)

	t.Run("TestGroupNamespacedGroupValidate", func(t *testing.T) {
		t.Run("Case=IPBlockWithPodSelectorDenied", func(t *testing.T) { testInvalidGroupIPBlockWithPodSelector(t) })
		t.Run("Case=IPBlockWithNamespaceSelectorDenied", func(t *testing.T) { testInvalidGroupIPBlockWithNSSelector(t) })
		t.Run("Case=ServiceRefWithPodSelectorDenied", func(t *testing.T) { testInvalidGroupServiceRefWithPodSelector(t) })
		t.Run("Case=ServiceRefWithNamespaceSelectorDenied", func(t *testing.T) { testInvalidGroupServiceRefWithNSSelector(t) })
		t.Run("Case=ServiceRefWithIPBlockDenied", func(t *testing.T) { testInvalidGroupServiceRefWithIPBlock(t) })
	})
	t.Run("TestGroupNamespacedGroupValidateChildGroup", func(t *testing.T) {
		createChildGroupForTest(t)
		t.Run("Case=ChildGroupWithPodSelectorDenied", func(t *testing.T) { testInvalidGroupChildGroupWithPodSelector(t) })
		t.Run("Case=ChildGroupWithPodServiceReferenceDenied", func(t *testing.T) { testInvalidGroupChildGroupWithServiceReference(t) })
		t.Run("Case=ChildGroupExceedMaxNestedLevel", func(t *testing.T) { testInvalidGroupMaxNestedLevel(t) })
		cleanupChildGroupForTest(t)
	})
	k8sUtils.Cleanup(namespaces) // clean up all cluster-scope resources, including CGs
}
