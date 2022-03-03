// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: dbspgraph.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of this step.
type Step_Type int32

const (
	Step_INVALID Step_Type = 0
	// This step corresponds to the PRE step of the executor callback.
	Step_PRE Step_Type = 1
	// This step corresponds to the POST step of the executor callback.
	Step_POST Step_Type = 2
	// This step corresponds to the POST_KEEP_RUNNING step of the executor
	// callback.
	Step_POST_KEEP_RUNNING Step_Type = 3
	// This step indicates to the master that the worker has successfully
	// completed the graph execution and is ready to perist the computation
	// results.
	Step_EXECUTED_GRAPH Step_Type = 4
	// This step indicates to the master that the worker has successfully
	// persisted the computation results.
	Step_PESISTED_RESULTS Step_Type = 5
	// This step indicates to the master that the worker has completed the job.
	Step_COMPLETED_JOB Step_Type = 6
)

// Enum value maps for Step_Type.
var (
	Step_Type_name = map[int32]string{
		0: "INVALID",
		1: "PRE",
		2: "POST",
		3: "POST_KEEP_RUNNING",
		4: "EXECUTED_GRAPH",
		5: "PESISTED_RESULTS",
		6: "COMPLETED_JOB",
	}
	Step_Type_value = map[string]int32{
		"INVALID":           0,
		"PRE":               1,
		"POST":              2,
		"POST_KEEP_RUNNING": 3,
		"EXECUTED_GRAPH":    4,
		"PESISTED_RESULTS":  5,
		"COMPLETED_JOB":     6,
	}
)

func (x Step_Type) Enum() *Step_Type {
	p := new(Step_Type)
	*p = x
	return p
}

func (x Step_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Step_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_dbspgraph_proto_enumTypes[0].Descriptor()
}

func (Step_Type) Type() protoreflect.EnumType {
	return &file_dbspgraph_proto_enumTypes[0]
}

func (x Step_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Step_Type.Descriptor instead.
func (Step_Type) EnumDescriptor() ([]byte, []int) {
	return file_dbspgraph_proto_rawDescGZIP(), []int{3, 0}
}

// WorkerPayload encapsulates the possible message types that a worker can
// send to a master node.
type WorkerPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*WorkerPayload_Step
	//	*WorkerPayload_RelayMessage
	Payload isWorkerPayload_Payload `protobuf_oneof:"payload"`
}

func (x *WorkerPayload) Reset() {
	*x = WorkerPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbspgraph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerPayload) ProtoMessage() {}

func (x *WorkerPayload) ProtoReflect() protoreflect.Message {
	mi := &file_dbspgraph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerPayload.ProtoReflect.Descriptor instead.
func (*WorkerPayload) Descriptor() ([]byte, []int) {
	return file_dbspgraph_proto_rawDescGZIP(), []int{0}
}

func (m *WorkerPayload) GetPayload() isWorkerPayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *WorkerPayload) GetStep() *Step {
	if x, ok := x.GetPayload().(*WorkerPayload_Step); ok {
		return x.Step
	}
	return nil
}

func (x *WorkerPayload) GetRelayMessage() *RelayMessage {
	if x, ok := x.GetPayload().(*WorkerPayload_RelayMessage); ok {
		return x.RelayMessage
	}
	return nil
}

type isWorkerPayload_Payload interface {
	isWorkerPayload_Payload()
}

type WorkerPayload_Step struct {
	Step *Step `protobuf:"bytes,1,opt,name=step,proto3,oneof"`
}

type WorkerPayload_RelayMessage struct {
	RelayMessage *RelayMessage `protobuf:"bytes,2,opt,name=relay_message,json=relayMessage,proto3,oneof"`
}

func (*WorkerPayload_Step) isWorkerPayload_Payload() {}

func (*WorkerPayload_RelayMessage) isWorkerPayload_Payload() {}

// MasterPayload encapsulates the possible message types that a master can
// send to a worker node.
type MasterPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*MasterPayload_JobDetails
	//	*MasterPayload_Step
	//	*MasterPayload_RelayMessage
	Payload isMasterPayload_Payload `protobuf_oneof:"payload"`
}

func (x *MasterPayload) Reset() {
	*x = MasterPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbspgraph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterPayload) ProtoMessage() {}

