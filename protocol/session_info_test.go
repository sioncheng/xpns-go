package protocol

import "testing"

func TestSessionInfo(t *testing.T) {
	sessionInfo := SessionInfo{
		Acid:      "user-123",
		Server:    "server-1",
		Port:      8080,
		Type:      SessionInfoLogon,
		Timestamp: 123,
	}

	if sessionInfo.Type == 1 {
		t.Log("ok")
	} else {
		t.Error("err")
	}
}
