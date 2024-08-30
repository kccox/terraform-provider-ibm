// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams

import (
	"context"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/eventstreams-go-sdk/pkg/schemaregistryv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	validCompatibilityConfigValues = []string{"NONE", "FULL", "FULL_TRANSITIVE", "FORWARD", "FORWARD_TRANSITIVE", "BACKWARD", "BACKWARD_TRANSITIVE"}
)

// The global compatibility rule in an Event Streams service instance.
// The ID is the CRN with the last two components "schema-global-compatibility:".
// The rule is the schema compatibility rule, one of the validRules values.
func ResourceIBMEventStreamsSchemaGlobalCompatibilityRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMEventStreamsSchemaGlobalCompatibilityRuleUpdate,
		ReadContext:   resourceIBMEventStreamsSchemaGlobalCompatibilityRuleRead,
		UpdateContext: resourceIBMEventStreamsSchemaGlobalCompatibilityRuleUpdate,
		DeleteContext: resourceIBMEventStreamsSchemaGlobalCompatibilityRuleDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"resource_instance_id": {
				Type:        schema.TypeString,
				Description: "The ID or the CRN of the Event Streams service instance",
				Required:    true,
				ForceNew:    true,
			},
			"config": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(validCompatibilityConfigValues, true),
				Description:  "The value of the global schema compatibility rule",
			},
		},
	}
}

func resourceIBMEventStreamsSchemaGlobalCompatibilityRuleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return dataSourceIBMEventStreamsSchemaGlobalCompatibilityRuleRead(context, d, meta)
}

// The global compatibility rule is always defined in the schema registry,
// so create and update have the same behavior
func resourceIBMEventStreamsSchemaGlobalCompatibilityRuleUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schemaregistryClient, err := meta.(conns.ClientSession).ESschemaRegistrySession()
	if err != nil {
		return diag.FromErr(err)
	}
	adminURL, _, err := getSchemaRuleInstanceURL(d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	schemaregistryClient.SetServiceURL(adminURL)

	updateOpts := &schemaregistryv1.UpdateGlobalRuleOptions{}
	updateOpts.SetType("COMPATIBILITY")
	updateOpts.SetRule("COMPATIBILITY")
	updateOpts.SetConfig(d.Get("config").(string))

	_, _, err = schemaregistryClient.UpdateGlobalRuleWithContext(context, updateOpts)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceIBMEventStreamsSchemaGlobalCompatibilityRuleRead(context, d, meta)
}

// The global rule can't be deleted, so just clear the ID
func resourceIBMEventStreamsSchemaGlobalCompatibilityRuleDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}