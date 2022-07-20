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

type DataSetDataLakeGen2Observation struct {
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataSetDataLakeGen2Parameters struct {

	// +kubebuilder:validation:Optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.DataLakeGen2FileSystem
	// +kubebuilder:validation:Optional
	FileSystemName *string `json:"fileSystemName,omitempty" tf:"file_system_name,omitempty"`

	// +kubebuilder:validation:Optional
	FileSystemNameRef *v1.Reference `json:"fileSystemNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FileSystemNameSelector *v1.Selector `json:"fileSystemNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// +crossplane:generate:reference:type=DataShare
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ShareID *string `json:"shareId,omitempty" tf:"share_id,omitempty"`

	// +kubebuilder:validation:Optional
	ShareIDRef *v1.Reference `json:"shareIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ShareIDSelector *v1.Selector `json:"shareIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

// DataSetDataLakeGen2Spec defines the desired state of DataSetDataLakeGen2
type DataSetDataLakeGen2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetDataLakeGen2Parameters `json:"forProvider"`
}

// DataSetDataLakeGen2Status defines the observed state of DataSetDataLakeGen2.
type DataSetDataLakeGen2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetDataLakeGen2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetDataLakeGen2 is the Schema for the DataSetDataLakeGen2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetDataLakeGen2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetDataLakeGen2Spec   `json:"spec"`
	Status            DataSetDataLakeGen2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetDataLakeGen2List contains a list of DataSetDataLakeGen2s
type DataSetDataLakeGen2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetDataLakeGen2 `json:"items"`
}

// Repository type metadata.
var (
	DataSetDataLakeGen2_Kind             = "DataSetDataLakeGen2"
	DataSetDataLakeGen2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetDataLakeGen2_Kind}.String()
	DataSetDataLakeGen2_KindAPIVersion   = DataSetDataLakeGen2_Kind + "." + CRDGroupVersion.String()
	DataSetDataLakeGen2_GroupVersionKind = CRDGroupVersion.WithKind(DataSetDataLakeGen2_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetDataLakeGen2{}, &DataSetDataLakeGen2List{})
}
