//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Daemonjob) DeepCopyInto(out *Daemonjob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Daemonjob.
func (in *Daemonjob) DeepCopy() *Daemonjob {
	if in == nil {
		return nil
	}
	out := new(Daemonjob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Daemonjob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonjobList) DeepCopyInto(out *DaemonjobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Daemonjob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonjobList.
func (in *DaemonjobList) DeepCopy() *DaemonjobList {
	if in == nil {
		return nil
	}
	out := new(DaemonjobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DaemonjobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonjobSpec) DeepCopyInto(out *DaemonjobSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonjobSpec.
func (in *DaemonjobSpec) DeepCopy() *DaemonjobSpec {
	if in == nil {
		return nil
	}
	out := new(DaemonjobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonjobStatus) DeepCopyInto(out *DaemonjobStatus) {
	*out = *in
	if in.LastRun != nil {
		in, out := &in.LastRun, &out.LastRun
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonjobStatus.
func (in *DaemonjobStatus) DeepCopy() *DaemonjobStatus {
	if in == nil {
		return nil
	}
	out := new(DaemonjobStatus)
	in.DeepCopyInto(out)
	return out
}
