package structures

//Broadcast struct defines the structure used by BBC sounds
type Broadcast struct {
	ID    string `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
	Title Title  `json:"titles"`
}