func (x *MasterPayload) ProtoReflect() protoreflect.Message {
	mi := &file_dbspgraph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterPayload.ProtoReflect.Descriptor instead.
func (*MasterPayload) Descriptor() ([]byte, []int) {
	return file_dbspgraph_proto_rawDescGZIP(), []int{1}
}

func (m *MasterPayload) GetPayload() isMasterPayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *MasterPayload) GetJobDetails() *JobDetails {
	if x, ok := x.GetPayload().(*MasterPayload_JobDetails); ok {
		return x.JobDetails
	}
	return nil
}

func (x *MasterPayload) GetStep() *Step {
	if x, ok := x.GetPayload().(*MasterPayload_Step); ok {
		return x.Step
	}
	return nil
}

func (x *MasterPayload) GetRelayMessage() *RelayMessage {
	if x, ok := x.GetPayload().(*MasterPayload_RelayMessage); ok {
		return x.RelayMessage
	}
	return nil
}

type isMasterPayload_Payload interface {
	isMasterPayload_Payload()
}

type MasterPayload_JobDetails struct {
	JobDetails *JobDetails `protobuf:"bytes,1,opt,name=job_details,json=jobDetails,proto3,oneof"`
}

type MasterPayload_Step struct {
	Step *Step `protobuf:"bytes,2,opt,name=step,proto3,oneof"`
}

type MasterPayload_RelayMessage struct {
	RelayMessage *RelayMessage `protobuf:"bytes,3,opt,name=relay_message,json=relayMessage,proto3,oneof"`
}

func (*MasterPayload_JobDetails) isMasterPayload_Payload() {}

func (*MasterPayload_Step) isMasterPayload_Payload() {}

func (*MasterPayload_RelayMessage) isMasterPayload_Payload() {}

// JobDetails describes a job assigned by a master node to a worker.
type JobDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique ID for the job.
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// The creation time for the job.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The [from, to) UUID range assigned to the worker. Note that from is
	// inclusive and to is exclusive.
	PartitionFromUuid []byte `protobuf:"bytes,3,opt,name=partition_from_uuid,json=partitionFromUuid,proto3" json:"partition_from_uuid,omitempty"`
	PartitionToUuid   []byte `protobuf:"bytes,4,opt,name=partition_to_uuid,json=partitionToUuid,proto3" json:"partition_to_uuid,omitempty"`
}

func (x *JobDetails) Reset() {
	*x = JobDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbspgraph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDetails) ProtoMessage() {}

func (x *JobDetails) ProtoReflect() protoreflect.Message {
	mi := &file_dbspgraph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDetails.ProtoReflect.Descriptor instead.
func (*JobDetails) Descriptor() ([]byte, []int) {
	return file_dbspgraph_proto_rawDescGZIP(), []int{2}
}

func (x *JobDetails) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobDetails) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *JobDetails) GetPartitionFromUuid() []byte {
	if x != nil {
		return x.PartitionFromUuid
	}
	return nil
}

func (x *JobDetails) GetPartitionToUuid() []byte {
	if x != nil {
		return x.PartitionToUuid
	}
	return nil
}

// Step describes the current state of a worker or a master. Workers send a
// Step message with their current state to enter a synchronization barrier
// and wait for the other workers. Once all workers reach the barrier, the
// master node (depending on the step type) processes the individual worker
// steps to update its global state and broadcasts a new Step message (with
// a matching step type) to notify the workers that they can exit the barrier.
type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of this step.
	Type Step_Type `protobuf:"varint,1,opt,name=type,proto3,enum=proto.Step_Type" json:"type,omitempty"`
	// Workers use this field to submit their local aggregator delta values wen
	// reaching the POST step. The master collects the deltas, aggregates them to
	// its own aggregator values and broadcasts the global aggregator values in
	// the response. Workers must then *overwrite* their local aggregator values
	// with the values provided by the master.
	AggregatorValues map[string]*anypb.Any `protobuf:"bytes,2,rep,name=aggregator_values,json=aggregatorValues,proto3" json:"aggregator_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Workers use this field to submit their local active-in-step count when
	// reaching the POST_KEEP_RUNNING step. The step response broadcasted by
	// the master uses the same field to specify the global active-in-step count
	// that the workers should pass to the graph executor callbacks.
	ActiveInStep int64 `protobuf:"varint,3,opt,name=activeInStep,proto3" json:"activeInStep,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbspgraph_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_dbspgraph_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_dbspgraph_proto_rawDescGZIP(), []int{3}
}

