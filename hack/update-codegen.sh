#!/usr/bin/env bash

# Copyright 2019 Antrea Authors
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
set -o pipefail

function echoerr {
    >&2 echo "$@"
}

ANTREA_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/../" && pwd )"
IMAGE_NAME="antrea/codegen:kubernetes-1.29.2-build.1"

# We will use git clone to make a working copy of the repository into a
# temporary directory. This requires that all changes have been committed
# otherwise the generated code may not be up-to-date. It is anyway good practice
# to commit changes before auto-generating code, in case there is an issue with
# the script.
# We run these checks here instead of inside the Docker container for 2 reasons:
# speed (bind mounts are slow) and to avoid having to add the source repository
# to the "safe" list in the Git config when starting the container (required
# because of user mismatch).
if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -n "${git_status}" ]]; then
  echoerr "!!! Dirty tree. Clean up and try again."
  exit 1
fi
# It is very common to have untracked files in a repository, so we only give a
# warning if we find untracked Golang source files.
untracked=$(git ls-files --others --exclude-standard '**.go')
if [ -n "$untracked" ]; then
    echoerr "The following Golang files are untracked, and will be ignored by code generation:"
    echoerr "$untracked"
fi

function docker_run() {
  # Silence CLI suggestions.
  export DOCKER_CLI_HINTS=false
  docker pull ${IMAGE_NAME}
  set -x
  ANTREA_SRC_PATH="/mnt/antrea"
  docker run --rm \
		-e GOPROXY=${GOPROXY} \
		-e HTTP_PROXY=${HTTP_PROXY} \
		-e HTTPS_PROXY=${HTTPS_PROXY} \
		-w ${ANTREA_SRC_PATH} \
                --mount type=bind,source=${ANTREA_ROOT},target=${ANTREA_SRC_PATH} \
                --mount type=volume,source=antrea-codegen-gopkgmod,target=/go/pkg/mod \
                --mount type=volume,source=antrea-codegen-gocache,target=/root/.cache/go-build \
		"${IMAGE_NAME}" bash -c "$@"
}

# Combine hack/update-codegen-dockerized.sh and the arguments to the script as a
# single argument to the docker_run function.
docker_run "hack/update-codegen-dockerized.sh $@"
