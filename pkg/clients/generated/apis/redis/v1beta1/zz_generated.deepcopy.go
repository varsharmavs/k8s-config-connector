//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMaintenancePolicy) DeepCopyInto(out *InstanceMaintenancePolicy) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.WeeklyMaintenanceWindow != nil {
		in, out := &in.WeeklyMaintenanceWindow, &out.WeeklyMaintenanceWindow
		*out = make([]InstanceWeeklyMaintenanceWindow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMaintenancePolicy.
func (in *InstanceMaintenancePolicy) DeepCopy() *InstanceMaintenancePolicy {
	if in == nil {
		return nil
	}
	out := new(InstanceMaintenancePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMaintenanceScheduleStatus) DeepCopyInto(out *InstanceMaintenanceScheduleStatus) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.ScheduleDeadlineTime != nil {
		in, out := &in.ScheduleDeadlineTime, &out.ScheduleDeadlineTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMaintenanceScheduleStatus.
func (in *InstanceMaintenanceScheduleStatus) DeepCopy() *InstanceMaintenanceScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceMaintenanceScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceNodesStatus) DeepCopyInto(out *InstanceNodesStatus) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceNodesStatus.
func (in *InstanceNodesStatus) DeepCopy() *InstanceNodesStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceNodesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancePersistenceConfig) DeepCopyInto(out *InstancePersistenceConfig) {
	*out = *in
	if in.PersistenceMode != nil {
		in, out := &in.PersistenceMode, &out.PersistenceMode
		*out = new(string)
		**out = **in
	}
	if in.RdbNextSnapshotTime != nil {
		in, out := &in.RdbNextSnapshotTime, &out.RdbNextSnapshotTime
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotPeriod != nil {
		in, out := &in.RdbSnapshotPeriod, &out.RdbSnapshotPeriod
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotStartTime != nil {
		in, out := &in.RdbSnapshotStartTime, &out.RdbSnapshotStartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancePersistenceConfig.
func (in *InstancePersistenceConfig) DeepCopy() *InstancePersistenceConfig {
	if in == nil {
		return nil
	}
	out := new(InstancePersistenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceServerCaCertsStatus) DeepCopyInto(out *InstanceServerCaCertsStatus) {
	*out = *in
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ExpireTime != nil {
		in, out := &in.ExpireTime, &out.ExpireTime
		*out = new(string)
		**out = **in
	}
	if in.SerialNumber != nil {
		in, out := &in.SerialNumber, &out.SerialNumber
		*out = new(string)
		**out = **in
	}
	if in.Sha1Fingerprint != nil {
		in, out := &in.Sha1Fingerprint, &out.Sha1Fingerprint
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceServerCaCertsStatus.
func (in *InstanceServerCaCertsStatus) DeepCopy() *InstanceServerCaCertsStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceServerCaCertsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStartTime) DeepCopyInto(out *InstanceStartTime) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(int)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(int)
		**out = **in
	}
	if in.Nanos != nil {
		in, out := &in.Nanos, &out.Nanos
		*out = new(int)
		**out = **in
	}
	if in.Seconds != nil {
		in, out := &in.Seconds, &out.Seconds
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStartTime.
func (in *InstanceStartTime) DeepCopy() *InstanceStartTime {
	if in == nil {
		return nil
	}
	out := new(InstanceStartTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceWeeklyMaintenanceWindow) DeepCopyInto(out *InstanceWeeklyMaintenanceWindow) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceWeeklyMaintenanceWindow.
func (in *InstanceWeeklyMaintenanceWindow) DeepCopy() *InstanceWeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(InstanceWeeklyMaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisInstance) DeepCopyInto(out *RedisInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisInstance.
func (in *RedisInstance) DeepCopy() *RedisInstance {
	if in == nil {
		return nil
	}
	out := new(RedisInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisInstanceList) DeepCopyInto(out *RedisInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisInstanceList.
func (in *RedisInstanceList) DeepCopy() *RedisInstanceList {
	if in == nil {
		return nil
	}
	out := new(RedisInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisInstanceSpec) DeepCopyInto(out *RedisInstanceSpec) {
	*out = *in
	if in.AlternativeLocationId != nil {
		in, out := &in.AlternativeLocationId, &out.AlternativeLocationId
		*out = new(string)
		**out = **in
	}
	if in.AuthEnabled != nil {
		in, out := &in.AuthEnabled, &out.AuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthString != nil {
		in, out := &in.AuthString, &out.AuthString
		*out = new(string)
		**out = **in
	}
	if in.AuthorizedNetworkRef != nil {
		in, out := &in.AuthorizedNetworkRef, &out.AuthorizedNetworkRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.ConnectMode != nil {
		in, out := &in.ConnectMode, &out.ConnectMode
		*out = new(string)
		**out = **in
	}
	if in.CustomerManagedKeyRef != nil {
		in, out := &in.CustomerManagedKeyRef, &out.CustomerManagedKeyRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.LocationId != nil {
		in, out := &in.LocationId, &out.LocationId
		*out = new(string)
		**out = **in
	}
	if in.MaintenancePolicy != nil {
		in, out := &in.MaintenancePolicy, &out.MaintenancePolicy
		*out = new(InstanceMaintenancePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistenceConfig != nil {
		in, out := &in.PersistenceConfig, &out.PersistenceConfig
		*out = new(InstancePersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadReplicasMode != nil {
		in, out := &in.ReadReplicasMode, &out.ReadReplicasMode
		*out = new(string)
		**out = **in
	}
	if in.RedisConfigs != nil {
		in, out := &in.RedisConfigs, &out.RedisConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int)
		**out = **in
	}
	if in.ReservedIpRange != nil {
		in, out := &in.ReservedIpRange, &out.ReservedIpRange
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.SecondaryIpRange != nil {
		in, out := &in.SecondaryIpRange, &out.SecondaryIpRange
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionMode != nil {
		in, out := &in.TransitEncryptionMode, &out.TransitEncryptionMode
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisInstanceSpec.
func (in *RedisInstanceSpec) DeepCopy() *RedisInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(RedisInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisInstanceStatus) DeepCopyInto(out *RedisInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.CurrentLocationId != nil {
		in, out := &in.CurrentLocationId, &out.CurrentLocationId
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceSchedule != nil {
		in, out := &in.MaintenanceSchedule, &out.MaintenanceSchedule
		*out = make([]InstanceMaintenanceScheduleStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]InstanceNodesStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.PersistenceIamIdentity != nil {
		in, out := &in.PersistenceIamIdentity, &out.PersistenceIamIdentity
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.ReadEndpoint != nil {
		in, out := &in.ReadEndpoint, &out.ReadEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ReadEndpointPort != nil {
		in, out := &in.ReadEndpointPort, &out.ReadEndpointPort
		*out = new(int)
		**out = **in
	}
	if in.ServerCaCerts != nil {
		in, out := &in.ServerCaCerts, &out.ServerCaCerts
		*out = make([]InstanceServerCaCertsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisInstanceStatus.
func (in *RedisInstanceStatus) DeepCopy() *RedisInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(RedisInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
