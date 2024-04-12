// Copyright 2023 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgcp/monitoring/dashboard/v1/table.proto

package dashboardpb

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum for metric metric_visualization
type TimeSeriesTable_MetricVisualization int32

const (
	// Unspecified state
	TimeSeriesTable_METRIC_VISUALIZATION_UNSPECIFIED TimeSeriesTable_MetricVisualization = 0
	// Default text rendering
	TimeSeriesTable_NUMBER TimeSeriesTable_MetricVisualization = 1
	// Horizontal bar rendering
	TimeSeriesTable_BAR TimeSeriesTable_MetricVisualization = 2
)

// Enum value maps for TimeSeriesTable_MetricVisualization.
var (
	TimeSeriesTable_MetricVisualization_name = map[int32]string{
		0: "METRIC_VISUALIZATION_UNSPECIFIED",
		1: "NUMBER",
		2: "BAR",
	}
	TimeSeriesTable_MetricVisualization_value = map[string]int32{
		"METRIC_VISUALIZATION_UNSPECIFIED": 0,
		"NUMBER":                           1,
		"BAR":                              2,
	}
)

func (x TimeSeriesTable_MetricVisualization) Enum() *TimeSeriesTable_MetricVisualization {
	p := new(TimeSeriesTable_MetricVisualization)
	*p = x
	return p
}

func (x TimeSeriesTable_MetricVisualization) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeSeriesTable_MetricVisualization) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_monitoring_dashboard_v1_table_proto_enumTypes[0].Descriptor()
}

func (TimeSeriesTable_MetricVisualization) Type() protoreflect.EnumType {
	return &file_mockgcp_monitoring_dashboard_v1_table_proto_enumTypes[0]
}

func (x TimeSeriesTable_MetricVisualization) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeSeriesTable_MetricVisualization.Descriptor instead.
func (TimeSeriesTable_MetricVisualization) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescGZIP(), []int{0, 0}
}

// A table that displays time series data.
type TimeSeriesTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The data displayed in this table.
	DataSets []*TimeSeriesTable_TableDataSet `protobuf:"bytes,1,rep,name=data_sets,json=dataSets,proto3" json:"data_sets,omitempty"`
	// Optional. Store rendering strategy
	MetricVisualization TimeSeriesTable_MetricVisualization `protobuf:"varint,2,opt,name=metric_visualization,json=metricVisualization,proto3,enum=mockgcp.monitoring.dashboard.v1.TimeSeriesTable_MetricVisualization" json:"metric_visualization,omitempty"`
	// Optional. The list of the persistent column settings for the table.
	ColumnSettings []*TimeSeriesTable_ColumnSettings `protobuf:"bytes,4,rep,name=column_settings,json=columnSettings,proto3" json:"column_settings,omitempty"`
}

func (x *TimeSeriesTable) Reset() {
	*x = TimeSeriesTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesTable) ProtoMessage() {}

func (x *TimeSeriesTable) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesTable.ProtoReflect.Descriptor instead.
func (*TimeSeriesTable) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescGZIP(), []int{0}
}

func (x *TimeSeriesTable) GetDataSets() []*TimeSeriesTable_TableDataSet {
	if x != nil {
		return x.DataSets
	}
	return nil
}

func (x *TimeSeriesTable) GetMetricVisualization() TimeSeriesTable_MetricVisualization {
	if x != nil {
		return x.MetricVisualization
	}
	return TimeSeriesTable_METRIC_VISUALIZATION_UNSPECIFIED
}

func (x *TimeSeriesTable) GetColumnSettings() []*TimeSeriesTable_ColumnSettings {
	if x != nil {
		return x.ColumnSettings
	}
	return nil
}

// Groups a time series query definition with table options.
type TimeSeriesTable_TableDataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Fields for querying time series data from the
	// Stackdriver metrics API.
	TimeSeriesQuery *TimeSeriesQuery `protobuf:"bytes,1,opt,name=time_series_query,json=timeSeriesQuery,proto3" json:"time_series_query,omitempty"`
	// Optional. A template string for naming `TimeSeries` in the resulting data
	// set. This should be a string with interpolations of the form
	// `${label_name}`, which will resolve to the label's value i.e.
	// "${resource.labels.project_id}."
	TableTemplate string `protobuf:"bytes,2,opt,name=table_template,json=tableTemplate,proto3" json:"table_template,omitempty"`
	// Optional. The lower bound on data point frequency for this data set,
	// implemented by specifying the minimum alignment period to use in a time
	// series query For example, if the data is published once every 10 minutes,
	// the `min_alignment_period` should be at least 10 minutes. It would not
	// make sense to fetch and align data at one minute intervals.
	MinAlignmentPeriod *duration.Duration `protobuf:"bytes,3,opt,name=min_alignment_period,json=minAlignmentPeriod,proto3" json:"min_alignment_period,omitempty"`
	// Optional. Table display options for configuring how the table is
	// rendered.
	TableDisplayOptions *TableDisplayOptions `protobuf:"bytes,4,opt,name=table_display_options,json=tableDisplayOptions,proto3" json:"table_display_options,omitempty"`
}

func (x *TimeSeriesTable_TableDataSet) Reset() {
	*x = TimeSeriesTable_TableDataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesTable_TableDataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesTable_TableDataSet) ProtoMessage() {}

func (x *TimeSeriesTable_TableDataSet) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesTable_TableDataSet.ProtoReflect.Descriptor instead.
func (*TimeSeriesTable_TableDataSet) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TimeSeriesTable_TableDataSet) GetTimeSeriesQuery() *TimeSeriesQuery {
	if x != nil {
		return x.TimeSeriesQuery
	}
	return nil
}

