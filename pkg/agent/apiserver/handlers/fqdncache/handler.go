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
	"net/http"
	"net/url"

	"k8s.io/klog/v2"

	agentapi "antrea.io/antrea/pkg/agent/apis"
	agentquerier "antrea.io/antrea/pkg/agent/querier"
	"antrea.io/antrea/pkg/querier"
)

func HandleFunc(aq agentquerier.AgentQuerier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fqdnFilter := newFilterFromURLQuery(r.URL.Query())
		dnsEntryCache := aq.GetNetworkPolicyInfoQuerier().GetFQDNCache(fqdnFilter)
		resp := []agentapi.FQDNCacheResponse{}
		for _, entry := range dnsEntryCache {
			resp = append(resp, agentapi.FQDNCacheResponse{
				FqdnName:       entry.FqdnName,
				IpAddress:      entry.IpAddress.String(),
				ExpirationTime: entry.ExpirationTime,
			})
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
			klog.ErrorS(err, "Failed to encode response")
		}
	}
}

func newFilterFromURLQuery(query url.Values) querier.FQDNCacheFilter {
	return querier.FQDNCacheFilter{DomainName: query.Get("domain")}
}
