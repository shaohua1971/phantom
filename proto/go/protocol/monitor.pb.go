// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor.proto

package monitor

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

type MONITOR_MESSAGE_TYPE int32

const (
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_NONE              MONITOR_MESSAGE_TYPE = 0
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_HELLO             MONITOR_MESSAGE_TYPE = 30
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_REGISTER          MONITOR_MESSAGE_TYPE = 31
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_PHANTOM              MONITOR_MESSAGE_TYPE = 32
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_LEDGER            MONITOR_MESSAGE_TYPE = 33
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_SYSTEM            MONITOR_MESSAGE_TYPE = 34
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_ALERT             MONITOR_MESSAGE_TYPE = 35
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_NOTICE            MONITOR_MESSAGE_TYPE = 36
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_ACCOUNT_EXCEPTION MONITOR_MESSAGE_TYPE = 37
	MONITOR_MESSAGE_TYPE_MONITOR_MSGTYPE_ERROR             MONITOR_MESSAGE_TYPE = 39
)

var MONITOR_MESSAGE_TYPE_name = map[int32]string{
	0:  "MONITOR_MSGTYPE_NONE",
	30: "MONITOR_MSGTYPE_HELLO",
	31: "MONITOR_MSGTYPE_REGISTER",
	32: "MONITOR_MSGTYPE_PHANTOM",
	33: "MONITOR_MSGTYPE_LEDGER",
	34: "MONITOR_MSGTYPE_SYSTEM",
	35: "MONITOR_MSGTYPE_ALERT",
	36: "MONITOR_MSGTYPE_NOTICE",
	37: "MONITOR_MSGTYPE_ACCOUNT_EXCEPTION",
	39: "MONITOR_MSGTYPE_ERROR",
}
var MONITOR_MESSAGE_TYPE_value = map[string]int32{
	"MONITOR_MSGTYPE_NONE":              0,
	"MONITOR_MSGTYPE_HELLO":             30,
	"MONITOR_MSGTYPE_REGISTER":          31,
	"MONITOR_MSGTYPE_PHANTOM":              32,
	"MONITOR_MSGTYPE_LEDGER":            33,
	"MONITOR_MSGTYPE_SYSTEM":            34,
	"MONITOR_MSGTYPE_ALERT":             35,
	"MONITOR_MSGTYPE_NOTICE":            36,
	"MONITOR_MSGTYPE_ACCOUNT_EXCEPTION": 37,
	"MONITOR_MSGTYPE_ERROR":             39,
}

func (x MONITOR_MESSAGE_TYPE) String() string {
	return proto.EnumName(MONITOR_MESSAGE_TYPE_name, int32(x))
}
func (MONITOR_MESSAGE_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{0}
}

