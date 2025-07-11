// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v5.29.3
// source: host/cache/cache.proto

package cache

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to store a string value
type SetStringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                                  // Cache key
	Value      string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`                              // String value to store
	TtlSeconds int64  `protobuf:"varint,3,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"` // TTL in seconds, 0 means use default
}

func (x *SetStringRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SetStringRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetStringRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SetStringRequest) GetTtlSeconds() int64 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

// Request to store an integer value
type SetIntRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                                  // Cache key
	Value      int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`                             // Integer value to store
	TtlSeconds int64  `protobuf:"varint,3,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"` // TTL in seconds, 0 means use default
}

func (x *SetIntRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SetIntRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetIntRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SetIntRequest) GetTtlSeconds() int64 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

// Request to store a float value
type SetFloatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                                  // Cache key
	Value      float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`                            // Float value to store
	TtlSeconds int64   `protobuf:"varint,3,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"` // TTL in seconds, 0 means use default
}

func (x *SetFloatRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SetFloatRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetFloatRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SetFloatRequest) GetTtlSeconds() int64 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

// Request to store a byte slice value
type SetBytesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                                  // Cache key
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`                              // Byte slice value to store
	TtlSeconds int64  `protobuf:"varint,3,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"` // TTL in seconds, 0 means use default
}

func (x *SetBytesRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SetBytesRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetBytesRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SetBytesRequest) GetTtlSeconds() int64 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

// Response after setting a value
type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Whether the operation was successful
}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request to get a value
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // Cache key
}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Response containing a string value
type GetStringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool   `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"` // Whether the key exists
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`    // The string value (if exists is true)
}

func (x *GetStringResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetStringResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

func (x *GetStringResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Response containing an integer value
type GetIntResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool  `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"` // Whether the key exists
	Value  int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`   // The integer value (if exists is true)
}

func (x *GetIntResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetIntResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

func (x *GetIntResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Response containing a float value
type GetFloatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool    `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"` // Whether the key exists
	Value  float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`  // The float value (if exists is true)
}

func (x *GetFloatResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetFloatResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

func (x *GetFloatResponse) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Response containing a byte slice value
type GetBytesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool   `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"` // Whether the key exists
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`    // The byte slice value (if exists is true)
}

func (x *GetBytesResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetBytesResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

func (x *GetBytesResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Request to remove a value
type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // Cache key
}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *RemoveRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Response after removing a value
type RemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Whether the operation was successful
}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *RemoveResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request to check if a key exists
type HasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` // Cache key
}

func (x *HasRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *HasRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Response indicating if a key exists
type HasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"` // Whether the key exists
}

func (x *HasResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *HasResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

// go:plugin type=host version=1
type CacheService interface {
	// Set a string value in the cache
	SetString(context.Context, *SetStringRequest) (*SetResponse, error)
	// Get a string value from the cache
	GetString(context.Context, *GetRequest) (*GetStringResponse, error)
	// Set an integer value in the cache
	SetInt(context.Context, *SetIntRequest) (*SetResponse, error)
	// Get an integer value from the cache
	GetInt(context.Context, *GetRequest) (*GetIntResponse, error)
	// Set a float value in the cache
	SetFloat(context.Context, *SetFloatRequest) (*SetResponse, error)
	// Get a float value from the cache
	GetFloat(context.Context, *GetRequest) (*GetFloatResponse, error)
	// Set a byte slice value in the cache
	SetBytes(context.Context, *SetBytesRequest) (*SetResponse, error)
	// Get a byte slice value from the cache
	GetBytes(context.Context, *GetRequest) (*GetBytesResponse, error)
	// Remove a value from the cache
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	// Check if a key exists in the cache
	Has(context.Context, *HasRequest) (*HasResponse, error)
}
