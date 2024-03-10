// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ticchain/tictactoe/game.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Game struct {
	GameId       uint64 `protobuf:"varint,1,opt,name=gameId,proto3" json:"gameId,omitempty"`
	AdminAddress string `protobuf:"bytes,2,opt,name=adminAddress,proto3" json:"adminAddress,omitempty"`
	Status       string `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Board        string `protobuf:"bytes,4,opt,name=board,proto3" json:"board,omitempty"`
	NextMove     string `protobuf:"bytes,5,opt,name=NextMove,proto3" json:"NextMove,omitempty"`
	Winner       string `protobuf:"bytes,6,opt,name=Winner,proto3" json:"Winner,omitempty"`
	// First player
	FirstPlayer *Player `protobuf:"bytes,7,opt,name=first_player,json=firstPlayer,proto3" json:"first_player,omitempty"`
	// Opponent (Second player)
	SecondPlayer *Player `protobuf:"bytes,8,opt,name=second_player,json=secondPlayer,proto3" json:"second_player,omitempty"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ec163e8bdf27d1, []int{0}
}
func (m *Game) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Game.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return m.Size()
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *Game) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
	}
	return ""
}

func (m *Game) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Game) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *Game) GetNextMove() string {
	if m != nil {
		return m.NextMove
	}
	return ""
}

func (m *Game) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *Game) GetFirstPlayer() *Player {
	if m != nil {
		return m.FirstPlayer
	}
	return nil
}

func (m *Game) GetSecondPlayer() *Player {
	if m != nil {
		return m.SecondPlayer
	}
	return nil
}

