// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage/types.proto

package storage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type State struct {
	ElectronReady        bool                      `protobuf:"varint,1,opt,name=electronReady,proto3" json:"electronReady,omitempty"`
	ElectronPath         string                    `protobuf:"bytes,2,opt,name=electronPath,proto3" json:"electronPath,omitempty"`
	DownloadList         map[string]*DownloadState `protobuf:"bytes,3,rep,name=downloadList,proto3" json:"downloadList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserStorageMap       map[string]*UserStorage   `protobuf:"bytes,4,rep,name=userStorageMap,proto3" json:"userStorageMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UserStorageId        string                    `protobuf:"bytes,5,opt,name=userStorageId,proto3" json:"userStorageId,omitempty"`
	Token                string                    `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	Settings             *StateSetting             `protobuf:"bytes,7,opt,name=settings,proto3" json:"settings,omitempty"`
	PanCookie            []*Cookies                `protobuf:"bytes,8,rep,name=panCookie,proto3" json:"panCookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_04d23f1c63b4f870, []int{0}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetElectronReady() bool {
	if m != nil {
		return m.ElectronReady
	}
	return false
}

func (m *State) GetElectronPath() string {
	if m != nil {
		return m.ElectronPath
	}
	return ""
}

func (m *State) GetDownloadList() map[string]*DownloadState {
	if m != nil {
		return m.DownloadList
	}
	return nil
}

func (m *State) GetUserStorageMap() map[string]*UserStorage {
	if m != nil {
		return m.UserStorageMap
	}
	return nil
}

func (m *State) GetUserStorageId() string {
	if m != nil {
		return m.UserStorageId
	}
	return ""
}

func (m *State) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *State) GetSettings() *StateSetting {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *State) GetPanCookie() []*Cookies {
	if m != nil {
		return m.PanCookie
	}
	return nil
}

