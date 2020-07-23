// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionSelector) DeepCopyInto(out *CustomResourceDefinitionSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionSelector.
func (in *CustomResourceDefinitionSelector) DeepCopy() *CustomResourceDefinitionSelector {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckType) DeepCopyInto(out *DuckType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckType.
func (in *DuckType) DeepCopy() *DuckType {
	if in == nil {
		return nil
	}
	out := new(DuckType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DuckType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckTypeList) DeepCopyInto(out *DuckTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DuckType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckTypeList.
func (in *DuckTypeList) DeepCopy() *DuckTypeList {
	if in == nil {
		return nil
	}
	out := new(DuckTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DuckTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckTypeNames) DeepCopyInto(out *DuckTypeNames) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckTypeNames.
func (in *DuckTypeNames) DeepCopy() *DuckTypeNames {
	if in == nil {
		return nil
	}
	out := new(DuckTypeNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckTypeSpec) DeepCopyInto(out *DuckTypeSpec) {
	*out = *in
	out.Names = in.Names
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]DuckVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]CustomResourceDefinitionSelector, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckTypeSpec.
func (in *DuckTypeSpec) DeepCopy() *DuckTypeSpec {
	if in == nil {
		return nil
	}
	out := new(DuckTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckTypeStatus) DeepCopyInto(out *DuckTypeStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.DuckList != nil {
		in, out := &in.DuckList, &out.DuckList
		*out = make([]FoundDuck, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckTypeStatus.
func (in *DuckTypeStatus) DeepCopy() *DuckTypeStatus {
	if in == nil {
		return nil
	}
	out := new(DuckTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckVersion) DeepCopyInto(out *DuckVersion) {
	*out = *in
	if in.Refs != nil {
		in, out := &in.Refs, &out.Refs
		*out = make([]GroupVersionResourceKind, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalPrinterColumns != nil {
		in, out := &in.AdditionalPrinterColumns, &out.AdditionalPrinterColumns
		*out = make([]v1.CustomResourceColumnDefinition, len(*in))
		copy(*out, *in)
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(v1.CustomResourceValidation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckVersion.
func (in *DuckVersion) DeepCopy() *DuckVersion {
	if in == nil {
		return nil
	}
	out := new(DuckVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundDuck) DeepCopyInto(out *FoundDuck) {
	*out = *in
	out.Ref = in.Ref
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundDuck.
func (in *FoundDuck) DeepCopy() *FoundDuck {
	if in == nil {
		return nil
	}
	out := new(FoundDuck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionResourceKind) DeepCopyInto(out *GroupVersionResourceKind) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionResourceKind.
func (in *GroupVersionResourceKind) DeepCopy() *GroupVersionResourceKind {
	if in == nil {
		return nil
	}
	out := new(GroupVersionResourceKind)
	in.DeepCopyInto(out)
	return out
}