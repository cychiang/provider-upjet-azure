//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureDescriptionInitParameters) DeepCopyInto(out *CaptureDescriptionInitParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(DestinationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Encoding != nil {
		in, out := &in.Encoding, &out.Encoding
		*out = new(string)
		**out = **in
	}
	if in.IntervalInSeconds != nil {
		in, out := &in.IntervalInSeconds, &out.IntervalInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.SizeLimitInBytes != nil {
		in, out := &in.SizeLimitInBytes, &out.SizeLimitInBytes
		*out = new(float64)
		**out = **in
	}
	if in.SkipEmptyArchives != nil {
		in, out := &in.SkipEmptyArchives, &out.SkipEmptyArchives
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureDescriptionInitParameters.
func (in *CaptureDescriptionInitParameters) DeepCopy() *CaptureDescriptionInitParameters {
	if in == nil {
		return nil
	}
	out := new(CaptureDescriptionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureDescriptionObservation) DeepCopyInto(out *CaptureDescriptionObservation) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(DestinationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Encoding != nil {
		in, out := &in.Encoding, &out.Encoding
		*out = new(string)
		**out = **in
	}
	if in.IntervalInSeconds != nil {
		in, out := &in.IntervalInSeconds, &out.IntervalInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.SizeLimitInBytes != nil {
		in, out := &in.SizeLimitInBytes, &out.SizeLimitInBytes
		*out = new(float64)
		**out = **in
	}
	if in.SkipEmptyArchives != nil {
		in, out := &in.SkipEmptyArchives, &out.SkipEmptyArchives
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureDescriptionObservation.
func (in *CaptureDescriptionObservation) DeepCopy() *CaptureDescriptionObservation {
	if in == nil {
		return nil
	}
	out := new(CaptureDescriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureDescriptionParameters) DeepCopyInto(out *CaptureDescriptionParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(DestinationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Encoding != nil {
		in, out := &in.Encoding, &out.Encoding
		*out = new(string)
		**out = **in
	}
	if in.IntervalInSeconds != nil {
		in, out := &in.IntervalInSeconds, &out.IntervalInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.SizeLimitInBytes != nil {
		in, out := &in.SizeLimitInBytes, &out.SizeLimitInBytes
		*out = new(float64)
		**out = **in
	}
	if in.SkipEmptyArchives != nil {
		in, out := &in.SkipEmptyArchives, &out.SkipEmptyArchives
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureDescriptionParameters.
func (in *CaptureDescriptionParameters) DeepCopy() *CaptureDescriptionParameters {
	if in == nil {
		return nil
	}
	out := new(CaptureDescriptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationInitParameters) DeepCopyInto(out *DestinationInitParameters) {
	*out = *in
	if in.ArchiveNameFormat != nil {
		in, out := &in.ArchiveNameFormat, &out.ArchiveNameFormat
		*out = new(string)
		**out = **in
	}
	if in.BlobContainerName != nil {
		in, out := &in.BlobContainerName, &out.BlobContainerName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationInitParameters.
func (in *DestinationInitParameters) DeepCopy() *DestinationInitParameters {
	if in == nil {
		return nil
	}
	out := new(DestinationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationObservation) DeepCopyInto(out *DestinationObservation) {
	*out = *in
	if in.ArchiveNameFormat != nil {
		in, out := &in.ArchiveNameFormat, &out.ArchiveNameFormat
		*out = new(string)
		**out = **in
	}
	if in.BlobContainerName != nil {
		in, out := &in.BlobContainerName, &out.BlobContainerName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationObservation.
func (in *DestinationObservation) DeepCopy() *DestinationObservation {
	if in == nil {
		return nil
	}
	out := new(DestinationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationParameters) DeepCopyInto(out *DestinationParameters) {
	*out = *in
	if in.ArchiveNameFormat != nil {
		in, out := &in.ArchiveNameFormat, &out.ArchiveNameFormat
		*out = new(string)
		**out = **in
	}
	if in.BlobContainerName != nil {
		in, out := &in.BlobContainerName, &out.BlobContainerName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationParameters.
func (in *DestinationParameters) DeepCopy() *DestinationParameters {
	if in == nil {
		return nil
	}
	out := new(DestinationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHub) DeepCopyInto(out *EventHub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHub.
func (in *EventHub) DeepCopy() *EventHub {
	if in == nil {
		return nil
	}
	out := new(EventHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventHub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubInitParameters) DeepCopyInto(out *EventHubInitParameters) {
	*out = *in
	if in.CaptureDescription != nil {
		in, out := &in.CaptureDescription, &out.CaptureDescription
		*out = new(CaptureDescriptionInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.MessageRetention != nil {
		in, out := &in.MessageRetention, &out.MessageRetention
		*out = new(float64)
		**out = **in
	}
	if in.PartitionCount != nil {
		in, out := &in.PartitionCount, &out.PartitionCount
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubInitParameters.
func (in *EventHubInitParameters) DeepCopy() *EventHubInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventHubInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubList) DeepCopyInto(out *EventHubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventHub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubList.
func (in *EventHubList) DeepCopy() *EventHubList {
	if in == nil {
		return nil
	}
	out := new(EventHubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventHubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubNamespace) DeepCopyInto(out *EventHubNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubNamespace.
func (in *EventHubNamespace) DeepCopy() *EventHubNamespace {
	if in == nil {
		return nil
	}
	out := new(EventHubNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventHubNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubNamespaceInitParameters) DeepCopyInto(out *EventHubNamespaceInitParameters) {
	*out = *in
	if in.AutoInflateEnabled != nil {
		in, out := &in.AutoInflateEnabled, &out.AutoInflateEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(float64)
		**out = **in
	}
	if in.DedicatedClusterID != nil {
		in, out := &in.DedicatedClusterID, &out.DedicatedClusterID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalAuthenticationEnabled != nil {
		in, out := &in.LocalAuthenticationEnabled, &out.LocalAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaximumThroughputUnits != nil {
		in, out := &in.MaximumThroughputUnits, &out.MaximumThroughputUnits
		*out = new(float64)
		**out = **in
	}
	if in.MinimumTLSVersion != nil {
		in, out := &in.MinimumTLSVersion, &out.MinimumTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.NetworkRulesets != nil {
		in, out := &in.NetworkRulesets, &out.NetworkRulesets
		*out = new(NetworkRulesetsInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ZoneRedundant != nil {
		in, out := &in.ZoneRedundant, &out.ZoneRedundant
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubNamespaceInitParameters.
func (in *EventHubNamespaceInitParameters) DeepCopy() *EventHubNamespaceInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventHubNamespaceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubNamespaceList) DeepCopyInto(out *EventHubNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventHubNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubNamespaceList.
func (in *EventHubNamespaceList) DeepCopy() *EventHubNamespaceList {
	if in == nil {
		return nil
	}
	out := new(EventHubNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventHubNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubNamespaceObservation) DeepCopyInto(out *EventHubNamespaceObservation) {
	*out = *in
	if in.AutoInflateEnabled != nil {
		in, out := &in.AutoInflateEnabled, &out.AutoInflateEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(float64)
		**out = **in
	}
	if in.DedicatedClusterID != nil {
		in, out := &in.DedicatedClusterID, &out.DedicatedClusterID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalAuthenticationEnabled != nil {
		in, out := &in.LocalAuthenticationEnabled, &out.LocalAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaximumThroughputUnits != nil {
		in, out := &in.MaximumThroughputUnits, &out.MaximumThroughputUnits
		*out = new(float64)
		**out = **in
	}
	if in.MinimumTLSVersion != nil {
		in, out := &in.MinimumTLSVersion, &out.MinimumTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.NetworkRulesets != nil {
		in, out := &in.NetworkRulesets, &out.NetworkRulesets
		*out = new(NetworkRulesetsObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ZoneRedundant != nil {
		in, out := &in.ZoneRedundant, &out.ZoneRedundant
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubNamespaceObservation.
func (in *EventHubNamespaceObservation) DeepCopy() *EventHubNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(EventHubNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubNamespaceParameters) DeepCopyInto(out *EventHubNamespaceParameters) {
	*out = *in
	if in.AutoInflateEnabled != nil {
		in, out := &in.AutoInflateEnabled, &out.AutoInflateEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(float64)
		**out = **in
	}
	if in.DedicatedClusterID != nil {
		in, out := &in.DedicatedClusterID, &out.DedicatedClusterID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(IdentityParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalAuthenticationEnabled != nil {
		in, out := &in.LocalAuthenticationEnabled, &out.LocalAuthenticationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaximumThroughputUnits != nil {
		in, out := &in.MaximumThroughputUnits, &out.MaximumThroughputUnits
		*out = new(float64)
		**out = **in
	}
	if in.MinimumTLSVersion != nil {
		in, out := &in.MinimumTLSVersion, &out.MinimumTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.NetworkRulesets != nil {
		in, out := &in.NetworkRulesets, &out.NetworkRulesets
		*out = new(NetworkRulesetsParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ZoneRedundant != nil {
		in, out := &in.ZoneRedundant, &out.ZoneRedundant
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubNamespaceParameters.
func (in *EventHubNamespaceParameters) DeepCopy() *EventHubNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(EventHubNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubNamespaceSpec) DeepCopyInto(out *EventHubNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubNamespaceSpec.
func (in *EventHubNamespaceSpec) DeepCopy() *EventHubNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(EventHubNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubNamespaceStatus) DeepCopyInto(out *EventHubNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubNamespaceStatus.
func (in *EventHubNamespaceStatus) DeepCopy() *EventHubNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(EventHubNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubObservation) DeepCopyInto(out *EventHubObservation) {
	*out = *in
	if in.CaptureDescription != nil {
		in, out := &in.CaptureDescription, &out.CaptureDescription
		*out = new(CaptureDescriptionObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MessageRetention != nil {
		in, out := &in.MessageRetention, &out.MessageRetention
		*out = new(float64)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.PartitionCount != nil {
		in, out := &in.PartitionCount, &out.PartitionCount
		*out = new(float64)
		**out = **in
	}
	if in.PartitionIds != nil {
		in, out := &in.PartitionIds, &out.PartitionIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubObservation.
func (in *EventHubObservation) DeepCopy() *EventHubObservation {
	if in == nil {
		return nil
	}
	out := new(EventHubObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubParameters) DeepCopyInto(out *EventHubParameters) {
	*out = *in
	if in.CaptureDescription != nil {
		in, out := &in.CaptureDescription, &out.CaptureDescription
		*out = new(CaptureDescriptionParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.MessageRetention != nil {
		in, out := &in.MessageRetention, &out.MessageRetention
		*out = new(float64)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.NamespaceNameRef != nil {
		in, out := &in.NamespaceNameRef, &out.NamespaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceNameSelector != nil {
		in, out := &in.NamespaceNameSelector, &out.NamespaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PartitionCount != nil {
		in, out := &in.PartitionCount, &out.PartitionCount
		*out = new(float64)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubParameters.
func (in *EventHubParameters) DeepCopy() *EventHubParameters {
	if in == nil {
		return nil
	}
	out := new(EventHubParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubSpec) DeepCopyInto(out *EventHubSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubSpec.
func (in *EventHubSpec) DeepCopy() *EventHubSpec {
	if in == nil {
		return nil
	}
	out := new(EventHubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventHubStatus) DeepCopyInto(out *EventHubStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventHubStatus.
func (in *EventHubStatus) DeepCopy() *EventHubStatus {
	if in == nil {
		return nil
	}
	out := new(EventHubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRuleInitParameters) DeepCopyInto(out *IPRuleInitParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.IPMask != nil {
		in, out := &in.IPMask, &out.IPMask
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRuleInitParameters.
func (in *IPRuleInitParameters) DeepCopy() *IPRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(IPRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRuleObservation) DeepCopyInto(out *IPRuleObservation) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.IPMask != nil {
		in, out := &in.IPMask, &out.IPMask
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRuleObservation.
func (in *IPRuleObservation) DeepCopy() *IPRuleObservation {
	if in == nil {
		return nil
	}
	out := new(IPRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRuleParameters) DeepCopyInto(out *IPRuleParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.IPMask != nil {
		in, out := &in.IPMask, &out.IPMask
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRuleParameters.
func (in *IPRuleParameters) DeepCopy() *IPRuleParameters {
	if in == nil {
		return nil
	}
	out := new(IPRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityInitParameters) DeepCopyInto(out *IdentityInitParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityInitParameters.
func (in *IdentityInitParameters) DeepCopy() *IdentityInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.IdentityIds != nil {
		in, out := &in.IdentityIds, &out.IdentityIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRulesetsInitParameters) DeepCopyInto(out *NetworkRulesetsInitParameters) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRule != nil {
		in, out := &in.IPRule, &out.IPRule
		*out = make([]IPRuleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TrustedServiceAccessEnabled != nil {
		in, out := &in.TrustedServiceAccessEnabled, &out.TrustedServiceAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.VirtualNetworkRule != nil {
		in, out := &in.VirtualNetworkRule, &out.VirtualNetworkRule
		*out = make([]VirtualNetworkRuleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRulesetsInitParameters.
func (in *NetworkRulesetsInitParameters) DeepCopy() *NetworkRulesetsInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkRulesetsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRulesetsObservation) DeepCopyInto(out *NetworkRulesetsObservation) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRule != nil {
		in, out := &in.IPRule, &out.IPRule
		*out = make([]IPRuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TrustedServiceAccessEnabled != nil {
		in, out := &in.TrustedServiceAccessEnabled, &out.TrustedServiceAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.VirtualNetworkRule != nil {
		in, out := &in.VirtualNetworkRule, &out.VirtualNetworkRule
		*out = make([]VirtualNetworkRuleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRulesetsObservation.
func (in *NetworkRulesetsObservation) DeepCopy() *NetworkRulesetsObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkRulesetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRulesetsParameters) DeepCopyInto(out *NetworkRulesetsParameters) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IPRule != nil {
		in, out := &in.IPRule, &out.IPRule
		*out = make([]IPRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TrustedServiceAccessEnabled != nil {
		in, out := &in.TrustedServiceAccessEnabled, &out.TrustedServiceAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.VirtualNetworkRule != nil {
		in, out := &in.VirtualNetworkRule, &out.VirtualNetworkRule
		*out = make([]VirtualNetworkRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRulesetsParameters.
func (in *NetworkRulesetsParameters) DeepCopy() *NetworkRulesetsParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkRulesetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleInitParameters) DeepCopyInto(out *VirtualNetworkRuleInitParameters) {
	*out = *in
	if in.IgnoreMissingVirtualNetworkServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVirtualNetworkServiceEndpoint, &out.IgnoreMissingVirtualNetworkServiceEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleInitParameters.
func (in *VirtualNetworkRuleInitParameters) DeepCopy() *VirtualNetworkRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleObservation) DeepCopyInto(out *VirtualNetworkRuleObservation) {
	*out = *in
	if in.IgnoreMissingVirtualNetworkServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVirtualNetworkServiceEndpoint, &out.IgnoreMissingVirtualNetworkServiceEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleObservation.
func (in *VirtualNetworkRuleObservation) DeepCopy() *VirtualNetworkRuleObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleParameters) DeepCopyInto(out *VirtualNetworkRuleParameters) {
	*out = *in
	if in.IgnoreMissingVirtualNetworkServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVirtualNetworkServiceEndpoint, &out.IgnoreMissingVirtualNetworkServiceEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleParameters.
func (in *VirtualNetworkRuleParameters) DeepCopy() *VirtualNetworkRuleParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleParameters)
	in.DeepCopyInto(out)
	return out
}
