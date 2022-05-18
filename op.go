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

import (
	"context"
	"encoding/json"
	"fmt"
)

type APICaller interface {
	Do(ctx context.Context, method, uri string, body interface{}) ([]byte, error)
	RootURL() string
}

var _ API = (*Op)(nil)

// Op implements WebAccelAPI interface
type Op struct {
	Client APICaller
}

// NewOp creates new Op instance
func NewOp(caller APICaller) API {
	return &Op{Client: caller}
}

// List is API call
func (o *Op) List(ctx context.Context) (*ListSitesResult, error) {
	url := o.Client.RootURL() + "site"

	// build request body
	var body interface{}

	// do request
	data, err := o.Client.Do(ctx, "GET", url, body)
	if err != nil {
		return nil, err
	}

	// build results
	var results ListSitesResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}
	return &results, nil
}

// Read is API call
func (o *Op) Read(ctx context.Context, id string) (*Site, error) {
	url := o.Client.RootURL() + fmt.Sprintf("site/%s", id)

	// build request body
	var body interface{}

	// do request
	data, err := o.Client.Do(ctx, "GET", url, body)
	if err != nil {
		return nil, err
	}

	// build results
	type readResult struct {
		Site *Site
	}
	var results readResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}
	return results.Site, nil
}

// ReadCertificate is API call
func (o *Op) ReadCertificate(ctx context.Context, id string) (*Certificates, error) {
	url := o.Client.RootURL() + fmt.Sprintf("site/%s/certificate", id)

	// build request body
	var body interface{}

	// do request
	data, err := o.Client.Do(ctx, "GET", url, body)
	if err != nil {
		return nil, err
	}

	type readCertificateResult struct {
		Certificate *Certificates
	}

	// build results
	var results readCertificateResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}
	return results.Certificate, nil
}

// CreateCertificate is API call
func (o *Op) CreateCertificate(ctx context.Context, id string, param *CreateOrUpdateCertificateRequest) (*Certificates, error) {
	url := o.Client.RootURL() + fmt.Sprintf("site/%s/certificate", id)

	// build request body
	type createRequest struct {
		Certificate *CreateOrUpdateCertificateRequest
	}
	body := &createRequest{Certificate: param}

	// do request
	data, err := o.Client.Do(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	type createCertificateResult struct {
		Certificate *Certificates
	}

	// build results
	var results createCertificateResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}
	return results.Certificate, nil
}

// UpdateCertificate is API call
func (o *Op) UpdateCertificate(ctx context.Context, id string, param *CreateOrUpdateCertificateRequest) (*Certificates, error) {
	url := o.Client.RootURL() + fmt.Sprintf("site/%s/certificate", id)

	// build request body
	type updateRequest struct {
		Certificate *CreateOrUpdateCertificateRequest
	}
	body := &updateRequest{Certificate: param}

	// do request
	data, err := o.Client.Do(ctx, "PUT", url, body)
	if err != nil {
		return nil, err
	}

	type updateCertificateResult struct {
		Certificate *Certificates
	}

	// build results
	var results updateCertificateResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}
	return results.Certificate, nil
}

// DeleteCertificate is API call
func (o *Op) DeleteCertificate(ctx context.Context, id string) error {
	url := o.Client.RootURL() + fmt.Sprintf("site/%s/certificate", id)

	// build request body
	var body interface{}

	// do request
	_, err := o.Client.Do(ctx, "DELETE", url, body)
	return err
}

// DeleteAllCache is API call
func (o *Op) DeleteAllCache(ctx context.Context, param *DeleteAllCacheRequest) error {
	url := o.Client.RootURL() + "deleteallcache"

	// build request body
	type deleteAllCacheRequest struct {
		Site *DeleteAllCacheRequest
	}
	body := &deleteAllCacheRequest{Site: param}

	// do request
	_, err := o.Client.Do(ctx, "POST", url, body)
	return err
}

// DeleteCache is API call
func (o *Op) DeleteCache(ctx context.Context, param *DeleteCacheRequest) ([]*DeleteCacheResult, error) {
	url := o.Client.RootURL() + "deletecache"

	// do request
	data, err := o.Client.Do(ctx, "POST", url, param)
	if err != nil {
		return nil, err
	}

	// build results
	type deleteCacheResult struct {
		Results []*DeleteCacheResult
	}
	var results deleteCacheResult
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}
	return results.Results, nil
}
