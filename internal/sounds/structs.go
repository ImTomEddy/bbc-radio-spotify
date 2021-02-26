package sounds

//Broadcast is the structure of the data provided by the BBC Sounds API
type Broadcast struct {
	ID    string `json:"id"`
	Title Title  `json:"titles"`
}

//Title is the structure of the data provided by the BBC Sounds API
type Title struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

//DataPacket structure defines a combo of the Broadcast and Song
type DataPacket struct {
	Broadcast Broadcast
	Song      Title
}
