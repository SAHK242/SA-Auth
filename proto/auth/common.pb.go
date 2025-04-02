// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.0
// source: auth/common.proto

package auth

import (
	gcommon "auth/proto/gcommon"
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

type EmployeeType int32

const (
	EmployeeType_EMPLOYEE_TYPE_UNSPECIFIED EmployeeType = 0
	EmployeeType_EMPLOYEE_TYPE_DOCTOR      EmployeeType = 1
	EmployeeType_EMPLOYEE_TYPE_NURSE       EmployeeType = 2
)

// Enum value maps for EmployeeType.
var (
	EmployeeType_name = map[int32]string{
		0: "EMPLOYEE_TYPE_UNSPECIFIED",
		1: "EMPLOYEE_TYPE_DOCTOR",
		2: "EMPLOYEE_TYPE_NURSE",
	}
	EmployeeType_value = map[string]int32{
		"EMPLOYEE_TYPE_UNSPECIFIED": 0,
		"EMPLOYEE_TYPE_DOCTOR":      1,
		"EMPLOYEE_TYPE_NURSE":       2,
	}
)

func (x EmployeeType) Enum() *EmployeeType {
	p := new(EmployeeType)
	*p = x
	return p
}

func (x EmployeeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmployeeType) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_common_proto_enumTypes[0].Descriptor()
}

func (EmployeeType) Type() protoreflect.EnumType {
	return &file_auth_common_proto_enumTypes[0]
}

func (x EmployeeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmployeeType.Descriptor instead.
func (EmployeeType) EnumDescriptor() ([]byte, []int) {
	return file_auth_common_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FirstName   string               `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName    string               `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	State       gcommon.AccountState `protobuf:"varint,4,opt,name=state,proto3,enum=common.AccountState" json:"state"`
	Username    string               `protobuf:"bytes,5,opt,name=username,proto3" json:"username"`
	Code        string               `protobuf:"bytes,6,opt,name=code,proto3" json:"code"`
	Gender      gcommon.Gender       `protobuf:"varint,7,opt,name=gender,proto3,enum=common.Gender" json:"gender"`
	DateOfBirth int64                `protobuf:"varint,8,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth"`
	Address     string               `protobuf:"bytes,9,opt,name=address,proto3" json:"address"`
	StartDate   int64                `protobuf:"varint,10,opt,name=start_date,json=startDate,proto3" json:"start_date"` // Onboard date
	EndDate     int64                `protobuf:"varint,11,opt,name=end_date,json=endDate,proto3" json:"end_date"`       // Off board date
	PhoneNumber string               `protobuf:"bytes,12,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number"`
	DegreeName  string               `protobuf:"bytes,13,opt,name=degree_name,json=degreeName,proto3" json:"degree_name"`
	DegreeYear  int32                `protobuf:"varint,14,opt,name=degree_year,json=degreeYear,proto3" json:"degree_year"`
	Department  *Department          `protobuf:"bytes,15,opt,name=department,proto3" json:"department"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_auth_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_common_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetState() gcommon.AccountState {
	if x != nil {
		return x.State
	}
	return gcommon.AccountState(0)
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *User) GetGender() gcommon.Gender {
	if x != nil {
		return x.Gender
	}
	return gcommon.Gender(0)
}

func (x *User) GetDateOfBirth() int64 {
	if x != nil {
		return x.DateOfBirth
	}
	return 0
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetStartDate() int64 {
	if x != nil {
		return x.StartDate
	}
	return 0
}

func (x *User) GetEndDate() int64 {
	if x != nil {
		return x.EndDate
	}
	return 0
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetDegreeName() string {
	if x != nil {
		return x.DegreeName
	}
	return ""
}

func (x *User) GetDegreeYear() int32 {
	if x != nil {
		return x.DegreeYear
	}
	return 0
}

func (x *User) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
}

func (x *Department) Reset() {
	*x = Department{}
	mi := &file_auth_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_auth_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_auth_common_proto_rawDescGZIP(), []int{1}
}

func (x *Department) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_auth_common_proto protoreflect.FileDescriptor

var file_auth_common_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe5, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x67, 0x72,
	0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65,
	0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x30, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x0a, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x60, 0x0a, 0x0c, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x45,
	0x4d, 0x50, 0x4c, 0x4f, 0x59, 0x45, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4d,
	0x50, 0x4c, 0x4f, 0x59, 0x45, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x43, 0x54,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4d, 0x50, 0x4c, 0x4f, 0x59, 0x45, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x52, 0x53, 0x45, 0x10, 0x02, 0x42, 0x0c, 0x5a,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_auth_common_proto_rawDescOnce sync.Once
	file_auth_common_proto_rawDescData = file_auth_common_proto_rawDesc
)

func file_auth_common_proto_rawDescGZIP() []byte {
	file_auth_common_proto_rawDescOnce.Do(func() {
		file_auth_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_common_proto_rawDescData)
	})
	return file_auth_common_proto_rawDescData
}

var file_auth_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_common_proto_goTypes = []any{
	(EmployeeType)(0),         // 0: auth.EmployeeType
	(*User)(nil),              // 1: auth.User
	(*Department)(nil),        // 2: auth.Department
	(gcommon.AccountState)(0), // 3: common.AccountState
	(gcommon.Gender)(0),       // 4: common.Gender
}
var file_auth_common_proto_depIdxs = []int32{
	3, // 0: auth.User.state:type_name -> common.AccountState
	4, // 1: auth.User.gender:type_name -> common.Gender
	2, // 2: auth.User.department:type_name -> auth.Department
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_auth_common_proto_init() }
func file_auth_common_proto_init() {
	if File_auth_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_common_proto_goTypes,
		DependencyIndexes: file_auth_common_proto_depIdxs,
		EnumInfos:         file_auth_common_proto_enumTypes,
		MessageInfos:      file_auth_common_proto_msgTypes,
	}.Build()
	File_auth_common_proto = out.File
	file_auth_common_proto_rawDesc = nil
	file_auth_common_proto_goTypes = nil
	file_auth_common_proto_depIdxs = nil
}
