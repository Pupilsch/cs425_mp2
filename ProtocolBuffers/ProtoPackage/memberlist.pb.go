// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: memberlist.proto

package ProtoPackage

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MessageType int32

const (
	MessageType_STANDARD MessageType = 0
	MessageType_JOINREQ  MessageType = 1
	MessageType_JOINREP  MessageType = 2
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "STANDARD",
		1: "JOINREQ",
		2: "JOINREP",
	}
	MessageType_value = map[string]int32{
		"STANDARD": 0,
		"JOINREQ":  1,
		"JOINREP":  2,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_memberlist_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_memberlist_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_memberlist_proto_rawDescGZIP(), []int{0}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeartbeatCounter int32                `protobuf:"varint,1,opt,name=HeartbeatCounter,proto3" json:"HeartbeatCounter,omitempty"`
	LastSeen         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=LastSeen,proto3" json:"LastSeen,omitempty"`
	IsLeaving        bool                 `protobuf:"varint,3,opt,name=IsLeaving,proto3" json:"IsLeaving,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_memberlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_memberlist_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetHeartbeatCounter() int32 {
	if x != nil {
		return x.HeartbeatCounter
	}
	return 0
}

func (x *Member) GetLastSeen() *timestamp.Timestamp {
	if x != nil {
		return x.LastSeen
	}
	return nil
}

func (x *Member) GetIsLeaving() bool {
	if x != nil {
		return x.IsLeaving
	}
	return false
}

type MembershipServiceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberList      map[string]*Member `protobuf:"bytes,1,rep,name=MemberList,proto3" json:"MemberList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Strategy        string             `protobuf:"bytes,2,opt,name=Strategy,proto3" json:"Strategy,omitempty"`
	StrategyCounter int32              `protobuf:"varint,3,opt,name=StrategyCounter,proto3" json:"StrategyCounter,omitempty"`
	Type            MessageType        `protobuf:"varint,4,opt,name=Type,proto3,enum=tutorial.MessageType" json:"Type,omitempty"`
}

func (x *MembershipServiceMessage) Reset() {
	*x = MembershipServiceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembershipServiceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipServiceMessage) ProtoMessage() {}

func (x *MembershipServiceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_memberlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipServiceMessage.ProtoReflect.Descriptor instead.
func (*MembershipServiceMessage) Descriptor() ([]byte, []int) {
	return file_memberlist_proto_rawDescGZIP(), []int{1}
}

func (x *MembershipServiceMessage) GetMemberList() map[string]*Member {
	if x != nil {
		return x.MemberList
	}
	return nil
}

func (x *MembershipServiceMessage) GetStrategy() string {
	if x != nil {
		return x.Strategy
	}
	return ""
}

func (x *MembershipServiceMessage) GetStrategyCounter() int32 {
	if x != nil {
		return x.StrategyCounter
	}
	return 0
}

func (x *MembershipServiceMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_STANDARD
}

var File_memberlist_proto protoreflect.FileDescriptor

var file_memberlist_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01,
	0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x49, 0x73, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x49, 0x73, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x22, 0xb0, 0x02, 0x0a, 0x18, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x52, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x4f, 0x0a, 0x0f,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x35, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4a, 0x4f,
	0x49, 0x4e, 0x52, 0x45, 0x51, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4a, 0x4f, 0x49, 0x4e, 0x52,
	0x45, 0x50, 0x10, 0x02, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memberlist_proto_rawDescOnce sync.Once
	file_memberlist_proto_rawDescData = file_memberlist_proto_rawDesc
)

func file_memberlist_proto_rawDescGZIP() []byte {
	file_memberlist_proto_rawDescOnce.Do(func() {
		file_memberlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_memberlist_proto_rawDescData)
	})
	return file_memberlist_proto_rawDescData
}

var file_memberlist_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_memberlist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_memberlist_proto_goTypes = []interface{}{
	(MessageType)(0),                 // 0: tutorial.MessageType
	(*Member)(nil),                   // 1: tutorial.Member
	(*MembershipServiceMessage)(nil), // 2: tutorial.MembershipServiceMessage
	nil,                              // 3: tutorial.MembershipServiceMessage.MemberListEntry
	(*timestamp.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_memberlist_proto_depIdxs = []int32{
	4, // 0: tutorial.Member.LastSeen:type_name -> google.protobuf.Timestamp
	3, // 1: tutorial.MembershipServiceMessage.MemberList:type_name -> tutorial.MembershipServiceMessage.MemberListEntry
	0, // 2: tutorial.MembershipServiceMessage.Type:type_name -> tutorial.MessageType
	1, // 3: tutorial.MembershipServiceMessage.MemberListEntry.value:type_name -> tutorial.Member
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_memberlist_proto_init() }
func file_memberlist_proto_init() {
	if File_memberlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memberlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_memberlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembershipServiceMessage); i {
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
			RawDescriptor: file_memberlist_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_memberlist_proto_goTypes,
		DependencyIndexes: file_memberlist_proto_depIdxs,
		EnumInfos:         file_memberlist_proto_enumTypes,
		MessageInfos:      file_memberlist_proto_msgTypes,
	}.Build()
	File_memberlist_proto = out.File
	file_memberlist_proto_rawDesc = nil
	file_memberlist_proto_goTypes = nil
	file_memberlist_proto_depIdxs = nil
}
