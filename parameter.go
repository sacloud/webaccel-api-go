// Copyright 2022 The webaccel-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webaccel

// DeleteAllCacheRequest .
type DeleteAllCacheRequest struct {
	Domain string
}

// DeleteCacheRequest .
type DeleteCacheRequest struct {
	URL []string
}

// DeleteCacheResult .
type DeleteCacheResult struct {
	URL    string
	Status int
	Result string
}

// ListSitesResult .
type ListSitesResult struct {
	Total int `json:",omitempty"` // Total count of target resources
	From  int `json:",omitempty"` // Current page number
	Count int `json:",omitempty"` // Count of current page

	Sites []*Site `json:",omitempty"`
}

// CreateOrUpdateCertificateRequest .
type CreateOrUpdateCertificateRequest struct {
	CertificateChain string
	Key              string
}
