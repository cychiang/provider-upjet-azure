/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyDefinitionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PolicyDefinitionIDRef,
		Selector:     mg.Spec.ForProvider.PolicyDefinitionIDSelector,
		To: reference.To{
			List:    &PolicyDefinitionList{},
			Managed: &PolicyDefinition{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyDefinitionID")
	}
	mg.Spec.ForProvider.PolicyDefinitionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyDefinitionIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResourceGroupIDRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupIDSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupID")
	}
	mg.Spec.ForProvider.ResourceGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupIDRef = rsp.ResolvedReference

	return nil
}
