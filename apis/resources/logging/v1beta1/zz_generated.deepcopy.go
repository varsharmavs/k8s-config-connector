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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingLogMetric) DeepCopyInto(out *LoggingLogMetric) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingLogMetric.
func (in *LoggingLogMetric) DeepCopy() *LoggingLogMetric {
	if in == nil {
		return nil
	}
	out := new(LoggingLogMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingLogMetric) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingLogMetricList) DeepCopyInto(out *LoggingLogMetricList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoggingLogMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingLogMetricList.
func (in *LoggingLogMetricList) DeepCopy() *LoggingLogMetricList {
	if in == nil {
		return nil
	}
	out := new(LoggingLogMetricList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingLogMetricList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingLogMetricSpec) DeepCopyInto(out *LoggingLogMetricSpec) {
	*out = *in
	if in.LoggingLogBucketRef != nil {
		in, out := &in.LoggingLogBucketRef, &out.LoggingLogBucketRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.BucketOptions != nil {
		in, out := &in.BucketOptions, &out.BucketOptions
		*out = new(LogmetricBucketOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.LabelExtractors != nil {
		in, out := &in.LabelExtractors, &out.LabelExtractors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MetricDescriptor != nil {
		in, out := &in.MetricDescriptor, &out.MetricDescriptor
		*out = new(LogmetricMetricDescriptor)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ValueExtractor != nil {
		in, out := &in.ValueExtractor, &out.ValueExtractor
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingLogMetricSpec.
func (in *LoggingLogMetricSpec) DeepCopy() *LoggingLogMetricSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingLogMetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingLogMetricStatus) DeepCopyInto(out *LoggingLogMetricStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = (*in).DeepCopy()
	}
	if in.MetricDescriptor != nil {
		in, out := &in.MetricDescriptor, &out.MetricDescriptor
		*out = new(LogmetricMetricDescriptorStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingLogMetricStatus.
func (in *LoggingLogMetricStatus) DeepCopy() *LoggingLogMetricStatus {
	if in == nil {
		return nil
	}
	out := new(LoggingLogMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogmetricBucketOptions) DeepCopyInto(out *LogmetricBucketOptions) {
	*out = *in
	if in.ExplicitBuckets != nil {
		in, out := &in.ExplicitBuckets, &out.ExplicitBuckets
		*out = new(LogmetricExplicitBuckets)
		(*in).DeepCopyInto(*out)
	}
	if in.ExponentialBuckets != nil {
		in, out := &in.ExponentialBuckets, &out.ExponentialBuckets
		*out = new(LogmetricExponentialBuckets)
		(*in).DeepCopyInto(*out)
	}
	if in.LinearBuckets != nil {
		in, out := &in.LinearBuckets, &out.LinearBuckets
		*out = new(LogmetricLinearBuckets)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogmetricBucketOptions.
func (in *LogmetricBucketOptions) DeepCopy() *LogmetricBucketOptions {
	if in == nil {
		return nil
	}
	out := new(LogmetricBucketOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogmetricExplicitBuckets) DeepCopyInto(out *LogmetricExplicitBuckets) {
	*out = *in
	if in.Bounds != nil {
		in, out := &in.Bounds, &out.Bounds
		*out = make([]float64, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogmetricExplicitBuckets.
func (in *LogmetricExplicitBuckets) DeepCopy() *LogmetricExplicitBuckets {
	if in == nil {
		return nil
	}
	out := new(LogmetricExplicitBuckets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogmetricExponentialBuckets) DeepCopyInto(out *LogmetricExponentialBuckets) {
	*out = *in
	if in.GrowthFactor != nil {
		in, out := &in.GrowthFactor, &out.GrowthFactor
		*out = new(float64)
		**out = **in
	}
	if in.NumFiniteBuckets != nil {
		in, out := &in.NumFiniteBuckets, &out.NumFiniteBuckets
		*out = new(int64)
		**out = **in
	}
	if in.Scale != nil {
		in, out := &in.Scale, &out.Scale
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogmetricExponentialBuckets.
func (in *LogmetricExponentialBuckets) DeepCopy() *LogmetricExponentialBuckets {
	if in == nil {
		return nil
	}
	out := new(LogmetricExponentialBuckets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogmetricLabels) DeepCopyInto(out *LogmetricLabels) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.ValueType != nil {
		in, out := &in.ValueType, &out.ValueType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogmetricLabels.
func (in *LogmetricLabels) DeepCopy() *LogmetricLabels {
	if in == nil {
		return nil
	}
	out := new(LogmetricLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogmetricLinearBuckets) DeepCopyInto(out *LogmetricLinearBuckets) {
	*out = *in
	if in.NumFiniteBuckets != nil {
		in, out := &in.NumFiniteBuckets, &out.NumFiniteBuckets
		*out = new(int64)
		**out = **in
	}
	if in.Offset != nil {
		in, out := &in.Offset, &out.Offset
		*out = new(float64)
		**out = **in
	}
	if in.Width != nil {
		in, out := &in.Width, &out.Width
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogmetricLinearBuckets.
func (in *LogmetricLinearBuckets) DeepCopy() *LogmetricLinearBuckets {
	if in == nil {
		return nil
	}
	out := new(LogmetricLinearBuckets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogmetricMetadata) DeepCopyInto(out *LogmetricMetadata) {
	*out = *in
	if in.IngestDelay != nil {
		in, out := &in.IngestDelay, &out.IngestDelay
		*out = new(string)
		**out = **in
	}
	if in.SamplePeriod != nil {
		in, out := &in.SamplePeriod, &out.SamplePeriod
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogmetricMetadata.
func (in *LogmetricMetadata) DeepCopy() *LogmetricMetadata {
	if in == nil {
		return nil
	}
	out := new(LogmetricMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogmetricMetricDescriptor) DeepCopyInto(out *LogmetricMetricDescriptor) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]LogmetricLabels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LaunchStage != nil {
		in, out := &in.LaunchStage, &out.LaunchStage
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(LogmetricMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricKind != nil {
		in, out := &in.MetricKind, &out.MetricKind
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	if in.ValueType != nil {
		in, out := &in.ValueType, &out.ValueType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogmetricMetricDescriptor.
func (in *LogmetricMetricDescriptor) DeepCopy() *LogmetricMetricDescriptor {
	if in == nil {
		return nil
	}
	out := new(LogmetricMetricDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogmetricMetricDescriptorStatus) DeepCopyInto(out *LogmetricMetricDescriptorStatus) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.MonitoredResourceTypes != nil {
		in, out := &in.MonitoredResourceTypes, &out.MonitoredResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogmetricMetricDescriptorStatus.
func (in *LogmetricMetricDescriptorStatus) DeepCopy() *LogmetricMetricDescriptorStatus {
	if in == nil {
		return nil
	}
	out := new(LogmetricMetricDescriptorStatus)
	in.DeepCopyInto(out)
	return out
}