type StateSetting struct {
	DownloadSegSize      int64    `protobuf:"varint,1,opt,name=downloadSegSize,proto3" json:"downloadSegSize,omitempty"`
	DownloadCoroutine    int64    `protobuf:"varint,2,opt,name=downloadCoroutine,proto3" json:"downloadCoroutine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateSetting) Reset()         { *m = StateSetting{} }
func (m *StateSetting) String() string { return proto.CompactTextString(m) }
func (*StateSetting) ProtoMessage()    {}
func (*StateSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_04d23f1c63b4f870, []int{0, 2}
}
func (m *StateSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateSetting.Unmarshal(m, b)
}
func (m *StateSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateSetting.Marshal(b, m, deterministic)
}
func (dst *StateSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateSetting.Merge(dst, src)
}
func (m *StateSetting) XXX_Size() int {
	return xxx_messageInfo_StateSetting.Size(m)
}
func (m *StateSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_StateSetting.DiscardUnknown(m)
}

var xxx_messageInfo_StateSetting proto.InternalMessageInfo

func (m *StateSetting) GetDownloadSegSize() int64 {
	if m != nil {
		return m.DownloadSegSize
	}
	return 0
}

func (m *StateSetting) GetDownloadCoroutine() int64 {
	if m != nil {
		return m.DownloadCoroutine
	}
	return 0
}

type DownloadState struct {
	Fid                  string         `protobuf:"bytes,1,opt,name=fid,proto3" json:"fid,omitempty"`
	Filename             string         `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	TempName             string         `protobuf:"bytes,3,opt,name=tempName,proto3" json:"tempName,omitempty"`
	LocalPath            string         `protobuf:"bytes,4,opt,name=localPath,proto3" json:"localPath,omitempty"`
	Size                 int64          `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Seg                  []*DownloadSeg `protobuf:"bytes,6,rep,name=seg,proto3" json:"seg,omitempty"`
	Downloading          bool           `protobuf:"varint,7,opt,name=downloading,proto3" json:"downloading,omitempty"`
	Finished             bool           `protobuf:"varint,8,opt,name=finished,proto3" json:"finished,omitempty"`
	Link                 string         `protobuf:"bytes,9,opt,name=link,proto3" json:"link,omitempty"`
	UserAgent            string         `protobuf:"bytes,10,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	Stopping             bool           `protobuf:"varint,11,opt,name=stopping,proto3" json:"stopping,omitempty"`
	FinishedSize         int64          `protobuf:"varint,12,opt,name=finishedSize,proto3" json:"finishedSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DownloadState) Reset()         { *m = DownloadState{} }
func (m *DownloadState) String() string { return proto.CompactTextString(m) }
func (*DownloadState) ProtoMessage()    {}
func (*DownloadState) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_04d23f1c63b4f870, []int{1}
}
func (m *DownloadState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadState.Unmarshal(m, b)
}
func (m *DownloadState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadState.Marshal(b, m, deterministic)
}
func (dst *DownloadState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadState.Merge(dst, src)
}
func (m *DownloadState) XXX_Size() int {
	return xxx_messageInfo_DownloadState.Size(m)
}
func (m *DownloadState) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadState.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadState proto.InternalMessageInfo

func (m *DownloadState) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func (m *DownloadState) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *DownloadState) GetTempName() string {
	if m != nil {
		return m.TempName
	}
	return ""
}

func (m *DownloadState) GetLocalPath() string {
	if m != nil {
		return m.LocalPath
	}
	return ""
}

func (m *DownloadState) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *DownloadState) GetSeg() []*DownloadSeg {
	if m != nil {
		return m.Seg
	}
	return nil
}

func (m *DownloadState) GetDownloading() bool {
	if m != nil {
		return m.Downloading
	}
	return false
}

func (m *DownloadState) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

func (m *DownloadState) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *DownloadState) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *DownloadState) GetStopping() bool {
	if m != nil {
		return m.Stopping
	}
	return false
}

func (m *DownloadState) GetFinishedSize() int64 {
	if m != nil {
		return m.FinishedSize
	}
	return 0
}

type DownloadSeg struct {
	Start                int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Len                  int64    `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	Finish               bool     `protobuf:"varint,3,opt,name=finish,proto3" json:"finish,omitempty"`
	Distributed          bool     `protobuf:"varint,4,opt,name=distributed,proto3" json:"distributed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadSeg) Reset()         { *m = DownloadSeg{} }
func (m *DownloadSeg) String() string { return proto.CompactTextString(m) }
func (*DownloadSeg) ProtoMessage()    {}
func (*DownloadSeg) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_04d23f1c63b4f870, []int{2}
}
func (m *DownloadSeg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadSeg.Unmarshal(m, b)
}
func (m *DownloadSeg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadSeg.Marshal(b, m, deterministic)
}
func (dst *DownloadSeg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadSeg.Merge(dst, src)
}
func (m *DownloadSeg) XXX_Size() int {
	return xxx_messageInfo_DownloadSeg.Size(m)
}
func (m *DownloadSeg) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadSeg.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadSeg proto.InternalMessageInfo

func (m *DownloadSeg) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DownloadSeg) GetLen() int64 {
	if m != nil {
		return m.Len
	}
	return 0
}

func (m *DownloadSeg) GetFinish() bool {
	if m != nil {
		return m.Finish
	}
	return false
}

func (m *DownloadSeg) GetDistributed() bool {
	if m != nil {
		return m.Distributed
	}
	return false
}

type UserStorage struct {
	DataBucket           map[string]string `protobuf:"bytes,1,rep,name=DataBucket,proto3" json:"DataBucket,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserStorage) Reset()         { *m = UserStorage{} }
func (m *UserStorage) String() string { return proto.CompactTextString(m) }
func (*UserStorage) ProtoMessage()    {}
func (*UserStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_04d23f1c63b4f870, []int{3}
}
func (m *UserStorage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserStorage.Unmarshal(m, b)
}
func (m *UserStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserStorage.Marshal(b, m, deterministic)
}
func (dst *UserStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserStorage.Merge(dst, src)
}
func (m *UserStorage) XXX_Size() int {
	return xxx_messageInfo_UserStorage.Size(m)
}
func (m *UserStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserStorage.DiscardUnknown(m)
}

var xxx_messageInfo_UserStorage proto.InternalMessageInfo

func (m *UserStorage) GetDataBucket() map[string]string {
	if m != nil {
		return m.DataBucket
	}
	return nil
}

type Cookies struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Domain               string   `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Expires              int64    `protobuf:"varint,4,opt,name=expires,proto3" json:"expires,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cookies) Reset()         { *m = Cookies{} }
func (m *Cookies) String() string { return proto.CompactTextString(m) }
func (*Cookies) ProtoMessage()    {}
func (*Cookies) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_04d23f1c63b4f870, []int{4}
}
func (m *Cookies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookies.Unmarshal(m, b)
}
func (m *Cookies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookies.Marshal(b, m, deterministic)
}
func (dst *Cookies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookies.Merge(dst, src)
}
func (m *Cookies) XXX_Size() int {
	return xxx_messageInfo_Cookies.Size(m)
}
func (m *Cookies) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookies.DiscardUnknown(m)
}

var xxx_messageInfo_Cookies proto.InternalMessageInfo

