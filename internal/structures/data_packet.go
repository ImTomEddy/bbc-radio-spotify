package structures

//DataPacket structure defines a combo of the Broadcast and Song
type DataPacket struct {
	Broadcast Broadcast
	Song      Title
}
