// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.3
// source: models/comment.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type CommentBasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId      int64                  `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	MarkdownId     int64                  `protobuf:"varint,2,opt,name=markdown_id,json=markdownId,proto3" json:"markdown_id,omitempty"`
	UserId         int64                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ParentId       int64                  `protobuf:"varint,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	RootId         int64                  `protobuf:"varint,5,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	CommentContent string                 `protobuf:"bytes,6,opt,name=comment_content,json=commentContent,proto3" json:"comment_content,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CommentBasicInfo) Reset() {
	*x = CommentBasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentBasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentBasicInfo) ProtoMessage() {}

func (x *CommentBasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentBasicInfo.ProtoReflect.Descriptor instead.
func (*CommentBasicInfo) Descriptor() ([]byte, []int) {
	return file_models_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentBasicInfo) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *CommentBasicInfo) GetMarkdownId() int64 {
	if x != nil {
		return x.MarkdownId
	}
	return 0
}

func (x *CommentBasicInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentBasicInfo) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CommentBasicInfo) GetRootId() int64 {
	if x != nil {
		return x.RootId
	}
	return 0
}

func (x *CommentBasicInfo) GetCommentContent() string {
	if x != nil {
		return x.CommentContent
	}
	return ""
}

func (x *CommentBasicInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CommentCountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId  int64 `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	LikeCount  int32 `protobuf:"varint,2,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	ReplyCount int32 `protobuf:"varint,3,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"`
	IsLiked    bool  `protobuf:"varint,4,opt,name=is_liked,json=isLiked,proto3" json:"is_liked,omitempty"`
	IsDisliked bool  `protobuf:"varint,5,opt,name=is_disliked,json=isDisliked,proto3" json:"is_disliked,omitempty"`
	IsShared   bool  `protobuf:"varint,6,opt,name=is_shared,json=isShared,proto3" json:"is_shared,omitempty"`
	IsReported bool  `protobuf:"varint,7,opt,name=is_reported,json=isReported,proto3" json:"is_reported,omitempty"`
}

func (x *CommentCountInfo) Reset() {
	*x = CommentCountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentCountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentCountInfo) ProtoMessage() {}

func (x *CommentCountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentCountInfo.ProtoReflect.Descriptor instead.
func (*CommentCountInfo) Descriptor() ([]byte, []int) {
	return file_models_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentCountInfo) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *CommentCountInfo) GetLikeCount() int32 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *CommentCountInfo) GetReplyCount() int32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

func (x *CommentCountInfo) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

func (x *CommentCountInfo) GetIsDisliked() bool {
	if x != nil {
		return x.IsDisliked
	}
	return false
}

func (x *CommentCountInfo) GetIsShared() bool {
	if x != nil {
		return x.IsShared
	}
	return false
}

func (x *CommentCountInfo) GetIsReported() bool {
	if x != nil {
		return x.IsReported
	}
	return false
}

type ListCommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarkdownId     int64                  `protobuf:"varint,1,opt,name=markdown_id,json=markdownId,proto3" json:"markdown_id,omitempty"`
	ParentId       int64                  `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Username       string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Pusername      string                 `protobuf:"bytes,4,opt,name=pusername,proto3" json:"pusername,omitempty"`
	CommentContent string                 `protobuf:"bytes,5,opt,name=comment_content,json=commentContent,proto3" json:"comment_content,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LikeCount      int64                  `protobuf:"varint,7,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	ReplyCount     int64                  `protobuf:"varint,8,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"`
	IsLiked        bool                   `protobuf:"varint,9,opt,name=is_liked,json=isLiked,proto3" json:"is_liked,omitempty"`
	IsDisliked     bool                   `protobuf:"varint,10,opt,name=is_disliked,json=isDisliked,proto3" json:"is_disliked,omitempty"`
	IsShared       bool                   `protobuf:"varint,11,opt,name=is_shared,json=isShared,proto3" json:"is_shared,omitempty"`
	IsReported     bool                   `protobuf:"varint,12,opt,name=is_reported,json=isReported,proto3" json:"is_reported,omitempty"`
	CommentId      int64                  `protobuf:"varint,13,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
}

func (x *ListCommentInfo) Reset() {
	*x = ListCommentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentInfo) ProtoMessage() {}

func (x *ListCommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentInfo.ProtoReflect.Descriptor instead.
func (*ListCommentInfo) Descriptor() ([]byte, []int) {
	return file_models_comment_proto_rawDescGZIP(), []int{2}
}

func (x *ListCommentInfo) GetMarkdownId() int64 {
	if x != nil {
		return x.MarkdownId
	}
	return 0
}

func (x *ListCommentInfo) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *ListCommentInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ListCommentInfo) GetPusername() string {
	if x != nil {
		return x.Pusername
	}
	return ""
}

func (x *ListCommentInfo) GetCommentContent() string {
	if x != nil {
		return x.CommentContent
	}
	return ""
}

func (x *ListCommentInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListCommentInfo) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *ListCommentInfo) GetReplyCount() int64 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

func (x *ListCommentInfo) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

func (x *ListCommentInfo) GetIsDisliked() bool {
	if x != nil {
		return x.IsDisliked
	}
	return false
}

func (x *ListCommentInfo) GetIsShared() bool {
	if x != nil {
		return x.IsShared
	}
	return false
}

func (x *ListCommentInfo) GetIsReported() bool {
	if x != nil {
		return x.IsReported
	}
	return false
}

func (x *ListCommentInfo) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

type ListAdminCommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId      int64                  `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	CommentContent string                 `protobuf:"bytes,2,opt,name=comment_content,json=commentContent,proto3" json:"comment_content,omitempty"`
	Username       string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email          string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ListAdminCommentInfo) Reset() {
	*x = ListAdminCommentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdminCommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdminCommentInfo) ProtoMessage() {}

func (x *ListAdminCommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdminCommentInfo.ProtoReflect.Descriptor instead.
func (*ListAdminCommentInfo) Descriptor() ([]byte, []int) {
	return file_models_comment_proto_rawDescGZIP(), []int{3}
}

func (x *ListAdminCommentInfo) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *ListAdminCommentInfo) GetCommentContent() string {
	if x != nil {
		return x.CommentContent
	}
	return ""
}

func (x *ListAdminCommentInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ListAdminCommentInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ListAdminCommentInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_models_comment_proto protoreflect.FileDescriptor

var file_models_comment_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x10,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xeb, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x69, 0x6b,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x69,
	0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x69, 0x73, 0x6c, 0x69,
	0x6b, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x22, 0xc6, 0x03, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x6b,
	0x64, 0x6f, 0x77, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xcb, 0x01, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x2f, 0x7a, 0x62,
	0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_comment_proto_rawDescOnce sync.Once
	file_models_comment_proto_rawDescData = file_models_comment_proto_rawDesc
)

func file_models_comment_proto_rawDescGZIP() []byte {
	file_models_comment_proto_rawDescOnce.Do(func() {
		file_models_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_comment_proto_rawDescData)
	})
	return file_models_comment_proto_rawDescData
}

var file_models_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_models_comment_proto_goTypes = []interface{}{
	(*CommentBasicInfo)(nil),      // 0: pb.CommentBasicInfo
	(*CommentCountInfo)(nil),      // 1: pb.CommentCountInfo
	(*ListCommentInfo)(nil),       // 2: pb.ListCommentInfo
	(*ListAdminCommentInfo)(nil),  // 3: pb.ListAdminCommentInfo
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_models_comment_proto_depIdxs = []int32{
	4, // 0: pb.CommentBasicInfo.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: pb.ListCommentInfo.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: pb.ListAdminCommentInfo.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_comment_proto_init() }
func file_models_comment_proto_init() {
	if File_models_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentBasicInfo); i {
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
		file_models_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentCountInfo); i {
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
		file_models_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentInfo); i {
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
		file_models_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAdminCommentInfo); i {
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
			RawDescriptor: file_models_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_comment_proto_goTypes,
		DependencyIndexes: file_models_comment_proto_depIdxs,
		MessageInfos:      file_models_comment_proto_msgTypes,
	}.Build()
	File_models_comment_proto = out.File
	file_models_comment_proto_rawDesc = nil
	file_models_comment_proto_goTypes = nil
	file_models_comment_proto_depIdxs = nil
}
