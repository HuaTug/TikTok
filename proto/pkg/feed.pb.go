// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feed.proto

package message

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DouyinFeedRequest struct {
	LatestTime int64 `protobuf:"varint,1,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
}

func (m *DouyinFeedRequest) Reset()         { *m = DouyinFeedRequest{} }
func (m *DouyinFeedRequest) String() string { return proto.CompactTextString(m) }
func (*DouyinFeedRequest) ProtoMessage()    {}
func (*DouyinFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{0}
}
func (m *DouyinFeedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DouyinFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DouyinFeedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DouyinFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DouyinFeedRequest.Merge(m, src)
}
func (m *DouyinFeedRequest) XXX_Size() int {
	return m.Size()
}
func (m *DouyinFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DouyinFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DouyinFeedRequest proto.InternalMessageInfo

func (m *DouyinFeedRequest) GetLatestTime() int64 {
	if m != nil {
		return m.LatestTime
	}
	return 0
}

type DouyinFeedResponse struct {
	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
	NextTime   int64    `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3" json:"next_time,omitempty"`
}

func (m *DouyinFeedResponse) Reset()         { *m = DouyinFeedResponse{} }
func (m *DouyinFeedResponse) String() string { return proto.CompactTextString(m) }
func (*DouyinFeedResponse) ProtoMessage()    {}
func (*DouyinFeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7a672c1337cb5ac, []int{1}
}
func (m *DouyinFeedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DouyinFeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DouyinFeedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DouyinFeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DouyinFeedResponse.Merge(m, src)
}
func (m *DouyinFeedResponse) XXX_Size() int {
	return m.Size()
}
func (m *DouyinFeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DouyinFeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DouyinFeedResponse proto.InternalMessageInfo

func (m *DouyinFeedResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DouyinFeedResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *DouyinFeedResponse) GetVideoList() []*Video {
	if m != nil {
		return m.VideoList
	}
	return nil
}

func (m *DouyinFeedResponse) GetNextTime() int64 {
	if m != nil {
		return m.NextTime
	}
	return 0
}

func init() {
	proto.RegisterType((*DouyinFeedRequest)(nil), "douyin.core.douyin_feed_request")
	proto.RegisterType((*DouyinFeedResponse)(nil), "douyin.core.douyin_feed_response")
}

func init() { proto.RegisterFile("feed.proto", fileDescriptor_d7a672c1337cb5ac) }

var fileDescriptor_d7a672c1337cb5ac = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0x6f, 0xac, 0x8a, 0x3d, 0x05, 0x87, 0xe8, 0x50, 0x14, 0x63, 0xb9, 0x53, 0xa7, 0x88,
	0x0a, 0x2e, 0x6e, 0xba, 0xea, 0x52, 0xc4, 0xc1, 0x25, 0xd4, 0x9b, 0x63, 0x09, 0xde, 0x36, 0xb5,
	0x27, 0x15, 0x7d, 0x0b, 0x9f, 0xc2, 0x67, 0x71, 0xbc, 0xa3, 0xa3, 0xb4, 0x2f, 0x22, 0x69, 0xee,
	0xa0, 0xe3, 0xf9, 0xfe, 0xff, 0x87, 0x8f, 0x03, 0xf0, 0x84, 0xa8, 0x65, 0xdb, 0x59, 0x67, 0x79,
	0xa2, 0x6d, 0xff, 0x6e, 0x1a, 0xb9, 0xb0, 0x1d, 0x1e, 0x40, 0x4f, 0xd8, 0x85, 0x60, 0x7e, 0x01,
	0x7b, 0x21, 0x52, 0xbe, 0xad, 0x3a, 0x7c, 0xe9, 0x91, 0x1c, 0x3f, 0x86, 0x64, 0x59, 0x3a, 0x24,
	0xa7, 0x9c, 0xa9, 0x31, 0x65, 0x19, 0xcb, 0xa3, 0x02, 0x02, 0xba, 0x33, 0x35, 0xce, 0x3f, 0x19,
	0xec, 0xff, 0x1f, 0x52, 0x6b, 0x1b, 0x42, 0xbf, 0x24, 0x57, 0xba, 0x9e, 0xd4, 0xc2, 0xea, 0xb0,
	0xdc, 0x2a, 0x20, 0xa0, 0x6b, 0xab, 0x91, 0x1f, 0xc1, 0xfa, 0x52, 0x35, 0x55, 0xe9, 0x46, 0xc6,
	0xf2, 0xb8, 0x88, 0x03, 0xb9, 0xa5, 0x8a, 0x9f, 0x02, 0xbc, 0x1a, 0x8d, 0x56, 0x2d, 0x0d, 0xb9,
	0x34, 0xca, 0xa2, 0x3c, 0x39, 0xe3, 0xf2, 0x8f, 0xbe, 0xbc, 0xf7, 0x71, 0x11, 0x4f, 0xad, 0x1b,
	0x43, 0x8e, 0x1f, 0x42, 0xdc, 0xe0, 0xdb, 0x5a, 0x75, 0x73, 0x52, 0xdd, 0xf1, 0xc0, 0x8b, 0x5e,
	0xe5, 0x5f, 0x83, 0x60, 0xab, 0x41, 0xb0, 0x9f, 0x41, 0xb0, 0x8f, 0x51, 0xcc, 0x56, 0xa3, 0x98,
	0x7d, 0x8f, 0x62, 0xf6, 0xb0, 0x2b, 0xe5, 0x49, 0xfb, 0x5c, 0x5d, 0xd6, 0x48, 0x54, 0x56, 0xf8,
	0xb8, 0x3d, 0x7d, 0xe4, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xa5, 0x19, 0xcc, 0x38, 0x01,
	0x00, 0x00,
}

