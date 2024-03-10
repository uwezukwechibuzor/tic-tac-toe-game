package types

// ValidateBasic is used for validating the packet
func (p IbcgamePacketData) ValidateBasic() error {
	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcgamePacketData) GetBytes() ([]byte, error) {
	var modulePacket TttibcPacketData

	modulePacket.Packet = &TttibcPacketData_IbcgamePacket{&p}

	return modulePacket.Marshal()
}