// Message representing a player.
type Player struct {
	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Symbol   string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Points   uint64 `protobuf:"varint,3,opt,name=points,proto3" json:"points,omitempty"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ec163e8bdf27d1, []int{1}
}
func (m *Player) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Player.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return m.Size()
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *Player) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Player) GetPoints() uint64 {
	if m != nil {
		return m.Points
	}
	return 0
}

func init() {
	proto.RegisterType((*Game)(nil), "ticchain.tictactoe.Game")
	proto.RegisterType((*Player)(nil), "ticchain.tictactoe.Player")
}

func init() { proto.RegisterFile("ticchain/tictactoe/game.proto", fileDescriptor_22ec163e8bdf27d1) }

var fileDescriptor_22ec163e8bdf27d1 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x9b, 0xfe, 0xd3, 0xf9, 0xdb, 0xb4, 0x6e, 0x82, 0x48, 0x68, 0x71, 0x28, 0x5d, 0x75,
	0xe3, 0x14, 0x14, 0x97, 0x22, 0xba, 0x91, 0x2e, 0x14, 0x19, 0x11, 0xc1, 0x4d, 0x49, 0x27, 0x51,
	0x03, 0x9d, 0x64, 0x48, 0xae, 0xd2, 0xbe, 0x85, 0x0f, 0xe1, 0xc3, 0xb8, 0xec, 0xd2, 0xa5, 0xb4,
	0x2f, 0x22, 0x93, 0xa4, 0x45, 0x71, 0xe1, 0xf2, 0x3b, 0xe7, 0x9e, 0x43, 0x72, 0x2f, 0xde, 0x07,
	0x99, 0xe7, 0x4f, 0x4c, 0xaa, 0x11, 0xc8, 0x1c, 0x58, 0x0e, 0x5a, 0x8c, 0x1e, 0x59, 0x21, 0xd2,
	0xd2, 0x68, 0xd0, 0x84, 0x6c, 0xec, 0x74, 0x6b, 0x0f, 0xde, 0xea, 0x38, 0xba, 0x60, 0x85, 0x20,
	0x7b, 0x38, 0xae, 0x46, 0xc7, 0x9c, 0xa2, 0x3e, 0x1a, 0x46, 0x59, 0x20, 0x32, 0xc0, 0x1d, 0xc6,
	0x0b, 0xa9, 0xce, 0x38, 0x37, 0xc2, 0x5a, 0x5a, 0xef, 0xa3, 0x61, 0x2b, 0xfb, 0xa1, 0x55, 0xd9,
	0x1b, 0x60, 0xf0, 0x6c, 0xe9, 0x3f, 0xe7, 0x06, 0x22, 0xbb, 0xb8, 0x31, 0xd5, 0xcc, 0x70, 0x1a,
	0x39, 0xd9, 0x03, 0xe9, 0xe2, 0xe6, 0x95, 0x98, 0xc3, 0xa5, 0x7e, 0x11, 0xb4, 0xe1, 0x8c, 0x2d,
	0x57, 0x4d, 0x77, 0x52, 0x29, 0x61, 0x68, 0xec, 0x9b, 0x3c, 0x91, 0x13, 0xdc, 0x79, 0x90, 0xc6,
	0xc2, 0xa4, 0x9c, 0xb1, 0x85, 0x30, 0xf4, 0x7f, 0x1f, 0x0d, 0xdb, 0x87, 0xdd, 0xf4, 0xf7, 0x8f,
	0xd2, 0x6b, 0x37, 0x91, 0xb5, 0xdd, 0xbc, 0x07, 0x72, 0x8a, 0x77, 0xac, 0xc8, 0xb5, 0xe2, 0x9b,
	0x7c, 0xf3, 0xcf, 0x7c, 0xc7, 0x07, 0x3c, 0x0d, 0x6e, 0x71, 0x1c, 0xaa, 0x7a, 0xb8, 0xe5, 0x3b,
	0x26, 0xd2, 0xaf, 0xaa, 0x95, 0x35, 0xbd, 0x30, 0xe6, 0xd5, 0xf3, 0xed, 0xa2, 0x98, 0xea, 0x59,
	0x58, 0x53, 0xa0, 0x4a, 0x2f, 0xb5, 0x54, 0xe0, 0x17, 0x14, 0x65, 0x81, 0xce, 0x8f, 0xdf, 0x57,
	0x09, 0x5a, 0xae, 0x12, 0xf4, 0xb9, 0x4a, 0xd0, 0xeb, 0x3a, 0xa9, 0x2d, 0xd7, 0x49, 0xed, 0x63,
	0x9d, 0xd4, 0xee, 0x7b, 0x20, 0xf3, 0x03, 0x7f, 0xcb, 0xf9, 0xb7, 0x6b, 0xc2, 0xa2, 0x14, 0x76,
	0x1a, 0xbb, 0x7b, 0x1e, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xf1, 0xa3, 0xac, 0xf0, 0x01,
	0x00, 0x00,
}

func (m *Game) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Game) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Game) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SecondPlayer != nil {
		{
			size, err := m.SecondPlayer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGame(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.FirstPlayer != nil {
		{
			size, err := m.FirstPlayer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGame(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.NextMove) > 0 {
		i -= len(m.NextMove)
		copy(dAtA[i:], m.NextMove)
		i = encodeVarintGame(dAtA, i, uint64(len(m.NextMove)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Board) > 0 {
		i -= len(m.Board)
		copy(dAtA[i:], m.Board)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Board)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AdminAddress) > 0 {
		i -= len(m.AdminAddress)
		copy(dAtA[i:], m.AdminAddress)
		i = encodeVarintGame(dAtA, i, uint64(len(m.AdminAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.GameId != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Player) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Player) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Player) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Points != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Points))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PlayerId) > 0 {
		i -= len(m.PlayerId)
		copy(dAtA[i:], m.PlayerId)
		i = encodeVarintGame(dAtA, i, uint64(len(m.PlayerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Game) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameId != 0 {
		n += 1 + sovGame(uint64(m.GameId))
	}
	l = len(m.AdminAddress)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Board)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.NextMove)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	if m.FirstPlayer != nil {
		l = m.FirstPlayer.Size()
		n += 1 + l + sovGame(uint64(l))
	}
	if m.SecondPlayer != nil {
		l = m.SecondPlayer.Size()
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *Player) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PlayerId)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	if m.Points != 0 {
		n += 1 + sovGame(uint64(m.Points))
	}
	return n
}

func sovGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGame(x uint64) (n int) {
	return sovGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Game) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: Game: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Game: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Board = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextMove", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextMove = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstPlayer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FirstPlayer == nil {
				m.FirstPlayer = &Player{}
			}
			if err := m.FirstPlayer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondPlayer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SecondPlayer == nil {
				m.SecondPlayer = &Player{}
			}
			if err := m.SecondPlayer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func (m *Player) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: Player: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Player: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Points", wireType)
			}
			m.Points = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Points |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func skipGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
				return 0, ErrInvalidLengthGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGame = fmt.Errorf("proto: unexpected end of group")
)