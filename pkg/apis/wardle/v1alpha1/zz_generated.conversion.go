//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	wardle "code.alipay.com/ant-iac/karbour/pkg/apis/wardle"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Fischer)(nil), (*wardle.Fischer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Fischer_To_wardle_Fischer(a.(*Fischer), b.(*wardle.Fischer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.Fischer)(nil), (*Fischer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_Fischer_To_v1alpha1_Fischer(a.(*wardle.Fischer), b.(*Fischer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FischerList)(nil), (*wardle.FischerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FischerList_To_wardle_FischerList(a.(*FischerList), b.(*wardle.FischerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.FischerList)(nil), (*FischerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_FischerList_To_v1alpha1_FischerList(a.(*wardle.FischerList), b.(*FischerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Flunder)(nil), (*wardle.Flunder)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Flunder_To_wardle_Flunder(a.(*Flunder), b.(*wardle.Flunder), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.Flunder)(nil), (*Flunder)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_Flunder_To_v1alpha1_Flunder(a.(*wardle.Flunder), b.(*Flunder), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FlunderList)(nil), (*wardle.FlunderList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlunderList_To_wardle_FlunderList(a.(*FlunderList), b.(*wardle.FlunderList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.FlunderList)(nil), (*FlunderList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_FlunderList_To_v1alpha1_FlunderList(a.(*wardle.FlunderList), b.(*FlunderList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FlunderStatus)(nil), (*wardle.FlunderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(a.(*FlunderStatus), b.(*wardle.FlunderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.FlunderStatus)(nil), (*FlunderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(a.(*wardle.FlunderStatus), b.(*FlunderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*FlunderSpec)(nil), (*wardle.FlunderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec(a.(*FlunderSpec), b.(*wardle.FlunderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*wardle.FlunderSpec)(nil), (*FlunderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(a.(*wardle.FlunderSpec), b.(*FlunderSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Fischer_To_wardle_Fischer(in *Fischer, out *wardle.Fischer, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DisallowedFlunders = *(*[]string)(unsafe.Pointer(&in.DisallowedFlunders))
	return nil
}

// Convert_v1alpha1_Fischer_To_wardle_Fischer is an autogenerated conversion function.
func Convert_v1alpha1_Fischer_To_wardle_Fischer(in *Fischer, out *wardle.Fischer, s conversion.Scope) error {
	return autoConvert_v1alpha1_Fischer_To_wardle_Fischer(in, out, s)
}

func autoConvert_wardle_Fischer_To_v1alpha1_Fischer(in *wardle.Fischer, out *Fischer, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DisallowedFlunders = *(*[]string)(unsafe.Pointer(&in.DisallowedFlunders))
	return nil
}

// Convert_wardle_Fischer_To_v1alpha1_Fischer is an autogenerated conversion function.
func Convert_wardle_Fischer_To_v1alpha1_Fischer(in *wardle.Fischer, out *Fischer, s conversion.Scope) error {
	return autoConvert_wardle_Fischer_To_v1alpha1_Fischer(in, out, s)
}

func autoConvert_v1alpha1_FischerList_To_wardle_FischerList(in *FischerList, out *wardle.FischerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]wardle.Fischer)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_FischerList_To_wardle_FischerList is an autogenerated conversion function.
func Convert_v1alpha1_FischerList_To_wardle_FischerList(in *FischerList, out *wardle.FischerList, s conversion.Scope) error {
	return autoConvert_v1alpha1_FischerList_To_wardle_FischerList(in, out, s)
}

func autoConvert_wardle_FischerList_To_v1alpha1_FischerList(in *wardle.FischerList, out *FischerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Fischer)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_wardle_FischerList_To_v1alpha1_FischerList is an autogenerated conversion function.
func Convert_wardle_FischerList_To_v1alpha1_FischerList(in *wardle.FischerList, out *FischerList, s conversion.Scope) error {
	return autoConvert_wardle_FischerList_To_v1alpha1_FischerList(in, out, s)
}

func autoConvert_v1alpha1_Flunder_To_wardle_Flunder(in *Flunder, out *wardle.Flunder, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Flunder_To_wardle_Flunder is an autogenerated conversion function.
func Convert_v1alpha1_Flunder_To_wardle_Flunder(in *Flunder, out *wardle.Flunder, s conversion.Scope) error {
	return autoConvert_v1alpha1_Flunder_To_wardle_Flunder(in, out, s)
}

func autoConvert_wardle_Flunder_To_v1alpha1_Flunder(in *wardle.Flunder, out *Flunder, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_wardle_Flunder_To_v1alpha1_Flunder is an autogenerated conversion function.
func Convert_wardle_Flunder_To_v1alpha1_Flunder(in *wardle.Flunder, out *Flunder, s conversion.Scope) error {
	return autoConvert_wardle_Flunder_To_v1alpha1_Flunder(in, out, s)
}

func autoConvert_v1alpha1_FlunderList_To_wardle_FlunderList(in *FlunderList, out *wardle.FlunderList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]wardle.Flunder, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Flunder_To_wardle_Flunder(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_FlunderList_To_wardle_FlunderList is an autogenerated conversion function.
func Convert_v1alpha1_FlunderList_To_wardle_FlunderList(in *FlunderList, out *wardle.FlunderList, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlunderList_To_wardle_FlunderList(in, out, s)
}

func autoConvert_wardle_FlunderList_To_v1alpha1_FlunderList(in *wardle.FlunderList, out *FlunderList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Flunder, len(*in))
		for i := range *in {
			if err := Convert_wardle_Flunder_To_v1alpha1_Flunder(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_wardle_FlunderList_To_v1alpha1_FlunderList is an autogenerated conversion function.
func Convert_wardle_FlunderList_To_v1alpha1_FlunderList(in *wardle.FlunderList, out *FlunderList, s conversion.Scope) error {
	return autoConvert_wardle_FlunderList_To_v1alpha1_FlunderList(in, out, s)
}

func autoConvert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec(in *FlunderSpec, out *wardle.FlunderSpec, s conversion.Scope) error {
	// WARNING: in.Reference requires manual conversion: does not exist in peer-type
	// WARNING: in.ReferenceType requires manual conversion: inconvertible types (*k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1.ReferenceType vs k8s.io/sample-apiserver/pkg/apis/wardle.ReferenceType)
	return nil
}

func autoConvert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(in *wardle.FlunderSpec, out *FlunderSpec, s conversion.Scope) error {
	// WARNING: in.FlunderReference requires manual conversion: does not exist in peer-type
	// WARNING: in.FischerReference requires manual conversion: does not exist in peer-type
	// WARNING: in.ReferenceType requires manual conversion: inconvertible types (k8s.io/sample-apiserver/pkg/apis/wardle.ReferenceType vs *k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1.ReferenceType)
	return nil
}

func autoConvert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(in *FlunderStatus, out *wardle.FlunderStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus is an autogenerated conversion function.
func Convert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(in *FlunderStatus, out *wardle.FlunderStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_FlunderStatus_To_wardle_FlunderStatus(in, out, s)
}

func autoConvert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(in *wardle.FlunderStatus, out *FlunderStatus, s conversion.Scope) error {
	return nil
}

// Convert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus is an autogenerated conversion function.
func Convert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(in *wardle.FlunderStatus, out *FlunderStatus, s conversion.Scope) error {
	return autoConvert_wardle_FlunderStatus_To_v1alpha1_FlunderStatus(in, out, s)
}
