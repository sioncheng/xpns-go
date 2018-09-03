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

//ToJSONString export notification to json string
func (n *Notification) ToJSONString() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
