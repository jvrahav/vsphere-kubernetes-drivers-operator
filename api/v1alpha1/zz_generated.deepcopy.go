//go:build !ignore_autogenerated
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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPIStatus) DeepCopyInto(out *CPIStatus) {
	*out = *in
	if in.NodeStatus != nil {
		in, out := &in.NodeStatus, &out.NodeStatus
		*out = make(map[string]NodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPIStatus.
func (in *CPIStatus) DeepCopy() *CPIStatus {
	if in == nil {
		return nil
	}
	out := new(CPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIStatus) DeepCopyInto(out *CSIStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIStatus.
func (in *CSIStatus) DeepCopy() *CSIStatus {
	if in == nil {
		return nil
	}
	out := new(CSIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderConfig) DeepCopyInto(out *CloudProviderConfig) {
	*out = *in
	if in.VsphereCloudConfigs != nil {
		in, out := &in.VsphereCloudConfigs, &out.VsphereCloudConfigs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Topology = in.Topology
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderConfig.
func (in *CloudProviderConfig) DeepCopy() *CloudProviderConfig {
	if in == nil {
		return nil
	}
	out := new(CloudProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileVolume) DeepCopyInto(out *FileVolume) {
	*out = *in
	if in.VSanDataStoreUrl != nil {
		in, out := &in.VSanDataStoreUrl, &out.VSanDataStoreUrl
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NetPermissions != nil {
		in, out := &in.NetPermissions, &out.NetPermissions
		*out = make([]NetPermission, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileVolume.
func (in *FileVolume) DeepCopy() *FileVolume {
	if in == nil {
		return nil
	}
	out := new(FileVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetPermission) DeepCopyInto(out *NetPermission) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetPermission.
func (in *NetPermission) DeepCopy() *NetPermission {
	if in == nil {
		return nil
	}
	out := new(NetPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProviderConfig) DeepCopyInto(out *StorageProviderConfig) {
	*out = *in
	in.FileVolumes.DeepCopyInto(&out.FileVolumes)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProviderConfig.
func (in *StorageProviderConfig) DeepCopy() *StorageProviderConfig {
	if in == nil {
		return nil
	}
	out := new(StorageProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopologyInfo) DeepCopyInto(out *TopologyInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopologyInfo.
func (in *TopologyInfo) DeepCopy() *TopologyInfo {
	if in == nil {
		return nil
	}
	out := new(TopologyInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VDOConfig) DeepCopyInto(out *VDOConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VDOConfig.
func (in *VDOConfig) DeepCopy() *VDOConfig {
	if in == nil {
		return nil
	}
	out := new(VDOConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VDOConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VDOConfigList) DeepCopyInto(out *VDOConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VDOConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VDOConfigList.
func (in *VDOConfigList) DeepCopy() *VDOConfigList {
	if in == nil {
		return nil
	}
	out := new(VDOConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VDOConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VDOConfigSpec) DeepCopyInto(out *VDOConfigSpec) {
	*out = *in
	in.CloudProvider.DeepCopyInto(&out.CloudProvider)
	in.StorageProvider.DeepCopyInto(&out.StorageProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VDOConfigSpec.
func (in *VDOConfigSpec) DeepCopy() *VDOConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VDOConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VDOConfigStatus) DeepCopyInto(out *VDOConfigStatus) {
	*out = *in
	in.CPIStatus.DeepCopyInto(&out.CPIStatus)
	out.CSIStatus = in.CSIStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VDOConfigStatus.
func (in *VDOConfigStatus) DeepCopy() *VDOConfigStatus {
	if in == nil {
		return nil
	}
	out := new(VDOConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereCloudConfig) DeepCopyInto(out *VsphereCloudConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereCloudConfig.
func (in *VsphereCloudConfig) DeepCopy() *VsphereCloudConfig {
	if in == nil {
		return nil
	}
	out := new(VsphereCloudConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereCloudConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereCloudConfigList) DeepCopyInto(out *VsphereCloudConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VsphereCloudConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereCloudConfigList.
func (in *VsphereCloudConfigList) DeepCopy() *VsphereCloudConfigList {
	if in == nil {
		return nil
	}
	out := new(VsphereCloudConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VsphereCloudConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereCloudConfigSpec) DeepCopyInto(out *VsphereCloudConfigSpec) {
	*out = *in
	if in.DataCenters != nil {
		in, out := &in.DataCenters, &out.DataCenters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereCloudConfigSpec.
func (in *VsphereCloudConfigSpec) DeepCopy() *VsphereCloudConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VsphereCloudConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereCloudConfigStatus) DeepCopyInto(out *VsphereCloudConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereCloudConfigStatus.
func (in *VsphereCloudConfigStatus) DeepCopy() *VsphereCloudConfigStatus {
	if in == nil {
		return nil
	}
	out := new(VsphereCloudConfigStatus)
	in.DeepCopyInto(out)
	return out
}
