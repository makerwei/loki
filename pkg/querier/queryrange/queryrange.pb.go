// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/querier/queryrange/queryrange.proto

package queryrange

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	logproto "github.com/grafana/loki/pkg/logproto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LokiRequest struct {
	Query     string             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Limit     uint32             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Step      int64              `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	StartTs   time.Time          `protobuf:"bytes,4,opt,name=startTs,proto3,stdtime" json:"startTs"`
	EndTs     time.Time          `protobuf:"bytes,5,opt,name=endTs,proto3,stdtime" json:"endTs"`
	Direction logproto.Direction `protobuf:"varint,6,opt,name=direction,proto3,enum=logproto.Direction" json:"direction,omitempty"`
	Path      string             `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
}

func (m *LokiRequest) Reset()      { *m = LokiRequest{} }
func (*LokiRequest) ProtoMessage() {}
func (*LokiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b9d53b40d11902, []int{0}
}
func (m *LokiRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LokiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LokiRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LokiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LokiRequest.Merge(m, src)
}
func (m *LokiRequest) XXX_Size() int {
	return m.Size()
}
func (m *LokiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LokiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LokiRequest proto.InternalMessageInfo

func (m *LokiRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *LokiRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *LokiRequest) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *LokiRequest) GetStartTs() time.Time {
	if m != nil {
		return m.StartTs
	}
	return time.Time{}
}

func (m *LokiRequest) GetEndTs() time.Time {
	if m != nil {
		return m.EndTs
	}
	return time.Time{}
}

func (m *LokiRequest) GetDirection() logproto.Direction {
	if m != nil {
		return m.Direction
	}
	return logproto.FORWARD
}

func (m *LokiRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type LokiResponse struct {
	Status    string             `protobuf:"bytes,1,opt,name=Status,json=status,proto3" json:"status"`
	Data      LokiData           `protobuf:"bytes,2,opt,name=Data,json=data,proto3" json:"data,omitempty"`
	ErrorType string             `protobuf:"bytes,3,opt,name=ErrorType,json=errorType,proto3" json:"errorType,omitempty"`
	Error     string             `protobuf:"bytes,4,opt,name=Error,json=error,proto3" json:"error,omitempty"`
	Direction logproto.Direction `protobuf:"varint,5,opt,name=direction,proto3,enum=logproto.Direction" json:"direction,omitempty"`
	Limit     uint32             `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Version   uint32             `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *LokiResponse) Reset()      { *m = LokiResponse{} }
func (*LokiResponse) ProtoMessage() {}
func (*LokiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b9d53b40d11902, []int{1}
}
func (m *LokiResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LokiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LokiResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LokiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LokiResponse.Merge(m, src)
}
func (m *LokiResponse) XXX_Size() int {
	return m.Size()
}
func (m *LokiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LokiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LokiResponse proto.InternalMessageInfo

func (m *LokiResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *LokiResponse) GetData() LokiData {
	if m != nil {
		return m.Data
	}
	return LokiData{}
}

func (m *LokiResponse) GetErrorType() string {
	if m != nil {
		return m.ErrorType
	}
	return ""
}

func (m *LokiResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *LokiResponse) GetDirection() logproto.Direction {
	if m != nil {
		return m.Direction
	}
	return logproto.FORWARD
}

func (m *LokiResponse) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *LokiResponse) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type LokiData struct {
	ResultType string            `protobuf:"bytes,1,opt,name=ResultType,json=resultType,proto3" json:"resultType"`
	Result     []logproto.Stream `protobuf:"bytes,2,rep,name=Result,json=result,proto3" json:"result"`
}

func (m *LokiData) Reset()      { *m = LokiData{} }
func (*LokiData) ProtoMessage() {}
func (*LokiData) Descriptor() ([]byte, []int) {
	return fileDescriptor_51b9d53b40d11902, []int{2}
}
func (m *LokiData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LokiData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LokiData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LokiData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LokiData.Merge(m, src)
}
func (m *LokiData) XXX_Size() int {
	return m.Size()
}
func (m *LokiData) XXX_DiscardUnknown() {
	xxx_messageInfo_LokiData.DiscardUnknown(m)
}

var xxx_messageInfo_LokiData proto.InternalMessageInfo

func (m *LokiData) GetResultType() string {
	if m != nil {
		return m.ResultType
	}
	return ""
}

func (m *LokiData) GetResult() []logproto.Stream {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*LokiRequest)(nil), "queryrange.LokiRequest")
	proto.RegisterType((*LokiResponse)(nil), "queryrange.LokiResponse")
	proto.RegisterType((*LokiData)(nil), "queryrange.LokiData")
}

func init() {
	proto.RegisterFile("pkg/querier/queryrange/queryrange.proto", fileDescriptor_51b9d53b40d11902)
}

