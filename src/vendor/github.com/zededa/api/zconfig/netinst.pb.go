// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netinst.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ZNetworkInstType int32

const (
	ZNetworkInstType_ZNetInstFirst       ZNetworkInstType = 0
	ZNetworkInstType_ZnetInstSwitch      ZNetworkInstType = 1
	ZNetworkInstType_ZnetInstLocal       ZNetworkInstType = 2
	ZNetworkInstType_ZnetInstCloud       ZNetworkInstType = 3
	ZNetworkInstType_ZnetInstMesh        ZNetworkInstType = 4
	ZNetworkInstType_ZnetInstHoneyPot    ZNetworkInstType = 5
	ZNetworkInstType_ZnetInstTransparent ZNetworkInstType = 6
	ZNetworkInstType_ZNetInstLast        ZNetworkInstType = 255
)

var ZNetworkInstType_name = map[int32]string{
	0:   "ZNetInstFirst",
	1:   "ZnetInstSwitch",
	2:   "ZnetInstLocal",
	3:   "ZnetInstCloud",
	4:   "ZnetInstMesh",
	5:   "ZnetInstHoneyPot",
	6:   "ZnetInstTransparent",
	255: "ZNetInstLast",
}
var ZNetworkInstType_value = map[string]int32{
	"ZNetInstFirst":       0,
	"ZnetInstSwitch":      1,
	"ZnetInstLocal":       2,
	"ZnetInstCloud":       3,
	"ZnetInstMesh":        4,
	"ZnetInstHoneyPot":    5,
	"ZnetInstTransparent": 6,
	"ZNetInstLast":        255,
}

func (x ZNetworkInstType) String() string {
	return proto.EnumName(ZNetworkInstType_name, int32(x))
}
func (ZNetworkInstType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type AddressType int32

const (
	AddressType_First      AddressType = 0
	AddressType_IPV4       AddressType = 1
	AddressType_IPV6       AddressType = 2
	AddressType_CryptoIPV4 AddressType = 3
	AddressType_CryptoIPV6 AddressType = 4
	AddressType_Last       AddressType = 255
)

var AddressType_name = map[int32]string{
	0:   "First",
	1:   "IPV4",
	2:   "IPV6",
	3:   "CryptoIPV4",
	4:   "CryptoIPV6",
	255: "Last",
}
var AddressType_value = map[string]int32{
	"First":      0,
	"IPV4":       1,
	"IPV6":       2,
	"CryptoIPV4": 3,
	"CryptoIPV6": 4,
	"Last":       255,
}

func (x AddressType) String() string {
	return proto.EnumName(AddressType_name, int32(x))
}
func (AddressType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

// Network Instance Opaque config. In future we might add more fields here
// but idea is here. This is service specific configuration.
type NetworkInstanceOpaqueConfig struct {
	Oconfig string `protobuf:"bytes,1,opt,name=oconfig" json:"oconfig,omitempty"`
}

func (m *NetworkInstanceOpaqueConfig) Reset()                    { *m = NetworkInstanceOpaqueConfig{} }
func (m *NetworkInstanceOpaqueConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkInstanceOpaqueConfig) ProtoMessage()               {}
func (*NetworkInstanceOpaqueConfig) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *NetworkInstanceOpaqueConfig) GetOconfig() string {
	if m != nil {
		return m.Oconfig
	}
	return ""
}

// Lisp NetworkInstance config
type NetworkInstanceLispConfig struct {
	LispMSs             []*ZcServicePoint `protobuf:"bytes,1,rep,name=LispMSs" json:"LispMSs,omitempty"`
	LispInstanceId      uint32            `protobuf:"varint,2,opt,name=LispInstanceId" json:"LispInstanceId,omitempty"`
	Allocate            bool              `protobuf:"varint,3,opt,name=allocate" json:"allocate,omitempty"`
	Exportprivate       bool              `protobuf:"varint,4,opt,name=exportprivate" json:"exportprivate,omitempty"`
	Allocationprefix    []byte            `protobuf:"bytes,5,opt,name=allocationprefix,proto3" json:"allocationprefix,omitempty"`
	Allocationprefixlen uint32            `protobuf:"varint,6,opt,name=allocationprefixlen" json:"allocationprefixlen,omitempty"`
	// various configuration to dataPlane, lispers.net vs Zededa
	Experimental bool `protobuf:"varint,20,opt,name=experimental" json:"experimental,omitempty"`
}

func (m *NetworkInstanceLispConfig) Reset()                    { *m = NetworkInstanceLispConfig{} }
func (m *NetworkInstanceLispConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkInstanceLispConfig) ProtoMessage()               {}
func (*NetworkInstanceLispConfig) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *NetworkInstanceLispConfig) GetLispMSs() []*ZcServicePoint {
	if m != nil {
		return m.LispMSs
	}
	return nil
}