func (m *DouyinFeedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DouyinFeedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DouyinFeedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestTime != 0 {
		i = encodeVarintFeed(dAtA, i, uint64(m.LatestTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DouyinFeedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DouyinFeedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DouyinFeedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextTime != 0 {
		i = encodeVarintFeed(dAtA, i, uint64(m.NextTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.VideoList) > 0 {
		for iNdEx := len(m.VideoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VideoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeed(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StatusMsg) > 0 {
		i -= len(m.StatusMsg)
		copy(dAtA[i:], m.StatusMsg)
		i = encodeVarintFeed(dAtA, i, uint64(len(m.StatusMsg)))
		i--
		dAtA[i] = 0x12
	}
	if m.StatusCode != 0 {
		i = encodeVarintFeed(dAtA, i, uint64(m.StatusCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeed(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeed(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DouyinFeedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LatestTime != 0 {
		n += 1 + sovFeed(uint64(m.LatestTime))
	}
	return n
}

func (m *DouyinFeedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StatusCode != 0 {
		n += 1 + sovFeed(uint64(m.StatusCode))
	}
	l = len(m.StatusMsg)
	if l > 0 {
		n += 1 + l + sovFeed(uint64(l))
	}
	if len(m.VideoList) > 0 {
		for _, e := range m.VideoList {
			l = e.Size()
			n += 1 + l + sovFeed(uint64(l))
		}
	}
	if m.NextTime != 0 {
		n += 1 + sovFeed(uint64(m.NextTime))
	}
	return n
}

func sovFeed(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeed(x uint64) (n int) {
	return sovFeed(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DouyinFeedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeed
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
			return fmt.Errorf("proto: douyin_feed_request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: douyin_feed_request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestTime", wireType)
			}
			m.LatestTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeed
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
func (m *DouyinFeedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeed
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
			return fmt.Errorf("proto: douyin_feed_response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: douyin_feed_response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusCode", wireType)
			}
			m.StatusCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
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
				return ErrInvalidLengthFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
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
				return ErrInvalidLengthFeed
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VideoList = append(m.VideoList, &Video{})
			if err := m.VideoList[len(m.VideoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextTime", wireType)
			}
			m.NextTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeed
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
func skipFeed(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeed
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
					return 0, ErrIntOverflowFeed
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFeed
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
				return 0, ErrInvalidLengthFeed
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeed
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeed
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeed        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeed          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeed = fmt.Errorf("proto: unexpected end of group")
)