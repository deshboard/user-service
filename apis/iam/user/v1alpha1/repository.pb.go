// Code generated by protoc-gen-go.
// source: iam/user/v1alpha1/repository.proto
// DO NOT EDIT!

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NewUser struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *NewUser) Reset()                    { *m = NewUser{} }
func (m *NewUser) String() string            { return proto.CompactTextString(m) }
func (*NewUser) ProtoMessage()               {}
func (*NewUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *NewUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NewUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NewUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserCreated struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserCreated) Reset()                    { *m = UserCreated{} }
func (m *UserCreated) String() string            { return proto.CompactTextString(m) }
func (*UserCreated) ProtoMessage()               {}
func (*UserCreated) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *UserCreated) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUser struct {
	// Types that are valid to be assigned to UserKey:
	//	*GetUser_Id
	//	*GetUser_Username
	//	*GetUser_Email
	UserKey isGetUser_UserKey `protobuf_oneof:"user_key"`
}

func (m *GetUser) Reset()                    { *m = GetUser{} }
func (m *GetUser) String() string            { return proto.CompactTextString(m) }
func (*GetUser) ProtoMessage()               {}
func (*GetUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type isGetUser_UserKey interface {
	isGetUser_UserKey()
}

type GetUser_Id struct {
	Id int64 `protobuf:"varint,1,opt,name=id,oneof"`
}
type GetUser_Username struct {
	Username string `protobuf:"bytes,2,opt,name=username,oneof"`
}
type GetUser_Email struct {
	Email string `protobuf:"bytes,3,opt,name=email,oneof"`
}

func (*GetUser_Id) isGetUser_UserKey()       {}
func (*GetUser_Username) isGetUser_UserKey() {}
func (*GetUser_Email) isGetUser_UserKey()    {}

func (m *GetUser) GetUserKey() isGetUser_UserKey {
	if m != nil {
		return m.UserKey
	}
	return nil
}

func (m *GetUser) GetId() int64 {
	if x, ok := m.GetUserKey().(*GetUser_Id); ok {
		return x.Id
	}
	return 0
}

func (m *GetUser) GetUsername() string {
	if x, ok := m.GetUserKey().(*GetUser_Username); ok {
		return x.Username
	}
	return ""
}

func (m *GetUser) GetEmail() string {
	if x, ok := m.GetUserKey().(*GetUser_Email); ok {
		return x.Email
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GetUser) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GetUser_OneofMarshaler, _GetUser_OneofUnmarshaler, _GetUser_OneofSizer, []interface{}{
		(*GetUser_Id)(nil),
		(*GetUser_Username)(nil),
		(*GetUser_Email)(nil),
	}
}

func _GetUser_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GetUser)
	// user_key
	switch x := m.UserKey.(type) {
	case *GetUser_Id:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Id))
	case *GetUser_Username:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Username)
	case *GetUser_Email:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Email)
	case nil:
	default:
		return fmt.Errorf("GetUser.UserKey has unexpected type %T", x)
	}
	return nil
}

func _GetUser_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GetUser)
	switch tag {
	case 1: // user_key.id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.UserKey = &GetUser_Id{int64(x)}
		return true, err
	case 2: // user_key.username
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.UserKey = &GetUser_Username{x}
		return true, err
	case 3: // user_key.email
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.UserKey = &GetUser_Email{x}
		return true, err
	default:
		return false, nil
	}
}

