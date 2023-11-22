// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0--rc2
// source: protos/product.proto

package product

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

type QueryProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *QueryProductRequest) Reset() {
	*x = QueryProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductRequest) ProtoMessage() {}

func (x *QueryProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductRequest.ProtoReflect.Descriptor instead.
func (*QueryProductRequest) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{0}
}

func (x *QueryProductRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type UpdateProductNetworthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProductNetworthRequest) Reset() {
	*x = UpdateProductNetworthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductNetworthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductNetworthRequest) ProtoMessage() {}

func (x *UpdateProductNetworthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductNetworthRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductNetworthRequest) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{1}
}

type UpdateProductNetworthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64                      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *UpdateProductNetworthData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateProductNetworthResponse) Reset() {
	*x = UpdateProductNetworthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductNetworthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductNetworthResponse) ProtoMessage() {}

func (x *UpdateProductNetworthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductNetworthResponse.ProtoReflect.Descriptor instead.
func (*UpdateProductNetworthResponse) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProductNetworthResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateProductNetworthResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpdateProductNetworthResponse) GetData() *UpdateProductNetworthData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateProductNetworthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProductNetworthData) Reset() {
	*x = UpdateProductNetworthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductNetworthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductNetworthData) ProtoMessage() {}

func (x *UpdateProductNetworthData) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductNetworthData.ProtoReflect.Descriptor instead.
func (*UpdateProductNetworthData) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{3}
}

type QueryProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *QueryProductData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *QueryProductResponse) Reset() {
	*x = QueryProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductResponse) ProtoMessage() {}

func (x *QueryProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductResponse.ProtoReflect.Descriptor instead.
func (*QueryProductResponse) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{4}
}

func (x *QueryProductResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *QueryProductResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *QueryProductResponse) GetData() *QueryProductData {
	if x != nil {
		return x.Data
	}
	return nil
}

type QueryProductData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 产品代码
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// 产品名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 产品类型
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// 产品风险等级
	RiskLevel int32 `protobuf:"varint,4,opt,name=risk_level,json=riskLevel,proto3" json:"risk_level,omitempty"`
	// 产品状态
	Status int32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	// 所属公司
	Company string `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`
	// 产品经理
	Manager string `protobuf:"bytes,7,opt,name=manager,proto3" json:"manager,omitempty"`
	// 起购额度
	Buymin string `protobuf:"bytes,8,opt,name=buymin,proto3" json:"buymin,omitempty"`
	// 当前净值
	Networth string `protobuf:"bytes,9,opt,name=networth,proto3" json:"networth,omitempty"`
}

func (x *QueryProductData) Reset() {
	*x = QueryProductData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductData) ProtoMessage() {}

func (x *QueryProductData) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductData.ProtoReflect.Descriptor instead.
func (*QueryProductData) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{5}
}

func (x *QueryProductData) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryProductData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryProductData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *QueryProductData) GetRiskLevel() int32 {
	if x != nil {
		return x.RiskLevel
	}
	return 0
}

func (x *QueryProductData) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueryProductData) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *QueryProductData) GetManager() string {
	if x != nil {
		return x.Manager
	}
	return ""
}

func (x *QueryProductData) GetBuymin() string {
	if x != nil {
		return x.Buymin
	}
	return ""
}

func (x *QueryProductData) GetNetworth() string {
	if x != nil {
		return x.Networth
	}
	return ""
}

var File_protos_product_proto protoreflect.FileDescriptor

var file_protos_product_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x1e, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x75, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x63, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xed, 0x01, 0x0a, 0x10, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x69, 0x73, 0x6b, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x69, 0x73, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x79, 0x6d, 0x69, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x79, 0x6d, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x32, 0xa9, 0x01, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a,
	0x15, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75,
	0x6e, 0x64, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x01, 0x5a, 0x14, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_product_proto_rawDescOnce sync.Once
	file_protos_product_proto_rawDescData = file_protos_product_proto_rawDesc
)

func file_protos_product_proto_rawDescGZIP() []byte {
	file_protos_product_proto_rawDescOnce.Do(func() {
		file_protos_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_product_proto_rawDescData)
	})
	return file_protos_product_proto_rawDescData
}

var file_protos_product_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_product_proto_goTypes = []interface{}{
	(*QueryProductRequest)(nil),           // 0: QueryProductRequest
	(*UpdateProductNetworthRequest)(nil),  // 1: UpdateProductNetworthRequest
	(*UpdateProductNetworthResponse)(nil), // 2: UpdateProductNetworthResponse
	(*UpdateProductNetworthData)(nil),     // 3: UpdateProductNetworthData
	(*QueryProductResponse)(nil),          // 4: QueryProductResponse
	(*QueryProductData)(nil),              // 5: QueryProductData
}
var file_protos_product_proto_depIdxs = []int32{
	3, // 0: UpdateProductNetworthResponse.data:type_name -> UpdateProductNetworthData
	5, // 1: QueryProductResponse.data:type_name -> QueryProductData
	1, // 2: ProductService.updateProductNetworth:input_type -> UpdateProductNetworthRequest
	0, // 3: ProductService.queryProduct:input_type -> QueryProductRequest
	2, // 4: ProductService.updateProductNetworth:output_type -> UpdateProductNetworthResponse
	4, // 5: ProductService.queryProduct:output_type -> QueryProductResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_product_proto_init() }
func file_protos_product_proto_init() {
	if File_protos_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductRequest); i {
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
		file_protos_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductNetworthRequest); i {
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
		file_protos_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductNetworthResponse); i {
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
		file_protos_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductNetworthData); i {
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
		file_protos_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductResponse); i {
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
		file_protos_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductData); i {
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
			RawDescriptor: file_protos_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_product_proto_goTypes,
		DependencyIndexes: file_protos_product_proto_depIdxs,
		MessageInfos:      file_protos_product_proto_msgTypes,
	}.Build()
	File_protos_product_proto = out.File
	file_protos_product_proto_rawDesc = nil
	file_protos_product_proto_goTypes = nil
	file_protos_product_proto_depIdxs = nil
}