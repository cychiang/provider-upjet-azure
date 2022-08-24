/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IdentityObservation struct {

	// The Principal ID of the Policy Assignment for this Resource Group.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID of the Policy Assignment for this Resource Group.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// A list of User Managed Identity IDs which should be assigned to the Policy Definition.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Type of Managed Identity which should be added to this Policy Definition. Possible values are SystemAssigned and UserAssigned.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NonComplianceMessageObservation struct {
}

type NonComplianceMessageParameters struct {

	// The non-compliance message text. When assigning policy sets , unless policy_definition_reference_id is specified then this message will be the default for all policies.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// When assigning policy sets , this is the ID of the policy definition that the non-compliance message applies to.
	// +kubebuilder:validation:Optional
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`
}

type ResourceGroupPolicyAssignmentObservation struct {

	// The ID of the Resource Group Policy Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`
}

type ResourceGroupPolicyAssignmentParameters struct {

	// A description which should be used for this Policy Assignment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Display Name for this Policy Assignment.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Specifies if this Policy should be enforced or not?
	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Azure Region where the Policy Assignment should exist. Changing this forces a new Policy Assignment to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A JSON mapping of any Metadata for this Policy.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// One or more non_compliance_message blocks as defined below.
	// +kubebuilder:validation:Optional
	NonComplianceMessage []NonComplianceMessageParameters `json:"nonComplianceMessage,omitempty" tf:"non_compliance_message,omitempty"`

	// Specifies a list of Resource Scopes  within this Management Group which are excluded from this Policy.
	// +kubebuilder:validation:Optional
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// A JSON mapping of any Parameters for this Policy.
	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The ID of the Policy Definition or Policy Definition Set. Changing this forces a new Policy Assignment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/authorization/v1beta1.PolicyDefinition
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty" tf:"policy_definition_id,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyDefinitionIDRef *v1.Reference `json:"policyDefinitionIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PolicyDefinitionIDSelector *v1.Selector `json:"policyDefinitionIdSelector,omitempty" tf:"-"`

	// The ID of the Resource Group where this Policy Assignment should be created. Changing this forces a new Policy Assignment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupIDRef *v1.Reference `json:"resourceGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupIDSelector *v1.Selector `json:"resourceGroupIdSelector,omitempty" tf:"-"`
}

// ResourceGroupPolicyAssignmentSpec defines the desired state of ResourceGroupPolicyAssignment
type ResourceGroupPolicyAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupPolicyAssignmentParameters `json:"forProvider"`
}

// ResourceGroupPolicyAssignmentStatus defines the observed state of ResourceGroupPolicyAssignment.
type ResourceGroupPolicyAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupPolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupPolicyAssignment is the Schema for the ResourceGroupPolicyAssignments API. Manages a Resource Group Policy Assignment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGroupPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupPolicyAssignmentSpec   `json:"spec"`
	Status            ResourceGroupPolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupPolicyAssignmentList contains a list of ResourceGroupPolicyAssignments
type ResourceGroupPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroupPolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroupPolicyAssignment_Kind             = "ResourceGroupPolicyAssignment"
	ResourceGroupPolicyAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroupPolicyAssignment_Kind}.String()
	ResourceGroupPolicyAssignment_KindAPIVersion   = ResourceGroupPolicyAssignment_Kind + "." + CRDGroupVersion.String()
	ResourceGroupPolicyAssignment_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroupPolicyAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroupPolicyAssignment{}, &ResourceGroupPolicyAssignmentList{})
}