type Hello struct {
	ServiceVersion       int64    `protobuf:"varint,1,opt,name=service_version,json=serviceVersion" json:"service_version,omitempty"`
	ConnectionTimeout    int64    `protobuf:"varint,2,opt,name=connection_timeout,json=connectionTimeout" json:"connection_timeout,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{0}
}
func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (dst *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(dst, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetServiceVersion() int64 {
	if m != nil {
		return m.ServiceVersion
	}
	return 0
}

func (m *Hello) GetConnectionTimeout() int64 {
	if m != nil {
		return m.ConnectionTimeout
	}
	return 0
}

func (m *Hello) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Register struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	BlockchainVersion    string   `protobuf:"bytes,2,opt,name=blockchain_version,json=blockchainVersion" json:"blockchain_version,omitempty"`
	DataVersion          int64    `protobuf:"varint,3,opt,name=data_version,json=dataVersion" json:"data_version,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{1}
}
func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (dst *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(dst, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Register) GetBlockchainVersion() string {
	if m != nil {
		return m.BlockchainVersion
	}
	return ""
}

func (m *Register) GetDataVersion() int64 {
	if m != nil {
		return m.DataVersion
	}
	return 0
}

func (m *Register) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Peer struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Delay                int64    `protobuf:"varint,2,opt,name=delay" json:"delay,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
	IpAddress            string   `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{2}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (dst *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(dst, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Peer) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *Peer) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Peer) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type GlueManager struct {
	SystemCurrentTime    string   `protobuf:"bytes,1,opt,name=system_current_time,json=systemCurrentTime" json:"system_current_time,omitempty"`
	ProcessUptime        string   `protobuf:"bytes,2,opt,name=process_uptime,json=processUptime" json:"process_uptime,omitempty"`
	SystemUptime         string   `protobuf:"bytes,3,opt,name=system_uptime,json=systemUptime" json:"system_uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlueManager) Reset()         { *m = GlueManager{} }
func (m *GlueManager) String() string { return proto.CompactTextString(m) }
func (*GlueManager) ProtoMessage()    {}
func (*GlueManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{3}
}
func (m *GlueManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlueManager.Unmarshal(m, b)
}
func (m *GlueManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlueManager.Marshal(b, m, deterministic)
}
func (dst *GlueManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlueManager.Merge(dst, src)
}
func (m *GlueManager) XXX_Size() int {
	return xxx_messageInfo_GlueManager.Size(m)
}
func (m *GlueManager) XXX_DiscardUnknown() {
	xxx_messageInfo_GlueManager.DiscardUnknown(m)
}

var xxx_messageInfo_GlueManager proto.InternalMessageInfo

func (m *GlueManager) GetSystemCurrentTime() string {
	if m != nil {
		return m.SystemCurrentTime
	}
	return ""
}

func (m *GlueManager) GetProcessUptime() string {
	if m != nil {
		return m.ProcessUptime
	}
	return ""
}

func (m *GlueManager) GetSystemUptime() string {
	if m != nil {
		return m.SystemUptime
	}
	return ""
}

type PeerManager struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peer_id,json=peerId" json:"peer_id,omitempty"`
	Peer                 []*Peer  `protobuf:"bytes,2,rep,name=peer" json:"peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerManager) Reset()         { *m = PeerManager{} }
func (m *PeerManager) String() string { return proto.CompactTextString(m) }
func (*PeerManager) ProtoMessage()    {}
func (*PeerManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{4}
}
func (m *PeerManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerManager.Unmarshal(m, b)
}
func (m *PeerManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerManager.Marshal(b, m, deterministic)
}
func (dst *PeerManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerManager.Merge(dst, src)
}
func (m *PeerManager) XXX_Size() int {
	return xxx_messageInfo_PeerManager.Size(m)
}
func (m *PeerManager) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerManager.DiscardUnknown(m)
}

var xxx_messageInfo_PeerManager proto.InternalMessageInfo

func (m *PeerManager) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *PeerManager) GetPeer() []*Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type PhantomStatus struct {
	GlueManager          *GlueManager `protobuf:"bytes,1,opt,name=glue_manager,json=glueManager" json:"glue_manager,omitempty"`
	PeerManager          *PeerManager `protobuf:"bytes,2,opt,name=peer_manager,json=peerManager" json:"peer_manager,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PhantomStatus) Reset()         { *m = PhantomStatus{} }
func (m *PhantomStatus) String() string { return proto.CompactTextString(m) }
func (*PhantomStatus) ProtoMessage()    {}
func (*PhantomStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{5}
}
func (m *PhantomStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhantomStatus.Unmarshal(m, b)
}
func (m *PhantomStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhantomStatus.Marshal(b, m, deterministic)
}
func (dst *PhantomStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhantomStatus.Merge(dst, src)
}
func (m *PhantomStatus) XXX_Size() int {
	return xxx_messageInfo_PhantomStatus.Size(m)
}
func (m *PhantomStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PhantomStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PhantomStatus proto.InternalMessageInfo

func (m *PhantomStatus) GetGlueManager() *GlueManager {
	if m != nil {
		return m.GlueManager
	}
	return nil
}

func (m *PhantomStatus) GetPeerManager() *PeerManager {
	if m != nil {
		return m.PeerManager
	}
	return nil
}

type LedgerStatus struct {
	LedgerHeader         *LedgerHeader `protobuf:"bytes,1,opt,name=ledger_header,json=ledgerHeader" json:"ledger_header,omitempty"`
	TransactionSize      int64         `protobuf:"varint,2,opt,name=transaction_size,json=transactionSize" json:"transaction_size,omitempty"`
	AccountCount         int64         `protobuf:"varint,3,opt,name=account_count,json=accountCount" json:"account_count,omitempty"`
	Timestamp            int64         `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LedgerStatus) Reset()         { *m = LedgerStatus{} }
func (m *LedgerStatus) String() string { return proto.CompactTextString(m) }
func (*LedgerStatus) ProtoMessage()    {}
func (*LedgerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{6}
}
func (m *LedgerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerStatus.Unmarshal(m, b)
}
func (m *LedgerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerStatus.Marshal(b, m, deterministic)
}
func (dst *LedgerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerStatus.Merge(dst, src)
}
func (m *LedgerStatus) XXX_Size() int {
	return xxx_messageInfo_LedgerStatus.Size(m)
}
func (m *LedgerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerStatus proto.InternalMessageInfo

func (m *LedgerStatus) GetLedgerHeader() *LedgerHeader {
	if m != nil {
		return m.LedgerHeader
	}
	return nil
}

func (m *LedgerStatus) GetTransactionSize() int64 {
	if m != nil {
		return m.TransactionSize
	}
	return 0
}

func (m *LedgerStatus) GetAccountCount() int64 {
	if m != nil {
		return m.AccountCount
	}
	return 0
}

func (m *LedgerStatus) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type SystemProperty struct {
	HostName             string   `protobuf:"bytes,1,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	OsVersion            string   `protobuf:"bytes,2,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	StartupTime          int64    `protobuf:"varint,3,opt,name=startup_time,json=startupTime" json:"startup_time,omitempty"`
	OsBit                string   `protobuf:"bytes,4,opt,name=os_bit,json=osBit" json:"os_bit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemProperty) Reset()         { *m = SystemProperty{} }
func (m *SystemProperty) String() string { return proto.CompactTextString(m) }
func (*SystemProperty) ProtoMessage()    {}
func (*SystemProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{7}
}
func (m *SystemProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemProperty.Unmarshal(m, b)
}
func (m *SystemProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemProperty.Marshal(b, m, deterministic)
}
func (dst *SystemProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemProperty.Merge(dst, src)
}
func (m *SystemProperty) XXX_Size() int {
	return xxx_messageInfo_SystemProperty.Size(m)
}
func (m *SystemProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemProperty.DiscardUnknown(m)
}

var xxx_messageInfo_SystemProperty proto.InternalMessageInfo

func (m *SystemProperty) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *SystemProperty) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *SystemProperty) GetStartupTime() int64 {
	if m != nil {
		return m.StartupTime
	}
	return 0
}

func (m *SystemProperty) GetOsBit() string {
	if m != nil {
		return m.OsBit
	}
	return ""
}

type SystemResource struct {
	Available            int64    `protobuf:"varint,1,opt,name=available" json:"available,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	UsedPercent          float64  `protobuf:"fixed64,3,opt,name=used_percent,json=usedPercent" json:"used_percent,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemResource) Reset()         { *m = SystemResource{} }
func (m *SystemResource) String() string { return proto.CompactTextString(m) }
func (*SystemResource) ProtoMessage()    {}
func (*SystemResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{8}
}
func (m *SystemResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemResource.Unmarshal(m, b)
}
func (m *SystemResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemResource.Marshal(b, m, deterministic)
}
func (dst *SystemResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemResource.Merge(dst, src)
}
func (m *SystemResource) XXX_Size() int {
	return xxx_messageInfo_SystemResource.Size(m)
}
func (m *SystemResource) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemResource.DiscardUnknown(m)
}

var xxx_messageInfo_SystemResource proto.InternalMessageInfo

func (m *SystemResource) GetAvailable() int64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *SystemResource) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *SystemResource) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

