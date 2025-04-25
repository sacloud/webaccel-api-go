// Copyright 2022-2025 The webaccel-api-go authors
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

// OriginType
const (
	OriginTypesWebServer     = "0" // ウェブサーバ
	OriginTypesObjectStorage = "1" // オブジェクトストレージ
)

// RequestProtocol
const (
	RequestProtocolsHttpAndHttps    = "0" // http/https
	RequestProtocolsHttpsOnly       = "1" // httpsのみ
	RequestProtocolsRedirectToHttps = "2" // httpsに」リダイレクト
)

// OriginProtocol
const (
	OriginProtocolsHttp  = "http"
	OriginProtocolsHttps = "https"
)

// VarySupport
const (
	VarySupportDisabled = "0" // 無効
	VarySupportEnabled  = "1" // 有効
)

// DocIndex
const (
	DocIndexDisabled = "0" // 無効
	DocIndexEnabled  = "1" // 有効
)

// NormalizeAE
const (
	NormalizeAEGz   = "1" // gzipに正規化
	NormalizeAEBzGz = "3" // bzとgzipの組に正規化
)
