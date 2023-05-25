package models

type SMS struct {
	To 		[]string `json:"to"`
	From 	string 	 `json:"from"`
	Message string 	 `json:"message"`
}

type Email struct {
	To      []string `json:"to"`
	From    string   `json:"from"`
	Message []byte   `json:"message"`
}

type ChannelMessage struct {
	Medium 		 string `json:"medium"`
	Type		 string `json:"type"`
	Notification []byte `json:"notification"`
}