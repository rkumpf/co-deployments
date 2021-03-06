// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisableInfo) DeepCopyInto(out *DisableInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisableInfo.
func (in *DisableInfo) DeepCopy() *DisableInfo {
	if in == nil {
		return nil
	}
	out := new(DisableInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPECSIDriver) DeepCopyInto(out *HPECSIDriver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPECSIDriver.
func (in *HPECSIDriver) DeepCopy() *HPECSIDriver {
	if in == nil {
		return nil
	}
	out := new(HPECSIDriver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HPECSIDriver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPECSIDriverList) DeepCopyInto(out *HPECSIDriverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HPECSIDriver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPECSIDriverList.
func (in *HPECSIDriverList) DeepCopy() *HPECSIDriverList {
	if in == nil {
		return nil
	}
	out := new(HPECSIDriverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HPECSIDriverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPECSIDriverSpec) DeepCopyInto(out *HPECSIDriverSpec) {
	*out = *in
	out.Disable = in.Disable
	out.Iscsi = in.Iscsi
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPECSIDriverSpec.
func (in *HPECSIDriverSpec) DeepCopy() *HPECSIDriverSpec {
	if in == nil {
		return nil
	}
	out := new(HPECSIDriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPECSIDriverStatus) DeepCopyInto(out *HPECSIDriverStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HelmAppCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeployedRelease != nil {
		in, out := &in.DeployedRelease, &out.DeployedRelease
		*out = new(HelmAppRelease)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPECSIDriverStatus.
func (in *HPECSIDriverStatus) DeepCopy() *HPECSIDriverStatus {
	if in == nil {
		return nil
	}
	out := new(HPECSIDriverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmAppCondition) DeepCopyInto(out *HelmAppCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmAppCondition.
func (in *HelmAppCondition) DeepCopy() *HelmAppCondition {
	if in == nil {
		return nil
	}
	out := new(HelmAppCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmAppRelease) DeepCopyInto(out *HelmAppRelease) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmAppRelease.
func (in *HelmAppRelease) DeepCopy() *HelmAppRelease {
	if in == nil {
		return nil
	}
	out := new(HelmAppRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IscsiInfo) DeepCopyInto(out *IscsiInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IscsiInfo.
func (in *IscsiInfo) DeepCopy() *IscsiInfo {
	if in == nil {
		return nil
	}
	out := new(IscsiInfo)
	in.DeepCopyInto(out)
	return out
}
