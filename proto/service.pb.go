// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package addservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ReadyMessage struct {
	Ready                int32    `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadyMessage) Reset()         { *m = ReadyMessage{} }
func (m *ReadyMessage) String() string { return proto.CompactTextString(m) }
func (*ReadyMessage) ProtoMessage()    {}
func (*ReadyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ReadyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadyMessage.Unmarshal(m, b)
}
func (m *ReadyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadyMessage.Marshal(b, m, deterministic)
}
func (m *ReadyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadyMessage.Merge(m, src)
}
func (m *ReadyMessage) XXX_Size() int {
	return xxx_messageInfo_ReadyMessage.Size(m)
}
func (m *ReadyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ReadyMessage proto.InternalMessageInfo

func (m *ReadyMessage) GetReady() int32 {
	if m != nil {
		return m.Ready
	}
	return 0
}

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ListQueuedMessage struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQueuedMessage) Reset()         { *m = ListQueuedMessage{} }
func (m *ListQueuedMessage) String() string { return proto.CompactTextString(m) }
func (*ListQueuedMessage) ProtoMessage()    {}
func (*ListQueuedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *ListQueuedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQueuedMessage.Unmarshal(m, b)
}
func (m *ListQueuedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQueuedMessage.Marshal(b, m, deterministic)
}
func (m *ListQueuedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQueuedMessage.Merge(m, src)
}
func (m *ListQueuedMessage) XXX_Size() int {
	return xxx_messageInfo_ListQueuedMessage.Size(m)
}
func (m *ListQueuedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQueuedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ListQueuedMessage proto.InternalMessageInfo

func (m *ListQueuedMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ReceivedQueuedMessageList struct {
	List                 map[string]bool `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReceivedQueuedMessageList) Reset()         { *m = ReceivedQueuedMessageList{} }
