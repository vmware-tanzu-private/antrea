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
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	queriertest "antrea.io/antrea/pkg/agent/querier/testing"
	"antrea.io/antrea/pkg/agent/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestFqdnCacheQuery(t *testing.T) {
	tests := []struct {
		name             string
		expectedStatus   int
		expectedResponse []types.DnsCacheEntry
	}{
		{
			name:           "FQDN cache exists",
			expectedStatus: http.StatusOK,
			expectedResponse: []types.DnsCacheEntry{
				{
					FqdnName:       "google.com",
					IpAddress:      net.ParseIP("10.0.0.1"),
					ExpirationTime: time.Date(2025, 12, 25, 15, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name:             "FQDN cache does not exist",
			expectedStatus:   http.StatusOK,
			expectedResponse: []types.DnsCacheEntry{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			q := queriertest.NewMockAgentQuerier(ctrl)
			q.EXPECT().GetFqdnCache().Return(tt.expectedResponse)
			handler := HandleFunc(q)
			req, err := http.NewRequest(http.MethodGet, "", nil)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req)
			assert.Equal(t, tt.expectedStatus, recorder.Code)
			if tt.expectedStatus == http.StatusOK {
				var received []map[string]interface{}
				err = json.Unmarshal(recorder.Body.Bytes(), &received)
				require.NoError(t, err)
				if len(received) > 0 {
					parsedTime, err := time.Parse(time.RFC3339, received[0]["expirationTime"].(string))
					require.NoError(t, err)
					assert.Equal(t, tt.expectedResponse, []types.DnsCacheEntry{{
						FqdnName:       received[0]["fqdnName"].(string),
						IpAddress:      net.ParseIP(received[0]["ipAddress"].(string)),
						ExpirationTime: parsedTime,
					}})
				}
			}
		})
	}
}