var fileDescriptor_51b9d53b40d11902 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xc1, 0x8e, 0xd3, 0x3c,
	0x10, 0x8e, 0xbb, 0x4d, 0xba, 0x71, 0xff, 0xbf, 0x20, 0x77, 0x05, 0x51, 0x91, 0x9c, 0xa8, 0x17,
	0x82, 0x04, 0xa9, 0x28, 0x20, 0x21, 0x0e, 0x08, 0x45, 0xcb, 0x8d, 0x93, 0xb7, 0x2f, 0x90, 0x6e,
	0x4d, 0x36, 0xda, 0xa6, 0xce, 0xda, 0x0e, 0x52, 0x6f, 0x3c, 0x42, 0x1f, 0x83, 0x57, 0xe0, 0x0d,
	0xf6, 0xd8, 0xe3, 0x9e, 0x02, 0x4d, 0x2f, 0xa8, 0xa7, 0x7d, 0x04, 0x14, 0xbb, 0x69, 0xc3, 0x0d,
	0x4e, 0x9e, 0xf9, 0x66, 0xbe, 0xf1, 0x37, 0xdf, 0xc0, 0xa7, 0xd9, 0x75, 0x3c, 0xba, 0xc9, 0x29,
	0x4f, 0x28, 0x57, 0xef, 0x92, 0x47, 0x8b, 0x98, 0x36, 0xc2, 0x20, 0xe3, 0x4c, 0x32, 0x04, 0x8f,
	0xc8, 0xe0, 0x45, 0x9c, 0xc8, 0xab, 0x7c, 0x1a, 0x5c, 0xb2, 0x74, 0x14, 0xb3, 0x98, 0x8d, 0x54,
	0xcb, 0x34, 0xff, 0xac, 0x32, 0x95, 0xa8, 0x48, 0x53, 0x07, 0x4f, 0xaa, 0x3f, 0xe6, 0x2c, 0xd6,
	0x85, 0x3a, 0xd8, 0x17, 0xdd, 0x98, 0xb1, 0x78, 0x4e, 0x8f, 0x23, 0x64, 0x92, 0x52, 0x21, 0xa3,
	0x34, 0xd3, 0x0d, 0xc3, 0x55, 0x0b, 0x76, 0x3f, 0xb1, 0xeb, 0x84, 0xd0, 0x9b, 0x9c, 0x0a, 0x89,
	0xce, 0xa0, 0xa9, 0xa4, 0x38, 0xc0, 0x03, 0xbe, 0x4d, 0x74, 0x52, 0xa1, 0xf3, 0x24, 0x4d, 0xa4,
	0xd3, 0xf2, 0x80, 0xff, 0x3f, 0xd1, 0x09, 0x42, 0xb0, 0x2d, 0x24, 0xcd, 0x9c, 0x13, 0x0f, 0xf8,
	0x27, 0x44, 0xc5, 0xe8, 0x3d, 0xec, 0x08, 0x19, 0x71, 0x39, 0x11, 0x4e, 0xdb, 0x03, 0x7e, 0x77,
	0x3c, 0x08, 0xb4, 0x84, 0xa0, 0x96, 0x10, 0x4c, 0x6a, 0x09, 0xe1, 0xe9, 0x6d, 0xe1, 0x1a, 0xab,
	0x1f, 0x2e, 0x20, 0x35, 0x09, 0xbd, 0x83, 0x26, 0x5d, 0xcc, 0x26, 0xc2, 0x31, 0xff, 0x81, 0xad,
	0x29, 0xe8, 0x25, 0xb4, 0x67, 0x09, 0xa7, 0x97, 0x32, 0x61, 0x0b, 0xc7, 0xf2, 0x80, 0xdf, 0x1b,
	0xf7, 0x83, 0x83, 0x21, 0xe7, 0x75, 0x89, 0x1c, 0xbb, 0xaa, 0x15, 0xb2, 0x48, 0x5e, 0x39, 0x1d,
	0xb5, 0xad, 0x8a, 0x87, 0xdf, 0x5b, 0xf0, 0x3f, 0x6d, 0x89, 0xc8, 0xd8, 0x42, 0x50, 0x34, 0x84,
	0xd6, 0x85, 0x8c, 0x64, 0x2e, 0xb4, 0x29, 0x21, 0xdc, 0x15, 0xae, 0x25, 0x14, 0x42, 0xf6, 0x2f,
	0xfa, 0x00, 0xdb, 0xe7, 0x91, 0x8c, 0x94, 0x41, 0xdd, 0xf1, 0x59, 0xd0, 0xb8, 0x70, 0x35, 0xab,
	0xaa, 0x85, 0x8f, 0x2a, 0xc1, 0xbb, 0xc2, 0xed, 0xcd, 0x22, 0x19, 0x3d, 0x67, 0x69, 0x22, 0x69,
	0x9a, 0xc9, 0x25, 0x69, 0x57, 0x39, 0x7a, 0x03, 0xed, 0x8f, 0x9c, 0x33, 0x3e, 0x59, 0x66, 0x54,
	0x59, 0x6a, 0x87, 0x8f, 0x77, 0x85, 0xdb, 0xa7, 0x35, 0xd8, 0x60, 0xd8, 0x07, 0x10, 0x3d, 0x83,
	0xa6, 0xa2, 0x29, 0xbb, 0xed, 0xb0, 0xbf, 0x2b, 0xdc, 0x07, 0xaa, 0xda, 0x68, 0x37, 0x15, 0xf0,
	0xa7, 0x3f, 0xe6, 0x5f, 0xf9, 0x73, 0x38, 0xbc, 0xd5, 0x3c, 0xbc, 0x03, 0x3b, 0x5f, 0x28, 0x17,
	0xd5, 0x98, 0x8e, 0xc2, 0xeb, 0x74, 0x28, 0xe1, 0x69, 0xbd, 0x2e, 0x0a, 0x20, 0x24, 0x54, 0xe4,
	0x73, 0xa9, 0x36, 0xd2, 0xd6, 0xf5, 0x76, 0x85, 0x0b, 0xf9, 0x01, 0x25, 0x8d, 0x18, 0xbd, 0x85,
	0x96, 0xee, 0x77, 0x5a, 0xde, 0x89, 0xdf, 0x1d, 0x3f, 0x3c, 0x6a, 0xbb, 0x90, 0x9c, 0x46, 0x69,
	0xd8, 0xdb, 0x1b, 0x68, 0x69, 0x16, 0xd9, 0xbf, 0xe1, 0xeb, 0xf5, 0x06, 0x1b, 0x77, 0x1b, 0x6c,
	0xdc, 0x6f, 0x30, 0xf8, 0x5a, 0x62, 0xf0, 0xad, 0xc4, 0xe0, 0xb6, 0xc4, 0x60, 0x5d, 0x62, 0xf0,
	0xb3, 0xc4, 0xe0, 0x57, 0x89, 0x8d, 0xfb, 0x12, 0x83, 0xd5, 0x16, 0x1b, 0xeb, 0x2d, 0x36, 0xee,
	0xb6, 0xd8, 0x98, 0x5a, 0x6a, 0xf6, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0xee, 0xaf,
	0xb4, 0xa5, 0x03, 0x00, 0x00,
}

