// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: controller/api/v1/warning.proto

package api

import (
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

// FieldWarning contains warning information on a per field basis.
type FieldWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The warning regarding the provided field name.
	Warning string `protobuf:"bytes,2,opt,name=warning,proto3" json:"warning,omitempty"`
}

func (x *FieldWarning) Reset() {
	*x = FieldWarning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_v1_warning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldWarning) ProtoMessage() {}

func (x *FieldWarning) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_v1_warning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldWarning.ProtoReflect.Descriptor instead.
func (*FieldWarning) Descriptor() ([]byte, []int) {
	return file_controller_api_v1_warning_proto_rawDescGZIP(), []int{0}
}

func (x *FieldWarning) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FieldWarning) GetWarning() string {
	if x != nil {
		return x.Warning
	}
	return ""
}

// ActionWarning contains warning information regarding a specific action.
type ActionWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the action
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The warning regarding the provided action
	Warning string `protobuf:"bytes,2,opt,name=warning,proto3" json:"warning,omitempty"`
}

func (x *ActionWarning) Reset() {
	*x = ActionWarning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_v1_warning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionWarning) ProtoMessage() {}

func (x *ActionWarning) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_v1_warning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionWarning.ProtoReflect.Descriptor instead.
func (*ActionWarning) Descriptor() ([]byte, []int) {
	return file_controller_api_v1_warning_proto_rawDescGZIP(), []int{1}
}

func (x *ActionWarning) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActionWarning) GetWarning() string {
	if x != nil {
		return x.Warning
	}
	return ""
}

// BehaviorWarning contains a warning about Boundary behavior which may be
// surprising or not intuitive.
type BehaviorWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The warning text regarding the behavior
	Warning string `protobuf:"bytes,1,opt,name=warning,proto3" json:"warning,omitempty"`
}

func (x *BehaviorWarning) Reset() {
	*x = BehaviorWarning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_v1_warning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BehaviorWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BehaviorWarning) ProtoMessage() {}

func (x *BehaviorWarning) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_v1_warning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BehaviorWarning.ProtoReflect.Descriptor instead.
func (*BehaviorWarning) Descriptor() ([]byte, []int) {
	return file_controller_api_v1_warning_proto_rawDescGZIP(), []int{2}
}

func (x *BehaviorWarning) GetWarning() string {
	if x != nil {
		return x.Warning
	}
	return ""
}

// A warning in the Boundary system.
type Warning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Types that are assignable to Warning:
	//
	//	*Warning_RequestField
	//	*Warning_Action
	//	*Warning_Behavior
	Warning isWarning_Warning `protobuf_oneof:"warning"`
}

func (x *Warning) Reset() {
	*x = Warning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_v1_warning_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warning) ProtoMessage() {}

func (x *Warning) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_v1_warning_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warning.ProtoReflect.Descriptor instead.
func (*Warning) Descriptor() ([]byte, []int) {
	return file_controller_api_v1_warning_proto_rawDescGZIP(), []int{3}
}

func (x *Warning) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (m *Warning) GetWarning() isWarning_Warning {
	if m != nil {
		return m.Warning
	}
	return nil
}

func (x *Warning) GetRequestField() *FieldWarning {
	if x, ok := x.GetWarning().(*Warning_RequestField); ok {
		return x.RequestField
	}
	return nil
}

func (x *Warning) GetAction() *ActionWarning {
	if x, ok := x.GetWarning().(*Warning_Action); ok {
		return x.Action
	}
	return nil
}

func (x *Warning) GetBehavior() *BehaviorWarning {
	if x, ok := x.GetWarning().(*Warning_Behavior); ok {
		return x.Behavior
	}
	return nil
}

type isWarning_Warning interface {
	isWarning_Warning()
}

type Warning_RequestField struct {
	RequestField *FieldWarning `protobuf:"bytes,2,opt,name=request_field,json=request_fields,proto3,oneof"`
}

