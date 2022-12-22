// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderSamlBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckIAMWorkforcePoolWorkforcePoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderSamlBasicExample(context),
			},
			{
				ResourceName:            "google_iam_workforce_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "workforce_pool_id", "provider_id"},
			},
		},
	})
}

func testAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderSamlBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workforce_pool" "pool" {
  provider = google-beta

  workforce_pool_id = "tf-test-example-pool%{random_suffix}"
  parent            = "organizations/%{org_id}"
  location          = "global"
}

resource "google_iam_workforce_pool_provider" "example" {
  provider = google-beta

  workforce_pool_id  = google_iam_workforce_pool.pool.workforce_pool_id
  location           = google_iam_workforce_pool.pool.location
  provider_id        = "tf-test-example-prvdr%{random_suffix}"
  attribute_mapping  = {
    "google.subject" = "assertion.sub"
  }
  saml {
    idp_metadata_xml = "<?xml version=\"1.0\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://test.com\"><md:IDPSSODescriptor protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"> <md:KeyDescriptor use=\"signing\"><ds:KeyInfo xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><ds:X509Data><ds:X509Certificate>MIIDpDCCAoygAwIBAgIGAX7/5qPhMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMMCmRldi00NTg0MjExHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMjIwMjE2MDAxOTEyWhcNMzIwMjE2MDAyMDEyWjCBkjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoMBE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNDU4NDIxMRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxrBl7GKz52cRpxF9xCsirnRuMxnhFBaUrsHqAQrLqWmdlpNYZTVg+T9iQ+aq/iE68L+BRZcZniKIvW58wqqS0ltXVvIkXuDSvnvnkkI5yMIVErR20K8jSOKQm1FmK+fgAJ4koshFiu9oLiqu0Ejc0DuL3/XRsb4RuxjktKTb1khgBBtb+7idEk0sFR0RPefAweXImJkDHDm7SxjDwGJUubbqpdTxasPr0W+AHI1VUzsUsTiHAoyb0XDkYqHfDzhj/ZdIEl4zHQ3bEZvlD984ztAnmX2SuFLLKfXeAAGHei8MMixJvwxYkkPeYZ/5h8WgBZPP4heS2CPjwYExt29L8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQARjJFz++a9Z5IQGFzsZMrX2EDR5ML4xxUiQkbhld1S1PljOLcYFARDmUC2YYHOueU4ee8Jid9nPGEUebV/4Jok+b+oQh+dWMgiWjSLI7h5q4OYZ3VJtdlVwgMFt2iz+/4yBKMUZ50g3Qgg36vE34us+eKitg759JgCNsibxn0qtJgSPm0sgP2L6yTaLnoEUbXBRxCwynTSkp9ZijZqEzbhN0e2dWv7Rx/nfpohpDP6vEiFImKFHpDSv3M/5de1ytQzPFrZBYt9WlzlYwE1aD9FHCxdd+rWgYMVVoRaRmndpV/Rq3QUuDuFJtaoX11bC7ExkOpg9KstZzA63i3VcfYv</ds:X509Certificate></ds:X509Data></ds:KeyInfo></md:KeyDescriptor><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://test.com/sso\"/></md:IDPSSODescriptor></md:EntityDescriptor>"
  }
}
`, context)
}

func TestAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderSamlFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckIAMWorkforcePoolWorkforcePoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderSamlFullExample(context),
			},
			{
				ResourceName:            "google_iam_workforce_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "workforce_pool_id", "provider_id"},
			},
		},
	})
}

func testAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderSamlFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workforce_pool" "pool" {
  provider = google-beta

  workforce_pool_id = "tf-test-example-pool%{random_suffix}"
  parent            = "organizations/%{org_id}"
  location          = "global"
}

resource "google_iam_workforce_pool_provider" "example" {
  provider = google-beta

  workforce_pool_id   = google_iam_workforce_pool.pool.workforce_pool_id
  location            = google_iam_workforce_pool.pool.location
  provider_id         = "tf-test-example-prvdr%{random_suffix}"
  attribute_mapping   = {
    "google.subject"  = "assertion.sub"
  }
  saml {
    idp_metadata_xml  = "<?xml version=\"1.0\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://test.com\"><md:IDPSSODescriptor protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"> <md:KeyDescriptor use=\"signing\"><ds:KeyInfo xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><ds:X509Data><ds:X509Certificate>MIIDpDCCAoygAwIBAgIGAX7/5qPhMA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEzARBgNVBAMMCmRldi00NTg0MjExHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMjIwMjE2MDAxOTEyWhcNMzIwMjE2MDAyMDEyWjCBkjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoMBE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRMwEQYDVQQDDApkZXYtNDU4NDIxMRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxrBl7GKz52cRpxF9xCsirnRuMxnhFBaUrsHqAQrLqWmdlpNYZTVg+T9iQ+aq/iE68L+BRZcZniKIvW58wqqS0ltXVvIkXuDSvnvnkkI5yMIVErR20K8jSOKQm1FmK+fgAJ4koshFiu9oLiqu0Ejc0DuL3/XRsb4RuxjktKTb1khgBBtb+7idEk0sFR0RPefAweXImJkDHDm7SxjDwGJUubbqpdTxasPr0W+AHI1VUzsUsTiHAoyb0XDkYqHfDzhj/ZdIEl4zHQ3bEZvlD984ztAnmX2SuFLLKfXeAAGHei8MMixJvwxYkkPeYZ/5h8WgBZPP4heS2CPjwYExt29L8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQARjJFz++a9Z5IQGFzsZMrX2EDR5ML4xxUiQkbhld1S1PljOLcYFARDmUC2YYHOueU4ee8Jid9nPGEUebV/4Jok+b+oQh+dWMgiWjSLI7h5q4OYZ3VJtdlVwgMFt2iz+/4yBKMUZ50g3Qgg36vE34us+eKitg759JgCNsibxn0qtJgSPm0sgP2L6yTaLnoEUbXBRxCwynTSkp9ZijZqEzbhN0e2dWv7Rx/nfpohpDP6vEiFImKFHpDSv3M/5de1ytQzPFrZBYt9WlzlYwE1aD9FHCxdd+rWgYMVVoRaRmndpV/Rq3QUuDuFJtaoX11bC7ExkOpg9KstZzA63i3VcfYv</ds:X509Certificate></ds:X509Data></ds:KeyInfo></md:KeyDescriptor><md:SingleSignOnService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\" Location=\"https://test.com/sso\"/></md:IDPSSODescriptor></md:EntityDescriptor>"
  }
  display_name        = "Display name"
  description         = "A sample SAML workforce pool provider."
  disabled            = false
  attribute_condition = "true"
}
`, context)
}

func TestAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderOidcBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckIAMWorkforcePoolWorkforcePoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderOidcBasicExample(context),
			},
			{
				ResourceName:            "google_iam_workforce_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "workforce_pool_id", "provider_id"},
			},
		},
	})
}

func testAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderOidcBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workforce_pool" "pool" {
  provider = google-beta

  workforce_pool_id = "tf-test-example-pool%{random_suffix}"
  parent            = "organizations/%{org_id}"
  location          = "global"
}

resource "google_iam_workforce_pool_provider" "example" {
  provider = google-beta

  workforce_pool_id  = google_iam_workforce_pool.pool.workforce_pool_id
  location           = google_iam_workforce_pool.pool.location
  provider_id        = "tf-test-example-prvdr%{random_suffix}"
  attribute_mapping  = {
    "google.subject" = "assertion.sub"
  }
  oidc {
    issuer_uri       = "https://accounts.google.com"
    client_id        = "client-id"
  }
}
`, context)
}

func TestAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderOidcFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckIAMWorkforcePoolWorkforcePoolProviderDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderOidcFullExample(context),
			},
			{
				ResourceName:            "google_iam_workforce_pool_provider.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "workforce_pool_id", "provider_id"},
			},
		},
	})
}

func testAccIAMWorkforcePoolWorkforcePoolProvider_iamWorkforcePoolProviderOidcFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workforce_pool" "pool" {
  provider = google-beta

  workforce_pool_id = "tf-test-example-pool%{random_suffix}"
  parent            = "organizations/%{org_id}"
  location          = "global"
}

resource "google_iam_workforce_pool_provider" "example" {
  provider = google-beta

  workforce_pool_id   = google_iam_workforce_pool.pool.workforce_pool_id
  location            = google_iam_workforce_pool.pool.location
  provider_id         = "tf-test-example-prvdr%{random_suffix}"
  attribute_mapping   = {
    "google.subject"  = "assertion.sub"
  }
  oidc {
    issuer_uri        = "https://accounts.google.com"
    client_id         = "client-id"
  }
  display_name        = "Display name"
  description         = "A sample OIDC workforce pool provider."
  disabled            = false
  attribute_condition = "true"
}
`, context)
}

func testAccCheckIAMWorkforcePoolWorkforcePoolProviderDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iam_workforce_pool_provider" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{IAMWorkforcePoolBasePath}}locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("IAMWorkforcePoolWorkforcePoolProvider still exists at %s", url)
			}
		}

		return nil
	}
}
