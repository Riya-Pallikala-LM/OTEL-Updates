# Copyright Splunk Inc.
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

#!/bin/bash

set -euo pipefail

BOSH_DIRECTOR_DIR="./bosh-env/virtualbox"
BOSH_DIRECTOR_DEPLOYMENT_DIR="bosh-deployment"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ -d $BOSH_DIRECTOR_DIR/$BOSH_DIRECTOR_DEPLOYMENT_DIR ]; then
  echo "Bosh director already exists"
  echo "If you want to build again, run make reinstall-director"
  exit 0
fi

mkdir -p $BOSH_DIRECTOR_DIR
cd $BOSH_DIRECTOR_DIR

git clone --depth 1 https://github.com/cloudfoundry/bosh-deployment.git

./$BOSH_DIRECTOR_DEPLOYMENT_DIR/virtualbox/create-env.sh

source .envrc
bosh -e vbox env

bosh -e vbox update-cloud-config bosh-deployment/warden/cloud-config.yml
bosh upload-stemcell --sha1 d44dc2d1b3f8415b41160ad4f82bc9d30b8dfdce \
https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-bionic-go_agent?v=1.71

cd $SCRIPT_DIR