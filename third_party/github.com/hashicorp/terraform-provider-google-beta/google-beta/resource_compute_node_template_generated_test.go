// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeNodeTemplate_nodeTemplateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNodeTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeTemplate_nodeTemplateBasicExample(context),
			},
			{
				ResourceName:            "google_compute_node_template.template",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeNodeTemplate_nodeTemplateBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_node_template" "template" {
  name      = "tf-test-soletenant-tmpl%{random_suffix}"
  region    = "us-central1"
  node_type = "n1-node-96-624"
}
`, context)
}

func TestAccComputeNodeTemplate_nodeTemplateServerBindingExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNodeTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNodeTemplate_nodeTemplateServerBindingExample(context),
			},
			{
				ResourceName:            "google_compute_node_template.template",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeNodeTemplate_nodeTemplateServerBindingExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_compute_node_types" "central1a" {
  zone     = "us-central1-a"
}

resource "google_compute_node_template" "template" {
  name      = "tf-test-soletenant-with-licenses%{random_suffix}"
  region    = "us-central1"
  node_type = "n1-node-96-624"

  node_affinity_labels = {
    foo = "baz"
  }

  server_binding {
    type = "RESTART_NODE_ON_MINIMAL_SERVERS"
  }
}
`, context)
}

func testAccCheckComputeNodeTemplateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_node_template" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeNodeTemplate still exists at %s", url)
			}
		}

		return nil
	}
}
