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

// Code generated by deepcopy-gen. DO NOT EDIT.

package awsplugin

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.Oidc != nil {
		in, out := &in.Oidc, &out.Oidc
		*out = new(OIDC)
		**out = **in
	}
	if in.Cognito != nil {
		in, out := &in.Cognito, &out.Cognito
		*out = new(Coginito)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsPluginSpec) DeepCopyInto(out *AwsPluginSpec) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(Auth)
		(*in).DeepCopyInto(*out)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnablePodIamPolicy != nil {
		in, out := &in.EnablePodIamPolicy, &out.EnablePodIamPolicy
		*out = new(bool)
		**out = **in
	}
	if in.EnableNodeGroupLog != nil {
		in, out := &in.EnableNodeGroupLog, &out.EnableNodeGroupLog
		*out = new(bool)
		**out = **in
	}
	if in.ManagedCluster != nil {
		in, out := &in.ManagedCluster, &out.ManagedCluster
		*out = new(bool)
		**out = **in
	}
	if in.ManagedRelationDatabase != nil {
		in, out := &in.ManagedRelationDatabase, &out.ManagedRelationDatabase
		*out = new(RelationDatabaseConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedObjectStorage != nil {
		in, out := &in.ManagedObjectStorage, &out.ManagedObjectStorage
		*out = new(ObjectStorageConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsPluginSpec.
func (in *AwsPluginSpec) DeepCopy() *AwsPluginSpec {
	if in == nil {
		return nil
	}
	out := new(AwsPluginSpec)
	in.DeepCopyInto(out)
	return out
}

//DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	*out = *in
	//if in.Password != "" {
	//	in, out := &in.Password, &out.Password
	//	*out = ""
	//	**out = **in
	//}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuth.
func (in *BasicAuth) DeepCopy() *BasicAuth {
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Coginito) DeepCopyInto(out *Coginito) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Coginito.
func (in *Coginito) DeepCopy() *Coginito {
	if in == nil {
		return nil
	}
	out := new(Coginito)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KfAwsPlugin) DeepCopyInto(out *KfAwsPlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KfAwsPlugin.
func (in *KfAwsPlugin) DeepCopy() *KfAwsPlugin {
	if in == nil {
		return nil
	}
	out := new(KfAwsPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KfAwsPlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDC) DeepCopyInto(out *OIDC) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDC.
func (in *OIDC) DeepCopy() *OIDC {
	if in == nil {
		return nil
	}
	out := new(OIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageConfig) DeepCopyInto(out *ObjectStorageConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageConfig.
func (in *ObjectStorageConfig) DeepCopy() *ObjectStorageConfig {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelationDatabaseConfig) DeepCopyInto(out *RelationDatabaseConfig) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelationDatabaseConfig.
func (in *RelationDatabaseConfig) DeepCopy() *RelationDatabaseConfig {
	if in == nil {
		return nil
	}
	out := new(RelationDatabaseConfig)
	in.DeepCopyInto(out)
	return out
}