func (x *TimeSeriesTable_TableDataSet) GetTableTemplate() string {
	if x != nil {
		return x.TableTemplate
	}
	return ""
}

func (x *TimeSeriesTable_TableDataSet) GetMinAlignmentPeriod() *duration.Duration {
	if x != nil {
		return x.MinAlignmentPeriod
	}
	return nil
}

func (x *TimeSeriesTable_TableDataSet) GetTableDisplayOptions() *TableDisplayOptions {
	if x != nil {
		return x.TableDisplayOptions
	}
	return nil
}

// The persistent settings for a table's columns.
type TimeSeriesTable_ColumnSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The id of the column.
	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	// Required. Whether the column should be visible on page load.
	Visible bool `protobuf:"varint,2,opt,name=visible,proto3" json:"visible,omitempty"`
}

func (x *TimeSeriesTable_ColumnSettings) Reset() {
	*x = TimeSeriesTable_ColumnSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeriesTable_ColumnSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesTable_ColumnSettings) ProtoMessage() {}

func (x *TimeSeriesTable_ColumnSettings) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesTable_ColumnSettings.ProtoReflect.Descriptor instead.
func (*TimeSeriesTable_ColumnSettings) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TimeSeriesTable_ColumnSettings) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *TimeSeriesTable_ColumnSettings) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

var File_mockgcp_monitoring_dashboard_v1_table_proto protoreflect.FileDescriptor

var file_mockgcp_monitoring_dashboard_v1_table_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d,
	0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b,
	0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x06, 0x0a, 0x0f,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x5f, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x73,
	0x12, 0x7c, 0x0a, 0x14, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x69, 0x73, 0x75, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x44,
	0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x13, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d,
	0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63,
	0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0xde, 0x02,
	0x0a, 0x0c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x61,
	0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x6f, 0x63, 0x6b,
	0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x2a, 0x0a, 0x0e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a,
	0x14, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x6d, 0x69, 0x6e,
	0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12,
	0x6d, 0x0a, 0x15, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x13, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x4c,
	0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x1b, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x1d, 0x0a,
	0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x50, 0x0a, 0x13,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x56, 0x49,
	0x53, 0x55, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d,
	0x42, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41, 0x52, 0x10, 0x02, 0x42, 0xf4,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x70, 0x62,
	0x3b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x70, 0x62, 0xaa, 0x02, 0x24, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescOnce sync.Once
	file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescData = file_mockgcp_monitoring_dashboard_v1_table_proto_rawDesc
)

func file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescGZIP() []byte {
	file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescOnce.Do(func() {
		file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescData)
	})
	return file_mockgcp_monitoring_dashboard_v1_table_proto_rawDescData
}

var file_mockgcp_monitoring_dashboard_v1_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mockgcp_monitoring_dashboard_v1_table_proto_goTypes = []interface{}{
	(TimeSeriesTable_MetricVisualization)(0), // 0: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.MetricVisualization
	(*TimeSeriesTable)(nil),                  // 1: mockgcp.monitoring.dashboard.v1.TimeSeriesTable
	(*TimeSeriesTable_TableDataSet)(nil),     // 2: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet
	(*TimeSeriesTable_ColumnSettings)(nil),   // 3: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.ColumnSettings
	(*TimeSeriesQuery)(nil),                  // 4: mockgcp.monitoring.dashboard.v1.TimeSeriesQuery
	(*duration.Duration)(nil),                // 5: google.protobuf.Duration
	(*TableDisplayOptions)(nil),              // 6: mockgcp.monitoring.dashboard.v1.TableDisplayOptions
}
var file_mockgcp_monitoring_dashboard_v1_table_proto_depIdxs = []int32{
	2, // 0: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.data_sets:type_name -> mockgcp.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet
	0, // 1: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.metric_visualization:type_name -> mockgcp.monitoring.dashboard.v1.TimeSeriesTable.MetricVisualization
	3, // 2: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.column_settings:type_name -> mockgcp.monitoring.dashboard.v1.TimeSeriesTable.ColumnSettings
	4, // 3: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet.time_series_query:type_name -> mockgcp.monitoring.dashboard.v1.TimeSeriesQuery
	5, // 4: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet.min_alignment_period:type_name -> google.protobuf.Duration
	6, // 5: mockgcp.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet.table_display_options:type_name -> mockgcp.monitoring.dashboard.v1.TableDisplayOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_mockgcp_monitoring_dashboard_v1_table_proto_init() }
func file_mockgcp_monitoring_dashboard_v1_table_proto_init() {
	if File_mockgcp_monitoring_dashboard_v1_table_proto != nil {
		return
	}
	file_mockgcp_monitoring_dashboard_v1_metrics_proto_init()
	file_mockgcp_monitoring_dashboard_v1_table_display_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesTable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesTable_TableDataSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesTable_ColumnSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mockgcp_monitoring_dashboard_v1_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_monitoring_dashboard_v1_table_proto_goTypes,
		DependencyIndexes: file_mockgcp_monitoring_dashboard_v1_table_proto_depIdxs,
		EnumInfos:         file_mockgcp_monitoring_dashboard_v1_table_proto_enumTypes,
		MessageInfos:      file_mockgcp_monitoring_dashboard_v1_table_proto_msgTypes,
	}.Build()
	File_mockgcp_monitoring_dashboard_v1_table_proto = out.File
	file_mockgcp_monitoring_dashboard_v1_table_proto_rawDesc = nil
	file_mockgcp_monitoring_dashboard_v1_table_proto_goTypes = nil
	file_mockgcp_monitoring_dashboard_v1_table_proto_depIdxs = nil
}