func (m *ReceivedQueuedMessageList) String() string { return proto.CompactTextString(m) }
func (*ReceivedQueuedMessageList) ProtoMessage()    {}
func (*ReceivedQueuedMessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ReceivedQueuedMessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedQueuedMessageList.Unmarshal(m, b)
}
func (m *ReceivedQueuedMessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedQueuedMessageList.Marshal(b, m, deterministic)
}
func (m *ReceivedQueuedMessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedQueuedMessageList.Merge(m, src)
}
func (m *ReceivedQueuedMessageList) XXX_Size() int {
	return xxx_messageInfo_ReceivedQueuedMessageList.Size(m)
}
func (m *ReceivedQueuedMessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedQueuedMessageList.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedQueuedMessageList proto.InternalMessageInfo

func (m *ReceivedQueuedMessageList) GetList() map[string]bool {
	if m != nil {
		return m.List
	}
	return nil
}

// Publisher --> Broker
type TopicMessage struct {
	SentOn               string   `protobuf:"bytes,1,opt,name=SentOn,proto3" json:"SentOn,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicMessage) Reset()         { *m = TopicMessage{} }
func (m *TopicMessage) String() string { return proto.CompactTextString(m) }
func (*TopicMessage) ProtoMessage()    {}
func (*TopicMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *TopicMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicMessage.Unmarshal(m, b)
}
func (m *TopicMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicMessage.Marshal(b, m, deterministic)
}
func (m *TopicMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMessage.Merge(m, src)
}
func (m *TopicMessage) XXX_Size() int {
	return xxx_messageInfo_TopicMessage.Size(m)
}
func (m *TopicMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMessage proto.InternalMessageInfo

func (m *TopicMessage) GetSentOn() string {
	if m != nil {
		return m.SentOn
	}
	return ""
}

func (m *TopicMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// Acking function: Broker --> Publisher
type QueuedMessage struct {
	SentOn               string   `protobuf:"bytes,1,opt,name=SentOn,proto3" json:"SentOn,omitempty"`
	Ack                  bool     `protobuf:"varint,2,opt,name=ack,proto3" json:"ack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueuedMessage) Reset()         { *m = QueuedMessage{} }
func (m *QueuedMessage) String() string { return proto.CompactTextString(m) }
func (*QueuedMessage) ProtoMessage()    {}
func (*QueuedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *QueuedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueuedMessage.Unmarshal(m, b)
}
func (m *QueuedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueuedMessage.Marshal(b, m, deterministic)
}
func (m *QueuedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueuedMessage.Merge(m, src)
}
func (m *QueuedMessage) XXX_Size() int {
	return xxx_messageInfo_QueuedMessage.Size(m)
}
func (m *QueuedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QueuedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QueuedMessage proto.InternalMessageInfo

func (m *QueuedMessage) GetSentOn() string {
	if m != nil {
		return m.SentOn
	}
	return ""
}

func (m *QueuedMessage) GetAck() bool {
	if m != nil {
		return m.Ack
	}
	return false
}

// Subscriber --> Broker
type TopicSubscription struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicSubscription) Reset()         { *m = TopicSubscription{} }
func (m *TopicSubscription) String() string { return proto.CompactTextString(m) }
func (*TopicSubscription) ProtoMessage()    {}
func (*TopicSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *TopicSubscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicSubscription.Unmarshal(m, b)
}
func (m *TopicSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicSubscription.Marshal(b, m, deterministic)
}
func (m *TopicSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicSubscription.Merge(m, src)
}
func (m *TopicSubscription) XXX_Size() int {
	return xxx_messageInfo_TopicSubscription.Size(m)
}
func (m *TopicSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_TopicSubscription proto.InternalMessageInfo

func (m *TopicSubscription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TopicSubscription) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

// Broker --> Subscriber
type SubscriberDeliveredMessage struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Ack                  bool     `protobuf:"varint,2,opt,name=ack,proto3" json:"ack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriberDeliveredMessage) Reset()         { *m = SubscriberDeliveredMessage{} }
func (m *SubscriberDeliveredMessage) String() string { return proto.CompactTextString(m) }
func (*SubscriberDeliveredMessage) ProtoMessage()    {}
func (*SubscriberDeliveredMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *SubscriberDeliveredMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberDeliveredMessage.Unmarshal(m, b)
}
func (m *SubscriberDeliveredMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberDeliveredMessage.Marshal(b, m, deterministic)
}
func (m *SubscriberDeliveredMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberDeliveredMessage.Merge(m, src)
}
func (m *SubscriberDeliveredMessage) XXX_Size() int {
	return xxx_messageInfo_SubscriberDeliveredMessage.Size(m)
}
func (m *SubscriberDeliveredMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberDeliveredMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberDeliveredMessage proto.InternalMessageInfo

func (m *SubscriberDeliveredMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *SubscriberDeliveredMessage) GetAck() bool {
	if m != nil {
		return m.Ack
	}
	return false
}

func init() {
	proto.RegisterType((*ReadyMessage)(nil), "addservice.ReadyMessage")
	proto.RegisterType((*Message)(nil), "addservice.Message")
	proto.RegisterType((*ListQueuedMessage)(nil), "addservice.ListQueuedMessage")
	proto.RegisterType((*ReceivedQueuedMessageList)(nil), "addservice.ReceivedQueuedMessageList")
	proto.RegisterMapType((map[string]bool)(nil), "addservice.ReceivedQueuedMessageList.ListEntry")
	proto.RegisterType((*TopicMessage)(nil), "addservice.TopicMessage")
	proto.RegisterType((*QueuedMessage)(nil), "addservice.QueuedMessage")
	proto.RegisterType((*TopicSubscription)(nil), "addservice.TopicSubscription")
	proto.RegisterType((*SubscriberDeliveredMessage)(nil), "addservice.SubscriberDeliveredMessage")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0x4a, 0xf3, 0x40,
	0x10, 0x26, 0x49, 0xdb, 0xff, 0xef, 0xd8, 0x82, 0x5d, 0x45, 0xda, 0x40, 0x41, 0x82, 0x87, 0x5e,
	0x45, 0xa8, 0xe0, 0xa1, 0xe0, 0x85, 0xa7, 0xbb, 0x8a, 0x9a, 0xfa, 0x02, 0x49, 0x76, 0x90, 0xd0,
	0x9a, 0x94, 0xcd, 0xa6, 0xd8, 0x47, 0x11, 0x7c, 0x58, 0xb3, 0x9b, 0x6d, 0xdd, 0x50, 0x1a, 0xbc,
	0x49, 0xe6, 0xf4, 0xcd, 0x7c, 0xdf, 0x64, 0x02, 0xed, 0x14, 0xd9, 0x22, 0x0a, 0xd1, 0x9d, 0xb3,
	0x84, 0x27, 0x04, 0x7c, 0x4a, 0x55, 0xc4, 0x39, 0x82, 0x96, 0x87, 0x3e, 0x5d, 0x3e, 0x61, 0x9a,
	0xfa, 0xef, 0x48, 0xf6, 0xa1, 0xce, 0x84, 0xdf, 0x35, 0x0e, 0x8d, 0x41, 0xdd, 0x2b, 0x1c, 0xa7,
	0x0f, 0xff, 0x56, 0x05, 0x04, 0x6a, 0x1c, 0x3f, 0xb9, 0xcc, 0x37, 0x3d, 0x69, 0x3b, 0xa7, 0xd0,
	0x19, 0x47, 0x29, 0x7f, 0xcd, 0x30, 0x43, 0x5a, 0x55, 0xf8, 0x65, 0x40, 0xcf, 0xc3, 0x10, 0xa3,
	0x05, 0xd2, 0x52, 0xb5, 0x80, 0x93, 0x7b, 0xa8, 0xcd, 0xf2, 0x77, 0x8e, 0xb0, 0x06, 0x3b, 0xc3,
	0x33, 0xf7, 0x97, 0xa6, 0xbb, 0x15, 0xe4, 0x8a, 0xc7, 0x63, 0xcc, 0xd9, 0xd2, 0x93, 0x60, 0xfb,
	0x12, 0x9a, 0xeb, 0x10, 0xd9, 0x05, 0x6b, 0x8a, 0x4b, 0x45, 0x41, 0x98, 0x42, 0xdf, 0xc2, 0x9f,
	0x65, 0xd8, 0x35, 0xf3, 0xd8, 0x7f, 0xaf, 0x70, 0x46, 0xe6, 0x95, 0xe1, 0x8c, 0xa0, 0xf5, 0x96,
	0xcc, 0xa3, 0x70, 0xc5, 0xff, 0x00, 0x1a, 0x13, 0x8c, 0xf9, 0x73, 0xac, 0xe0, 0xca, 0x5b, 0xeb,
	0x32, 0x35, 0x5d, 0xd7, 0xd0, 0x2e, 0x8b, 0xdf, 0x06, 0xce, 0x09, 0xf9, 0xe1, 0x54, 0x0d, 0x17,
	0xa6, 0x73, 0x03, 0x1d, 0x39, 0x76, 0x92, 0x05, 0x69, 0xc8, 0xa2, 0x39, 0x8f, 0x12, 0x39, 0x23,
	0xf6, 0x3f, 0x70, 0x35, 0x43, 0xd8, 0x82, 0x39, 0x17, 0x85, 0xaa, 0x63, 0xe1, 0x38, 0x77, 0x60,
	0x2b, 0x64, 0x80, 0xec, 0x01, 0x67, 0xf9, 0x92, 0x58, 0xe5, 0x37, 0xd8, 0xa4, 0x30, 0xfc, 0x36,
	0x01, 0x6e, 0x29, 0x9d, 0x14, 0xbb, 0x26, 0x23, 0xb0, 0x5e, 0xb2, 0x80, 0x74, 0xf5, 0xfd, 0xeb,
	0x9b, 0xb1, 0x7b, 0x7a, 0xa6, 0xac, 0x7b, 0x0c, 0x56, 0x4e, 0x87, 0xf4, 0x37, 0xb0, 0xba, 0x3c,
	0xfb, 0x44, 0x4f, 0x57, 0xd0, 0x1f, 0x43, 0x4d, 0x1e, 0x46, 0xa9, 0xdd, 0xc6, 0xa5, 0xd9, 0xc7,
	0x7f, 0xba, 0x14, 0x72, 0x01, 0x75, 0x21, 0x11, 0xcb, 0xca, 0xf4, 0xeb, 0xb7, 0xf7, 0xf4, 0x8c,
	0x0a, 0x06, 0x0d, 0xf9, 0xd7, 0x9c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x2c, 0xf0, 0xb0,
	0x46, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	Pub(ctx context.Context, in *TopicMessage, opts ...grpc.CallOption) (*QueuedMessage, error)
	Sub(ctx context.Context, in *TopicSubscription, opts ...grpc.CallOption) (*SubscriberDeliveredMessage, error)
	List(ctx context.Context, in *ListQueuedMessage, opts ...grpc.CallOption) (*ReceivedQueuedMessageList, error)
	Serve(ctx context.Context, in *ReadyMessage, opts ...grpc.CallOption) (*Message, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Pub(ctx context.Context, in *TopicMessage, opts ...grpc.CallOption) (*QueuedMessage, error) {
	out := new(QueuedMessage)
	err := c.cc.Invoke(ctx, "/addservice.AddService/Pub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Sub(ctx context.Context, in *TopicSubscription, opts ...grpc.CallOption) (*SubscriberDeliveredMessage, error) {
	out := new(SubscriberDeliveredMessage)
	err := c.cc.Invoke(ctx, "/addservice.AddService/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) List(ctx context.Context, in *ListQueuedMessage, opts ...grpc.CallOption) (*ReceivedQueuedMessageList, error) {
	out := new(ReceivedQueuedMessageList)
	err := c.cc.Invoke(ctx, "/addservice.AddService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Serve(ctx context.Context, in *ReadyMessage, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/addservice.AddService/Serve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	Pub(context.Context, *TopicMessage) (*QueuedMessage, error)
	Sub(context.Context, *TopicSubscription) (*SubscriberDeliveredMessage, error)
	List(context.Context, *ListQueuedMessage) (*ReceivedQueuedMessageList, error)
	Serve(context.Context, *ReadyMessage) (*Message, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) Pub(ctx context.Context, req *TopicMessage) (*QueuedMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pub not implemented")
}
func (*UnimplementedAddServiceServer) Sub(ctx context.Context, req *TopicSubscription) (*SubscriberDeliveredMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (*UnimplementedAddServiceServer) List(ctx context.Context, req *ListQueuedMessage) (*ReceivedQueuedMessageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAddServiceServer) Serve(ctx context.Context, req *ReadyMessage) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Serve not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Pub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Pub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addservice.AddService/Pub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Pub(ctx, req.(*TopicMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addservice.AddService/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Sub(ctx, req.(*TopicSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQueuedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addservice.AddService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).List(ctx, req.(*ListQueuedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Serve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Serve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addservice.AddService/Serve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Serve(ctx, req.(*ReadyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addservice.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pub",
			Handler:    _AddService_Pub_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _AddService_Sub_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AddService_List_Handler,
		},
		{
			MethodName: "Serve",
			Handler:    _AddService_Serve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
