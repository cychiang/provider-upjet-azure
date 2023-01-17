/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ResourceGroupCostManagementExport
func (mg *ResourceGroupCostManagementExport) GetTerraformResourceType() string {
	return "azurerm_resource_group_cost_management_export"
}

// GetConnectionDetailsMapping for this ResourceGroupCostManagementExport
func (tr *ResourceGroupCostManagementExport) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ResourceGroupCostManagementExport
func (tr *ResourceGroupCostManagementExport) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ResourceGroupCostManagementExport
func (tr *ResourceGroupCostManagementExport) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ResourceGroupCostManagementExport
func (tr *ResourceGroupCostManagementExport) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ResourceGroupCostManagementExport
func (tr *ResourceGroupCostManagementExport) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ResourceGroupCostManagementExport
func (tr *ResourceGroupCostManagementExport) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ResourceGroupCostManagementExport using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ResourceGroupCostManagementExport) LateInitialize(attrs []byte) (bool, error) {
	params := &ResourceGroupCostManagementExportParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ResourceGroupCostManagementExport) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubscriptionCostManagementExport
func (mg *SubscriptionCostManagementExport) GetTerraformResourceType() string {
	return "azurerm_subscription_cost_management_export"
}

// GetConnectionDetailsMapping for this SubscriptionCostManagementExport
func (tr *SubscriptionCostManagementExport) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubscriptionCostManagementExport
func (tr *SubscriptionCostManagementExport) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubscriptionCostManagementExport
func (tr *SubscriptionCostManagementExport) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubscriptionCostManagementExport
func (tr *SubscriptionCostManagementExport) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubscriptionCostManagementExport
func (tr *SubscriptionCostManagementExport) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubscriptionCostManagementExport
func (tr *SubscriptionCostManagementExport) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SubscriptionCostManagementExport using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubscriptionCostManagementExport) LateInitialize(attrs []byte) (bool, error) {
	params := &SubscriptionCostManagementExportParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubscriptionCostManagementExport) GetTerraformSchemaVersion() int {
	return 0
}
