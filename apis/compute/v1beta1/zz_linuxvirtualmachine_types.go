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

type AdditionalCapabilitiesObservation struct {
}

type AdditionalCapabilitiesParameters struct {

	// Should the capacity to enable Data Disks of the UltraSSD_LRS storage account type be supported on this Virtual Machine? Defaults to false.
	// +kubebuilder:validation:Optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled,omitempty"`
}

type AdminSSHKeyObservation struct {
}

type AdminSSHKeyParameters struct {

	// The Public Key which should be used for authentication, which needs to be at least 2048-bit and in ssh-rsa format. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`

	// The Username for which this Public SSH Key should be configured. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type BootDiagnosticsObservation struct {
}

type BootDiagnosticsParameters struct {

	// The Primary/Secondary Endpoint for the Azure Storage Account which should be used to store Boot Diagnostics, including Console Output and Screenshots from the Hypervisor.
	// +kubebuilder:validation:Optional
	StorageAccountURI *string `json:"storageAccountUri,omitempty" tf:"storage_account_uri,omitempty"`
}

type CertificateObservation struct {
}

type CertificateParameters struct {

	// The Secret URL of a Key Vault Certificate.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type DiffDiskSettingsObservation struct {
}

type DiffDiskSettingsParameters struct {

	// Specifies the Ephemeral Disk Settings for the OS Disk. At this time the only possible value is Local. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Option *string `json:"option" tf:"option,omitempty"`

	// Specifies where to store the Ephemeral Disk. Possible values are CacheDisk and ResourceDisk. Defaults to CacheDisk. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Placement *string `json:"placement,omitempty" tf:"placement,omitempty"`
}

type GalleryApplicationObservation struct {
}

type GalleryApplicationParameters struct {

	// Specifies the URI to an Azure Blob that will replace the default configuration for the package if provided.
	// +kubebuilder:validation:Optional
	ConfigurationBlobURI *string `json:"configurationBlobUri,omitempty" tf:"configuration_blob_uri,omitempty"`

	// Specifies the order in which the packages have to be installed. Possible values are between 0 and 2,147,483,647.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Specifies a passthrough value for more generic context. This field can be any valid string value.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// Specifies the Gallery Application Version resource ID.
	// +kubebuilder:validation:Required
	VersionID *string `json:"versionId" tf:"version_id,omitempty"`
}