func (this *LokiRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LokiRequest)
	if !ok {
		that2, ok := that.(LokiRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Query != that1.Query {
		return false
	}
	if this.Limit != that1.Limit {
		return false
	}
	if this.Step != that1.Step {
		return false
	}
	if !this.StartTs.Equal(that1.StartTs) {
		return false
	}
	if !this.EndTs.Equal(that1.EndTs) {
		return false
	}
	if this.Direction != that1.Direction {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	return true
}
func (this *LokiResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LokiResponse)
	if !ok {
		that2, ok := that.(LokiResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !this.Data.Equal(&that1.Data) {
		return false
	}
	if this.ErrorType != that1.ErrorType {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	if this.Direction != that1.Direction {
		return false
	}
	if this.Limit != that1.Limit {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (this *LokiData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LokiData)
	if !ok {
		that2, ok := that.(LokiData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ResultType != that1.ResultType {
		return false
	}
	if len(this.Result) != len(that1.Result) {
		return false
	}
	for i := range this.Result {
		if !this.Result[i].Equal(&that1.Result[i]) {
			return false
		}
	}
	return true
}
func (this *LokiRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&queryrange.LokiRequest{")
	s = append(s, "Query: "+fmt.Sprintf("%#v", this.Query)+",\n")
	s = append(s, "Limit: "+fmt.Sprintf("%#v", this.Limit)+",\n")
	s = append(s, "Step: "+fmt.Sprintf("%#v", this.Step)+",\n")
	s = append(s, "StartTs: "+fmt.Sprintf("%#v", this.StartTs)+",\n")
	s = append(s, "EndTs: "+fmt.Sprintf("%#v", this.EndTs)+",\n")
	s = append(s, "Direction: "+fmt.Sprintf("%#v", this.Direction)+",\n")
	s = append(s, "Path: "+fmt.Sprintf("%#v", this.Path)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LokiResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&queryrange.LokiResponse{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Data: "+strings.Replace(this.Data.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "ErrorType: "+fmt.Sprintf("%#v", this.ErrorType)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "Direction: "+fmt.Sprintf("%#v", this.Direction)+",\n")
	s = append(s, "Limit: "+fmt.Sprintf("%#v", this.Limit)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LokiData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&queryrange.LokiData{")
	s = append(s, "ResultType: "+fmt.Sprintf("%#v", this.ResultType)+",\n")
	if this.Result != nil {
		vs := make([]*logproto.Stream, len(this.Result))
		for i := range vs {
			vs[i] = &this.Result[i]
		}
		s = append(s, "Result: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQueryrange(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *LokiRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LokiRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LokiRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintQueryrange(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Direction != 0 {
		i = encodeVarintQueryrange(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTs, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTs):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintQueryrange(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTs, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTs):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintQueryrange(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if m.Step != 0 {
		i = encodeVarintQueryrange(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x18
	}
	if m.Limit != 0 {
		i = encodeVarintQueryrange(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintQueryrange(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LokiResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LokiResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LokiResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintQueryrange(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x38
	}
	if m.Limit != 0 {
		i = encodeVarintQueryrange(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x30
	}
	if m.Direction != 0 {
		i = encodeVarintQueryrange(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintQueryrange(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ErrorType) > 0 {
		i -= len(m.ErrorType)
		copy(dAtA[i:], m.ErrorType)
		i = encodeVarintQueryrange(dAtA, i, uint64(len(m.ErrorType)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryrange(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintQueryrange(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LokiData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LokiData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LokiData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		for iNdEx := len(m.Result) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Result[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryrange(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ResultType) > 0 {
		i -= len(m.ResultType)
		copy(dAtA[i:], m.ResultType)
		i = encodeVarintQueryrange(dAtA, i, uint64(len(m.ResultType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryrange(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryrange(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LokiRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovQueryrange(uint64(l))
	}
	if m.Limit != 0 {
		n += 1 + sovQueryrange(uint64(m.Limit))
	}
	if m.Step != 0 {
		n += 1 + sovQueryrange(uint64(m.Step))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTs)
	n += 1 + l + sovQueryrange(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTs)
	n += 1 + l + sovQueryrange(uint64(l))
	if m.Direction != 0 {
		n += 1 + sovQueryrange(uint64(m.Direction))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovQueryrange(uint64(l))
	}
	return n
}

func (m *LokiResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovQueryrange(uint64(l))
	}
	l = m.Data.Size()
	n += 1 + l + sovQueryrange(uint64(l))
	l = len(m.ErrorType)
	if l > 0 {
		n += 1 + l + sovQueryrange(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovQueryrange(uint64(l))
	}
	if m.Direction != 0 {
		n += 1 + sovQueryrange(uint64(m.Direction))
	}
	if m.Limit != 0 {
		n += 1 + sovQueryrange(uint64(m.Limit))
	}
	if m.Version != 0 {
		n += 1 + sovQueryrange(uint64(m.Version))
	}
	return n
}

func (m *LokiData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ResultType)
	if l > 0 {
		n += 1 + l + sovQueryrange(uint64(l))
	}
	if len(m.Result) > 0 {
		for _, e := range m.Result {
			l = e.Size()
			n += 1 + l + sovQueryrange(uint64(l))
		}
	}
	return n
}

func sovQueryrange(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryrange(x uint64) (n int) {
	return sovQueryrange(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LokiRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LokiRequest{`,
		`Query:` + fmt.Sprintf("%v", this.Query) + `,`,
		`Limit:` + fmt.Sprintf("%v", this.Limit) + `,`,
		`Step:` + fmt.Sprintf("%v", this.Step) + `,`,
		`StartTs:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.StartTs), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`EndTs:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.EndTs), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`Direction:` + fmt.Sprintf("%v", this.Direction) + `,`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LokiResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LokiResponse{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Data:` + strings.Replace(strings.Replace(this.Data.String(), "LokiData", "LokiData", 1), `&`, ``, 1) + `,`,
		`ErrorType:` + fmt.Sprintf("%v", this.ErrorType) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`Direction:` + fmt.Sprintf("%v", this.Direction) + `,`,
		`Limit:` + fmt.Sprintf("%v", this.Limit) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LokiData) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForResult := "[]Stream{"
	for _, f := range this.Result {
		repeatedStringForResult += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForResult += "}"
	s := strings.Join([]string{`&LokiData{`,
		`ResultType:` + fmt.Sprintf("%v", this.ResultType) + `,`,
		`Result:` + repeatedStringForResult + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQueryrange(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LokiRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryrange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LokiRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LokiRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTs, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTs, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= logproto.Direction(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryrange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryrange
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueryrange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LokiResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryrange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LokiResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LokiResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= logproto.Direction(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryrange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryrange
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueryrange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LokiData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryrange
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LokiData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LokiData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQueryrange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryrange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result, logproto.Stream{})
			if err := m.Result[len(m.Result)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryrange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQueryrange
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQueryrange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQueryrange(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryrange
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQueryrange
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQueryrange
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQueryrange
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQueryrange
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipQueryrange(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQueryrange
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthQueryrange = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryrange   = fmt.Errorf("proto: integer overflow")
)
