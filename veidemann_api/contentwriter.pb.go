// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contentwriter.proto

package veidemann_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RecordType int32

const (
	RecordType_WARCINFO     RecordType = 0
	RecordType_RESPONSE     RecordType = 1
	RecordType_RESOURCE     RecordType = 2
	RecordType_REQUEST      RecordType = 3
	RecordType_METADATA     RecordType = 4
	RecordType_REVISIT      RecordType = 5
	RecordType_CONVERSION   RecordType = 6
	RecordType_CONTINUATION RecordType = 7
)

var RecordType_name = map[int32]string{
	0: "WARCINFO",
	1: "RESPONSE",
	2: "RESOURCE",
	3: "REQUEST",
	4: "METADATA",
	5: "REVISIT",
	6: "CONVERSION",
	7: "CONTINUATION",
}

var RecordType_value = map[string]int32{
	"WARCINFO":     0,
	"RESPONSE":     1,
	"RESOURCE":     2,
	"REQUEST":      3,
	"METADATA":     4,
	"REVISIT":      5,
	"CONVERSION":   6,
	"CONTINUATION": 7,
}

func (x RecordType) String() string {
	return proto.EnumName(RecordType_name, int32(x))
}

func (RecordType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e59044a96414a19, []int{0}
}

type Data struct {
	RecordNum            int32    `protobuf:"varint,1,opt,name=record_num,json=recordNum,proto3" json:"record_num,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e59044a96414a19, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetRecordNum() int32 {
	if m != nil {
		return m.RecordNum
	}
	return 0
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WriteRequestMeta struct {
	ExecutionId    string                                 `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	TargetUri      string                                 `protobuf:"bytes,2,opt,name=target_uri,json=targetUri,proto3" json:"target_uri,omitempty"`
	RecordMeta     map[int32]*WriteRequestMeta_RecordMeta `protobuf:"bytes,3,rep,name=record_meta,json=recordMeta,proto3" json:"record_meta,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FetchTimeStamp *timestamp.Timestamp                   `protobuf:"bytes,4,opt,name=fetch_time_stamp,json=fetchTimeStamp,proto3" json:"fetch_time_stamp,omitempty"`
	IpAddress      string                                 `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// The http status code. Only relevant for response records.
	StatusCode           int32    `protobuf:"varint,6,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequestMeta) Reset()         { *m = WriteRequestMeta{} }
func (m *WriteRequestMeta) String() string { return proto.CompactTextString(m) }
func (*WriteRequestMeta) ProtoMessage()    {}
func (*WriteRequestMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e59044a96414a19, []int{1}
}

func (m *WriteRequestMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequestMeta.Unmarshal(m, b)
}
func (m *WriteRequestMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequestMeta.Marshal(b, m, deterministic)
}
func (m *WriteRequestMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequestMeta.Merge(m, src)
}
func (m *WriteRequestMeta) XXX_Size() int {
	return xxx_messageInfo_WriteRequestMeta.Size(m)
}
func (m *WriteRequestMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequestMeta.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequestMeta proto.InternalMessageInfo

func (m *WriteRequestMeta) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *WriteRequestMeta) GetTargetUri() string {
	if m != nil {
		return m.TargetUri
	}
	return ""
}

func (m *WriteRequestMeta) GetRecordMeta() map[int32]*WriteRequestMeta_RecordMeta {
	if m != nil {
		return m.RecordMeta
	}
	return nil
}

func (m *WriteRequestMeta) GetFetchTimeStamp() *timestamp.Timestamp {
	if m != nil {
		return m.FetchTimeStamp
	}
	return nil
}

func (m *WriteRequestMeta) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *WriteRequestMeta) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

type WriteRequestMeta_RecordMeta struct {
	RecordNum            int32      `protobuf:"varint,1,opt,name=record_num,json=recordNum,proto3" json:"record_num,omitempty"`
	Type                 RecordType `protobuf:"varint,2,opt,name=type,proto3,enum=veidemann.api.RecordType" json:"type,omitempty"`
	RecordContentType    string     `protobuf:"bytes,3,opt,name=record_content_type,json=recordContentType,proto3" json:"record_content_type,omitempty"`
	PayloadContentType   string     `protobuf:"bytes,4,opt,name=payload_content_type,json=payloadContentType,proto3" json:"payload_content_type,omitempty"`
	BlockDigest          string     `protobuf:"bytes,5,opt,name=block_digest,json=blockDigest,proto3" json:"block_digest,omitempty"`
	PayloadDigest        string     `protobuf:"bytes,6,opt,name=payload_digest,json=payloadDigest,proto3" json:"payload_digest,omitempty"`
	Size                 int64      `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	WarcRefersTo         string     `protobuf:"bytes,8,opt,name=warc_refers_to,json=warcRefersTo,proto3" json:"warc_refers_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WriteRequestMeta_RecordMeta) Reset()         { *m = WriteRequestMeta_RecordMeta{} }
func (m *WriteRequestMeta_RecordMeta) String() string { return proto.CompactTextString(m) }
func (*WriteRequestMeta_RecordMeta) ProtoMessage()    {}
func (*WriteRequestMeta_RecordMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e59044a96414a19, []int{1, 0}
}

func (m *WriteRequestMeta_RecordMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequestMeta_RecordMeta.Unmarshal(m, b)
}
func (m *WriteRequestMeta_RecordMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequestMeta_RecordMeta.Marshal(b, m, deterministic)
}
func (m *WriteRequestMeta_RecordMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequestMeta_RecordMeta.Merge(m, src)
}
func (m *WriteRequestMeta_RecordMeta) XXX_Size() int {
	return xxx_messageInfo_WriteRequestMeta_RecordMeta.Size(m)
}
func (m *WriteRequestMeta_RecordMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequestMeta_RecordMeta.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequestMeta_RecordMeta proto.InternalMessageInfo

func (m *WriteRequestMeta_RecordMeta) GetRecordNum() int32 {
	if m != nil {
		return m.RecordNum
	}
	return 0
}

func (m *WriteRequestMeta_RecordMeta) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_WARCINFO
}

func (m *WriteRequestMeta_RecordMeta) GetRecordContentType() string {
	if m != nil {
		return m.RecordContentType
	}
	return ""
}

func (m *WriteRequestMeta_RecordMeta) GetPayloadContentType() string {
	if m != nil {
		return m.PayloadContentType
	}
	return ""
}

func (m *WriteRequestMeta_RecordMeta) GetBlockDigest() string {
	if m != nil {
		return m.BlockDigest
	}
	return ""
}

func (m *WriteRequestMeta_RecordMeta) GetPayloadDigest() string {
	if m != nil {
		return m.PayloadDigest
	}
	return ""
}

func (m *WriteRequestMeta_RecordMeta) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *WriteRequestMeta_RecordMeta) GetWarcRefersTo() string {
	if m != nil {
		return m.WarcRefersTo
	}
	return ""
}

type WriteRequest struct {
	// Types that are valid to be assigned to Value:
	//	*WriteRequest_Meta
	//	*WriteRequest_Header
	//	*WriteRequest_Payload
	//	*WriteRequest_Cancel
	Value                isWriteRequest_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e59044a96414a19, []int{2}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

type isWriteRequest_Value interface {
	isWriteRequest_Value()
}

type WriteRequest_Meta struct {
	Meta *WriteRequestMeta `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type WriteRequest_Header struct {
	Header *Data `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

type WriteRequest_Payload struct {
	Payload *Data `protobuf:"bytes,3,opt,name=payload,proto3,oneof"`
}

type WriteRequest_Cancel struct {
	Cancel string `protobuf:"bytes,4,opt,name=cancel,proto3,oneof"`
}

func (*WriteRequest_Meta) isWriteRequest_Value() {}

func (*WriteRequest_Header) isWriteRequest_Value() {}

func (*WriteRequest_Payload) isWriteRequest_Value() {}

func (*WriteRequest_Cancel) isWriteRequest_Value() {}

func (m *WriteRequest) GetValue() isWriteRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *WriteRequest) GetMeta() *WriteRequestMeta {
	if x, ok := m.GetValue().(*WriteRequest_Meta); ok {
		return x.Meta
	}
	return nil
}

func (m *WriteRequest) GetHeader() *Data {
	if x, ok := m.GetValue().(*WriteRequest_Header); ok {
		return x.Header
	}
	return nil
}

func (m *WriteRequest) GetPayload() *Data {
	if x, ok := m.GetValue().(*WriteRequest_Payload); ok {
		return x.Payload
	}
	return nil
}

func (m *WriteRequest) GetCancel() string {
	if x, ok := m.GetValue().(*WriteRequest_Cancel); ok {
		return x.Cancel
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WriteRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WriteRequest_OneofMarshaler, _WriteRequest_OneofUnmarshaler, _WriteRequest_OneofSizer, []interface{}{
		(*WriteRequest_Meta)(nil),
		(*WriteRequest_Header)(nil),
		(*WriteRequest_Payload)(nil),
		(*WriteRequest_Cancel)(nil),
	}
}

func _WriteRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WriteRequest)
	// value
	switch x := m.Value.(type) {
	case *WriteRequest_Meta:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Meta); err != nil {
			return err
		}
	case *WriteRequest_Header:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Header); err != nil {
			return err
		}
	case *WriteRequest_Payload:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Payload); err != nil {
			return err
		}
	case *WriteRequest_Cancel:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Cancel)
	case nil:
	default:
		return fmt.Errorf("WriteRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _WriteRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WriteRequest)
	switch tag {
	case 1: // value.meta
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WriteRequestMeta)
		err := b.DecodeMessage(msg)
		m.Value = &WriteRequest_Meta{msg}
		return true, err
	case 2: // value.header
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Data)
		err := b.DecodeMessage(msg)
		m.Value = &WriteRequest_Header{msg}
		return true, err
	case 3: // value.payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Data)
		err := b.DecodeMessage(msg)
		m.Value = &WriteRequest_Payload{msg}
		return true, err
	case 4: // value.cancel
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &WriteRequest_Cancel{x}
		return true, err
	default:
		return false, nil
	}
}

func _WriteRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WriteRequest)
	// value
	switch x := m.Value.(type) {
	case *WriteRequest_Meta:
		s := proto.Size(x.Meta)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WriteRequest_Header:
		s := proto.Size(x.Header)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WriteRequest_Payload:
		s := proto.Size(x.Payload)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WriteRequest_Cancel:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Cancel)))
		n += len(x.Cancel)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type WriteResponseMeta struct {
	RecordMeta           map[int32]*WriteResponseMeta_RecordMeta `protobuf:"bytes,1,rep,name=record_meta,json=recordMeta,proto3" json:"record_meta,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *WriteResponseMeta) Reset()         { *m = WriteResponseMeta{} }
func (m *WriteResponseMeta) String() string { return proto.CompactTextString(m) }
func (*WriteResponseMeta) ProtoMessage()    {}
func (*WriteResponseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e59044a96414a19, []int{3}
}

func (m *WriteResponseMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponseMeta.Unmarshal(m, b)
}
func (m *WriteResponseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponseMeta.Marshal(b, m, deterministic)
}
func (m *WriteResponseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponseMeta.Merge(m, src)
}
func (m *WriteResponseMeta) XXX_Size() int {
	return xxx_messageInfo_WriteResponseMeta.Size(m)
}
func (m *WriteResponseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponseMeta proto.InternalMessageInfo

func (m *WriteResponseMeta) GetRecordMeta() map[int32]*WriteResponseMeta_RecordMeta {
	if m != nil {
		return m.RecordMeta
	}
	return nil
}

type WriteResponseMeta_RecordMeta struct {
	RecordNum            int32      `protobuf:"varint,1,opt,name=record_num,json=recordNum,proto3" json:"record_num,omitempty"`
	Type                 RecordType `protobuf:"varint,2,opt,name=type,proto3,enum=veidemann.api.RecordType" json:"type,omitempty"`
	WarcId               string     `protobuf:"bytes,3,opt,name=warc_id,json=warcId,proto3" json:"warc_id,omitempty"`
	StorageRef           string     `protobuf:"bytes,4,opt,name=storage_ref,json=storageRef,proto3" json:"storage_ref,omitempty"`
	BlockDigest          string     `protobuf:"bytes,5,opt,name=block_digest,json=blockDigest,proto3" json:"block_digest,omitempty"`
	PayloadDigest        string     `protobuf:"bytes,6,opt,name=payload_digest,json=payloadDigest,proto3" json:"payload_digest,omitempty"`
	WarcRefersTo         string     `protobuf:"bytes,7,opt,name=warc_refers_to,json=warcRefersTo,proto3" json:"warc_refers_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *WriteResponseMeta_RecordMeta) Reset()         { *m = WriteResponseMeta_RecordMeta{} }
func (m *WriteResponseMeta_RecordMeta) String() string { return proto.CompactTextString(m) }
func (*WriteResponseMeta_RecordMeta) ProtoMessage()    {}
func (*WriteResponseMeta_RecordMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e59044a96414a19, []int{3, 0}
}

func (m *WriteResponseMeta_RecordMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponseMeta_RecordMeta.Unmarshal(m, b)
}
func (m *WriteResponseMeta_RecordMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponseMeta_RecordMeta.Marshal(b, m, deterministic)
}
func (m *WriteResponseMeta_RecordMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponseMeta_RecordMeta.Merge(m, src)
}
func (m *WriteResponseMeta_RecordMeta) XXX_Size() int {
	return xxx_messageInfo_WriteResponseMeta_RecordMeta.Size(m)
}
func (m *WriteResponseMeta_RecordMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponseMeta_RecordMeta.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponseMeta_RecordMeta proto.InternalMessageInfo

func (m *WriteResponseMeta_RecordMeta) GetRecordNum() int32 {
	if m != nil {
		return m.RecordNum
	}
	return 0
}

func (m *WriteResponseMeta_RecordMeta) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_WARCINFO
}

func (m *WriteResponseMeta_RecordMeta) GetWarcId() string {
	if m != nil {
		return m.WarcId
	}
	return ""
}

func (m *WriteResponseMeta_RecordMeta) GetStorageRef() string {
	if m != nil {
		return m.StorageRef
	}
	return ""
}

func (m *WriteResponseMeta_RecordMeta) GetBlockDigest() string {
	if m != nil {
		return m.BlockDigest
	}
	return ""
}

func (m *WriteResponseMeta_RecordMeta) GetPayloadDigest() string {
	if m != nil {
		return m.PayloadDigest
	}
	return ""
}

func (m *WriteResponseMeta_RecordMeta) GetWarcRefersTo() string {
	if m != nil {
		return m.WarcRefersTo
	}
	return ""
}

type WriteReply struct {
	Meta                 *WriteResponseMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *WriteReply) Reset()         { *m = WriteReply{} }
func (m *WriteReply) String() string { return proto.CompactTextString(m) }
func (*WriteReply) ProtoMessage()    {}
func (*WriteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e59044a96414a19, []int{4}
}

func (m *WriteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteReply.Unmarshal(m, b)
}
func (m *WriteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteReply.Marshal(b, m, deterministic)
}
func (m *WriteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteReply.Merge(m, src)
}
func (m *WriteReply) XXX_Size() int {
	return xxx_messageInfo_WriteReply.Size(m)
}
func (m *WriteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteReply.DiscardUnknown(m)
}

var xxx_messageInfo_WriteReply proto.InternalMessageInfo

func (m *WriteReply) GetMeta() *WriteResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterEnum("veidemann.api.RecordType", RecordType_name, RecordType_value)
	proto.RegisterType((*Data)(nil), "veidemann.api.Data")
	proto.RegisterType((*WriteRequestMeta)(nil), "veidemann.api.WriteRequestMeta")
	proto.RegisterMapType((map[int32]*WriteRequestMeta_RecordMeta)(nil), "veidemann.api.WriteRequestMeta.RecordMetaEntry")
	proto.RegisterType((*WriteRequestMeta_RecordMeta)(nil), "veidemann.api.WriteRequestMeta.RecordMeta")
	proto.RegisterType((*WriteRequest)(nil), "veidemann.api.WriteRequest")
	proto.RegisterType((*WriteResponseMeta)(nil), "veidemann.api.WriteResponseMeta")
	proto.RegisterMapType((map[int32]*WriteResponseMeta_RecordMeta)(nil), "veidemann.api.WriteResponseMeta.RecordMetaEntry")
	proto.RegisterType((*WriteResponseMeta_RecordMeta)(nil), "veidemann.api.WriteResponseMeta.RecordMeta")
	proto.RegisterType((*WriteReply)(nil), "veidemann.api.WriteReply")
}

func init() { proto.RegisterFile("contentwriter.proto", fileDescriptor_1e59044a96414a19) }

var fileDescriptor_1e59044a96414a19 = []byte{
	// 862 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x15, 0xad, 0xaf, 0x78, 0x24, 0xab, 0xcc, 0xba, 0x68, 0x18, 0x05, 0x85, 0x15, 0xa1, 0x05,
	0x84, 0x14, 0xa1, 0x03, 0xb5, 0x05, 0x9a, 0xa0, 0x87, 0xca, 0x32, 0x8b, 0xe8, 0x10, 0xc9, 0x59,
	0xd1, 0x09, 0xd0, 0x0b, 0xb1, 0x26, 0x47, 0x32, 0x1b, 0x8a, 0x64, 0xc9, 0x65, 0x52, 0xf6, 0xd4,
	0x43, 0x7f, 0x58, 0x8f, 0x05, 0xfa, 0x0b, 0xda, 0x5f, 0x53, 0xec, 0x87, 0x6a, 0x4b, 0x71, 0xe2,
	0x18, 0x68, 0x6f, 0xbb, 0x6f, 0xe6, 0x0d, 0x67, 0xdf, 0xbe, 0xe1, 0xc2, 0xbe, 0x9f, 0xc4, 0x1c,
	0x63, 0xfe, 0x26, 0x0b, 0x39, 0x66, 0x76, 0x9a, 0x25, 0x3c, 0x21, 0x7b, 0xaf, 0x31, 0x0c, 0x70,
	0xc5, 0xe2, 0xd8, 0x66, 0x69, 0xd8, 0x3d, 0x58, 0x26, 0xc9, 0x32, 0xc2, 0x43, 0x19, 0x3c, 0x2b,
	0x16, 0x87, 0x3c, 0x5c, 0x61, 0xce, 0xd9, 0x2a, 0x55, 0xf9, 0xdd, 0x7b, 0xdb, 0x09, 0xb8, 0x4a,
	0x79, 0xa9, 0x82, 0xfd, 0xc7, 0x50, 0x3b, 0x66, 0x9c, 0x91, 0x4f, 0x01, 0x32, 0xf4, 0x93, 0x2c,
	0xf0, 0xe2, 0x62, 0x65, 0x19, 0x3d, 0x63, 0x50, 0xa7, 0xbb, 0x0a, 0x99, 0x16, 0x2b, 0x42, 0xa0,
	0x16, 0x30, 0xce, 0xac, 0x9d, 0x9e, 0x31, 0x68, 0x53, 0xb9, 0xee, 0xff, 0x5d, 0x07, 0xf3, 0xa5,
	0x68, 0x8c, 0xe2, 0x4f, 0x05, 0xe6, 0xfc, 0x19, 0x72, 0x46, 0xee, 0x43, 0x1b, 0x7f, 0x46, 0xbf,
	0xe0, 0x61, 0x12, 0x7b, 0x61, 0x20, 0x2b, 0xed, 0xd2, 0xd6, 0xbf, 0xd8, 0x24, 0x10, 0x9f, 0xe2,
	0x2c, 0x5b, 0x22, 0xf7, 0x8a, 0x2c, 0x94, 0x15, 0x77, 0xe9, 0xae, 0x42, 0x4e, 0xb3, 0x90, 0x9c,
	0x40, 0x4b, 0x77, 0xb2, 0x42, 0xce, 0xac, 0x6a, 0xaf, 0x3a, 0x68, 0x0d, 0x0f, 0xed, 0x8d, 0x43,
	0xdb, 0xdb, 0xdf, 0xb5, 0xa9, 0xa4, 0x88, 0xa5, 0x13, 0xf3, 0xac, 0xa4, 0xfa, 0x34, 0xb2, 0xa7,
	0x63, 0x30, 0x17, 0xc8, 0xfd, 0x73, 0x4f, 0x28, 0xe3, 0x49, 0x69, 0xac, 0x5a, 0xcf, 0x18, 0xb4,
	0x86, 0x5d, 0x5b, 0x69, 0x63, 0xaf, 0xb5, 0xb1, 0xdd, 0xb5, 0x78, 0xb4, 0x23, 0x39, 0x62, 0x3f,
	0x17, 0x7b, 0xd1, 0x76, 0x98, 0x7a, 0x2c, 0x08, 0x32, 0xcc, 0x73, 0xab, 0xae, 0xda, 0x0e, 0xd3,
	0x91, 0x02, 0xc8, 0x01, 0xb4, 0x72, 0xce, 0x78, 0x91, 0x7b, 0x7e, 0x12, 0xa0, 0xd5, 0x90, 0x0a,
	0x82, 0x82, 0xc6, 0x49, 0x80, 0xdd, 0xdf, 0x77, 0x00, 0x2e, 0xba, 0xbc, 0x4e, 0xf0, 0x87, 0x50,
	0xe3, 0x65, 0x8a, 0x52, 0x9e, 0xce, 0xf0, 0xee, 0xd6, 0xf1, 0x55, 0x1d, 0xb7, 0x4c, 0x91, 0xca,
	0x34, 0x62, 0xc3, 0xbe, 0xae, 0xa6, 0x1d, 0xe3, 0x49, 0x76, 0x55, 0x76, 0x79, 0x5b, 0x85, 0xc6,
	0x2a, 0x22, 0x58, 0xe4, 0x11, 0x7c, 0x9c, 0xb2, 0x32, 0x4a, 0xd8, 0x16, 0xa1, 0x26, 0x09, 0x44,
	0xc7, 0x2e, 0x33, 0xee, 0x43, 0xfb, 0x2c, 0x4a, 0xfc, 0x57, 0x5e, 0x10, 0x2e, 0x31, 0xe7, 0x5a,
	0x80, 0x96, 0xc4, 0x8e, 0x25, 0x44, 0x3e, 0x87, 0xce, 0xba, 0xa8, 0x4e, 0x6a, 0xc8, 0xa4, 0x3d,
	0x8d, 0xea, 0x34, 0x02, 0xb5, 0x3c, 0xfc, 0x05, 0xad, 0x66, 0xcf, 0x18, 0x54, 0xa9, 0x5c, 0x93,
	0xcf, 0xa0, 0xf3, 0x86, 0x65, 0xbe, 0x97, 0xe1, 0x02, 0xb3, 0xdc, 0xe3, 0x89, 0x75, 0x4b, 0x52,
	0xdb, 0x02, 0xa5, 0x12, 0x74, 0x93, 0x6e, 0x08, 0x1f, 0x6d, 0xdd, 0x33, 0x31, 0xa1, 0xfa, 0x0a,
	0x4b, 0xad, 0x9f, 0x58, 0x92, 0xef, 0xa0, 0xfe, 0x9a, 0x45, 0x85, 0x92, 0xae, 0x35, 0x7c, 0xf0,
	0xe1, 0xce, 0xa1, 0x8a, 0xf8, 0x64, 0xe7, 0x1b, 0xa3, 0xff, 0x87, 0x01, 0xed, 0xcb, 0xa9, 0xe4,
	0x6b, 0xa8, 0x49, 0x3f, 0x1a, 0xb2, 0xea, 0xc1, 0x35, 0x55, 0x9f, 0x56, 0xa8, 0x4c, 0x27, 0x0f,
	0xa1, 0x71, 0x8e, 0x2c, 0xc0, 0x4c, 0xb7, 0xb3, 0xbf, 0x45, 0x14, 0xc3, 0xf7, 0xb4, 0x42, 0x75,
	0x12, 0x39, 0x84, 0xa6, 0x16, 0x4b, 0xde, 0xdd, 0x3b, 0xf3, 0xd7, 0x59, 0xc4, 0x82, 0x86, 0xcf,
	0x62, 0x1f, 0x23, 0x75, 0x75, 0xa2, 0x94, 0xda, 0x1f, 0x35, 0xb5, 0x0e, 0xfd, 0xbf, 0xaa, 0x70,
	0x5b, 0xf7, 0x97, 0xa7, 0x49, 0x9c, 0xa3, 0xf4, 0xdf, 0xf3, 0xcd, 0x31, 0x33, 0xe4, 0x98, 0x3d,
	0xba, 0xfa, 0x58, 0x17, 0xb4, 0xf7, 0xcd, 0x59, 0xf7, 0xb7, 0xff, 0xd3, 0xe1, 0x77, 0xa0, 0x29,
	0x1d, 0x12, 0x06, 0xda, 0xd5, 0x0d, 0xb1, 0x9d, 0x04, 0x6a, 0xf0, 0x92, 0x8c, 0x2d, 0x51, 0xb8,
	0x47, 0x3b, 0x18, 0x34, 0x44, 0x71, 0xf1, 0x1f, 0x3a, 0xf7, 0x6d, 0x97, 0x36, 0xaf, 0x70, 0xe9,
	0x8f, 0x1f, 0xe2, 0xd2, 0xd1, 0xa6, 0x4b, 0xbf, 0xb8, 0x81, 0xf0, 0x97, 0x6d, 0x7a, 0x04, 0xa0,
	0x53, 0xd3, 0xa8, 0x24, 0x5f, 0x6d, 0x78, 0xb4, 0x77, 0x5d, 0x4d, 0x65, 0xd1, 0x07, 0xbf, 0x1a,
	0xeb, 0x6b, 0x93, 0x83, 0xde, 0x86, 0x5b, 0x2f, 0x47, 0x74, 0x3c, 0x99, 0x7e, 0x3f, 0x33, 0x2b,
	0x62, 0x47, 0x9d, 0xf9, 0xc9, 0x6c, 0x3a, 0x77, 0x4c, 0x43, 0xef, 0x66, 0xa7, 0x74, 0xec, 0x98,
	0x3b, 0xa4, 0x05, 0x4d, 0xea, 0x3c, 0x3f, 0x75, 0xe6, 0xae, 0x59, 0x15, 0xa1, 0x67, 0x8e, 0x3b,
	0x3a, 0x1e, 0xb9, 0x23, 0xb3, 0xa6, 0x42, 0x2f, 0x26, 0xf3, 0x89, 0x6b, 0xd6, 0x49, 0x07, 0x60,
	0x3c, 0x9b, 0xbe, 0x70, 0xe8, 0x7c, 0x32, 0x9b, 0x9a, 0x0d, 0x62, 0x42, 0x7b, 0x3c, 0x9b, 0xba,
	0x93, 0xe9, 0xe9, 0xc8, 0x15, 0x48, 0x73, 0xf8, 0xa7, 0x01, 0x7b, 0xfa, 0x67, 0x23, 0xbb, 0xcc,
	0xc8, 0x18, 0xea, 0xf2, 0xd1, 0x23, 0xf7, 0xde, 0x33, 0x69, 0xdd, 0xbb, 0x57, 0x07, 0xd3, 0xa8,
	0xec, 0x57, 0x06, 0x06, 0x79, 0x0c, 0xf5, 0x45, 0x54, 0xe4, 0xe7, 0xe4, 0x93, 0xb7, 0xfe, 0xf3,
	0x8e, 0x78, 0x03, 0xbb, 0xef, 0xc0, 0xfb, 0x15, 0xf2, 0x04, 0x1a, 0x01, 0x46, 0xc8, 0xf1, 0xe6,
	0xdc, 0xa3, 0x6f, 0xe1, 0x4e, 0x9c, 0xd8, 0xf1, 0x99, 0x1d, 0xc7, 0x6c, 0xb3, 0xc1, 0x23, 0xb2,
	0x71, 0xca, 0x13, 0xc1, 0xfc, 0xe1, 0xe2, 0x3d, 0xf7, 0x58, 0x1a, 0x9e, 0x35, 0x64, 0xbd, 0x2f,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x60, 0x1c, 0xf5, 0xfc, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContentWriterClient is the client API for ContentWriter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentWriterClient interface {
	Write(ctx context.Context, opts ...grpc.CallOption) (ContentWriter_WriteClient, error)
	Flush(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type contentWriterClient struct {
	cc *grpc.ClientConn
}

func NewContentWriterClient(cc *grpc.ClientConn) ContentWriterClient {
	return &contentWriterClient{cc}
}

func (c *contentWriterClient) Write(ctx context.Context, opts ...grpc.CallOption) (ContentWriter_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ContentWriter_serviceDesc.Streams[0], "/veidemann.api.ContentWriter/write", opts...)
	if err != nil {
		return nil, err
	}
	x := &contentWriterWriteClient{stream}
	return x, nil
}

type ContentWriter_WriteClient interface {
	Send(*WriteRequest) error
	CloseAndRecv() (*WriteReply, error)
	grpc.ClientStream
}

type contentWriterWriteClient struct {
	grpc.ClientStream
}

func (x *contentWriterWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *contentWriterWriteClient) CloseAndRecv() (*WriteReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *contentWriterClient) Flush(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/veidemann.api.ContentWriter/flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentWriterClient) Delete(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/veidemann.api.ContentWriter/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentWriterServer is the server API for ContentWriter service.
type ContentWriterServer interface {
	Write(ContentWriter_WriteServer) error
	Flush(context.Context, *empty.Empty) (*empty.Empty, error)
	Delete(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterContentWriterServer(s *grpc.Server, srv ContentWriterServer) {
	s.RegisterService(&_ContentWriter_serviceDesc, srv)
}

func _ContentWriter_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ContentWriterServer).Write(&contentWriterWriteServer{stream})
}

type ContentWriter_WriteServer interface {
	SendAndClose(*WriteReply) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type contentWriterWriteServer struct {
	grpc.ServerStream
}

func (x *contentWriterWriteServer) SendAndClose(m *WriteReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *contentWriterWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ContentWriter_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentWriterServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.ContentWriter/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentWriterServer).Flush(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentWriter_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentWriterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.ContentWriter/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentWriterServer).Delete(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentWriter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.ContentWriter",
	HandlerType: (*ContentWriterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "flush",
			Handler:    _ContentWriter_Flush_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ContentWriter_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "write",
			Handler:       _ContentWriter_Write_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "contentwriter.proto",
}
