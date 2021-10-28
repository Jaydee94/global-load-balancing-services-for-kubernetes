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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceptedStatus) DeepCopyInto(out *AcceptedStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptedStatus.
func (in *AcceptedStatus) DeepCopy() *AcceptedStatus {
	if in == nil {
		return nil
	}
	out := new(AcceptedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfig) DeepCopyInto(out *BackendConfig) {
	*out = *in
	out.Service = in.Service
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfig.
func (in *BackendConfig) DeepCopy() *BackendConfig {
	if in == nil {
		return nil
	}
	out := new(BackendConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSet) DeepCopyInto(out *ClusterSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSet.
func (in *ClusterSet) DeepCopy() *ClusterSet {
	if in == nil {
		return nil
	}
	out := new(ClusterSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetList) DeepCopyInto(out *ClusterSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetList.
func (in *ClusterSetList) DeepCopy() *ClusterSetList {
	if in == nil {
		return nil
	}
	out := new(ClusterSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetSpec) DeepCopyInto(out *ClusterSetSpec) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]ClusterConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetSpec.
func (in *ClusterSetSpec) DeepCopy() *ClusterSetSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSetStatus) DeepCopyInto(out *ClusterSetStatus) {
	*out = *in
	if in.ServiceDiscovery != nil {
		in, out := &in.ServiceDiscovery, &out.ServiceDiscovery
		*out = make([]ServiceDiscoveryStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSetStatus.
func (in *ClusterSetStatus) DeepCopy() *ClusterSetStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBConfig) DeepCopyInto(out *GSLBConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBConfig.
func (in *GSLBConfig) DeepCopy() *GSLBConfig {
	if in == nil {
		return nil
	}
	out := new(GSLBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GSLBConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBConfigList) DeepCopyInto(out *GSLBConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GSLBConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBConfigList.
func (in *GSLBConfigList) DeepCopy() *GSLBConfigList {
	if in == nil {
		return nil
	}
	out := new(GSLBConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GSLBConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBConfigSpec) DeepCopyInto(out *GSLBConfigSpec) {
	*out = *in
	out.GSLBLeader = in.GSLBLeader
	if in.MemberClusters != nil {
		in, out := &in.MemberClusters, &out.MemberClusters
		*out = make([]MemberCluster, len(*in))
		copy(*out, *in)
	}
	if in.UseCustomGlobalFqdn != nil {
		in, out := &in.UseCustomGlobalFqdn, &out.UseCustomGlobalFqdn
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBConfigSpec.
func (in *GSLBConfigSpec) DeepCopy() *GSLBConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GSLBConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBConfigSpecList) DeepCopyInto(out *GSLBConfigSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GSLBConfigSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBConfigSpecList.
func (in *GSLBConfigSpecList) DeepCopy() *GSLBConfigSpecList {
	if in == nil {
		return nil
	}
	out := new(GSLBConfigSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GSLBConfigSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBConfigStatus) DeepCopyInto(out *GSLBConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBConfigStatus.
func (in *GSLBConfigStatus) DeepCopy() *GSLBConfigStatus {
	if in == nil {
		return nil
	}
	out := new(GSLBConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBHostRule) DeepCopyInto(out *GSLBHostRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBHostRule.
func (in *GSLBHostRule) DeepCopy() *GSLBHostRule {
	if in == nil {
		return nil
	}
	out := new(GSLBHostRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GSLBHostRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBHostRuleList) DeepCopyInto(out *GSLBHostRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GSLBHostRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBHostRuleList.
func (in *GSLBHostRuleList) DeepCopy() *GSLBHostRuleList {
	if in == nil {
		return nil
	}
	out := new(GSLBHostRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GSLBHostRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBHostRuleSpec) DeepCopyInto(out *GSLBHostRuleSpec) {
	*out = *in
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(int)
		**out = **in
	}
	if in.PoolAlgorithmSettings != nil {
		in, out := &in.PoolAlgorithmSettings, &out.PoolAlgorithmSettings
		*out = new(PoolAlgorithmSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.SitePersistence != nil {
		in, out := &in.SitePersistence, &out.SitePersistence
		*out = new(SitePersistence)
		**out = **in
	}
	if in.ThirdPartyMembers != nil {
		in, out := &in.ThirdPartyMembers, &out.ThirdPartyMembers
		*out = make([]ThirdPartyMember, len(*in))
		copy(*out, *in)
	}
	if in.HealthMonitorRefs != nil {
		in, out := &in.HealthMonitorRefs, &out.HealthMonitorRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TrafficSplit != nil {
		in, out := &in.TrafficSplit, &out.TrafficSplit
		*out = make([]TrafficSplitElem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBHostRuleSpec.
func (in *GSLBHostRuleSpec) DeepCopy() *GSLBHostRuleSpec {
	if in == nil {
		return nil
	}
	out := new(GSLBHostRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBHostRuleStatus) DeepCopyInto(out *GSLBHostRuleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBHostRuleStatus.
func (in *GSLBHostRuleStatus) DeepCopy() *GSLBHostRuleStatus {
	if in == nil {
		return nil
	}
	out := new(GSLBHostRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GSLBLeader) DeepCopyInto(out *GSLBLeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GSLBLeader.
func (in *GSLBLeader) DeepCopy() *GSLBLeader {
	if in == nil {
		return nil
	}
	out := new(GSLBLeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoFallback) DeepCopyInto(out *GeoFallback) {
	*out = *in
	if in.HashMask != nil {
		in, out := &in.HashMask, &out.HashMask
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoFallback.
func (in *GeoFallback) DeepCopy() *GeoFallback {
	if in == nil {
		return nil
	}
	out := new(GeoFallback)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressStatus) DeepCopyInto(out *IngressStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressStatus.
func (in *IngressStatus) DeepCopy() *IngressStatus {
	if in == nil {
		return nil
	}
	out := new(IngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]IngressStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberCluster) DeepCopyInto(out *MemberCluster) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberCluster.
func (in *MemberCluster) DeepCopy() *MemberCluster {
	if in == nil {
		return nil
	}
	out := new(MemberCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterIngress) DeepCopyInto(out *MultiClusterIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterIngress.
func (in *MultiClusterIngress) DeepCopy() *MultiClusterIngress {
	if in == nil {
		return nil
	}
	out := new(MultiClusterIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterIngressList) DeepCopyInto(out *MultiClusterIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterIngressList.
func (in *MultiClusterIngressList) DeepCopy() *MultiClusterIngressList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterIngressSpec) DeepCopyInto(out *MultiClusterIngressSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]BackendConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterIngressSpec.
func (in *MultiClusterIngressSpec) DeepCopy() *MultiClusterIngressSpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterIngressStatus) DeepCopyInto(out *MultiClusterIngressStatus) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterIngressStatus.
func (in *MultiClusterIngressStatus) DeepCopy() *MultiClusterIngressStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterIngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolAlgorithmSettings) DeepCopyInto(out *PoolAlgorithmSettings) {
	*out = *in
	if in.HashMask != nil {
		in, out := &in.HashMask, &out.HashMask
		*out = new(int)
		**out = **in
	}
	if in.FallbackAlgorithm != nil {
		in, out := &in.FallbackAlgorithm, &out.FallbackAlgorithm
		*out = new(GeoFallback)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolAlgorithmSettings.
func (in *PoolAlgorithmSettings) DeepCopy() *PoolAlgorithmSettings {
	if in == nil {
		return nil
	}
	out := new(PoolAlgorithmSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryStatus) DeepCopyInto(out *ServiceDiscoveryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryStatus.
func (in *ServiceDiscoveryStatus) DeepCopy() *ServiceDiscoveryStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SitePersistence) DeepCopyInto(out *SitePersistence) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SitePersistence.
func (in *SitePersistence) DeepCopy() *SitePersistence {
	if in == nil {
		return nil
	}
	out := new(SitePersistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThirdPartyMember) DeepCopyInto(out *ThirdPartyMember) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThirdPartyMember.
func (in *ThirdPartyMember) DeepCopy() *ThirdPartyMember {
	if in == nil {
		return nil
	}
	out := new(ThirdPartyMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficSplitElem) DeepCopyInto(out *TrafficSplitElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficSplitElem.
func (in *TrafficSplitElem) DeepCopy() *TrafficSplitElem {
	if in == nil {
		return nil
	}
	out := new(TrafficSplitElem)
	in.DeepCopyInto(out)
	return out
}
