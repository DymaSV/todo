// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: todo.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Done bool   `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Task) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type TaskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskList) Reset() {
	*x = TaskList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskList) ProtoMessage() {}

func (x *TaskList) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskList.ProtoReflect.Descriptor instead.
func (*TaskList) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TaskList) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f,
	0x64, 0x6f, 0x22, 0x2e, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f,
	0x6e, 0x65, 0x22, 0x2c, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x2d, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x24, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_todo_proto_goTypes = []interface{}{
	(*Task)(nil),     // 0: todo.Task
	(*TaskList)(nil), // 1: todo.TaskList
	(*Void)(nil),     // 2: todo.Void
}
var file_todo_proto_depIdxs = []int32{
	0, // 0: todo.TaskList.tasks:type_name -> todo.Task
	2, // 1: todo.Tasks.List:input_type -> todo.Void
	1, // 2: todo.Tasks.List:output_type -> todo.TaskList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskList); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TasksClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error)
}

type tasksClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksClient(cc grpc.ClientConnInterface) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error) {
	out := new(TaskList)
	err := c.cc.Invoke(ctx, "/todo.Tasks/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
type TasksServer interface {
	List(context.Context, *Void) (*TaskList, error)
}

// UnimplementedTasksServer can be embedded to have forward compatible implementations.
type UnimplementedTasksServer struct {
}

func (*UnimplementedTasksServer) List(context.Context, *Void) (*TaskList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Tasks/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Tasks_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
