// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: proto/user/v1/user.proto

package v1

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

type UserMobileLoginResponseUserGender int32

const (
	//    option allow_alias = true;
	UserMobileLoginResponse_UNKNOWN UserMobileLoginResponseUserGender = 0
	UserMobileLoginResponse_MAN     UserMobileLoginResponseUserGender = 1
	UserMobileLoginResponse_FEMALE  UserMobileLoginResponseUserGender = 2
)

// Enum value maps for UserMobileLoginResponseUserGender.
var (
	UserMobileLoginResponseUserGender_name = map[int32]string{
		0: "UNKNOWN",
		1: "MAN",
		2: "FEMALE",
	}
	UserMobileLoginResponseUserGender_value = map[string]int32{
		"UNKNOWN": 0,
		"MAN":     1,
		"FEMALE":  2,
	}
)

func (x UserMobileLoginResponseUserGender) Enum() *UserMobileLoginResponseUserGender {
	p := new(UserMobileLoginResponseUserGender)
	*p = x
	return p
}

func (x UserMobileLoginResponseUserGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserMobileLoginResponseUserGender) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_v1_user_proto_enumTypes[0].Descriptor()
}

func (UserMobileLoginResponseUserGender) Type() protoreflect.EnumType {
	return &file_proto_user_v1_user_proto_enumTypes[0]
}

func (x UserMobileLoginResponseUserGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserMobileLoginResponseUserGender.Descriptor instead.
func (UserMobileLoginResponseUserGender) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_v1_user_proto_rawDescGZIP(), []int{1, 0}
}

//用户登录
type UserMobileLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` //运营商token。SDK返回的运营商TOKEN，有效期：移动2分钟、电信10分钟、联通30分钟，一次有效。
}

func (x *UserMobileLoginRequest) Reset() {
	*x = UserMobileLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMobileLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMobileLoginRequest) ProtoMessage() {}

func (x *UserMobileLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMobileLoginRequest.ProtoReflect.Descriptor instead.
func (*UserMobileLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserMobileLoginRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

//用户登录响应
type UserMobileLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientToken  string                            `protobuf:"bytes,1,opt,name=clientToken,proto3" json:"clientToken,omitempty"`                                               //玩家的登录凭据
	ExpireAt     int32                             `protobuf:"varint,2,opt,name=expireAt,proto3" json:"expireAt,omitempty"`                                                    //凭据的过期时间
	UserName     string                            `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`                                                     //玩家账号
	Nickname     string                            `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`                                                     //昵称
	Gender       UserMobileLoginResponseUserGender `protobuf:"varint,5,opt,name=gender,proto3,enum=user.service.v1.UserMobileLoginResponseUserGender" json:"gender,omitempty"` //性别
	Avatar       string                            `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`                                                         //头像
	IsFirstLogin int32                             `protobuf:"zigzag32,7,opt,name=isFirstLogin,proto3" json:"isFirstLogin,omitempty"`                                          //是否为首次登陆
}

func (x *UserMobileLoginResponse) Reset() {
	*x = UserMobileLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMobileLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMobileLoginResponse) ProtoMessage() {}

func (x *UserMobileLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMobileLoginResponse.ProtoReflect.Descriptor instead.
func (*UserMobileLoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserMobileLoginResponse) GetClientToken() string {
	if x != nil {
		return x.ClientToken
	}
	return ""
}

func (x *UserMobileLoginResponse) GetExpireAt() int32 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *UserMobileLoginResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserMobileLoginResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserMobileLoginResponse) GetGender() UserMobileLoginResponseUserGender {
	if x != nil {
		return x.Gender
	}
	return UserMobileLoginResponse_UNKNOWN
}

func (x *UserMobileLoginResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserMobileLoginResponse) GetIsFirstLogin() int32 {
	if x != nil {
		return x.IsFirstLogin
	}
	return 0
}

var File_proto_user_v1_user_proto protoreflect.FileDescriptor

var file_proto_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x2e, 0x0a, 0x16, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc8, 0x02, 0x0a, 0x17,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0c, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x2e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x18, 0x5a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_v1_user_proto_rawDescOnce sync.Once
	file_proto_user_v1_user_proto_rawDescData = file_proto_user_v1_user_proto_rawDesc
)

func file_proto_user_v1_user_proto_rawDescGZIP() []byte {
	file_proto_user_v1_user_proto_rawDescOnce.Do(func() {
		file_proto_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_v1_user_proto_rawDescData)
	})
	return file_proto_user_v1_user_proto_rawDescData
}

var file_proto_user_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_user_v1_user_proto_goTypes = []interface{}{
	(UserMobileLoginResponseUserGender)(0), // 0: user.service.v1.UserMobileLoginResponse.userGender
	(*UserMobileLoginRequest)(nil),         // 1: user.service.v1.UserMobileLoginRequest
	(*UserMobileLoginResponse)(nil),        // 2: user.service.v1.UserMobileLoginResponse
}
var file_proto_user_v1_user_proto_depIdxs = []int32{
	0, // 0: user.service.v1.UserMobileLoginResponse.gender:type_name -> user.service.v1.UserMobileLoginResponse.userGender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_user_v1_user_proto_init() }
func file_proto_user_v1_user_proto_init() {
	if File_proto_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMobileLoginRequest); i {
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
		file_proto_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMobileLoginResponse); i {
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
			RawDescriptor: file_proto_user_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_user_v1_user_proto_goTypes,
		DependencyIndexes: file_proto_user_v1_user_proto_depIdxs,
		EnumInfos:         file_proto_user_v1_user_proto_enumTypes,
		MessageInfos:      file_proto_user_v1_user_proto_msgTypes,
	}.Build()
	File_proto_user_v1_user_proto = out.File
	file_proto_user_v1_user_proto_rawDesc = nil
	file_proto_user_v1_user_proto_goTypes = nil
	file_proto_user_v1_user_proto_depIdxs = nil
}