func _GetUser_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GetUser)
	// user_key
	switch x := m.UserKey.(type) {
	case *GetUser_Id:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Id))
	case *GetUser_Username:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Username)))
		n += len(x.Username)
	case *GetUser_Email:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Email)))
		n += len(x.Email)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type User struct {
	Id                int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username          string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email             string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	EncryptedPassword string `protobuf:"bytes,4,opt,name=encrypted_password,json=encryptedPassword" json:"encrypted_password,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetEncryptedPassword() string {
	if m != nil {
		return m.EncryptedPassword
	}
	return ""
}

func init() {
	proto.RegisterType((*NewUser)(nil), "deshboard.iam.user.v1alpha1.NewUser")
	proto.RegisterType((*UserCreated)(nil), "deshboard.iam.user.v1alpha1.UserCreated")
	proto.RegisterType((*GetUser)(nil), "deshboard.iam.user.v1alpha1.GetUser")
	proto.RegisterType((*User)(nil), "deshboard.iam.user.v1alpha1.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserRepository service

type UserRepositoryClient interface {
	// Create a new user and return it's primary identifier
	Create(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*UserCreated, error)
	// Return a User based on some search criteria
	Get(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*User, error)
}

type userRepositoryClient struct {
	cc *grpc.ClientConn
}

func NewUserRepositoryClient(cc *grpc.ClientConn) UserRepositoryClient {
	return &userRepositoryClient{cc}
}

func (c *userRepositoryClient) Create(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*UserCreated, error) {
	out := new(UserCreated)
	err := grpc.Invoke(ctx, "/deshboard.iam.user.v1alpha1.UserRepository/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRepositoryClient) Get(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/deshboard.iam.user.v1alpha1.UserRepository/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserRepository service

type UserRepositoryServer interface {
	// Create a new user and return it's primary identifier
	Create(context.Context, *NewUser) (*UserCreated, error)
	// Return a User based on some search criteria
	Get(context.Context, *GetUser) (*User, error)
}

func RegisterUserRepositoryServer(s *grpc.Server, srv UserRepositoryServer) {
	s.RegisterService(&_UserRepository_serviceDesc, srv)
}

func _UserRepository_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRepositoryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deshboard.iam.user.v1alpha1.UserRepository/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRepositoryServer).Create(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRepository_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRepositoryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deshboard.iam.user.v1alpha1.UserRepository/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRepositoryServer).Get(ctx, req.(*GetUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRepository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deshboard.iam.user.v1alpha1.UserRepository",
	HandlerType: (*UserRepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserRepository_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserRepository_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam/user/v1alpha1/repository.proto",
}

func init() { proto.RegisterFile("iam/user/v1alpha1/repository.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0xdd, 0xda, 0xb9, 0xe9, 0x15, 0x86, 0x06, 0x91, 0x52, 0x15, 0xb4, 0xf8, 0xb0, 0x17, 0x53,
	0xa6, 0xff, 0x60, 0x3e, 0x6c, 0x4f, 0x43, 0x0a, 0xa2, 0x08, 0x32, 0x32, 0x73, 0x61, 0xc1, 0x75,
	0x29, 0x49, 0xb4, 0xf4, 0xaf, 0xf9, 0xeb, 0x24, 0xe9, 0x87, 0x55, 0x70, 0x7b, 0xbc, 0x39, 0xe7,
	0x9e, 0x73, 0xcf, 0xbd, 0x81, 0x48, 0xb0, 0x34, 0xfe, 0xd0, 0xa8, 0xe2, 0xcf, 0x31, 0x5b, 0x67,
	0x2b, 0x36, 0x8e, 0x15, 0x66, 0x52, 0x0b, 0x23, 0x55, 0x41, 0x33, 0x25, 0x8d, 0x24, 0x67, 0x1c,
	0xf5, 0x6a, 0x29, 0x99, 0xe2, 0x54, 0xb0, 0x94, 0x5a, 0x36, 0xad, 0xd9, 0xd1, 0x13, 0x0c, 0xe6,
	0x98, 0x3f, 0x6a, 0x54, 0x24, 0x84, 0x7d, 0x8b, 0x6d, 0x58, 0x8a, 0x41, 0xf7, 0xb2, 0x3b, 0x3a,
	0x48, 0x9a, 0x9a, 0x9c, 0xc0, 0x1e, 0xa6, 0x4c, 0xac, 0x03, 0xcf, 0x01, 0x65, 0x61, 0x3b, 0x32,
	0xa6, 0x75, 0x2e, 0x15, 0x0f, 0xfc, 0xb2, 0xa3, 0xae, 0xa3, 0x0b, 0x38, 0xb4, 0xaa, 0xf7, 0x0a,
	0x99, 0x41, 0x4e, 0x86, 0xe0, 0x09, 0xee, 0x64, 0xfd, 0xc4, 0x13, 0x3c, 0x7a, 0x85, 0xc1, 0x14,
	0x8d, 0xf3, 0x3d, 0xfa, 0x81, 0x66, 0x1d, 0x0b, 0x92, 0xf3, 0xd6, 0x24, 0xce, 0x70, 0xd6, 0x69,
	0xcd, 0x72, 0x5a, 0xcf, 0xe2, 0x57, 0x50, 0x59, 0x4e, 0xa0, 0xec, 0x5a, 0xbc, 0x63, 0x11, 0xe5,
	0xd0, 0x73, 0xda, 0x7f, 0x6c, 0x7f, 0x65, 0xf4, 0xfe, 0xcb, 0xe8, 0xb7, 0x33, 0xde, 0x00, 0xc1,
	0xcd, 0x9b, 0x2a, 0x32, 0x83, 0x7c, 0xd1, 0xa4, 0xed, 0x39, 0xca, 0x71, 0x83, 0x3c, 0x54, 0xc0,
	0xed, 0x57, 0x17, 0x86, 0xd6, 0x39, 0x69, 0xae, 0x40, 0x9e, 0xa1, 0x5f, 0x6e, 0x81, 0x5c, 0xd3,
	0x2d, 0xa7, 0xa0, 0xd5, 0x1d, 0xc2, 0xd1, 0x56, 0x56, 0x7b, 0xa9, 0x73, 0xf0, 0xa7, 0x68, 0x76,
	0xc8, 0x56, 0x6b, 0x0e, 0xaf, 0x76, 0xca, 0x4e, 0xfa, 0x2f, 0x3d, 0xfb, 0xba, 0xec, 0xbb, 0x8f,
	0x73, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x52, 0x7b, 0xa8, 0x1e, 0x5e, 0x02, 0x00, 0x00,
}
