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

type WorkspaceObservation struct {

	// The Log Analytics Workspace ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Workspace (or Customer) ID for the Log Analytics Workspace.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type WorkspaceParameters struct {

	// Is Customer Managed Storage mandatory for query management?
	// +kubebuilder:validation:Optional
	CmkForQueryForced *bool `json:"cmkForQueryForced,omitempty" tf:"cmk_for_query_forced,omitempty"`

	// The workspace daily quota for ingestion in GB.  Defaults to -1 (unlimited) if omitted.
	// +kubebuilder:validation:Optional
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty" tf:"daily_quota_gb,omitempty"`

	// Should the Log Analytics Workspace support ingestion over the Public Internet? Defaults to true.
	// +kubebuilder:validation:Optional
	InternetIngestionEnabled *bool `json:"internetIngestionEnabled,omitempty" tf:"internet_ingestion_enabled,omitempty"`

	// Should the Log Analytics Workspace support querying over the Public Internet? Defaults to true.
	// +kubebuilder:validation:Optional
	InternetQueryEnabled *bool `json:"internetQueryEnabled,omitempty" tf:"internet_query_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The capacity reservation level in GB for this workspace.  Must be in increments of 100  between 100 and 5000.
	// +kubebuilder:validation:Optional
	ReservationCapacityInGbPerDay *float64 `json:"reservationCapacityInGbPerDay,omitempty" tf:"reservation_capacity_in_gb_per_day,omitempty"`

	// The name of the resource group in which the Log Analytics workspace is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The workspace data retention in days. Possible values are either 7 (Free Tier only) or range between 30 and 730.
	// +kubebuilder:validation:Optional
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// Specifies the SKU of the Log Analytics Workspace. Possible values are Free, PerNode, Premium, Standard, Standalone, Unlimited, CapacityReservation, and PerGB2018 (new SKU as of 2018-04-03). Defaults to PerGB2018.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workspace is the Schema for the Workspaces API. Manages a Log Analytics (formally Operational Insights) Workspace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
