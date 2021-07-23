// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: movie.proto

package pb

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

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tconst      string `protobuf:"bytes,1,opt,name=tconst,proto3" json:"tconst,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ReleaseDate int64  `protobuf:"varint,3,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Country     string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Runtime     int32  `protobuf:"varint,5,opt,name=runtime,proto3" json:"runtime,omitempty"`
	MpaRating   string `protobuf:"bytes,6,opt,name=mpa_rating,json=mpaRating,proto3" json:"mpa_rating,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetTconst() string {
	if x != nil {
		return x.Tconst
	}
	return ""
}

func (x *Movie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Movie) GetReleaseDate() int64 {
	if x != nil {
		return x.ReleaseDate
	}
	return 0
}

func (x *Movie) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Movie) GetRuntime() int32 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

func (x *Movie) GetMpaRating() string {
	if x != nil {
		return x.MpaRating
	}
	return ""
}

type PutMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *PutMovieRequest) Reset() {
	*x = PutMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutMovieRequest) ProtoMessage() {}

func (x *PutMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutMovieRequest.ProtoReflect.Descriptor instead.
func (*PutMovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *PutMovieRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type PutMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PutMovieResponse) Reset() {
	*x = PutMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutMovieResponse) ProtoMessage() {}

func (x *PutMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutMovieResponse.ProtoReflect.Descriptor instead.
func (*PutMovieResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{2}
}

func (x *PutMovieResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PutMovieResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tconst string `protobuf:"bytes,1,opt,name=tconst,proto3" json:"tconst,omitempty"`
}

func (x *GetMovieRequest) Reset() {
	*x = GetMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieRequest) ProtoMessage() {}

func (x *GetMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieRequest.ProtoReflect.Descriptor instead.
func (*GetMovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{3}
}

func (x *GetMovieRequest) GetTconst() string {
	if x != nil {
		return x.Tconst
	}
	return ""
}

type GetMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Movie      *Movie `protobuf:"bytes,3,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *GetMovieResponse) Reset() {
	*x = GetMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieResponse) ProtoMessage() {}

func (x *GetMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieResponse.ProtoReflect.Descriptor instead.
func (*GetMovieResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{4}
}

func (x *GetMovieResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetMovieResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMovieResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type GetMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryName string `protobuf:"bytes,1,opt,name=query_name,json=queryName,proto3" json:"query_name,omitempty"`
}

func (x *GetMoviesRequest) Reset() {
	*x = GetMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoviesRequest) ProtoMessage() {}

func (x *GetMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoviesRequest.ProtoReflect.Descriptor instead.
func (*GetMoviesRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{5}
}

func (x *GetMoviesRequest) GetQueryName() string {
	if x != nil {
		return x.QueryName
	}
	return ""
}

type GetMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Movies     []*Movie `protobuf:"bytes,3,rep,name=movies,proto3" json:"movies,omitempty"`
}

func (x *GetMoviesResponse) Reset() {
	*x = GetMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoviesResponse) ProtoMessage() {}

func (x *GetMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoviesResponse.ProtoReflect.Descriptor instead.
func (*GetMoviesResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{6}
}

func (x *GetMoviesResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetMoviesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMoviesResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

type DeleteMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tconst string `protobuf:"bytes,1,opt,name=tconst,proto3" json:"tconst,omitempty"`
}

func (x *DeleteMovieRequest) Reset() {
	*x = DeleteMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovieRequest) ProtoMessage() {}

func (x *DeleteMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovieRequest.ProtoReflect.Descriptor instead.
func (*DeleteMovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMovieRequest) GetTconst() string {
	if x != nil {
		return x.Tconst
	}
	return ""
}

type DeleteMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteMovieResponse) Reset() {
	*x = DeleteMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMovieResponse) ProtoMessage() {}

func (x *DeleteMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMovieResponse.ProtoReflect.Descriptor instead.
func (*DeleteMovieResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMovieResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DeleteMovieResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xa9, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x70, 0x61, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x70, 0x61, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x32, 0x0a,
	0x0f, 0x50, 0x75, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x22, 0x4d, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x31, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x71,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21,
	0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x73, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x22,
	0x50, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0xf2, 0x01, 0x0a, 0x08, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x76, 0x63, 0x12, 0x35,
	0x0a, 0x08, 0x50, 0x75, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x75, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_movie_proto_goTypes = []interface{}{
	(*Movie)(nil),               // 0: pb.Movie
	(*PutMovieRequest)(nil),     // 1: pb.PutMovieRequest
	(*PutMovieResponse)(nil),    // 2: pb.PutMovieResponse
	(*GetMovieRequest)(nil),     // 3: pb.GetMovieRequest
	(*GetMovieResponse)(nil),    // 4: pb.GetMovieResponse
	(*GetMoviesRequest)(nil),    // 5: pb.GetMoviesRequest
	(*GetMoviesResponse)(nil),   // 6: pb.GetMoviesResponse
	(*DeleteMovieRequest)(nil),  // 7: pb.DeleteMovieRequest
	(*DeleteMovieResponse)(nil), // 8: pb.DeleteMovieResponse
}
var file_movie_proto_depIdxs = []int32{
	0, // 0: pb.PutMovieRequest.movie:type_name -> pb.Movie
	0, // 1: pb.GetMovieResponse.movie:type_name -> pb.Movie
	0, // 2: pb.GetMoviesResponse.movies:type_name -> pb.Movie
	1, // 3: pb.MovieSvc.PutMovie:input_type -> pb.PutMovieRequest
	3, // 4: pb.MovieSvc.GetMovie:input_type -> pb.GetMovieRequest
	5, // 5: pb.MovieSvc.GetMovies:input_type -> pb.GetMoviesRequest
	7, // 6: pb.MovieSvc.DeleteMovie:input_type -> pb.DeleteMovieRequest
	2, // 7: pb.MovieSvc.PutMovie:output_type -> pb.PutMovieResponse
	4, // 8: pb.MovieSvc.GetMovie:output_type -> pb.GetMovieResponse
	6, // 9: pb.MovieSvc.GetMovies:output_type -> pb.GetMoviesResponse
	8, // 10: pb.MovieSvc.DeleteMovie:output_type -> pb.DeleteMovieResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutMovieRequest); i {
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
		file_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutMovieResponse); i {
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
		file_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieRequest); i {
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
		file_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieResponse); i {
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
		file_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoviesRequest); i {
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
		file_movie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoviesResponse); i {
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
		file_movie_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMovieRequest); i {
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
		file_movie_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMovieResponse); i {
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
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}
