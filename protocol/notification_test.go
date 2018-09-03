package protocol

import (
	"testing"
)

func TestNotificationToJSONString(t *testing.T) {
	ext := make(map[string]interface{})
	ext["order_id"] = 1

	noti := Notification{
		UniqID: "noti-1234567890",
		To:     "user-123",
		Title:  "title",
		Body:   "body",
		Ext:    ext,
	}

	notiJSONString, err := noti.ToJSONString()

	if err != nil {
		t.Error(err)
	}

	expectJSONString := `{"uniqId":"noti-1234567890","to":"user-123","title":"title","body":"body","ext":{"order_id":1}}`

	if expectJSONString != notiJSONString {
		t.Error("excpect " + expectJSONString + " but " + notiJSONString)
	}

	t.Log("ok")
}
