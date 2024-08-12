// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: health.proto

package health

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

type GenerateHealthRecommendationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId             string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RecommendationType string `protobuf:"bytes,2,opt,name=recommendation_type,json=recommendationType,proto3" json:"recommendation_type,omitempty"`
	Description        string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Priority           int32  `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *GenerateHealthRecommendationsReq) Reset() {
	*x = GenerateHealthRecommendationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateHealthRecommendationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateHealthRecommendationsReq) ProtoMessage() {}

func (x *GenerateHealthRecommendationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateHealthRecommendationsReq.ProtoReflect.Descriptor instead.
func (*GenerateHealthRecommendationsReq) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateHealthRecommendationsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GenerateHealthRecommendationsReq) GetRecommendationType() string {
	if x != nil {
		return x.RecommendationType
	}
	return ""
}

func (x *GenerateHealthRecommendationsReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GenerateHealthRecommendationsReq) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type GenerateHealthRecommendationsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message bool `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GenerateHealthRecommendationsRes) Reset() {
	*x = GenerateHealthRecommendationsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateHealthRecommendationsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateHealthRecommendationsRes) ProtoMessage() {}

func (x *GenerateHealthRecommendationsRes) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateHealthRecommendationsRes.ProtoReflect.Descriptor instead.
func (*GenerateHealthRecommendationsRes) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateHealthRecommendationsRes) GetMessage() bool {
	if x != nil {
		return x.Message
	}
	return false
}

type HealthRecommendation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName          string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName           string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	RecommendationType string `protobuf:"bytes,3,opt,name=recommendation_type,json=recommendationType,proto3" json:"recommendation_type,omitempty"`
	Description        string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Priority           int32  `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *HealthRecommendation) Reset() {
	*x = HealthRecommendation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRecommendation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRecommendation) ProtoMessage() {}

func (x *HealthRecommendation) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRecommendation.ProtoReflect.Descriptor instead.
func (*HealthRecommendation) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{2}
}

func (x *HealthRecommendation) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *HealthRecommendation) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *HealthRecommendation) GetRecommendationType() string {
	if x != nil {
		return x.RecommendationType
	}
	return ""
}

func (x *HealthRecommendation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HealthRecommendation) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type GetRealtimeHealthMonitoringReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetRealtimeHealthMonitoringReq) Reset() {
	*x = GetRealtimeHealthMonitoringReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRealtimeHealthMonitoringReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealtimeHealthMonitoringReq) ProtoMessage() {}

func (x *GetRealtimeHealthMonitoringReq) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealtimeHealthMonitoringReq.ProtoReflect.Descriptor instead.
func (*GetRealtimeHealthMonitoringReq) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{3}
}

func (x *GetRealtimeHealthMonitoringReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetRealtimeHealthMonitoringRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName          string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName           string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	RecommendationType string `protobuf:"bytes,3,opt,name=recommendation_type,json=recommendationType,proto3" json:"recommendation_type,omitempty"`
	Description        string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Priority           int32  `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *GetRealtimeHealthMonitoringRes) Reset() {
	*x = GetRealtimeHealthMonitoringRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRealtimeHealthMonitoringRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealtimeHealthMonitoringRes) ProtoMessage() {}

func (x *GetRealtimeHealthMonitoringRes) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealtimeHealthMonitoringRes.ProtoReflect.Descriptor instead.
func (*GetRealtimeHealthMonitoringRes) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{4}
}

func (x *GetRealtimeHealthMonitoringRes) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetRealtimeHealthMonitoringRes) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetRealtimeHealthMonitoringRes) GetRecommendationType() string {
	if x != nil {
		return x.RecommendationType
	}
	return ""
}

func (x *GetRealtimeHealthMonitoringRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetRealtimeHealthMonitoringRes) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type GetDailyHealthSummaryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Date      string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetDailyHealthSummaryReq) Reset() {
	*x = GetDailyHealthSummaryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDailyHealthSummaryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDailyHealthSummaryReq) ProtoMessage() {}

func (x *GetDailyHealthSummaryReq) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDailyHealthSummaryReq.ProtoReflect.Descriptor instead.
func (*GetDailyHealthSummaryReq) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{5}
}

func (x *GetDailyHealthSummaryReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetDailyHealthSummaryReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetDailyHealthSummaryReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetDailyHealthSummaryReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetDailyHealthSummaryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recommendations *HealthRecommendation `protobuf:"bytes,1,opt,name=recommendations,proto3" json:"recommendations,omitempty"`
}

func (x *GetDailyHealthSummaryRes) Reset() {
	*x = GetDailyHealthSummaryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDailyHealthSummaryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDailyHealthSummaryRes) ProtoMessage() {}

func (x *GetDailyHealthSummaryRes) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDailyHealthSummaryRes.ProtoReflect.Descriptor instead.
func (*GetDailyHealthSummaryRes) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{6}
}

func (x *GetDailyHealthSummaryRes) GetRecommendations() *HealthRecommendation {
	if x != nil {
		return x.Recommendations
	}
	return nil
}

type GetWeeklyHealthSummaryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartDate string `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *GetWeeklyHealthSummaryReq) Reset() {
	*x = GetWeeklyHealthSummaryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeeklyHealthSummaryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeeklyHealthSummaryReq) ProtoMessage() {}

func (x *GetWeeklyHealthSummaryReq) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeeklyHealthSummaryReq.ProtoReflect.Descriptor instead.
func (*GetWeeklyHealthSummaryReq) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{7}
}

func (x *GetWeeklyHealthSummaryReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetWeeklyHealthSummaryReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetWeeklyHealthSummaryReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetWeeklyHealthSummaryReq) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetWeeklyHealthSummaryReq) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type GetWeeklyHealthSummaryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recommendations []*HealthRecommendation `protobuf:"bytes,1,rep,name=recommendations,proto3" json:"recommendations,omitempty"`
}

func (x *GetWeeklyHealthSummaryRes) Reset() {
	*x = GetWeeklyHealthSummaryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeeklyHealthSummaryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeeklyHealthSummaryRes) ProtoMessage() {}

func (x *GetWeeklyHealthSummaryRes) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeeklyHealthSummaryRes.ProtoReflect.Descriptor instead.
func (*GetWeeklyHealthSummaryRes) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{8}
}

func (x *GetWeeklyHealthSummaryRes) GetRecommendations() []*HealthRecommendation {
	if x != nil {
		return x.Recommendations
	}
	return nil
}

var File_health_proto protoreflect.FileDescriptor

var file_health_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0xaa, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x22, 0x3c, 0x0a, 0x20, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xc1, 0x01, 0x0a, 0x14, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x39, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c,
	0x74, 0x69, 0x6d, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xcb, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x83,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x62, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x12, 0x46, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x63, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x12, 0x46, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xae, 0x03, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x73, 0x0a, 0x1d, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12,
	0x6d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x26,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x74,
	0x69, 0x6d, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x12, 0x5b,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_proto_rawDescOnce sync.Once
	file_health_proto_rawDescData = file_health_proto_rawDesc
)

func file_health_proto_rawDescGZIP() []byte {
	file_health_proto_rawDescOnce.Do(func() {
		file_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_proto_rawDescData)
	})
	return file_health_proto_rawDescData
}

var file_health_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_health_proto_goTypes = []any{
	(*GenerateHealthRecommendationsReq)(nil), // 0: health.GenerateHealthRecommendationsReq
	(*GenerateHealthRecommendationsRes)(nil), // 1: health.GenerateHealthRecommendationsRes
	(*HealthRecommendation)(nil),             // 2: health.HealthRecommendation
	(*GetRealtimeHealthMonitoringReq)(nil),   // 3: health.GetRealtimeHealthMonitoringReq
	(*GetRealtimeHealthMonitoringRes)(nil),   // 4: health.GetRealtimeHealthMonitoringRes
	(*GetDailyHealthSummaryReq)(nil),         // 5: health.GetDailyHealthSummaryReq
	(*GetDailyHealthSummaryRes)(nil),         // 6: health.GetDailyHealthSummaryRes
	(*GetWeeklyHealthSummaryReq)(nil),        // 7: health.GetWeeklyHealthSummaryReq
	(*GetWeeklyHealthSummaryRes)(nil),        // 8: health.GetWeeklyHealthSummaryRes
}
var file_health_proto_depIdxs = []int32{
	2, // 0: health.GetDailyHealthSummaryRes.recommendations:type_name -> health.HealthRecommendation
	2, // 1: health.GetWeeklyHealthSummaryRes.recommendations:type_name -> health.HealthRecommendation
	0, // 2: health.HealthCheck.GenerateHealthRecommendations:input_type -> health.GenerateHealthRecommendationsReq
	3, // 3: health.HealthCheck.GetRealtimeHealthMonitoring:input_type -> health.GetRealtimeHealthMonitoringReq
	5, // 4: health.HealthCheck.GetDailyHealthSummary:input_type -> health.GetDailyHealthSummaryReq
	7, // 5: health.HealthCheck.GetWeeklyHealthSummary:input_type -> health.GetWeeklyHealthSummaryReq
	1, // 6: health.HealthCheck.GenerateHealthRecommendations:output_type -> health.GenerateHealthRecommendationsRes
	4, // 7: health.HealthCheck.GetRealtimeHealthMonitoring:output_type -> health.GetRealtimeHealthMonitoringRes
	6, // 8: health.HealthCheck.GetDailyHealthSummary:output_type -> health.GetDailyHealthSummaryRes
	8, // 9: health.HealthCheck.GetWeeklyHealthSummary:output_type -> health.GetWeeklyHealthSummaryRes
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_health_proto_init() }
func file_health_proto_init() {
	if File_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateHealthRecommendationsReq); i {
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
		file_health_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateHealthRecommendationsRes); i {
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
		file_health_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HealthRecommendation); i {
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
		file_health_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetRealtimeHealthMonitoringReq); i {
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
		file_health_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetRealtimeHealthMonitoringRes); i {
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
		file_health_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetDailyHealthSummaryReq); i {
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
		file_health_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetDailyHealthSummaryRes); i {
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
		file_health_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetWeeklyHealthSummaryReq); i {
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
		file_health_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetWeeklyHealthSummaryRes); i {
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
			RawDescriptor: file_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_proto_goTypes,
		DependencyIndexes: file_health_proto_depIdxs,
		MessageInfos:      file_health_proto_msgTypes,
	}.Build()
	File_health_proto = out.File
	file_health_proto_rawDesc = nil
	file_health_proto_goTypes = nil
	file_health_proto_depIdxs = nil
}
