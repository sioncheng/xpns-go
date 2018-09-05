package protocol

import (
	"encoding/json"
)

//Notification struct
type Notification struct {
	UniqID string                 `json:"uniqId"`
	To     string                 `json:"to"`
	Title  string                 `json:"title"`
	Body   string                 `json:"body"`
	Ext    map[string]interface{} `json:"ext"`
}

//NilNotification as a standard nil notification
var NilNotification = Notification{
	UniqID: "",
}

//ToJSONString export notification to json string
func (n *Notification) ToJSONString() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

//ParseNotificationFromJSONString parse json string to notification struct if possible
func ParseNotificationFromJSONString(jsonString string) (Notification, error) {
	noti := Notification{}

	err := json.Unmarshal([]byte(jsonString), &noti)
	if err != nil {
		return NilNotification, err
	}

	return noti, nil
}
