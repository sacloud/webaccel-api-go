package webaccel

import (
	"testing"
)

func TestMapRequestProtocolToNickname(t *testing.T) {
	tt := []struct {
		Name        string
		Given       *Site
		Want        string
		ExpectError bool
	}{
		{
			"valid http+https",
			&Site{
				RequestProtocol: RequestProtocolsHttpAndHttps,
			},
			httpOrHttpsRequestProtocolNickname,
			false,
		},
		{
			"valid https",
			&Site{
				RequestProtocol: RequestProtocolsHttpsOnly,
			},
			httpsOnlyRequestProtocolNickname,
			false,
		},
		{
			"valid https-redirect",
			&Site{
				RequestProtocol: RequestProtocolsRedirectToHttps,
			},
			httpsRedirectedRequestProtocolNickname,
			false,
		},
		{
			"invalid request protocol",
			&Site{
				RequestProtocol: "NO-SUCH-RP",
			},
			"",
			true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			res, err := MapRequestProtocolToNickname(tc.Given)
			if tc.ExpectError {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
			} else if res != tc.Want {
				t.Fatalf("FAILED %s: got: %v\nwant: %v", tc.Name, res, tc.Want)
			}
		})
	}
}

func TestMapNormalizeAEValueToNickname(t *testing.T) {
	tt := []struct {
		Name        string
		Given       *Site
		Want        string
		ExpectError bool
	}{
		{
			"valid gzip",
			&Site{
				NormalizeAE: NormalizeAEGz,
			},
			gunzipCompressionNickname,
			false,
		},
		{
			"valid brotli",
			&Site{
				NormalizeAE: NormalizeAEBrGz,
			},
			brotliCompressionNickname,
			false,
		},
		{
			"invalid encoding",
			&Site{
				NormalizeAE: "3-NO-SUCH-NORMALIZE-AE-PARAM",
			},
			"",
			true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			res, err := MapNormalizeAEValueToNickname(tc.Given)
			if tc.ExpectError {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
			} else if res != tc.Want {
				t.Fatalf("FAILED %s: got: %v\nwant: %v", tc.Name, res, tc.Want)
			}
		})
	}
}