func (m *NetworkInstanceLispConfig) GetLispInstanceId() uint32 {
	if m != nil {
		return m.LispInstanceId
	}
	return 0
}

func (m *NetworkInstanceLispConfig) GetAllocate() bool {
	if m != nil {
		return m.Allocate
	}
	return false
}

func (m *NetworkInstanceLispConfig) GetExportprivate() bool {
	if m != nil {
		return m.Exportprivate
	}
	return false
}

func (m *NetworkInstanceLispConfig) GetAllocationprefix() []byte {
	if m != nil {
		return m.Allocationprefix
	}
	return nil
}

func (m *NetworkInstanceLispConfig) GetAllocationprefixlen() uint32 {
	if m != nil {
		return m.Allocationprefixlen
	}
	return 0
}

func (m *NetworkInstanceLispConfig) GetExperimental() bool {
	if m != nil {
		return m.Experimental
	}
	return false
}

type NetworkInstanceConfig struct {
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion" json:"uuidandversion,omitempty"`
	Displayname    string          `protobuf:"bytes,2,opt,name=displayname" json:"displayname,omitempty"`
	// instType - Type of network instance ( local, bridge etc )
	InstType ZNetworkInstType `protobuf:"varint,4,opt,name=instType,enum=ZNetworkInstType" json:"instType,omitempty"`
	// activate
	//  - True by default. If set to false ( inactivate), the network instance
	//    configuration is downloaded to the device, but the network instance
	//    itself is not created on the device.
	Activate bool `protobuf:"varint,5,opt,name=activate" json:"activate,omitempty"`
	// port - Only a single port is supported.
	//    This is used as the external connection for the network instance.
	//    This can be a physical (eth0 ) or logical port (vlan 0).
	//    The port name comes from DeviceConfig ( When it is supported in future).
	//    If the user needs multiple physical ports, Device config should be
	//    used to create a label for multiple physical ports.
	Port *Adapter `protobuf:"bytes,20,opt,name=port" json:"port,omitempty"`
	// cfg - Used to pass some feature-specific configuration to the
	//       network instance. For Ex: Lisp, StriongSwan etc
	Cfg *NetworkInstanceOpaqueConfig `protobuf:"bytes,30,opt,name=cfg" json:"cfg,omitempty"`
	// type of ipSpec
	IpType AddressType `protobuf:"varint,39,opt,name=ipType,enum=AddressType" json:"ipType,omitempty"`
	// network ip specification
	Ip *Ipspec `protobuf:"bytes,40,opt,name=ip" json:"ip,omitempty"`
	// static DNS entry, if we are running DNS/DHCP service
	Dns []*ZnetStaticDNSEntry `protobuf:"bytes,41,rep,name=dns" json:"dns,omitempty"`
}

func (m *NetworkInstanceConfig) Reset()                    { *m = NetworkInstanceConfig{} }
func (m *NetworkInstanceConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkInstanceConfig) ProtoMessage()               {}
func (*NetworkInstanceConfig) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *NetworkInstanceConfig) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *NetworkInstanceConfig) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *NetworkInstanceConfig) GetInstType() ZNetworkInstType {
	if m != nil {
		return m.InstType
	}
	return ZNetworkInstType_ZNetInstFirst
}

func (m *NetworkInstanceConfig) GetActivate() bool {
	if m != nil {
		return m.Activate
	}
	return false
}

func (m *NetworkInstanceConfig) GetPort() *Adapter {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *NetworkInstanceConfig) GetCfg() *NetworkInstanceOpaqueConfig {
	if m != nil {
		return m.Cfg
	}
	return nil
}

func (m *NetworkInstanceConfig) GetIpType() AddressType {
	if m != nil {
		return m.IpType
	}
	return AddressType_First
}

func (m *NetworkInstanceConfig) GetIp() *Ipspec {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *NetworkInstanceConfig) GetDns() []*ZnetStaticDNSEntry {
	if m != nil {
		return m.Dns
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkInstanceOpaqueConfig)(nil), "NetworkInstanceOpaqueConfig")
	proto.RegisterType((*NetworkInstanceLispConfig)(nil), "NetworkInstanceLispConfig")
	proto.RegisterType((*NetworkInstanceConfig)(nil), "NetworkInstanceConfig")
	proto.RegisterEnum("ZNetworkInstType", ZNetworkInstType_name, ZNetworkInstType_value)
	proto.RegisterEnum("AddressType", AddressType_name, AddressType_value)
}

