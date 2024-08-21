#!/usr/bin/env bash

# Copyright 2021 Antrea Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

GOPATH=`go env GOPATH`
ANTREA_PKG="antrea.io/antrea"

# We make a temporary working copy of the source repository using git clone. The
# copy is in the container's writable layer, which is much faster than a bind
# mount on non-Linux hosts (Docker Desktop, Colima).
ANTREA_SRC_PATH=$(pwd)
ANTREA_CODEGEN_PATH=/go/src/antrea.io/antrea
git clone ${ANTREA_SRC_PATH} ${ANTREA_CODEGEN_PATH}
pushd ${ANTREA_CODEGEN_PATH}

trap "popd; rm -rf ${ANTREA_CODEGEN_PATH}" EXIT

# Copy generated code back into the source repository, which was cloned at the
# beginning of this script to make a temporary working copy.
function copy_generated_code_to_source {
  git ls-files --modified --others --exclude-standard | while read file; do
      if [ -e $file ]; then
          cp $file ${ANTREA_SRC_PATH}/$file
      else
          rm ${ANTREA_SRC_PATH}/$file
      fi
  done
}

source hack/update-codegen-common.sh

MOCKGEN_TARGETS=(
  "multicluster/controllers/multicluster/leader MemberClusterStatusManager . mock_membercluster_status_manager.go"
  "multicluster/controllers/multicluster/commonarea CommonArea,RemoteCommonArea,ImportReconciler,RemoteCommonAreaGetter . mock_remote_common_area.go"
)

if [[ "$#" -eq 1 && $1 == "mockgen" ]]; then
  # Remove all files generated by MockGen.
  # We exclude multicluster/hack (otherwise we would remove ourselves) and multicluster/test which
  # includes a third-party generated mock.
  git grep --files-with-matches -e 'Code generated by MockGen\. DO NOT EDIT' -- multicluster ':(exclude)multicluster/hack' ':(exclude)multicluster/test' | xargs rm
  generate_mocks
  reset_year_change
  copy_generated_code_to_source
  exit 0
fi

# Remove all generated files.
git grep --files-with-matches -e 'Code generated by .* DO NOT EDIT' -- multicluster ':(exclude)multicluster/hack' ':(exclude)multicluster/test' | xargs rm

function generate_multicluster_client_code {
  # Generate clientset and apis code with K8s codegen tools.
  $GOPATH/bin/client-gen \
    --clientset-name versioned \
    --input-base "${ANTREA_PKG}/multicluster/apis" \
    --input "multicluster/v1alpha1" \
    --input "multicluster/v1alpha2" \
    --output-package "${ANTREA_PKG}/multicluster/pkg/client/clientset" \
    --go-header-file hack/boilerplate/license_header.go.txt
  
  # Generate listers with K8s codegen tools.
  $GOPATH/bin/lister-gen \
    --input-dirs "${ANTREA_PKG}/multicluster/apis/multicluster/v1alpha1" \
    --input-dirs "${ANTREA_PKG}/multicluster/apis/multicluster/v1alpha2" \
    --output-package "${ANTREA_PKG}/multicluster/pkg/client/listers" \
    --go-header-file hack/boilerplate/license_header.go.txt
  
  # Generate informers with K8s codegen tools.
  $GOPATH/bin/informer-gen \
    --input-dirs "${ANTREA_PKG}/multicluster/apis/multicluster/v1alpha1" \
    --input-dirs "${ANTREA_PKG}/multicluster/apis/multicluster/v1alpha2" \
    --versioned-clientset-package "${ANTREA_PKG}/multicluster/pkg/client/clientset/versioned" \
    --listers-package "${ANTREA_PKG}/multicluster/pkg/client/listers" \
    --output-package "${ANTREA_PKG}/multicluster/pkg/client/informers" \
    --go-header-file hack/boilerplate/license_header.go.txt

  $GOPATH/bin/controller-gen object:headerFile="hack/boilerplate/license_header.go.txt",year=$(date "+%Y") paths="./multicluster/..."
}

generate_multicluster_client_code
generate_mocks
reset_year_change
copy_generated_code_to_source
