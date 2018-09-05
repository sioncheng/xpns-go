package protocol

import "sync/atomic"

const (
	//Zero 0
	_ = iota
	//LoginCode login command code
	LoginCode
	//LogonCode logon command code
	LogonCode
	//NotificationCode notification command code
	NotificationCode
	//AckCode ack command code
	AckCode
	//HeartbeatCode heart beat command code
	HeartbeatCode
)

//JSONCommand json command object
type JSONCommand struct {
	SerialNumber  int64
	CommandType   int
	CommandObject interface{}
}

var serialNumber int64

//GetNewSerialNumber return atomic incremented command serial number
func GetNewSerialNumber() int64 {
	return atomic.AddInt64(&serialNumber, 1)
}
