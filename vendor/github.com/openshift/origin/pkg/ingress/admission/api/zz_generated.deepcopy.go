// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_IngressAdmissionConfig, InType: reflect.TypeOf(&IngressAdmissionConfig{})},
	)
}

// DeepCopy_api_IngressAdmissionConfig is an autogenerated deepcopy function.
func DeepCopy_api_IngressAdmissionConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressAdmissionConfig)
		out := out.(*IngressAdmissionConfig)
		*out = *in
		return nil
	}
}
