// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eth/v1alpha1/attestation.proto

package eth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Attestation struct {
	AggregationBits      []byte           `protobuf:"bytes,1,opt,name=aggregation_bits,json=aggregationBits,proto3" json:"aggregation_bits,omitempty"`
	Data                 *AttestationData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Signature            []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{0}
}

func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attestation.Unmarshal(m, b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
}
func (m *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Size() int {
	return xxx_messageInfo_Attestation.Size(m)
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

func (m *Attestation) GetAggregationBits() []byte {
	if m != nil {
		return m.AggregationBits
	}
	return nil
}

func (m *Attestation) GetData() *AttestationData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Attestation) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AggregateAttestationAndProof struct {
	AggregatorIndex      uint64       `protobuf:"varint,1,opt,name=aggregator_index,json=aggregatorIndex,proto3" json:"aggregator_index,omitempty"`
	Aggregate            *Attestation `protobuf:"bytes,3,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	SelectionProof       []byte       `protobuf:"bytes,2,opt,name=selection_proof,json=selectionProof,proto3" json:"selection_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AggregateAttestationAndProof) Reset()         { *m = AggregateAttestationAndProof{} }
func (m *AggregateAttestationAndProof) String() string { return proto.CompactTextString(m) }
func (*AggregateAttestationAndProof) ProtoMessage()    {}
func (*AggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{1}
}

func (m *AggregateAttestationAndProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateAttestationAndProof.Unmarshal(m, b)
}
func (m *AggregateAttestationAndProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateAttestationAndProof.Marshal(b, m, deterministic)
}
func (m *AggregateAttestationAndProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateAttestationAndProof.Merge(m, src)
}
func (m *AggregateAttestationAndProof) XXX_Size() int {
	return xxx_messageInfo_AggregateAttestationAndProof.Size(m)
}
func (m *AggregateAttestationAndProof) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateAttestationAndProof.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateAttestationAndProof proto.InternalMessageInfo

func (m *AggregateAttestationAndProof) GetAggregatorIndex() uint64 {
	if m != nil {
		return m.AggregatorIndex
	}
	return 0
}

func (m *AggregateAttestationAndProof) GetAggregate() *Attestation {
	if m != nil {
		return m.Aggregate
	}
	return nil
}

func (m *AggregateAttestationAndProof) GetSelectionProof() []byte {
	if m != nil {
		return m.SelectionProof
	}
	return nil
}

type AttestationData struct {
	Slot                 uint64      `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	CommitteeIndex       uint64      `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty"`
	BeaconBlockRoot      []byte      `protobuf:"bytes,3,opt,name=beacon_block_root,json=beaconBlockRoot,proto3" json:"beacon_block_root,omitempty"`
	Source               *Checkpoint `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Target               *Checkpoint `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AttestationData) Reset()         { *m = AttestationData{} }
func (m *AttestationData) String() string { return proto.CompactTextString(m) }
func (*AttestationData) ProtoMessage()    {}
func (*AttestationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{2}
}

func (m *AttestationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationData.Unmarshal(m, b)
}
func (m *AttestationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationData.Marshal(b, m, deterministic)
}
func (m *AttestationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationData.Merge(m, src)
}
func (m *AttestationData) XXX_Size() int {
	return xxx_messageInfo_AttestationData.Size(m)
}
func (m *AttestationData) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationData.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationData proto.InternalMessageInfo

func (m *AttestationData) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *AttestationData) GetCommitteeIndex() uint64 {
	if m != nil {
		return m.CommitteeIndex
	}
	return 0
}

func (m *AttestationData) GetBeaconBlockRoot() []byte {
	if m != nil {
		return m.BeaconBlockRoot
	}
	return nil
}

func (m *AttestationData) GetSource() *Checkpoint {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AttestationData) GetTarget() *Checkpoint {
	if m != nil {
		return m.Target
	}
	return nil
}

type Crosslink struct {
	Shard                uint64   `protobuf:"varint,1,opt,name=shard,proto3" json:"shard,omitempty"`
	ParentRoot           []byte   `protobuf:"bytes,2,opt,name=parent_root,json=parentRoot,proto3" json:"parent_root,omitempty"`
	StartEpoch           uint64   `protobuf:"varint,3,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch             uint64   `protobuf:"varint,4,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	DataRoot             []byte   `protobuf:"bytes,5,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crosslink) Reset()         { *m = Crosslink{} }
func (m *Crosslink) String() string { return proto.CompactTextString(m) }
func (*Crosslink) ProtoMessage()    {}
func (*Crosslink) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{3}
}

