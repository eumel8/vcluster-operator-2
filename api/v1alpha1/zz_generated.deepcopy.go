// +build !ignore_autogenerated

/*
Copyright 2021 Zachary Seguin

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualCluster) DeepCopyInto(out *VirtualCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualCluster.
func (in *VirtualCluster) DeepCopy() *VirtualCluster {
	if in == nil {
		return nil
	}
	out := new(VirtualCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterComponentSpec) DeepCopyInto(out *VirtualClusterComponentSpec) {
	*out = *in
	out.Image = in.Image
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterComponentSpec.
func (in *VirtualClusterComponentSpec) DeepCopy() *VirtualClusterComponentSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterImageSpec) DeepCopyInto(out *VirtualClusterImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterImageSpec.
func (in *VirtualClusterImageSpec) DeepCopy() *VirtualClusterImageSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterIngressSpec) DeepCopyInto(out *VirtualClusterIngressSpec) {
	*out = *in
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
	if in.TLSSecretName != nil {
		in, out := &in.TLSSecretName, &out.TLSSecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterIngressSpec.
func (in *VirtualClusterIngressSpec) DeepCopy() *VirtualClusterIngressSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterList) DeepCopyInto(out *VirtualClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterList.
func (in *VirtualClusterList) DeepCopy() *VirtualClusterList {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterNodeSyncSpec) DeepCopyInto(out *VirtualClusterNodeSyncSpec) {
	*out = *in
	if in.SelectorLabels != nil {
		in, out := &in.SelectorLabels, &out.SelectorLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterNodeSyncSpec.
func (in *VirtualClusterNodeSyncSpec) DeepCopy() *VirtualClusterNodeSyncSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterNodeSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterPodSchedulingSpec) DeepCopyInto(out *VirtualClusterPodSchedulingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterPodSchedulingSpec.
func (in *VirtualClusterPodSchedulingSpec) DeepCopy() *VirtualClusterPodSchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterPodSchedulingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterSpec) DeepCopyInto(out *VirtualClusterSpec) {
	*out = *in
	in.ControlPlane.DeepCopyInto(&out.ControlPlane)
	in.Syncer.DeepCopyInto(&out.Syncer)
	in.NodeSync.DeepCopyInto(&out.NodeSync)
	out.PodScheduling = in.PodScheduling
	out.Storage = in.Storage
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(VirtualClusterIngressSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterSpec.
func (in *VirtualClusterSpec) DeepCopy() *VirtualClusterSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterStatus) DeepCopyInto(out *VirtualClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterStatus.
func (in *VirtualClusterStatus) DeepCopy() *VirtualClusterStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterStorageSpec) DeepCopyInto(out *VirtualClusterStorageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterStorageSpec.
func (in *VirtualClusterStorageSpec) DeepCopy() *VirtualClusterStorageSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterStorageSpec)
	in.DeepCopyInto(out)
	return out
}
