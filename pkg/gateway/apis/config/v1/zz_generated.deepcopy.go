// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
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
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.FrontProxy != nil {
		in, out := &in.FrontProxy, &out.FrontProxy
		*out = new(FrontProxyComponent)
		**out = **in
	}
	if in.Passthrough != nil {
		in, out := &in.Passthrough, &out.Passthrough
		*out = new(PassthroughComponent)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Components) DeepCopyInto(out *Components) {
	*out = *in
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.Business != nil {
		in, out := &in.Business, &out.Business
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.Notify != nil {
		in, out := &in.Notify, &out.Notify
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.Monitor != nil {
		in, out := &in.Monitor, &out.Monitor
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.LogAgent != nil {
		in, out := &in.LogAgent, &out.LogAgent
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	if in.Mesh != nil {
		in, out := &in.Mesh, &out.Mesh
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Components.
func (in *Components) DeepCopy() *Components {
	if in == nil {
		return nil
	}
	out := new(Components)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleConfig) DeepCopyInto(out *ConsoleConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleConfig.
func (in *ConsoleConfig) DeepCopy() *ConsoleConfig {
	if in == nil {
		return nil
	}
	out := new(ConsoleConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontProxyComponent) DeepCopyInto(out *FrontProxyComponent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontProxyComponent.
func (in *FrontProxyComponent) DeepCopy() *FrontProxyComponent {
	if in == nil {
		return nil
	}
	out := new(FrontProxyComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfiguration) DeepCopyInto(out *GatewayConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Components.DeepCopyInto(&out.Components)
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(Registry)
		**out = **in
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(Auth)
		**out = **in
	}
	if in.ConsoleConfig != nil {
		in, out := &in.ConsoleConfig, &out.ConsoleConfig
		*out = new(ConsoleConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfiguration.
func (in *GatewayConfiguration) DeepCopy() *GatewayConfiguration {
	if in == nil {
		return nil
	}
	out := new(GatewayConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PassthroughComponent) DeepCopyInto(out *PassthroughComponent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PassthroughComponent.
func (in *PassthroughComponent) DeepCopy() *PassthroughComponent {
	if in == nil {
		return nil
	}
	out := new(PassthroughComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}
