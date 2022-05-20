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

import "context"

// API is interface for operate WebAccel resource
type API interface {
	List(ctx context.Context) (*ListSitesResult, error)
	Read(ctx context.Context, id string) (*Site, error)
	ReadCertificate(ctx context.Context, id string) (*Certificates, error)
	CreateCertificate(ctx context.Context, id string, param *CreateOrUpdateCertificateRequest) (*Certificates, error)
	UpdateCertificate(ctx context.Context, id string, param *CreateOrUpdateCertificateRequest) (*Certificates, error)
	DeleteCertificate(ctx context.Context, id string) error
	DeleteAllCache(ctx context.Context, param *DeleteAllCacheRequest) error
	DeleteCache(ctx context.Context, param *DeleteCacheRequest) ([]*DeleteCacheResult, error)
	MonthlyUsage(ctx context.Context, targetYM string) (*MonthlyUsageResults, error)
}
