/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha11 "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1"
	v1alpha1 "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1"
	v1alpha12 "github.com/fire-ant/provider-netbox/apis/virtualization/v1alpha1"
	common "github.com/fire-ant/provider-netbox/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Aggregate.
func (mg *Aggregate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RirID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RirIDRef,
		Selector:     mg.Spec.ForProvider.RirIDSelector,
		To: reference.To{
			List:    &RirList{},
			Managed: &Rir{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RirID")
	}
	mg.Spec.ForProvider.RirID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RirIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Asn.
func (mg *Asn) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RirID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RirIDRef,
		Selector:     mg.Spec.ForProvider.RirIDSelector,
		To: reference.To{
			List:    &RirList{},
			Managed: &Rir{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RirID")
	}
	mg.Spec.ForProvider.RirID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RirIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AvailableIPAdrress.
func (mg *AvailableIPAdrress) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.IPRangeID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IPRangeIDRef,
		Selector:     mg.Spec.ForProvider.IPRangeIDSelector,
		To: reference.To{
			List:    &IPRangeList{},
			Managed: &IPRange{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IPRangeID")
	}
	mg.Spec.ForProvider.IPRangeID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IPRangeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.InterfaceID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InterfaceIDRef,
		Selector:     mg.Spec.ForProvider.InterfaceIDSelector,
		To: reference.To{
			List:    &v1alpha11.DeviceInterfaceList{},
			Managed: &v1alpha11.DeviceInterface{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InterfaceID")
	}
	mg.Spec.ForProvider.InterfaceID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InterfaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.PrefixID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PrefixIDRef,
		Selector:     mg.Spec.ForProvider.PrefixIDSelector,
		To: reference.To{
			List:    &PrefixList{},
			Managed: &Prefix{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrefixID")
	}
	mg.Spec.ForProvider.PrefixID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrefixIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VrfID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VrfIDRef,
		Selector:     mg.Spec.ForProvider.VrfIDSelector,
		To: reference.To{
			List:    &VrfList{},
			Managed: &Vrf{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VrfID")
	}
	mg.Spec.ForProvider.VrfID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VrfIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AvailablePrefix.
func (mg *AvailablePrefix) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RoleID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RoleIDRef,
		Selector:     mg.Spec.ForProvider.RoleIDSelector,
		To: reference.To{
			List:    &IpamRoleList{},
			Managed: &IpamRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.SiteID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SiteIDRef,
		Selector:     mg.Spec.ForProvider.SiteIDSelector,
		To: reference.To{
			List:    &v1alpha11.SiteList{},
			Managed: &v1alpha11.Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VlanID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VlanIDRef,
		Selector:     mg.Spec.ForProvider.VlanIDSelector,
		To: reference.To{
			List:    &VlanList{},
			Managed: &Vlan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VlanID")
	}
	mg.Spec.ForProvider.VlanID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VlanIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VrfID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VrfIDRef,
		Selector:     mg.Spec.ForProvider.VrfIDSelector,
		To: reference.To{
			List:    &VrfList{},
			Managed: &Vrf{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VrfID")
	}
	mg.Spec.ForProvider.VrfID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VrfIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.ScopeID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ScopeIDRef,
		Selector:     mg.Spec.ForProvider.ScopeIDSelector,
		To: reference.To{
			List:    &v1alpha11.SiteList{},
			Managed: &v1alpha11.Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScopeID")
	}
	mg.Spec.ForProvider.ScopeID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScopeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IPAddress.
func (mg *IPAddress) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.InterfaceID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InterfaceIDRef,
		Selector:     mg.Spec.ForProvider.InterfaceIDSelector,
		To: reference.To{
			List:    &v1alpha11.DeviceInterfaceList{},
			Managed: &v1alpha11.DeviceInterface{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InterfaceID")
	}
	mg.Spec.ForProvider.InterfaceID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InterfaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VrfID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VrfIDRef,
		Selector:     mg.Spec.ForProvider.VrfIDSelector,
		To: reference.To{
			List:    &VrfList{},
			Managed: &Vrf{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VrfID")
	}
	mg.Spec.ForProvider.VrfID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VrfIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IPRange.
func (mg *IPRange) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RoleID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RoleIDRef,
		Selector:     mg.Spec.ForProvider.RoleIDSelector,
		To: reference.To{
			List:    &IpamRoleList{},
			Managed: &IpamRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VrfID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VrfIDRef,
		Selector:     mg.Spec.ForProvider.VrfIDSelector,
		To: reference.To{
			List:    &VrfList{},
			Managed: &Vrf{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VrfID")
	}
	mg.Spec.ForProvider.VrfID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VrfIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Prefix.
func (mg *Prefix) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RoleID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RoleIDRef,
		Selector:     mg.Spec.ForProvider.RoleIDSelector,
		To: reference.To{
			List:    &IpamRoleList{},
			Managed: &IpamRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.SiteID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SiteIDRef,
		Selector:     mg.Spec.ForProvider.SiteIDSelector,
		To: reference.To{
			List:    &v1alpha11.SiteList{},
			Managed: &v1alpha11.Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VlanID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VlanIDRef,
		Selector:     mg.Spec.ForProvider.VlanIDSelector,
		To: reference.To{
			List:    &VlanList{},
			Managed: &Vlan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VlanID")
	}
	mg.Spec.ForProvider.VlanID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VlanIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VrfID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VrfIDRef,
		Selector:     mg.Spec.ForProvider.VrfIDSelector,
		To: reference.To{
			List:    &VrfList{},
			Managed: &Vrf{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VrfID")
	}
	mg.Spec.ForProvider.VrfID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VrfIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.VirtualMachineID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualMachineIDRef,
		Selector:     mg.Spec.ForProvider.VirtualMachineIDSelector,
		To: reference.To{
			List:    &v1alpha12.MachineList{},
			Managed: &v1alpha12.Machine{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualMachineID")
	}
	mg.Spec.ForProvider.VirtualMachineID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualMachineIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Target.
func (mg *Target) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Vlan.
func (mg *Vlan) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.GroupID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GroupIDRef,
		Selector:     mg.Spec.ForProvider.GroupIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupID")
	}
	mg.Spec.ForProvider.GroupID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.RoleID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RoleIDRef,
		Selector:     mg.Spec.ForProvider.RoleIDSelector,
		To: reference.To{
			List:    &IpamRoleList{},
			Managed: &IpamRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.SiteID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SiteIDRef,
		Selector:     mg.Spec.ForProvider.SiteIDSelector,
		To: reference.To{
			List:    &v1alpha11.SiteList{},
			Managed: &v1alpha11.Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Vrf.
func (mg *Vrf) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1alpha1.TenantList{},
			Managed: &v1alpha1.Tenant{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}