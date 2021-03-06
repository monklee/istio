// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/stdio/config/config.proto

/*
	Package config is a generated protocol buffer package.

	It is generated from these files:
		mixer/adapter/stdio/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Stream is used to select between different log output sinks.
type Params_Stream int32

const (
	// Output to the Mixer process' standard output stream. This is the default value.
	STDOUT Params_Stream = 0
	// Output to the Mixer process' standard error stream.
	STDERR Params_Stream = 1
	// Output to a specific file.
	FILE Params_Stream = 2
	// Output to a specific rotating file, controlled by the various file rotation options.
	ROTATED_FILE Params_Stream = 3
)

var Params_Stream_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
	2: "FILE",
	3: "ROTATED_FILE",
}
var Params_Stream_value = map[string]int32{
	"STDOUT":       0,
	"STDERR":       1,
	"FILE":         2,
	"ROTATED_FILE": 3,
}

func (Params_Stream) EnumDescriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 0} }

// Importance level for individual items output by this adapter.
type Params_Level int32

const (
	INFO    Params_Level = 0
	WARNING Params_Level = 1
	ERROR   Params_Level = 2
)

var Params_Level_name = map[int32]string{
	0: "INFO",
	1: "WARNING",
	2: "ERROR",
}
var Params_Level_value = map[string]int32{
	"INFO":    0,
	"WARNING": 1,
	"ERROR":   2,
}

func (Params_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 1} }

// The stdio adapter makes it possible to capture log and metrics and output them to the Mixer process'
// standard output or standard error stream, or to a specific file.
type Params struct {
	// Selects which standard stream to write to for log entries.
	// STDERR is the default Stream.
	LogStream Params_Stream `protobuf:"varint,1,opt,name=log_stream,json=logStream,proto3,enum=adapter.stdio.config.Params_Stream" json:"log_stream,omitempty"`
	// Maps from severity strings as specified in LogEntry instances to
	// the set of levels supported by this adapter. This defaults to a map of
	//    "INFORMATIONAL" : INFO,
	//    "informational" : INFO,
	//    "INFO" : INFO,
	//    "info" : INFO,
	//    "WARNING" : WARNING,
	//    "warning" : WARNING,
	//    "WARN": WARNING,
	//    "warning": WARNING,
	//    "ERROR": ERROR,
	//    "error": ERROR,
	//    "ERR": ERROR,
	//    "err": ERROR,
	//    "FATAL": ERROR,
	//    "fatal": ERROR,
	SeverityLevels map[string]Params_Level `protobuf:"bytes,2,rep,name=severity_levels,json=severityLevels" json:"severity_levels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=adapter.stdio.config.Params_Level"`
	// The level to assign to metrics being output. Defaults to INFO.
	MetricLevel Params_Level `protobuf:"varint,3,opt,name=metric_level,json=metricLevel,proto3,enum=adapter.stdio.config.Params_Level" json:"metric_level,omitempty"`
	// Whether to output a console-friendly or json-friendly format. Defaults to true.
	OutputAsJson bool `protobuf:"varint,4,opt,name=output_as_json,json=outputAsJson,proto3" json:"output_as_json,omitempty"`
	// The minimum level to output, anything less than this level is ignored. Defaults to INFO (everything).
	OutputLevel Params_Level `protobuf:"varint,5,opt,name=output_level,json=outputLevel,proto3,enum=adapter.stdio.config.Params_Level" json:"output_level,omitempty"`
	// The file system path when outputting to a file or rotating file.
	//
	// When using rotated log files, this path is used as a foundational path. This is where log
	// output is normally saved. When a rotation needs to take place because the file got too big
	// or too old, then the file is renamed by appending a timestamp to the name. Such renamed
	// files are called backups. Once a backup has been created, output resumes to this path.
	OutputPath string `protobuf:"bytes,6,opt,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	// The maximum size in megabytes of a log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxMegabytesBeforeRotation int32 `protobuf:"varint,7,opt,name=max_megabytes_before_rotation,json=maxMegabytesBeforeRotation,proto3" json:"max_megabytes_before_rotation,omitempty"`
	// The maximum number of days to retain old rotated log files based on the
	// timestamp encoded in their filename. Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is to remove log files
	// older than 30 days. 0 indicates no limit.
	MaxDaysBeforeRotation int32 `protobuf:"varint,8,opt,name=max_days_before_rotation,json=maxDaysBeforeRotation,proto3" json:"max_days_before_rotation,omitempty"`
	// The maximum number of old rotated log files to retain.  The default
	// is to retain at most 1000 logs. 0 indicates no limit.
	MaxRotatedFiles int32 `protobuf:"varint,9,opt,name=max_rotated_files,json=maxRotatedFiles,proto3" json:"max_rotated_files,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func init() {
	proto.RegisterType((*Params)(nil), "adapter.stdio.config.Params")
	proto.RegisterEnum("adapter.stdio.config.Params_Stream", Params_Stream_name, Params_Stream_value)
	proto.RegisterEnum("adapter.stdio.config.Params_Level", Params_Level_name, Params_Level_value)
}
func (x Params_Stream) String() string {
	s, ok := Params_Stream_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x Params_Level) String() string {
	s, ok := Params_Level_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LogStream != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.LogStream))
	}
	if len(m.SeverityLevels) > 0 {
		for k, _ := range m.SeverityLevels {
			dAtA[i] = 0x12
			i++
			v := m.SeverityLevels[k]
			mapSize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + sovConfig(uint64(v))
			i = encodeVarintConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintConfig(dAtA, i, uint64(v))
		}
	}
	if m.MetricLevel != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MetricLevel))
	}
	if m.OutputAsJson {
		dAtA[i] = 0x20
		i++
		if m.OutputAsJson {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.OutputLevel != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.OutputLevel))
	}
	if len(m.OutputPath) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.OutputPath)))
		i += copy(dAtA[i:], m.OutputPath)
	}
	if m.MaxMegabytesBeforeRotation != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxMegabytesBeforeRotation))
	}
	if m.MaxDaysBeforeRotation != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxDaysBeforeRotation))
	}
	if m.MaxRotatedFiles != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MaxRotatedFiles))
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	if m.LogStream != 0 {
		n += 1 + sovConfig(uint64(m.LogStream))
	}
	if len(m.SeverityLevels) > 0 {
		for k, v := range m.SeverityLevels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + sovConfig(uint64(v))
			n += mapEntrySize + 1 + sovConfig(uint64(mapEntrySize))
		}
	}
	if m.MetricLevel != 0 {
		n += 1 + sovConfig(uint64(m.MetricLevel))
	}
	if m.OutputAsJson {
		n += 2
	}
	if m.OutputLevel != 0 {
		n += 1 + sovConfig(uint64(m.OutputLevel))
	}
	l = len(m.OutputPath)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.MaxMegabytesBeforeRotation != 0 {
		n += 1 + sovConfig(uint64(m.MaxMegabytesBeforeRotation))
	}
	if m.MaxDaysBeforeRotation != 0 {
		n += 1 + sovConfig(uint64(m.MaxDaysBeforeRotation))
	}
	if m.MaxRotatedFiles != 0 {
		n += 1 + sovConfig(uint64(m.MaxRotatedFiles))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	keysForSeverityLevels := make([]string, 0, len(this.SeverityLevels))
	for k, _ := range this.SeverityLevels {
		keysForSeverityLevels = append(keysForSeverityLevels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForSeverityLevels)
	mapStringForSeverityLevels := "map[string]Params_Level{"
	for _, k := range keysForSeverityLevels {
		mapStringForSeverityLevels += fmt.Sprintf("%v: %v,", k, this.SeverityLevels[k])
	}
	mapStringForSeverityLevels += "}"
	s := strings.Join([]string{`&Params{`,
		`LogStream:` + fmt.Sprintf("%v", this.LogStream) + `,`,
		`SeverityLevels:` + mapStringForSeverityLevels + `,`,
		`MetricLevel:` + fmt.Sprintf("%v", this.MetricLevel) + `,`,
		`OutputAsJson:` + fmt.Sprintf("%v", this.OutputAsJson) + `,`,
		`OutputLevel:` + fmt.Sprintf("%v", this.OutputLevel) + `,`,
		`OutputPath:` + fmt.Sprintf("%v", this.OutputPath) + `,`,
		`MaxMegabytesBeforeRotation:` + fmt.Sprintf("%v", this.MaxMegabytesBeforeRotation) + `,`,
		`MaxDaysBeforeRotation:` + fmt.Sprintf("%v", this.MaxDaysBeforeRotation) + `,`,
		`MaxRotatedFiles:` + fmt.Sprintf("%v", this.MaxRotatedFiles) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogStream", wireType)
			}
			m.LogStream = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogStream |= (Params_Stream(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeverityLevels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SeverityLevels == nil {
				m.SeverityLevels = make(map[string]Params_Level)
			}
			var mapkey string
			var mapvalue Params_Level
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthConfig
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (Params_Level(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConfig(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthConfig
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.SeverityLevels[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricLevel", wireType)
			}
			m.MetricLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MetricLevel |= (Params_Level(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputAsJson", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OutputAsJson = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputLevel", wireType)
			}
			m.OutputLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OutputLevel |= (Params_Level(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutputPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMegabytesBeforeRotation", wireType)
			}
			m.MaxMegabytesBeforeRotation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMegabytesBeforeRotation |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDaysBeforeRotation", wireType)
			}
			m.MaxDaysBeforeRotation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxDaysBeforeRotation |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRotatedFiles", wireType)
			}
			m.MaxRotatedFiles = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRotatedFiles |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/stdio/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0xc7, 0xe3, 0x76, 0xc9, 0xda, 0x67, 0xd3, 0x96, 0xd7, 0xef, 0x90, 0xa2, 0x49, 0x98, 0xaa,
	0x20, 0x51, 0x38, 0xa4, 0x68, 0x1c, 0x98, 0x10, 0x97, 0x4e, 0xcb, 0xd0, 0xd0, 0x58, 0x27, 0xaf,
	0x08, 0xc1, 0x25, 0x72, 0x57, 0x37, 0x0b, 0x24, 0x75, 0x15, 0xbb, 0x55, 0x73, 0xe3, 0x23, 0xf0,
	0x31, 0xf8, 0x28, 0x3b, 0xee, 0xc8, 0x91, 0x86, 0x0b, 0xc7, 0x1d, 0xf8, 0x00, 0x28, 0x76, 0x38,
	0x00, 0xd3, 0xb4, 0x53, 0x1e, 0xff, 0xf3, 0xfb, 0x3d, 0x8f, 0xe5, 0x38, 0xf0, 0x30, 0x8d, 0x17,
	0x3c, 0xeb, 0xb2, 0x11, 0x9b, 0x2a, 0x9e, 0x75, 0xa5, 0x1a, 0xc5, 0xa2, 0x7b, 0x26, 0x26, 0xe3,
	0x38, 0xaa, 0x1e, 0xfe, 0x34, 0x13, 0x4a, 0xe0, 0xad, 0x0a, 0xf1, 0x35, 0xe2, 0x9b, 0x77, 0xdb,
	0x5b, 0x91, 0x88, 0x84, 0x06, 0xba, 0x65, 0x65, 0xd8, 0xf6, 0x4f, 0x1b, 0x9c, 0x13, 0x96, 0xb1,
	0x54, 0xe2, 0x3d, 0x80, 0x44, 0x44, 0xa1, 0x54, 0x19, 0x67, 0xa9, 0x87, 0x5a, 0xa8, 0xb3, 0xb1,
	0x73, 0xdf, 0xbf, 0xae, 0x97, 0x6f, 0x0c, 0xff, 0x54, 0xa3, 0xb4, 0x99, 0x88, 0xc8, 0x94, 0xf8,
	0x1d, 0x6c, 0x4a, 0x3e, 0xe7, 0x59, 0xac, 0xf2, 0x30, 0xe1, 0x73, 0x9e, 0x48, 0xaf, 0xd6, 0xaa,
	0x77, 0xd6, 0x76, 0x9e, 0xdc, 0xdc, 0xa8, 0x72, 0x8e, 0xb4, 0x12, 0x4c, 0x54, 0x96, 0xd3, 0x0d,
	0xf9, 0x47, 0x88, 0x03, 0x58, 0x4f, 0xb9, 0xca, 0xe2, 0x33, 0xd3, 0xd8, 0xab, 0xeb, 0x0d, 0xb6,
	0x6f, 0xec, 0xab, 0x55, 0xba, 0x66, 0x3c, 0xbd, 0xc0, 0x0f, 0x60, 0x43, 0xcc, 0xd4, 0x74, 0xa6,
	0x42, 0x26, 0xc3, 0x0f, 0x52, 0x4c, 0xbc, 0x95, 0x16, 0xea, 0x34, 0xe8, 0xba, 0x49, 0x7b, 0xf2,
	0x95, 0x14, 0x93, 0x72, 0x58, 0x45, 0x99, 0x61, 0xf6, 0xed, 0x87, 0x19, 0xcf, 0x0c, 0xbb, 0x07,
	0xd5, 0x32, 0x9c, 0x32, 0x75, 0xee, 0x39, 0x2d, 0xd4, 0x69, 0x52, 0x30, 0xd1, 0x09, 0x53, 0xe7,
	0xb8, 0x07, 0x77, 0x53, 0xb6, 0x08, 0x53, 0x1e, 0xb1, 0x61, 0xae, 0xb8, 0x0c, 0x87, 0x7c, 0x2c,
	0x32, 0x1e, 0x66, 0x42, 0x31, 0x15, 0x8b, 0x89, 0xb7, 0xda, 0x42, 0x1d, 0x9b, 0x6e, 0xa7, 0x6c,
	0xf1, 0xfa, 0x37, 0xb3, 0xa7, 0x11, 0x5a, 0x11, 0xf8, 0x19, 0x78, 0x65, 0x8b, 0x11, 0xcb, 0xff,
	0xb5, 0x1b, 0xda, 0xbe, 0x93, 0xb2, 0xc5, 0x3e, 0xcb, 0xff, 0x16, 0x1f, 0xc3, 0x7f, 0xa5, 0xa8,
	0x61, 0x3e, 0x0a, 0xc7, 0x71, 0xc2, 0xa5, 0xd7, 0xd4, 0xc6, 0x66, 0xca, 0x16, 0xd4, 0xe4, 0x07,
	0x65, 0xbc, 0xcd, 0xe1, 0xff, 0x6b, 0xbe, 0x11, 0x76, 0xa1, 0xfe, 0x91, 0xe7, 0xfa, 0xae, 0x34,
	0x69, 0x59, 0xe2, 0x5d, 0xb0, 0xe7, 0x2c, 0x99, 0x71, 0xaf, 0x76, 0xeb, 0x13, 0x33, 0xc2, 0xf3,
	0xda, 0x2e, 0x6a, 0xbf, 0x00, 0xa7, 0xba, 0x48, 0x00, 0xce, 0xe9, 0x60, 0xbf, 0xff, 0x66, 0xe0,
	0x5a, 0x55, 0x1d, 0x50, 0xea, 0x22, 0xdc, 0x80, 0x95, 0x83, 0xc3, 0xa3, 0xc0, 0xad, 0x61, 0x17,
	0xd6, 0x69, 0x7f, 0xd0, 0x1b, 0x04, 0xfb, 0xa1, 0x4e, 0xea, 0xed, 0x47, 0x60, 0x9b, 0x63, 0x6f,
	0xc0, 0xca, 0xe1, 0xf1, 0x41, 0xdf, 0xb5, 0xf0, 0x1a, 0xac, 0xbe, 0xed, 0xd1, 0xe3, 0xc3, 0xe3,
	0x97, 0x2e, 0xc2, 0x4d, 0xb0, 0x03, 0x4a, 0xfb, 0xd4, 0xad, 0xed, 0xed, 0x5e, 0x2c, 0x89, 0x75,
	0xb9, 0x24, 0xd6, 0xd7, 0x25, 0xb1, 0xae, 0x96, 0xc4, 0xfa, 0x54, 0x10, 0xf4, 0xa5, 0x20, 0xd6,
	0x45, 0x41, 0xd0, 0x65, 0x41, 0xd0, 0xb7, 0x82, 0xa0, 0x1f, 0x05, 0xb1, 0xae, 0x0a, 0x82, 0x3e,
	0x7f, 0x27, 0xd6, 0x7b, 0xc7, 0x6c, 0x7d, 0xe8, 0xe8, 0xff, 0xe6, 0xe9, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x2c, 0xa2, 0x33, 0x8e, 0x03, 0x00, 0x00,
}