type Warning_Action struct {
	Action *ActionWarning `protobuf:"bytes,3,opt,name=action,proto3,oneof"`
}

type Warning_Behavior struct {
	Behavior *BehaviorWarning `protobuf:"bytes,4,opt,name=behavior,proto3,oneof"`
}

func (*Warning_RequestField) isWarning_Warning() {}

func (*Warning_Action) isWarning_Warning() {}

func (*Warning_Behavior) isWarning_Warning() {}

// Warning is returned by the JSON API when a warning occurs.
type WarningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Request-field-specific warning details.
	Warnings []*Warning `protobuf:"bytes,1,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *WarningResponse) Reset() {
	*x = WarningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_v1_warning_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarningResponse) ProtoMessage() {}

func (x *WarningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_v1_warning_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarningResponse.ProtoReflect.Descriptor instead.
func (*WarningResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_v1_warning_proto_rawDescGZIP(), []int{4}
}

func (x *WarningResponse) GetWarnings() []*Warning {
	if x != nil {
		return x.Warnings
	}
	return nil
}

var File_controller_api_v1_warning_proto protoreflect.FileDescriptor

var file_controller_api_v1_warning_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x22, 0x3c, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x57, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x22, 0x3d, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x22, 0x2b, 0x0a, 0x0f, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x57, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0xf0,
	0x01, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x48,
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x57,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x08, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x22, 0x49, 0x0a, 0x0f, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_v1_warning_proto_rawDescOnce sync.Once
	file_controller_api_v1_warning_proto_rawDescData = file_controller_api_v1_warning_proto_rawDesc
)

func file_controller_api_v1_warning_proto_rawDescGZIP() []byte {
	file_controller_api_v1_warning_proto_rawDescOnce.Do(func() {
		file_controller_api_v1_warning_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_v1_warning_proto_rawDescData)
	})
	return file_controller_api_v1_warning_proto_rawDescData
}

var file_controller_api_v1_warning_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_controller_api_v1_warning_proto_goTypes = []interface{}{
	(*FieldWarning)(nil),    // 0: controller.api.v1.FieldWarning
	(*ActionWarning)(nil),   // 1: controller.api.v1.ActionWarning
	(*BehaviorWarning)(nil), // 2: controller.api.v1.BehaviorWarning
	(*Warning)(nil),         // 3: controller.api.v1.Warning
	(*WarningResponse)(nil), // 4: controller.api.v1.WarningResponse
}
var file_controller_api_v1_warning_proto_depIdxs = []int32{
	0, // 0: controller.api.v1.Warning.request_field:type_name -> controller.api.v1.FieldWarning
	1, // 1: controller.api.v1.Warning.action:type_name -> controller.api.v1.ActionWarning
	2, // 2: controller.api.v1.Warning.behavior:type_name -> controller.api.v1.BehaviorWarning
	3, // 3: controller.api.v1.WarningResponse.warnings:type_name -> controller.api.v1.Warning
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_controller_api_v1_warning_proto_init() }
func file_controller_api_v1_warning_proto_init() {
	if File_controller_api_v1_warning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_v1_warning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldWarning); i {
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
		file_controller_api_v1_warning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionWarning); i {
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
		file_controller_api_v1_warning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BehaviorWarning); i {
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
		file_controller_api_v1_warning_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warning); i {
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
		file_controller_api_v1_warning_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarningResponse); i {
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
	file_controller_api_v1_warning_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Warning_RequestField)(nil),
		(*Warning_Action)(nil),
		(*Warning_Behavior)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_api_v1_warning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_v1_warning_proto_goTypes,
		DependencyIndexes: file_controller_api_v1_warning_proto_depIdxs,
		MessageInfos:      file_controller_api_v1_warning_proto_msgTypes,
	}.Build()
	File_controller_api_v1_warning_proto = out.File
	file_controller_api_v1_warning_proto_rawDesc = nil
	file_controller_api_v1_warning_proto_goTypes = nil
	file_controller_api_v1_warning_proto_depIdxs = nil
}
