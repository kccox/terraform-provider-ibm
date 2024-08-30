// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEventStreamsSchemaGlobalCompatibilityRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleResourceConfig(getTestInstanceName(mzrKey), "FORWARD"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleProperties("ibm_event_streams_schema_global_compatibility_rule.es_globalrule", "FORWARD"),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleResourceConfig(getTestInstanceName(mzrKey), "NONE"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleProperties("ibm_event_streams_schema_global_compatibility_rule.es_globalrule", "NONE"),
				),
			},
			{
				ResourceName:      "ibm_event_streams_schema_global_compatibility_rule.es_globalrule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEventStreamsSchemaGlobalCompatibilityRuleResourceConfig(instanceName string, rule string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	data "ibm_resource_instance" "es_instance" {
		resource_group_id = data.ibm_resource_group.group.id
		name              = "%s"
	}
	resource "ibm_event_streams_schema_global_compatibility_rule" "es_globalrule" {
		resource_instance_id = data.ibm_resource_instance.es_instance.id
		config = "%s"
	}`, instanceName, rule)
}
