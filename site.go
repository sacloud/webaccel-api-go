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

import "time"

// Site サイト
type Site struct {
	ID                 string
	Name               string
	DomainType         string `validate:"oneof=own_domain subdomain"`
	Domain             string
	Subdomain          string
	ASCIIDomain        string
	Origin             string
	HostHeader         string
	Status             string `validate:"oneof=enabled disabled"`
	HasCertificate     bool
	HasOldCertificate  bool
	GibSentInLastWeek  int64
	CertValidNotBefore int64
	CertValidNotAfter  int64
	CreatedAt          time.Time
}
