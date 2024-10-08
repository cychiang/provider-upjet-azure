// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package management

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
)

// Configure configures management group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_management_group", func(r *config.Resource) {
		r.Kind = "ManagementGroup"
	})
	p.AddResourceConfigurator("azurerm_management_group_subscription_association", func(r *config.Resource) {
		r.Kind = "ManagementGroupSubscriptionAssociation"
		r.References["management_group_id"] = config.Reference{
			TerraformName: "azurerm_management_group",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
		r.References["subscription_id"] = config.Reference{
			TerraformName: "azurerm_subscription",
			Extractor:     rconfig.ExtractResourceIDFuncPath,
		}
	})
}