type LinuxVirtualMachineIdentityObservation struct {

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type LinuxVirtualMachineIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Linux Virtual Machine.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Linux Virtual Machine. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LinuxVirtualMachineObservation struct {

	// The ID of the Linux Virtual Machine.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []LinuxVirtualMachineIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Primary Private IP Address assigned to this Virtual Machine.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// A list of Private IP Addresses assigned to this Virtual Machine.
	PrivateIPAddresses []*string `json:"privateIpAddresses,omitempty" tf:"private_ip_addresses,omitempty"`

	// The Primary Public IP Address assigned to this Virtual Machine.
	PublicIPAddress *string `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	// A list of the Public IP Addresses assigned to this Virtual Machine.
	PublicIPAddresses []*string `json:"publicIpAddresses,omitempty" tf:"public_ip_addresses,omitempty"`

	// A 128-bit identifier which uniquely identifies this Virtual Machine.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type LinuxVirtualMachineOsDiskObservation struct {
}

type LinuxVirtualMachineOsDiskParameters struct {

	// The Type of Caching which should be used for the Internal OS Disk. Possible values are None, ReadOnly and ReadWrite.
	// +kubebuilder:validation:Required
	Caching *string `json:"caching" tf:"caching,omitempty"`

	// A diff_disk_settings block as defined above. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DiffDiskSettings []DiffDiskSettingsParameters `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings,omitempty"`

	// The ID of the Disk Encryption Set which should be used to Encrypt this OS Disk. Conflicts with secure_vm_disk_encryption_set_id.
	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// The Size of the Internal OS Disk in GB, if you wish to vary from the size used in the image this Virtual Machine is sourced from.
	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// The name which should be used for the Internal OS Disk. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Disk Encryption Set which should be used to Encrypt this OS Disk when the Virtual Machine is a Confidential VM. Conflicts with disk_encryption_set_id. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecureVMDiskEncryptionSetID *string `json:"secureVmDiskEncryptionSetId,omitempty" tf:"secure_vm_disk_encryption_set_id,omitempty"`

	// Encryption Type when the Virtual Machine is a Confidential VM. Possible values are VMGuestStateOnly and DiskWithVMGuestState. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecurityEncryptionType *string `json:"securityEncryptionType,omitempty" tf:"security_encryption_type,omitempty"`

	// The Type of Storage Account which should back this the Internal OS Disk. Possible values are Standard_LRS, StandardSSD_LRS, Premium_LRS, StandardSSD_ZRS and Premium_ZRS. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// Should Write Accelerator be Enabled for this OS Disk? Defaults to false.
	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type LinuxVirtualMachineParameters struct {

	// A additional_capabilities block as defined below.
	// +kubebuilder:validation:Optional
	AdditionalCapabilities []AdditionalCapabilitiesParameters `json:"additionalCapabilities,omitempty" tf:"additional_capabilities,omitempty"`

	// The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// One or more admin_ssh_key blocks as defined below.
	// +kubebuilder:validation:Optional
	AdminSSHKey []AdminSSHKeyParameters `json:"adminSshKey,omitempty" tf:"admin_ssh_key,omitempty"`

	// The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// Should Extension Operations be allowed on this Virtual Machine?
	// +kubebuilder:validation:Optional
	AllowExtensionOperations *bool `json:"allowExtensionOperations,omitempty" tf:"allow_extension_operations,omitempty"`

	// Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AvailabilitySetID *string `json:"availabilitySetId,omitempty" tf:"availability_set_id,omitempty"`

	// A boot_diagnostics block as defined below.
	// +kubebuilder:validation:Optional
	BootDiagnostics []BootDiagnosticsParameters `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`

	// Specifies the ID of the Capacity Reservation Group which the Virtual Machine should be allocated to.
	// +kubebuilder:validation:Optional
	CapacityReservationGroupID *string `json:"capacityReservationGroupId,omitempty" tf:"capacity_reservation_group_id,omitempty"`

	// Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the name field. If the value of the name field is not a valid computer_name, then you must specify computer_name. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ComputerName *string `json:"computerName,omitempty" tf:"computer_name,omitempty"`

	// The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CustomDataSecretRef *v1.SecretKeySelector `json:"customDataSecretRef,omitempty" tf:"-"`

	// The ID of a Dedicated Host Group that this Linux Virtual Machine should be run within. Conflicts with dedicated_host_id.
	// +kubebuilder:validation:Optional
	DedicatedHostGroupID *string `json:"dedicatedHostGroupId,omitempty" tf:"dedicated_host_group_id,omitempty"`

	// The ID of a Dedicated Host where this machine should be run on. Conflicts with dedicated_host_group_id.
	// +kubebuilder:validation:Optional
	DedicatedHostID *string `json:"dedicatedHostId,omitempty" tf:"dedicated_host_id,omitempty"`

	// Should Password Authentication be disabled on this Virtual Machine? Defaults to true. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Linux Virtual Machine should exist. Changing this forces a new Linux Virtual Machine to be created.
	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
	// +kubebuilder:validation:Optional
	EncryptionAtHostEnabled *bool `json:"encryptionAtHostEnabled,omitempty" tf:"encryption_at_host_enabled,omitempty"`

	// Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. Possible values are Deallocate and Delete. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// Specifies the duration allocated for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. Defaults to 90 minutes (PT1H30M).
	// +kubebuilder:validation:Optional
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty" tf:"extensions_time_budget,omitempty"`

	// A gallery_application block as defined below.
	// +kubebuilder:validation:Optional
	GalleryApplication []GalleryApplicationParameters `json:"galleryApplication,omitempty" tf:"gallery_application,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []LinuxVirtualMachineIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the BYOL Type for this Virtual Machine. Possible values are RHEL_BYOS and SLES_BYOS.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the eviction_policy. Defaults to -1, which means that the Virtual Machine should not be evicted for price reasons.
	// +kubebuilder:validation:Optional
	MaxBidPrice *float64 `json:"maxBidPrice,omitempty" tf:"max_bid_price,omitempty"`

	// . A list of Network Interface IDs which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	// References to NetworkInterface in network to populate networkInterfaceIds.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIdsRefs []v1.Reference `json:"networkInterfaceIdsRefs,omitempty" tf:"-"`

	// Selector for a list of NetworkInterface in network to populate networkInterfaceIds.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIdsSelector *v1.Selector `json:"networkInterfaceIdsSelector,omitempty" tf:"-"`

	// A os_disk block as defined below.
	// +kubebuilder:validation:Required
	OsDisk []LinuxVirtualMachineOsDiskParameters `json:"osDisk" tf:"os_disk,omitempty"`

	// Specifies the mode of VM Guest Patching for the Virtual Machine. Possible values are AutomaticByPlatform or ImageDefault. Defaults to ImageDefault.
	// +kubebuilder:validation:Optional
	PatchAssessmentMode *string `json:"patchAssessmentMode,omitempty" tf:"patch_assessment_mode,omitempty"`

	// Specifies the mode of in-guest patching to this Linux Virtual Machine. Possible values are AutomaticByPlatform and ImageDefault. Defaults to ImageDefault. For more information on patch modes please see the product documentation.
	// +kubebuilder:validation:Optional
	PatchMode *string `json:"patchMode,omitempty" tf:"patch_mode,omitempty"`

	// A plan block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Plan []PlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// Specifies the Platform Fault Domain in which this Linux Virtual Machine should be created. Defaults to -1, which means this will be automatically assigned to a fault domain that best maintains balance across the available fault domains. Changing this forces a new Linux Virtual Machine to be created.
	// +kubebuilder:validation:Optional
	PlatformFaultDomain *float64 `json:"platformFaultDomain,omitempty" tf:"platform_fault_domain,omitempty"`

	// Specifies the priority of this Virtual Machine. Possible values are Regular and Spot. Defaults to Regular. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to true. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ProvisionVMAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`

	// The ID of the Proximity Placement Group which the Virtual Machine should be assigned to.
	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// One or more secret blocks as defined below.
	// +kubebuilder:validation:Optional
	Secret []SecretParameters `json:"secret,omitempty" tf:"secret,omitempty"`

	// Specifies whether secure boot should be enabled on the virtual machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty" tf:"secure_boot_enabled,omitempty"`

	// The SKU which should be used for this Virtual Machine, such as Standard_F2.
	// +kubebuilder:validation:Required
	Size *string `json:"size" tf:"size,omitempty"`

	// The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created. Possible Image ID types include Image IDs, Shared Image IDs, Shared Image Version IDs, Community Gallery Image IDs, Community Gallery Image Version IDs, Shared Gallery Image IDs and Shared Gallery Image Version IDs.
	// +kubebuilder:validation:Optional
	SourceImageID *string `json:"sourceImageId,omitempty" tf:"source_image_id,omitempty"`

	// A source_image_reference block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceImageReference []SourceImageReferenceParameters `json:"sourceImageReference,omitempty" tf:"source_image_reference,omitempty"`

	// A mapping of tags which should be assigned to this Virtual Machine.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A termination_notification block as defined below.
	// +kubebuilder:validation:Optional
	TerminationNotification []TerminationNotificationParameters `json:"terminationNotification,omitempty" tf:"termination_notification,omitempty"`

	// The Base64-Encoded User Data which should be used for this Virtual Machine.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	VirtualMachineScaleSetID *string `json:"virtualMachineScaleSetId,omitempty" tf:"virtual_machine_scale_set_id,omitempty"`

	// Specifies whether vTPM should be enabled on the virtual machine. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	VtpmEnabled *bool `json:"vtpmEnabled,omitempty" tf:"vtpm_enabled,omitempty"`

	// Specifies the Availability Zones in which this Linux Virtual Machine should be located. Changing this forces a new Linux Virtual Machine to be created.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PlanObservation struct {
}

type PlanParameters struct {

	// Specifies the Name of the Marketplace Image this Virtual Machine should be created from. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the Product of the Marketplace Image this Virtual Machine should be created from. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Product *string `json:"product" tf:"product,omitempty"`

	// Specifies the Publisher of the Marketplace Image this Virtual Machine should be created from. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`
}

type SecretObservation struct {
}

type SecretParameters struct {

	// One or more certificate blocks as defined above.
	// +kubebuilder:validation:Required
	Certificate []CertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// The ID of the Key Vault from which all Secrets should be sourced.
	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`
}

type SourceImageReferenceObservation struct {
}

type SourceImageReferenceParameters struct {

	// Specifies the offer of the image used to create the virtual machines.
	// +kubebuilder:validation:Required
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// Specifies the publisher of the image used to create the virtual machines.
	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// Specifies the SKU of the image used to create the virtual machines.
	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// Specifies the version of the image used to create the virtual machines.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type TerminationNotificationObservation struct {
}

type TerminationNotificationParameters struct {

	// Should the termination notification be enabled on this Virtual Machine? Defaults to false.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Length of time (in minutes, between 5 and 15) a notification to be sent to the VM on the instance metadata server till the VM gets deleted. The time duration should be specified in ISO 8601 format.
	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

// LinuxVirtualMachineSpec defines the desired state of LinuxVirtualMachine
type LinuxVirtualMachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinuxVirtualMachineParameters `json:"forProvider"`
}

// LinuxVirtualMachineStatus defines the observed state of LinuxVirtualMachine.
type LinuxVirtualMachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinuxVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinuxVirtualMachine is the Schema for the LinuxVirtualMachines API. Manages a Linux Virtual Machine.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinuxVirtualMachineSpec   `json:"spec"`
	Status            LinuxVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinuxVirtualMachineList contains a list of LinuxVirtualMachines
type LinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinuxVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	LinuxVirtualMachine_Kind             = "LinuxVirtualMachine"
	LinuxVirtualMachine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinuxVirtualMachine_Kind}.String()
	LinuxVirtualMachine_KindAPIVersion   = LinuxVirtualMachine_Kind + "." + CRDGroupVersion.String()
	LinuxVirtualMachine_GroupVersionKind = CRDGroupVersion.WithKind(LinuxVirtualMachine_Kind)
)

func init() {
	SchemeBuilder.Register(&LinuxVirtualMachine{}, &LinuxVirtualMachineList{})
}
