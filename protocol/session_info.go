package protocol

const (
	//Zero 0
	_ = iota
	//SessionInfoLogon  when client logon
	SessionInfoLogon
	//SessionInfoQuery when query user session info
	SessionInfoQuery
)

//SessionInfo user(acid) connected to which xpns server
type SessionInfo struct {
	Acid      string `json:"acid"`
	Server    string `json:"server"`
	Port      int    `json:"port"`
	Type      int    `json:"type"`
	Timestamp int64  `json:"timestamp"`
}