func init() { proto.RegisterFile("netinst.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdd, 0x4e, 0x1b, 0x39,
	0x1c, 0xc5, 0x77, 0xf2, 0x45, 0xf2, 0xcf, 0x07, 0xc6, 0xb0, 0x62, 0x96, 0x65, 0xb7, 0x51, 0x44,
	0xdb, 0x80, 0xd4, 0xa1, 0xa2, 0x15, 0xdc, 0x96, 0x42, 0xab, 0x46, 0x02, 0x8a, 0x26, 0x40, 0xa5,
	0xdc, 0x19, 0x8f, 0x01, 0xab, 0x89, 0xed, 0xda, 0x4e, 0x20, 0x3c, 0x56, 0xfb, 0x64, 0x7d, 0x82,
	0x56, 0xf6, 0xcc, 0xa0, 0x90, 0x56, 0xbd, 0xf3, 0xff, 0x77, 0x8e, 0xad, 0xe3, 0x23, 0xcf, 0x40,
	0x53, 0x30, 0xcb, 0x85, 0xb1, 0x91, 0xd2, 0xd2, 0xca, 0xb5, 0xc5, 0x84, 0x4d, 0xa8, 0x1c, 0x8d,
	0xa4, 0xc8, 0x40, 0x43, 0x30, 0x4b, 0x47, 0xd9, 0xd4, 0xd9, 0x83, 0x7f, 0x4f, 0x98, 0xbd, 0x95,
	0xfa, 0x73, 0x4f, 0x18, 0x4b, 0x04, 0x65, 0x1f, 0x15, 0xf9, 0x32, 0x66, 0x07, 0x52, 0x5c, 0xf1,
	0x6b, 0x1c, 0xc2, 0x82, 0xa4, 0x7e, 0x19, 0x06, 0xed, 0xa0, 0x5b, 0x8b, 0xf3, 0xb1, 0xf3, 0xb5,
	0x00, 0xff, 0xcc, 0xed, 0x3c, 0xe2, 0x46, 0x65, 0xfb, 0x36, 0x61, 0xc1, 0x4d, 0xc7, 0x7d, 0x13,
	0x06, 0xed, 0x62, 0xb7, 0xbe, 0xb3, 0x18, 0x0d, 0x68, 0x9f, 0xe9, 0x09, 0xa7, 0xec, 0x54, 0x72,
	0x61, 0xe3, 0x5c, 0xc7, 0xcf, 0xa0, 0xe5, 0x96, 0xf9, 0x21, 0xbd, 0x24, 0x2c, 0xb4, 0x83, 0x6e,
	0x33, 0x9e, 0xa3, 0x78, 0x0d, 0xaa, 0x64, 0x38, 0x94, 0x94, 0x58, 0x16, 0x16, 0xdb, 0x41, 0xb7,
	0x1a, 0x3f, 0xcc, 0x78, 0x03, 0x9a, 0xec, 0x4e, 0x49, 0x6d, 0x95, 0xe6, 0x13, 0x67, 0x28, 0x79,
	0xc3, 0x63, 0x88, 0xb7, 0x00, 0x65, 0x3b, 0xb8, 0x14, 0x4a, 0xb3, 0x2b, 0x7e, 0x17, 0x96, 0xdb,
	0x41, 0xb7, 0x11, 0xff, 0xc2, 0xf1, 0x4b, 0x58, 0x9e, 0x67, 0x43, 0x26, 0xc2, 0x8a, 0x8f, 0xf6,
	0x3b, 0x09, 0x77, 0xa0, 0xc1, 0xee, 0x14, 0xd3, 0x7c, 0xc4, 0x84, 0x25, 0xc3, 0x70, 0xc5, 0x47,
	0x78, 0xc4, 0x3a, 0xdf, 0x0b, 0xf0, 0xf7, 0x5c, 0x69, 0x59, 0x61, 0x7b, 0xd0, 0x1a, 0x8f, 0x79,
	0x42, 0x44, 0x32, 0x61, 0xda, 0x70, 0x29, 0x7c, 0xdf, 0xae, 0xb7, 0xf3, 0xf3, 0xde, 0x21, 0x11,
	0xc9, 0x45, 0x8a, 0xe3, 0x39, 0x1b, 0x6e, 0x43, 0x3d, 0xe1, 0x46, 0x0d, 0xc9, 0x54, 0x90, 0x11,
	0xf3, 0xdd, 0xd5, 0xe2, 0x59, 0x84, 0x5f, 0x40, 0xd5, 0xbd, 0x87, 0xb3, 0xa9, 0x4a, 0x7b, 0x69,
	0xed, 0x2c, 0x45, 0x83, 0x99, 0x14, 0x4e, 0x88, 0x1f, 0x2c, 0xbe, 0x67, 0x6a, 0xd3, 0x1a, 0xcb,
	0x59, 0xcf, 0xd9, 0x8c, 0xd7, 0xa1, 0xe4, 0x0a, 0xf5, 0x77, 0xab, 0xef, 0x54, 0xa3, 0xfd, 0x84,
	0x28, 0xcb, 0x74, 0xec, 0x29, 0x8e, 0xa0, 0x48, 0xaf, 0xae, 0xc3, 0xff, 0xbd, 0xb8, 0x1e, 0xfd,
	0xe1, 0x5d, 0xc5, 0xce, 0x88, 0x37, 0xa0, 0xc2, 0x95, 0x8f, 0xf5, 0xdc, 0xc7, 0x6a, 0x44, 0xfb,
	0x49, 0xa2, 0x99, 0x31, 0x3e, 0x51, 0xa6, 0xe1, 0x55, 0x28, 0x70, 0x15, 0x76, 0xfd, 0xa1, 0x0b,
	0x11, 0x57, 0x46, 0x31, 0x1a, 0x17, 0xb8, 0xc2, 0x4f, 0xa1, 0x98, 0x08, 0x13, 0x6e, 0xfa, 0xf7,
	0xb5, 0x1c, 0x0d, 0x04, 0xb3, 0x7d, 0x4b, 0x2c, 0xa7, 0x87, 0x27, 0xfd, 0x77, 0xc2, 0xea, 0x69,
	0xec, 0xf4, 0xad, 0x6f, 0x01, 0xa0, 0xf9, 0xeb, 0xe2, 0x25, 0x68, 0x3a, 0xe6, 0xe6, 0xf7, 0x5c,
	0x1b, 0x8b, 0xfe, 0xc2, 0x18, 0x5a, 0xee, 0x08, 0x87, 0xfa, 0xb7, 0xdc, 0xd2, 0x1b, 0x14, 0x78,
	0x5b, 0xc6, 0x8e, 0x24, 0x25, 0x43, 0x54, 0x98, 0x45, 0x07, 0x43, 0x39, 0x4e, 0x50, 0x11, 0x23,
	0x68, 0xe4, 0xe8, 0x98, 0x99, 0x1b, 0x54, 0xc2, 0x2b, 0x80, 0x72, 0xf2, 0x41, 0x0a, 0x36, 0x3d,
	0x95, 0x16, 0x95, 0xf1, 0x2a, 0x2c, 0xe7, 0xf4, 0x4c, 0x13, 0x61, 0x14, 0xd1, 0x4c, 0x58, 0x54,
	0xc1, 0x4b, 0xd0, 0xc8, 0xd3, 0x1c, 0x11, 0x63, 0xd1, 0x8f, 0x60, 0xeb, 0x13, 0xd4, 0x67, 0xca,
	0xc0, 0x35, 0x28, 0xe7, 0x39, 0xab, 0x50, 0xea, 0x9d, 0x5e, 0xbc, 0x46, 0x41, 0xb6, 0xda, 0x45,
	0x05, 0xdc, 0x02, 0x38, 0xd0, 0x53, 0x65, 0xa5, 0x57, 0x8a, 0x8f, 0xe6, 0x5d, 0x54, 0xc2, 0x35,
	0x28, 0x65, 0x07, 0xbf, 0x7d, 0x03, 0x4f, 0xa8, 0x1c, 0x45, 0xf7, 0x2c, 0x61, 0x09, 0x89, 0xa8,
	0xbb, 0x42, 0x34, 0x36, 0xe9, 0x97, 0x99, 0xfe, 0x13, 0x06, 0xff, 0x5d, 0x73, 0x7b, 0x33, 0xbe,
	0x8c, 0xa8, 0x1c, 0x6d, 0xa7, 0xbe, 0x6d, 0xa2, 0xf8, 0xf6, 0x7d, 0xfa, 0xe5, 0x5f, 0x56, 0xbc,
	0xeb, 0xd5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x55, 0x92, 0xeb, 0x69, 0x04, 0x00, 0x00,
}