func (m *Cookies) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Cookies) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Cookies) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Cookies) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*State)(nil), "storage.state")
	proto.RegisterMapType((map[string]*DownloadState)(nil), "storage.state.DownloadListEntry")
	proto.RegisterMapType((map[string]*UserStorage)(nil), "storage.state.UserStorageMapEntry")
	proto.RegisterType((*StateSetting)(nil), "storage.state.setting")
	proto.RegisterType((*DownloadState)(nil), "storage.downloadState")
	proto.RegisterType((*DownloadSeg)(nil), "storage.downloadSeg")
	proto.RegisterType((*UserStorage)(nil), "storage.userStorage")
	proto.RegisterMapType((map[string]string)(nil), "storage.userStorage.DataBucketEntry")
	proto.RegisterType((*Cookies)(nil), "storage.cookies")
}

func init() { proto.RegisterFile("storage/types.proto", fileDescriptor_types_04d23f1c63b4f870) }

var fileDescriptor_types_04d23f1c63b4f870 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0xea, 0x26, 0x71, 0x26, 0x2d, 0x6d, 0xb7, 0x51, 0xb5, 0x8a, 0x38, 0x58, 0x56, 0x85,
	0x22, 0x54, 0x19, 0xa9, 0x5c, 0x10, 0x12, 0x07, 0x68, 0x38, 0x80, 0x00, 0xa1, 0x8d, 0x50, 0xcf,
	0xdb, 0x78, 0xea, 0xae, 0xe2, 0xee, 0x5a, 0xde, 0x0d, 0x50, 0x1e, 0x83, 0x13, 0xef, 0xc7, 0x8b,
	0xa0, 0xdd, 0xb5, 0x1d, 0x3b, 0xcd, 0x81, 0x9b, 0xe7, 0x9b, 0x99, 0x6f, 0xbe, 0xf9, 0x59, 0xc3,
	0xa9, 0x36, 0xaa, 0xe4, 0x19, 0xbe, 0x30, 0x0f, 0x05, 0xea, 0xa4, 0x28, 0x95, 0x51, 0x64, 0x58,
	0x81, 0xf1, 0x9f, 0x3e, 0xf4, 0xb5, 0xe1, 0x06, 0xc9, 0x39, 0x1c, 0x62, 0x8e, 0x4b, 0x53, 0x2a,
	0xc9, 0x90, 0xa7, 0x0f, 0xb4, 0x17, 0xf5, 0x66, 0x21, 0xeb, 0x82, 0x24, 0x86, 0x83, 0x1a, 0xf8,
	0xca, 0xcd, 0x1d, 0xdd, 0x8b, 0x7a, 0xb3, 0x11, 0xeb, 0x60, 0x64, 0x0e, 0x07, 0xa9, 0xfa, 0x21,
	0x73, 0xc5, 0xd3, 0x4f, 0x42, 0x1b, 0x1a, 0x44, 0xc1, 0x6c, 0x7c, 0x19, 0x25, 0x55, 0xcd, 0xc4,
	0xd5, 0x4b, 0xe6, 0xad, 0x90, 0xf7, 0xd2, 0x94, 0x0f, 0xac, 0x93, 0x45, 0x3e, 0xc2, 0x93, 0xb5,
	0xc6, 0x72, 0xe1, 0x93, 0x3e, 0xf3, 0x82, 0xee, 0x3b, 0x9e, 0x78, 0x8b, 0xe7, 0x5b, 0x27, 0xc8,
	0x33, 0x6d, 0x65, 0xda, 0xde, 0x5a, 0xc8, 0x87, 0x94, 0xf6, 0x9d, 0xec, 0x2e, 0x48, 0x26, 0xd0,
	0x37, 0x6a, 0x85, 0x92, 0x0e, 0x9c, 0xd7, 0x1b, 0xe4, 0x12, 0x42, 0x8d, 0xc6, 0x08, 0x99, 0x69,
	0x3a, 0x8c, 0x7a, 0xb3, 0xf1, 0xe5, 0xd9, 0x96, 0x82, 0xca, 0xcd, 0x9a, 0x38, 0x92, 0xc0, 0xa8,
	0xe0, 0xf2, 0x4a, 0xa9, 0x95, 0x40, 0x1a, 0x3a, 0xd9, 0xc7, 0x4d, 0xd2, 0xd2, 0xc1, 0x9a, 0x6d,
	0x42, 0xa6, 0xd7, 0x70, 0xf2, 0x68, 0x1c, 0xe4, 0x18, 0x82, 0x15, 0xfa, 0x35, 0x8c, 0x98, 0xfd,
	0x24, 0x17, 0xd0, 0xff, 0xce, 0xf3, 0x35, 0xba, 0xa9, 0xb7, 0x75, 0xd4, 0x83, 0x5b, 0x58, 0x3d,
	0xcc, 0x07, 0xbd, 0xde, 0x7b, 0xd5, 0x9b, 0x5e, 0xc3, 0xe9, 0x8e, 0xf9, 0xec, 0xa0, 0x7e, 0xde,
	0xa5, 0x9e, 0x34, 0xd4, 0xad, 0x11, 0xb5, 0x89, 0x39, 0x0c, 0xab, 0x6e, 0xc9, 0x0c, 0x8e, 0x9a,
	0xfa, 0x98, 0x2d, 0xc4, 0x2f, 0x74, 0xc4, 0x01, 0xdb, 0x86, 0xc9, 0x05, 0x9c, 0xd4, 0xd0, 0x95,
	0x2a, 0xd5, 0xda, 0x08, 0xe9, 0x0b, 0x06, 0xec, 0xb1, 0x23, 0xfe, 0xbb, 0x07, 0x87, 0x9d, 0xc6,
	0xac, 0xec, 0x5b, 0x91, 0xd6, 0xb2, 0x6f, 0x45, 0x4a, 0xa6, 0x10, 0xde, 0x8a, 0x1c, 0x25, 0xbf,
	0xc7, 0xea, 0x14, 0x1b, 0xdb, 0xfa, 0x0c, 0xde, 0x17, 0x5f, 0xac, 0x2f, 0xf0, 0xbe, 0xda, 0x26,
	0x4f, 0x61, 0x94, 0xab, 0x25, 0xcf, 0xdd, 0x0d, 0xef, 0x3b, 0xe7, 0x06, 0x20, 0x04, 0xf6, 0xb5,
	0x6d, 0xa3, 0xef, 0xa4, 0xb9, 0x6f, 0xf2, 0x0c, 0x02, 0x8d, 0x19, 0x1d, 0xb8, 0x65, 0x4e, 0x1e,
	0x4f, 0x1e, 0x33, 0x66, 0x03, 0x48, 0x04, 0xe3, 0x1a, 0x13, 0x32, 0x73, 0x17, 0x13, 0xb2, 0x36,
	0xe4, 0x35, 0x4b, 0xa1, 0xef, 0x30, 0xa5, 0xa1, 0x73, 0x37, 0xb6, 0xad, 0x9c, 0x0b, 0xb9, 0xa2,
	0x23, 0x27, 0xc9, 0x7d, 0x5b, 0xad, 0x76, 0x09, 0x6f, 0x33, 0x94, 0x86, 0x82, 0xd7, 0xda, 0x00,
	0x96, 0x4d, 0x1b, 0x55, 0x14, 0xb6, 0xd8, 0xd8, 0xb3, 0xd5, 0xb6, 0x7d, 0xac, 0x35, 0xb3, 0x5b,
	0xcb, 0x81, 0xeb, 0xa7, 0x83, 0xc5, 0x6a, 0xa3, 0x77, 0x81, 0x99, 0x7d, 0x03, 0xda, 0xf0, 0xd2,
	0x54, 0x2b, 0xf4, 0x86, 0x1d, 0x7c, 0x8e, 0xb2, 0x5a, 0x95, 0xfd, 0x24, 0x67, 0x30, 0xf0, 0x34,
	0x6e, 0xb4, 0x21, 0xab, 0x2c, 0xd7, 0xbe, 0xd0, 0xa6, 0x14, 0x37, 0x6b, 0x83, 0xa9, 0x1b, 0xad,
	0x6d, 0x7f, 0x03, 0xc5, 0xbf, 0x7b, 0x30, 0x6e, 0x1d, 0x15, 0x99, 0x03, 0xcc, 0xb9, 0xe1, 0xef,
	0xd6, 0xcb, 0x15, 0xda, 0xb2, 0x76, 0xbe, 0xe7, 0xbb, 0xce, 0x2f, 0xd9, 0x84, 0xf9, 0x57, 0xde,
	0xca, 0x9b, 0xbe, 0x81, 0xa3, 0x2d, 0xf7, 0x8e, 0x23, 0x9f, 0xb4, 0x8f, 0x7c, 0xd4, 0x3a, 0xe7,
	0x78, 0x09, 0xc3, 0xea, 0x59, 0xfe, 0x6f, 0x9a, 0x9d, 0x40, 0xaa, 0xee, 0xb9, 0x90, 0xd5, 0x71,
	0x55, 0x16, 0xa1, 0x30, 0xc4, 0x9f, 0x85, 0x28, 0x51, 0xbb, 0xee, 0x03, 0x56, 0x9b, 0x37, 0x03,
	0xf7, 0xef, 0x7d, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x33, 0x56, 0xd5, 0xca, 0x92, 0x05, 0x00,
	0x00,
}
