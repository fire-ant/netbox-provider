//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Device) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceList) DeepCopyInto(out *DeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Device, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceList.
func (in *DeviceList) DeepCopy() *DeviceList {
	if in == nil {
		return nil
	}
	out := new(DeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceObservation) DeepCopyInto(out *DeviceObservation) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(float64)
		**out = **in
	}
	if in.Comments != nil {
		in, out := &in.Comments, &out.Comments
		*out = new(string)
		**out = **in
	}
	if in.CustomFields != nil {
		in, out := &in.CustomFields, &out.CustomFields
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DeviceTypeID != nil {
		in, out := &in.DeviceTypeID, &out.DeviceTypeID
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LocationID != nil {
		in, out := &in.LocationID, &out.LocationID
		*out = new(float64)
		**out = **in
	}
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(float64)
		**out = **in
	}
	if in.PrimaryIPv4 != nil {
		in, out := &in.PrimaryIPv4, &out.PrimaryIPv4
		*out = new(float64)
		**out = **in
	}
	if in.PrimaryIPv6 != nil {
		in, out := &in.PrimaryIPv6, &out.PrimaryIPv6
		*out = new(float64)
		**out = **in
	}
	if in.RackFace != nil {
		in, out := &in.RackFace, &out.RackFace
		*out = new(string)
		**out = **in
	}
	if in.RackID != nil {
		in, out := &in.RackID, &out.RackID
		*out = new(float64)
		**out = **in
	}
	if in.RackPosition != nil {
		in, out := &in.RackPosition, &out.RackPosition
		*out = new(float64)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(float64)
		**out = **in
	}
	if in.Serial != nil {
		in, out := &in.Serial, &out.Serial
		*out = new(string)
		**out = **in
	}
	if in.SiteID != nil {
		in, out := &in.SiteID, &out.SiteID
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceObservation.
func (in *DeviceObservation) DeepCopy() *DeviceObservation {
	if in == nil {
		return nil
	}
	out := new(DeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceParameters) DeepCopyInto(out *DeviceParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(float64)
		**out = **in
	}
	if in.Comments != nil {
		in, out := &in.Comments, &out.Comments
		*out = new(string)
		**out = **in
	}
	if in.CustomFields != nil {
		in, out := &in.CustomFields, &out.CustomFields
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DeviceTypeID != nil {
		in, out := &in.DeviceTypeID, &out.DeviceTypeID
		*out = new(float64)
		**out = **in
	}
	if in.LocationID != nil {
		in, out := &in.LocationID, &out.LocationID
		*out = new(float64)
		**out = **in
	}
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(float64)
		**out = **in
	}
	if in.RackFace != nil {
		in, out := &in.RackFace, &out.RackFace
		*out = new(string)
		**out = **in
	}
	if in.RackID != nil {
		in, out := &in.RackID, &out.RackID
		*out = new(float64)
		**out = **in
	}
	if in.RackPosition != nil {
		in, out := &in.RackPosition, &out.RackPosition
		*out = new(float64)
		**out = **in
	}
	if in.RoleID != nil {
		in, out := &in.RoleID, &out.RoleID
		*out = new(float64)
		**out = **in
	}
	if in.Serial != nil {
		in, out := &in.Serial, &out.Serial
		*out = new(string)
		**out = **in
	}
	if in.SiteID != nil {
		in, out := &in.SiteID, &out.SiteID
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceParameters.
func (in *DeviceParameters) DeepCopy() *DeviceParameters {
	if in == nil {
		return nil
	}
	out := new(DeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpec) DeepCopyInto(out *DeviceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpec.
func (in *DeviceSpec) DeepCopy() *DeviceSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceStatus) DeepCopyInto(out *DeviceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceStatus.
func (in *DeviceStatus) DeepCopy() *DeviceStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceStatus)
	in.DeepCopyInto(out)
	return out
}
