package structures

//Broadcast struct defines the structure used by BBC sounds
type Broadcast struct {
	ID    string `json:"id"`
	Title Title  `json:"titles"`
}