func (x *Step) GetType() Step_Type {
	if x != nil {
		return x.Type
	}
	return Step_INVALID
}

func (x *Step) GetAggregatorValues() map[string]*anypb.Any {
	if x != nil {
		return x.AggregatorValues
	}
	return nil
}

func (x *Step) GetActiveInStep() int64 {
	if x != nil {
		return x.ActiveInStep
	}
	return 0
}

// RelayMessage describes a graph message that should be relayed to a remote
// graph instance which is managed by another worker.
type RelayMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message destination UUID.
	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// The serialized message contents.
	Message *anypb.Any `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RelayMessage) Reset() {
	*x = RelayMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dbspgraph_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayMessage) ProtoMessage() {}

func (x *RelayMessage) ProtoReflect() protoreflect.Message {
	mi := &file_dbspgraph_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayMessage.ProtoReflect.Descriptor instead.
func (*RelayMessage) Descriptor() ([]byte, []int) {
	return file_dbspgraph_proto_rawDescGZIP(), []int{4}
}

func (x *RelayMessage) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *RelayMessage) GetMessage() *anypb.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_dbspgraph_proto protoreflect.FileDescriptor

var file_dbspgraph_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x62, 0x73, 0x70, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x65, 0x70,
	0x48, 0x00, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0xaf, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x34, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a,
	0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x6a, 0x6f, 0x62,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x48, 0x00, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x3a, 0x0a, 0x0d, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0xba, 0x01, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x6f, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x55, 0x75, 0x69, 0x64, 0x22, 0xf7,
	0x02, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a,
	0x11, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x53, 0x74, 0x65,
	0x70, 0x1a, 0x59, 0x0a, 0x15, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7a, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x52, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f,
	0x53, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x4b, 0x45, 0x45,
	0x50, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x10, 0x04, 0x12,
	0x14, 0x0a, 0x10, 0x50, 0x45, 0x53, 0x49, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x53, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x5f, 0x4a, 0x4f, 0x42, 0x10, 0x06, 0x22, 0x60, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a, 0x08, 0x4a, 0x6f,
	0x62, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dbspgraph_proto_rawDescOnce sync.Once
	file_dbspgraph_proto_rawDescData = file_dbspgraph_proto_rawDesc
)

func file_dbspgraph_proto_rawDescGZIP() []byte {
	file_dbspgraph_proto_rawDescOnce.Do(func() {
		file_dbspgraph_proto_rawDescData = protoimpl.X.CompressGZIP(file_dbspgraph_proto_rawDescData)
	})
	return file_dbspgraph_proto_rawDescData
}

var file_dbspgraph_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dbspgraph_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dbspgraph_proto_goTypes = []interface{}{
	(Step_Type)(0),                // 0: proto.Step.Type
	(*WorkerPayload)(nil),         // 1: proto.WorkerPayload
	(*MasterPayload)(nil),         // 2: proto.MasterPayload
	(*JobDetails)(nil),            // 3: proto.JobDetails
	(*Step)(nil),                  // 4: proto.Step
	(*RelayMessage)(nil),          // 5: proto.RelayMessage
	nil,                           // 6: proto.Step.AggregatorValuesEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 8: google.protobuf.Any
}
var file_dbspgraph_proto_depIdxs = []int32{
	4,  // 0: proto.WorkerPayload.step:type_name -> proto.Step
	5,  // 1: proto.WorkerPayload.relay_message:type_name -> proto.RelayMessage
	3,  // 2: proto.MasterPayload.job_details:type_name -> proto.JobDetails
	4,  // 3: proto.MasterPayload.step:type_name -> proto.Step
	5,  // 4: proto.MasterPayload.relay_message:type_name -> proto.RelayMessage
	7,  // 5: proto.JobDetails.created_at:type_name -> google.protobuf.Timestamp
	0,  // 6: proto.Step.type:type_name -> proto.Step.Type
	6,  // 7: proto.Step.aggregator_values:type_name -> proto.Step.AggregatorValuesEntry
	8,  // 8: proto.RelayMessage.message:type_name -> google.protobuf.Any
	8,  // 9: proto.Step.AggregatorValuesEntry.value:type_name -> google.protobuf.Any
	1,  // 10: proto.JobQueue.JobStream:input_type -> proto.WorkerPayload
	2,  // 11: proto.JobQueue.JobStream:output_type -> proto.MasterPayload
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_dbspgraph_proto_init() }
func file_dbspgraph_proto_init() {
	if File_dbspgraph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dbspgraph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerPayload); i {
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
		file_dbspgraph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterPayload); i {
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
		file_dbspgraph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDetails); i {
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
		file_dbspgraph_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_dbspgraph_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayMessage); i {
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
	file_dbspgraph_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WorkerPayload_Step)(nil),
		(*WorkerPayload_RelayMessage)(nil),
	}
	file_dbspgraph_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MasterPayload_JobDetails)(nil),
		(*MasterPayload_Step)(nil),
		(*MasterPayload_RelayMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dbspgraph_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dbspgraph_proto_goTypes,
		DependencyIndexes: file_dbspgraph_proto_depIdxs,
		EnumInfos:         file_dbspgraph_proto_enumTypes,
		MessageInfos:      file_dbspgraph_proto_msgTypes,
	}.Build()
	File_dbspgraph_proto = out.File
	file_dbspgraph_proto_rawDesc = nil
	file_dbspgraph_proto_goTypes = nil
	file_dbspgraph_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobQueueClient is the client API for JobQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobQueueClient interface {
	// JobStream establishes a bi-directional connection between a worker and a
	// master. The master eventually broadcasts a new job to the connected
	// workers and the job super-steps are executed in lock-step across all
	// workers.
	JobStream(ctx context.Context, opts ...grpc.CallOption) (JobQueue_JobStreamClient, error)
}

type jobQueueClient struct {
	cc grpc.ClientConnInterface
}

func NewJobQueueClient(cc grpc.ClientConnInterface) JobQueueClient {
	return &jobQueueClient{cc}
}

func (c *jobQueueClient) JobStream(ctx context.Context, opts ...grpc.CallOption) (JobQueue_JobStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JobQueue_serviceDesc.Streams[0], "/proto.JobQueue/JobStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobQueueJobStreamClient{stream}
	return x, nil
}

type JobQueue_JobStreamClient interface {
	Send(*WorkerPayload) error
	Recv() (*MasterPayload, error)
	grpc.ClientStream
}

type jobQueueJobStreamClient struct {
	grpc.ClientStream
}

func (x *jobQueueJobStreamClient) Send(m *WorkerPayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *jobQueueJobStreamClient) Recv() (*MasterPayload, error) {
	m := new(MasterPayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JobQueueServer is the server API for JobQueue service.
type JobQueueServer interface {
	// JobStream establishes a bi-directional connection between a worker and a
	// master. The master eventually broadcasts a new job to the connected
	// workers and the job super-steps are executed in lock-step across all
	// workers.
	JobStream(JobQueue_JobStreamServer) error
}

// UnimplementedJobQueueServer can be embedded to have forward compatible implementations.
type UnimplementedJobQueueServer struct {
}

func (*UnimplementedJobQueueServer) JobStream(JobQueue_JobStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method JobStream not implemented")
}

func RegisterJobQueueServer(s *grpc.Server, srv JobQueueServer) {
	s.RegisterService(&_JobQueue_serviceDesc, srv)
}

func _JobQueue_JobStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JobQueueServer).JobStream(&jobQueueJobStreamServer{stream})
}

type JobQueue_JobStreamServer interface {
	Send(*MasterPayload) error
	Recv() (*WorkerPayload, error)
	grpc.ServerStream
}

type jobQueueJobStreamServer struct {
	grpc.ServerStream
}

func (x *jobQueueJobStreamServer) Send(m *MasterPayload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *jobQueueJobStreamServer) Recv() (*WorkerPayload, error) {
	m := new(WorkerPayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _JobQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.JobQueue",
	HandlerType: (*JobQueueServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JobStream",
			Handler:       _JobQueue_JobStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dbspgraph.proto",
}