func (m *Crosslink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crosslink.Unmarshal(m, b)
}
func (m *Crosslink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crosslink.Marshal(b, m, deterministic)
}
func (m *Crosslink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crosslink.Merge(m, src)
}
func (m *Crosslink) XXX_Size() int {
	return xxx_messageInfo_Crosslink.Size(m)
}
func (m *Crosslink) XXX_DiscardUnknown() {
	xxx_messageInfo_Crosslink.DiscardUnknown(m)
}

var xxx_messageInfo_Crosslink proto.InternalMessageInfo

func (m *Crosslink) GetShard() uint64 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *Crosslink) GetParentRoot() []byte {
	if m != nil {
		return m.ParentRoot
	}
	return nil
}

func (m *Crosslink) GetStartEpoch() uint64 {
	if m != nil {
		return m.StartEpoch
	}
	return 0
}

func (m *Crosslink) GetEndEpoch() uint64 {
	if m != nil {
		return m.EndEpoch
	}
	return 0
}

func (m *Crosslink) GetDataRoot() []byte {
	if m != nil {
		return m.DataRoot
	}
	return nil
}

type Checkpoint struct {
	Epoch                uint64   `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Root                 []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{4}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Checkpoint) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Attestation)(nil), "ethereum.eth.v1alpha1.Attestation")
	proto.RegisterType((*AggregateAttestationAndProof)(nil), "ethereum.eth.v1alpha1.AggregateAttestationAndProof")
	proto.RegisterType((*AttestationData)(nil), "ethereum.eth.v1alpha1.AttestationData")
	proto.RegisterType((*Crosslink)(nil), "ethereum.eth.v1alpha1.Crosslink")
	proto.RegisterType((*Checkpoint)(nil), "ethereum.eth.v1alpha1.Checkpoint")
}

func init() { proto.RegisterFile("eth/v1alpha1/attestation.proto", fileDescriptor_7894c78397fc93a1) }

var fileDescriptor_7894c78397fc93a1 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0xe0, 0x54, 0xcd, 0xa4, 0x6a, 0xca, 0x8a, 0x4a, 0x41, 0x54, 0xb4, 0xf8, 0x00,
	0x85, 0x83, 0xa3, 0x16, 0xa9, 0x12, 0x70, 0x21, 0x29, 0x39, 0x70, 0x8b, 0x7c, 0x44, 0x95, 0xa2,
	0x8d, 0x3d, 0xd8, 0xab, 0x38, 0x1e, 0x6b, 0x77, 0x82, 0xe0, 0x35, 0xb8, 0x70, 0xe2, 0x05, 0x38,
	0xf3, 0x14, 0x3c, 0x0d, 0x8f, 0x80, 0x76, 0x6d, 0xc7, 0x16, 0x2a, 0xa2, 0x37, 0xcf, 0x3f, 0x3b,
	0xff, 0x7c, 0x33, 0xeb, 0x85, 0xc7, 0xc8, 0xd9, 0xe4, 0xd3, 0x85, 0xcc, 0xcb, 0x4c, 0x5e, 0x4c,
	0x24, 0x33, 0x1a, 0x96, 0xac, 0xa8, 0x08, 0x4b, 0x4d, 0x4c, 0xe2, 0x18, 0x39, 0x43, 0x8d, 0xdb,
	0x4d, 0x88, 0x9c, 0x85, 0xcd, 0xc1, 0xe0, 0xab, 0x07, 0xc3, 0x69, 0x7b, 0x58, 0x3c, 0x87, 0x23,
	0x99, 0xa6, 0x1a, 0x53, 0x17, 0x2e, 0x57, 0x8a, 0xcd, 0xd8, 0x3b, 0xf3, 0xce, 0x0f, 0xa2, 0x51,
	0x47, 0x9f, 0x29, 0x36, 0xe2, 0x35, 0xf8, 0x89, 0x64, 0x39, 0xee, 0x9d, 0x79, 0xe7, 0xc3, 0xcb,
	0xa7, 0xe1, 0xad, 0x0d, 0xc2, 0x8e, 0xf9, 0x3b, 0xc9, 0x32, 0x72, 0x35, 0xe2, 0x04, 0x06, 0x46,
	0xa5, 0x85, 0xe4, 0xad, 0xc6, 0xf1, 0x3d, 0xe7, 0xdf, 0x0a, 0xc1, 0x4f, 0x0f, 0x4e, 0xa6, 0x75,
	0x37, 0xec, 0x18, 0x4c, 0x8b, 0x64, 0xa1, 0x89, 0x3e, 0x76, 0x29, 0x49, 0x2f, 0x55, 0x91, 0xe0,
	0x67, 0x47, 0xe9, 0xb7, 0x94, 0xa4, 0xdf, 0x5b, 0x59, 0xbc, 0x85, 0x41, 0x23, 0x55, 0x9d, 0x86,
	0x97, 0xc1, 0xff, 0x51, 0xa3, 0xb6, 0x48, 0x3c, 0x83, 0x91, 0xc1, 0x1c, 0x63, 0xb7, 0x90, 0xd2,
	0xf6, 0x77, 0x23, 0x1f, 0x44, 0x87, 0x3b, 0xd9, 0x51, 0x05, 0xbf, 0x3d, 0x18, 0xfd, 0x35, 0xae,
	0x10, 0xe0, 0x9b, 0x9c, 0xb8, 0xa6, 0x73, 0xdf, 0xd6, 0x30, 0xa6, 0xcd, 0x46, 0x31, 0x23, 0xd6,
	0xf0, 0x3d, 0x97, 0x3e, 0xdc, 0xc9, 0x15, 0xfb, 0x0b, 0xb8, 0xbf, 0x42, 0x19, 0xdb, 0x7b, 0xc8,
	0x29, 0x5e, 0x2f, 0x35, 0x11, 0xd7, 0xdb, 0x1a, 0x55, 0x89, 0x99, 0xd5, 0x23, 0x22, 0x16, 0xaf,
	0x60, 0xcf, 0xd0, 0x56, 0xc7, 0x38, 0xf6, 0xdd, 0x90, 0x4f, 0xfe, 0x31, 0xe4, 0x75, 0x86, 0xf1,
	0xba, 0x24, 0x55, 0x70, 0x54, 0x17, 0xd8, 0x52, 0x96, 0x3a, 0x45, 0x1e, 0xf7, 0xef, 0x5c, 0x5a,
	0x15, 0x04, 0xdf, 0x3d, 0x18, 0x5c, 0x6b, 0x32, 0x26, 0x57, 0xc5, 0x5a, 0x3c, 0x80, 0xbe, 0xc9,
	0xa4, 0x4e, 0xea, 0x69, 0xab, 0x40, 0x9c, 0xc2, 0xb0, 0x94, 0x1a, 0x0b, 0xae, 0xf8, 0xab, 0xdd,
	0x41, 0x25, 0x39, 0xf4, 0x53, 0x18, 0x1a, 0x96, 0x9a, 0x97, 0x58, 0x52, 0x9c, 0xb9, 0x01, 0xfd,
	0x08, 0x9c, 0x34, 0xb7, 0x8a, 0x78, 0x04, 0x03, 0x2c, 0x92, 0x3a, 0xed, 0xbb, 0xf4, 0x3e, 0x16,
	0xc9, 0x2e, 0x69, 0x7f, 0xa9, 0xca, 0xbc, 0xef, 0xcc, 0xf7, 0xad, 0x60, 0xad, 0x83, 0x2b, 0x80,
	0x96, 0xda, 0xf2, 0x55, 0x1e, 0x35, 0x9f, 0x0b, 0xec, 0x15, 0x75, 0xc0, 0xdc, 0xf7, 0xec, 0x9b,
	0x07, 0x0f, 0x49, 0xa7, 0xb7, 0x2f, 0x62, 0x76, 0xd4, 0xb9, 0xe5, 0x85, 0x7d, 0x5d, 0x0b, 0xef,
	0xc3, 0x55, 0xaa, 0x38, 0xdb, 0xae, 0xc2, 0x98, 0x36, 0x93, 0x52, 0x7f, 0x31, 0x1b, 0xc9, 0x2a,
	0xce, 0xe5, 0xca, 0x4c, 0x1a, 0x0f, 0x59, 0x2a, 0x17, 0xec, 0x5e, 0xe9, 0x1b, 0xe4, 0xec, 0x47,
	0xef, 0x78, 0xde, 0xf4, 0x98, 0x77, 0x7a, 0xfc, 0x6a, 0xf5, 0x9b, 0x39, 0x67, 0x37, 0x8d, 0xbe,
	0xda, 0x73, 0xcf, 0xf9, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xcc, 0x5d, 0x27, 0xf0,
	0x03, 0x00, 0x00,
}
