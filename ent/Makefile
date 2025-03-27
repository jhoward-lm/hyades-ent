# --------------------------------------------------------------------
# SPDX-FileCopyrightText: Copyright © 2025 The DependencyTrack Authors
# SPDX-FileName: ent/Makefile
# SPDX-FileType: SOURCE
# SPDX-License-Identifier: Apache-2.0
# --------------------------------------------------------------------
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# --------------------------------------------------------------------
ENT_ROOT_DIR := ${abspath ${patsubst %/,%,${dir ${lastword ${MAKEFILE_LIST}}}}}

.PHONY: generate-ent describe-schemas clean-ent
.SILENT: generate-ent describe-schemas clean-ent

#@ Ent Utilities
clean-ent: # Clean generated ent code
	cd ${ENT_ROOT_DIR} && find . -type d -mindepth 1 -maxdepth 1 \
	  -not "(" -path "./schema*" -or -path "./template*" ")" \
	  -exec ${RM} -r {} \;

	cd ${ENT_ROOT_DIR} && find . -maxdepth 1 -name "*.go" \
	  -not "(" -name entc.go -or -name generate.go ")" \
	  -exec ${RM} {} \;

generate-ent: # Generate ent database types from schema
	printf "Generating ${CYAN}ent${RESET} database types with ${CYAN}go generate${RESET}..."
	go generate ${ENT_ROOT_DIR}
	printf "${GREEN}DONE${RESET}\n"

describe-schemas: # Print descriptions of generated ent schema types
	go run -mod=mod entgo.io/ent/cmd/ent describe ${ENT_ROOT_DIR}/schema