func (m *SystemResource) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Partition struct {
	TotalBytes           int64             `protobuf:"varint,1,opt,name=total_bytes,json=totalBytes" json:"total_bytes,omitempty"`
	Partition            []*SystemResource `protobuf:"bytes,2,rep,name=partition" json:"partition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Partition) Reset()         { *m = Partition{} }
func (m *Partition) String() string { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()    {}
func (*Partition) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{9}
}
func (m *Partition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partition.Unmarshal(m, b)
}
func (m *Partition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partition.Marshal(b, m, deterministic)
}
func (dst *Partition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partition.Merge(dst, src)
}
func (m *Partition) XXX_Size() int {
	return xxx_messageInfo_Partition.Size(m)
}
func (m *Partition) XXX_DiscardUnknown() {
	xxx_messageInfo_Partition.DiscardUnknown(m)
}

var xxx_messageInfo_Partition proto.InternalMessageInfo

func (m *Partition) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func (m *Partition) GetPartition() []*SystemResource {
	if m != nil {
		return m.Partition
	}
	return nil
}

type CPU struct {
	UsedPercent          float64  `protobuf:"fixed64,1,opt,name=used_percent,json=usedPercent" json:"used_percent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{10}
}
func (m *CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPU.Unmarshal(m, b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
}
func (dst *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(dst, src)
}
func (m *CPU) XXX_Size() int {
	return xxx_messageInfo_CPU.Size(m)
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetUsedPercent() float64 {
	if m != nil {
		return m.UsedPercent
	}
	return 0
}

type SystemStatus struct {
	Property             *SystemProperty `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Memory               *SystemResource `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	Cpu                  *CPU            `protobuf:"bytes,3,opt,name=cpu" json:"cpu,omitempty"`
	Partitions           *Partition      `protobuf:"bytes,4,opt,name=partitions" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SystemStatus) Reset()         { *m = SystemStatus{} }
func (m *SystemStatus) String() string { return proto.CompactTextString(m) }
func (*SystemStatus) ProtoMessage()    {}
func (*SystemStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{11}
}
func (m *SystemStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemStatus.Unmarshal(m, b)
}
func (m *SystemStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemStatus.Marshal(b, m, deterministic)
}
func (dst *SystemStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemStatus.Merge(dst, src)
}
func (m *SystemStatus) XXX_Size() int {
	return xxx_messageInfo_SystemStatus.Size(m)
}
func (m *SystemStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SystemStatus proto.InternalMessageInfo

func (m *SystemStatus) GetProperty() *SystemProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *SystemStatus) GetMemory() *SystemResource {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *SystemStatus) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *SystemStatus) GetPartitions() *Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

type AlertStatus struct {
	LedgerSequence       int64         `protobuf:"varint,1,opt,name=ledger_sequence,json=ledgerSequence" json:"ledger_sequence,omitempty"`
	NodeId               string        `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	System               *SystemStatus `protobuf:"bytes,3,opt,name=system" json:"system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AlertStatus) Reset()         { *m = AlertStatus{} }
func (m *AlertStatus) String() string { return proto.CompactTextString(m) }
func (*AlertStatus) ProtoMessage()    {}
func (*AlertStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_2647886f338c3a47, []int{12}
}
func (m *AlertStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertStatus.Unmarshal(m, b)
}
func (m *AlertStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertStatus.Marshal(b, m, deterministic)
}
func (dst *AlertStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertStatus.Merge(dst, src)
}
func (m *AlertStatus) XXX_Size() int {
	return xxx_messageInfo_AlertStatus.Size(m)
}
func (m *AlertStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AlertStatus proto.InternalMessageInfo

func (m *AlertStatus) GetLedgerSequence() int64 {
	if m != nil {
		return m.LedgerSequence
	}
	return 0
}

func (m *AlertStatus) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AlertStatus) GetSystem() *SystemStatus {
	if m != nil {
		return m.System
	}
	return nil
}

func init() {
	proto.RegisterType((*Hello)(nil), "monitor.Hello")
	proto.RegisterType((*Register)(nil), "monitor.Register")
	proto.RegisterType((*Peer)(nil), "monitor.Peer")
	proto.RegisterType((*GlueManager)(nil), "monitor.GlueManager")
	proto.RegisterType((*PeerManager)(nil), "monitor.PeerManager")
	proto.RegisterType((*PhantomStatus)(nil), "monitor.PhantomStatus")
	proto.RegisterType((*LedgerStatus)(nil), "monitor.LedgerStatus")
	proto.RegisterType((*SystemProperty)(nil), "monitor.SystemProperty")
	proto.RegisterType((*SystemResource)(nil), "monitor.SystemResource")
	proto.RegisterType((*Partition)(nil), "monitor.Partition")
	proto.RegisterType((*CPU)(nil), "monitor.CPU")
	proto.RegisterType((*SystemStatus)(nil), "monitor.SystemStatus")
	proto.RegisterType((*AlertStatus)(nil), "monitor.AlertStatus")
	proto.RegisterEnum("monitor.MONITOR_MESSAGE_TYPE", MONITOR_MESSAGE_TYPE_name, MONITOR_MESSAGE_TYPE_value)
}

func init() { proto.RegisterFile("monitor.proto", fileDescriptor_monitor_2647886f338c3a47) }

var fileDescriptor_monitor_2647886f338c3a47 = []byte{
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0x1b, 0x45,
	0x14, 0xc5, 0x76, 0xe2, 0xc6, 0x77, 0x9d, 0xc4, 0x19, 0x92, 0xd4, 0x94, 0xd2, 0x26, 0x5b, 0xa2,
	0x04, 0x24, 0x8c, 0x94, 0x0a, 0xf5, 0x81, 0xa7, 0xc4, 0xac, 0x12, 0x4b, 0x89, 0x6d, 0x8d, 0x1d,
	0x44, 0x9f, 0x56, 0xe3, 0xdd, 0x4b, 0xb2, 0xea, 0x7a, 0x67, 0x99, 0x99, 0x8d, 0x48, 0x25, 0xfa,
	0x88, 0x04, 0xbf, 0x87, 0x17, 0xfe, 0x01, 0x3f, 0x0b, 0xcd, 0x87, 0xd7, 0xae, 0x63, 0xc4, 0xcb,
	0x6a, 0xef, 0xb9, 0x67, 0xe6, 0x9e, 0xb9, 0x73, 0xe6, 0xc2, 0xe6, 0x94, 0x67, 0x89, 0xe2, 0xa2,
	0x93, 0x0b, 0xae, 0x38, 0x79, 0xe2, 0xc2, 0x67, 0x5e, 0x74, 0xc7, 0x92, 0xcc, 0xa2, 0xfe, 0x6f,
	0xb0, 0x7e, 0x89, 0x69, 0xca, 0xc9, 0x31, 0x6c, 0x4b, 0x14, 0xf7, 0x49, 0x84, 0xe1, 0x3d, 0x0a,
	0x99, 0xf0, 0xac, 0x5d, 0x39, 0xa8, 0x9c, 0xd4, 0xe8, 0x96, 0x83, 0x7f, 0xb4, 0x28, 0xf9, 0x06,
	0x48, 0xc4, 0xb3, 0x0c, 0x23, 0x95, 0xf0, 0x2c, 0x54, 0xc9, 0x14, 0x79, 0xa1, 0xda, 0x55, 0xc3,
	0xdd, 0x99, 0x67, 0xc6, 0x36, 0x41, 0x9e, 0x43, 0x43, 0x73, 0xa4, 0x62, 0xd3, 0xbc, 0x5d, 0x33,
	0xac, 0x39, 0xe0, 0xff, 0x59, 0x81, 0x0d, 0x8a, 0xb7, 0x89, 0x54, 0x28, 0xc8, 0x16, 0x54, 0x93,
	0xd8, 0x54, 0x6d, 0xd0, 0x6a, 0x12, 0xeb, 0x4a, 0x93, 0x94, 0x47, 0xef, 0x8c, 0xde, 0x52, 0x55,
	0xd5, 0xe4, 0x77, 0xe6, 0x99, 0x99, 0xb0, 0x43, 0x68, 0xc6, 0x4c, 0xb1, 0x92, 0x68, 0x8b, 0x79,
	0x1a, 0x9b, 0x51, 0x3e, 0x12, 0xb3, 0xb6, 0x2c, 0x26, 0x82, 0xb5, 0x21, 0xae, 0xd0, 0xb1, 0x0b,
	0xeb, 0x31, 0xa6, 0xec, 0xc1, 0x1d, 0xd2, 0x06, 0x64, 0x1f, 0xea, 0x2c, 0x52, 0xc9, 0x3d, 0x9a,
	0x42, 0x1b, 0xd4, 0x45, 0xe4, 0x0b, 0x80, 0x24, 0x0f, 0x59, 0x1c, 0x0b, 0x94, 0xd2, 0x14, 0x69,
	0xd0, 0x46, 0x92, 0x9f, 0x59, 0xc0, 0xff, 0xa3, 0x02, 0xde, 0x45, 0x5a, 0xe0, 0x35, 0xcb, 0xd8,
	0x2d, 0x0a, 0xd2, 0x81, 0x4f, 0xe5, 0x83, 0x54, 0x38, 0x0d, 0xa3, 0x42, 0x08, 0xcc, 0x94, 0x69,
	0xa9, 0xab, 0xbe, 0x63, 0x53, 0x5d, 0x9b, 0xd1, 0x2d, 0x25, 0x47, 0xb0, 0x95, 0x0b, 0x1e, 0xa1,
	0x94, 0x61, 0x91, 0x1b, 0xaa, 0x6d, 0xc8, 0xa6, 0x43, 0x6f, 0x0c, 0x48, 0x5e, 0xc1, 0xa6, 0xdb,
	0xd6, 0xb1, 0x6a, 0x86, 0xd5, 0xb4, 0xa0, 0x25, 0xf9, 0x3d, 0xf0, 0xf4, 0x81, 0x67, 0x52, 0x9e,
	0xc2, 0x93, 0x1c, 0x51, 0x84, 0xe5, 0xe1, 0xeb, 0x3a, 0xec, 0xc5, 0xe4, 0x10, 0xd6, 0xf4, 0x5f,
	0xbb, 0x7a, 0x50, 0x3b, 0xf1, 0x4e, 0x37, 0x3b, 0x33, 0x63, 0xe9, 0xc5, 0xd4, 0xa4, 0xfc, 0x0f,
	0x00, 0xe7, 0xc5, 0x94, 0x8f, 0x14, 0x53, 0x85, 0x24, 0x6f, 0xa0, 0x79, 0x9b, 0x16, 0x18, 0x4e,
	0xed, 0xce, 0x66, 0x3b, 0xef, 0x74, 0xb7, 0x5c, 0xb8, 0xd0, 0x00, 0xea, 0xdd, 0x2e, 0x74, 0xe3,
	0x0d, 0x34, 0x8d, 0x84, 0xd9, 0xc2, 0xea, 0xd2, 0xc2, 0x05, 0xb9, 0xd4, 0xcb, 0xe7, 0x81, 0xff,
	0x77, 0x05, 0x9a, 0x57, 0x18, 0xdf, 0xa2, 0x70, 0x12, 0xbe, 0x87, 0xcd, 0xd4, 0xc4, 0xe1, 0x1d,
	0xb2, 0xb8, 0xd4, 0xb0, 0x6f, 0x7d, 0x1f, 0xf1, 0xb4, 0x63, 0xe9, 0x97, 0x26, 0x4b, 0x9b, 0xe9,
	0x42, 0x44, 0xbe, 0x82, 0x96, 0x12, 0x2c, 0x93, 0xcc, 0x9a, 0x5c, 0x26, 0xef, 0xd1, 0x5d, 0xfe,
	0xf6, 0x02, 0x3e, 0x4a, 0xde, 0x9b, 0x46, 0xb3, 0x28, 0xe2, 0x45, 0xa6, 0x42, 0xf3, 0x75, 0xb6,
	0x6b, 0x3a, 0xb0, 0xab, 0x3f, 0xff, 0xe3, 0xbb, 0xdf, 0x2b, 0xb0, 0x35, 0x32, 0xf7, 0x32, 0x14,
	0x3c, 0x47, 0xa1, 0x1e, 0xc8, 0xe7, 0xd0, 0xb8, 0xe3, 0x52, 0x85, 0x19, 0x2b, 0xbd, 0xb0, 0xa1,
	0x81, 0x3e, 0x9b, 0x1a, 0x87, 0x71, 0xb9, 0xf4, 0x1e, 0x1a, 0x5c, 0x2e, 0xbc, 0x03, 0xa9, 0x98,
	0x50, 0x45, 0x1e, 0x96, 0x37, 0x5f, 0xa3, 0x9e, 0xc3, 0x8c, 0x89, 0xf6, 0xa0, 0xce, 0x65, 0x38,
	0x49, 0x94, 0xf3, 0xe7, 0x3a, 0x97, 0xe7, 0x89, 0xd2, 0xde, 0x74, 0x42, 0x28, 0x4a, 0x5e, 0x88,
	0x08, 0xb5, 0x72, 0x76, 0xcf, 0x92, 0x94, 0x4d, 0x52, 0x74, 0x03, 0x61, 0x0e, 0xe8, 0x97, 0xa1,
	0xb8, 0x62, 0xe9, 0xec, 0x65, 0x98, 0x40, 0x0b, 0x28, 0x24, 0xc6, 0x61, 0x8e, 0x22, 0x42, 0xd7,
	0x91, 0x0a, 0xf5, 0x34, 0x36, 0xb4, 0x10, 0x39, 0x00, 0x2f, 0x46, 0x19, 0x89, 0x24, 0xd7, 0x8d,
	0x74, 0x2a, 0x16, 0x21, 0x3f, 0x82, 0xc6, 0x90, 0x09, 0x95, 0xe8, 0x80, 0xbc, 0x04, 0xcf, 0x6c,
	0x1d, 0x4e, 0x1e, 0x14, 0x4a, 0xa7, 0x03, 0x0c, 0x74, 0xae, 0x11, 0xf2, 0x1d, 0x34, 0xf2, 0x19,
	0xdb, 0xd9, 0xf4, 0x69, 0x69, 0x9a, 0x8f, 0x8f, 0x44, 0xe7, 0x4c, 0xff, 0x04, 0x6a, 0xdd, 0xe1,
	0xcd, 0x23, 0xc1, 0x95, 0x47, 0x82, 0xfd, 0x7f, 0x2a, 0xd0, 0xb4, 0xfb, 0x38, 0x7f, 0xbd, 0x86,
	0x8d, 0xdc, 0xdd, 0x96, 0xb3, 0xd6, 0x72, 0xc1, 0xd9, 0x65, 0xd2, 0x92, 0x48, 0xbe, 0x85, 0xfa,
	0x14, 0xa7, 0x5c, 0x3c, 0x38, 0x63, 0xff, 0xa7, 0x46, 0x47, 0x23, 0x2f, 0xa0, 0x16, 0xe5, 0x85,
	0xe9, 0xa0, 0x77, 0xda, 0x2c, 0xd9, 0xdd, 0xe1, 0x0d, 0xd5, 0x09, 0x72, 0x0a, 0x50, 0x9e, 0xc6,
	0x0e, 0x1b, 0xef, 0x94, 0xcc, 0x5f, 0xcb, 0x2c, 0x45, 0x17, 0x58, 0xfe, 0x07, 0xf0, 0xce, 0x52,
	0x14, 0xca, 0x1d, 0xe4, 0x18, 0xb6, 0xdd, 0x43, 0x91, 0xf8, 0x4b, 0x81, 0x59, 0x34, 0xbb, 0xe7,
	0x2d, 0x0b, 0x8f, 0x1c, 0xaa, 0xc7, 0x43, 0xc6, 0x63, 0xd4, 0xe3, 0xc1, 0x7a, 0xae, 0xae, 0xc3,
	0x9e, 0x9e, 0xd3, 0x75, 0x3b, 0x56, 0x9c, 0xce, 0xbd, 0xa5, 0x53, 0xd9, 0x42, 0xd4, 0x91, 0xbe,
	0xfe, 0xab, 0x0a, 0xbb, 0xd7, 0x83, 0x7e, 0x6f, 0x3c, 0xa0, 0xe1, 0x75, 0x30, 0x1a, 0x9d, 0x5d,
	0x04, 0xe1, 0xf8, 0xed, 0x30, 0x20, 0xed, 0x05, 0x7c, 0x74, 0xa1, 0xa1, 0xb0, 0x3f, 0xe8, 0x07,
	0xad, 0x4f, 0xc8, 0x67, 0xb0, 0xb7, 0x9c, 0xb9, 0x0c, 0xae, 0xae, 0x06, 0xad, 0x17, 0xe4, 0x39,
	0xb4, 0x97, 0x53, 0x34, 0xb8, 0xe8, 0x8d, 0xc6, 0x01, 0x6d, 0xbd, 0x5c, 0xb5, 0xe5, 0xf9, 0xcd,
	0xf5, 0xa0, 0x75, 0x40, 0x9e, 0xc1, 0xfe, 0x72, 0xe6, 0x2a, 0xf8, 0xe1, 0x22, 0xa0, 0xad, 0xc3,
	0x55, 0xb9, 0xd1, 0xdb, 0xd1, 0x38, 0xb8, 0x6e, 0xf9, 0xab, 0xa4, 0x9c, 0x5d, 0x05, 0x74, 0xdc,
	0x7a, 0xb5, 0x6a, 0x59, 0x7f, 0x30, 0xee, 0x75, 0x83, 0xd6, 0x97, 0xe4, 0x08, 0x0e, 0x1f, 0x2d,
	0xeb, 0x76, 0x07, 0x37, 0xfd, 0x71, 0x18, 0xfc, 0xd4, 0x0d, 0x86, 0xe3, 0xde, 0xa0, 0xdf, 0x3a,
	0x5a, 0xb5, 0x7b, 0x40, 0xe9, 0x80, 0xb6, 0x8e, 0xcf, 0x7d, 0x38, 0x48, 0x78, 0x67, 0x52, 0x4c,
	0x79, 0x47, 0xc6, 0xef, 0x3a, 0x11, 0x17, 0xd8, 0xc1, 0x5f, 0x15, 0x66, 0xb1, 0x1d, 0x69, 0x93,
	0xe2, 0xe7, 0x49, 0xdd, 0xfc, 0xbd, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x20, 0x5e, 0x62, 0xb6,
	0xfb, 0x07, 0x00, 0x00,
}
