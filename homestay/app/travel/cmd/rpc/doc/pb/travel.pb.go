// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.9.0
// source: doc/travel.proto

package pb

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

type Homestay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title               string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	SubTitle            string `protobuf:"bytes,3,opt,name=subTitle,proto3" json:"subTitle,omitempty"`
	Banner              string `protobuf:"bytes,4,opt,name=banner,proto3" json:"banner,omitempty"`
	Info                string `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	PeopleNum           int64  `protobuf:"varint,6,opt,name=peopleNum,proto3" json:"peopleNum,omitempty"`                      //容纳人的数量
	HomestayBusinessId  int64  `protobuf:"varint,7,opt,name=homestayBusinessId,proto3" json:"homestayBusinessId,omitempty"`    //店铺id
	UserId              int64  `protobuf:"varint,8,opt,name=userId,proto3" json:"userId,omitempty"`                            //房东id
	RowState            int64  `protobuf:"varint,9,opt,name=rowState,proto3" json:"rowState,omitempty"`                        //0:下架，1：上架
	RowType             int64  `protobuf:"varint,10,opt,name=rowType,proto3" json:"rowType,omitempty"`                         //售卖类型：0：按房间预定，1：按人数预定
	FoodInfo            string `protobuf:"bytes,11,opt,name=foodInfo,proto3" json:"foodInfo,omitempty"`                        //餐食标准
	FoodPrice           int64  `protobuf:"varint,12,opt,name=foodPrice,proto3" json:"foodPrice,omitempty"`                     //餐食价格（份）
	HomestayPrice       int64  `protobuf:"varint,13,opt,name=homestayPrice,proto3" json:"homestayPrice,omitempty"`             //民宿价格
	MarketHomestayPrice int64  `protobuf:"varint,14,opt,name=marketHomestayPrice,proto3" json:"marketHomestayPrice,omitempty"` //民宿市场价格
}

func (x *Homestay) Reset() {
	*x = Homestay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_travel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Homestay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Homestay) ProtoMessage() {}

func (x *Homestay) ProtoReflect() protoreflect.Message {
	mi := &file_doc_travel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Homestay.ProtoReflect.Descriptor instead.
func (*Homestay) Descriptor() ([]byte, []int) {
	return file_doc_travel_proto_rawDescGZIP(), []int{0}
}

func (x *Homestay) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Homestay) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Homestay) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *Homestay) GetBanner() string {
	if x != nil {
		return x.Banner
	}
	return ""
}

func (x *Homestay) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *Homestay) GetPeopleNum() int64 {
	if x != nil {
		return x.PeopleNum
	}
	return 0
}

func (x *Homestay) GetHomestayBusinessId() int64 {
	if x != nil {
		return x.HomestayBusinessId
	}
	return 0
}

func (x *Homestay) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Homestay) GetRowState() int64 {
	if x != nil {
		return x.RowState
	}
	return 0
}

func (x *Homestay) GetRowType() int64 {
	if x != nil {
		return x.RowType
	}
	return 0
}

func (x *Homestay) GetFoodInfo() string {
	if x != nil {
		return x.FoodInfo
	}
	return ""
}

func (x *Homestay) GetFoodPrice() int64 {
	if x != nil {
		return x.FoodPrice
	}
	return 0
}

func (x *Homestay) GetHomestayPrice() int64 {
	if x != nil {
		return x.HomestayPrice
	}
	return 0
}

func (x *Homestay) GetMarketHomestayPrice() int64 {
	if x != nil {
		return x.MarketHomestayPrice
	}
	return 0
}

type HomestayDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *HomestayDetailReq) Reset() {
	*x = HomestayDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_travel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomestayDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomestayDetailReq) ProtoMessage() {}

func (x *HomestayDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_doc_travel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomestayDetailReq.ProtoReflect.Descriptor instead.
func (*HomestayDetailReq) Descriptor() ([]byte, []int) {
	return file_doc_travel_proto_rawDescGZIP(), []int{1}
}

func (x *HomestayDetailReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type HomestayDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Homestay *Homestay `protobuf:"bytes,1,opt,name=homestay,proto3" json:"homestay,omitempty"`
}

func (x *HomestayDetailResp) Reset() {
	*x = HomestayDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_travel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomestayDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomestayDetailResp) ProtoMessage() {}

func (x *HomestayDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_doc_travel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomestayDetailResp.ProtoReflect.Descriptor instead.
func (*HomestayDetailResp) Descriptor() ([]byte, []int) {
	return file_doc_travel_proto_rawDescGZIP(), []int{2}
}

func (x *HomestayDetailResp) GetHomestay() *Homestay {
	if x != nil {
		return x.Homestay
	}
	return nil
}

var File_doc_travel_proto protoreflect.FileDescriptor

var file_doc_travel_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x6f, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa6, 0x03, 0x0a, 0x08, 0x48, 0x6f, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x12,
	0x2e, 0x0a, 0x12, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x68, 0x6f, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x79, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x77, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x6f, 0x77, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6f,
	0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x6f,
	0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x13, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x23, 0x0a, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x08, 0x68, 0x6f,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x79, 0x32, 0x49, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x12, 0x3f,
	0x0a, 0x0e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x64, 0x6f, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_doc_travel_proto_rawDescOnce sync.Once
	file_doc_travel_proto_rawDescData = file_doc_travel_proto_rawDesc
)

func file_doc_travel_proto_rawDescGZIP() []byte {
	file_doc_travel_proto_rawDescOnce.Do(func() {
		file_doc_travel_proto_rawDescData = protoimpl.X.CompressGZIP(file_doc_travel_proto_rawDescData)
	})
	return file_doc_travel_proto_rawDescData
}

var file_doc_travel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_doc_travel_proto_goTypes = []interface{}{
	(*Homestay)(nil),           // 0: pb.Homestay
	(*HomestayDetailReq)(nil),  // 1: pb.HomestayDetailReq
	(*HomestayDetailResp)(nil), // 2: pb.HomestayDetailResp
}
var file_doc_travel_proto_depIdxs = []int32{
	0, // 0: pb.HomestayDetailResp.homestay:type_name -> pb.Homestay
	1, // 1: pb.travel.homestayDetail:input_type -> pb.HomestayDetailReq
	2, // 2: pb.travel.homestayDetail:output_type -> pb.HomestayDetailResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_doc_travel_proto_init() }
func file_doc_travel_proto_init() {
	if File_doc_travel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doc_travel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Homestay); i {
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
		file_doc_travel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomestayDetailReq); i {
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
		file_doc_travel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomestayDetailResp); i {
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
			RawDescriptor: file_doc_travel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doc_travel_proto_goTypes,
		DependencyIndexes: file_doc_travel_proto_depIdxs,
		MessageInfos:      file_doc_travel_proto_msgTypes,
	}.Build()
	File_doc_travel_proto = out.File
	file_doc_travel_proto_rawDesc = nil
	file_doc_travel_proto_goTypes = nil
	file_doc_travel_proto_depIdxs = nil
}
