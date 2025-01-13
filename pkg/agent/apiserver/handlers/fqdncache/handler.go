// Copyright 2025 Antrea Authors
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

package fqdncache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"k8s.io/klog/v2"

	agentquerier "antrea.io/antrea/pkg/agent/querier"
	"antrea.io/antrea/pkg/querier"
)

func HandleFunc(aq agentquerier.AgentQuerier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fqdnFilter, err := newFilterFromURLQuery(r.URL.Query())
		if err != nil {
			klog.ErrorS(err, "Failed to create filter from query")
		}
		dnsEntryCache := aq.GetFqdnCache(fqdnFilter)
		if err := json.NewEncoder(w).Encode(dnsEntryCache); err != nil {
			http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
			klog.ErrorS(err, "Failed to encode response")
		}
	}
}

func newFilterFromURLQuery(query url.Values) (querier.FQDNCacheFilter, error) {
	fmt.Printf("query: %v\n", query)
	return querier.FQDNCacheFilter{}, nil
}
