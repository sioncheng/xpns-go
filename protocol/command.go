package protocol

const (
	//MagicByteHigh high byte of mgic bytes
	MagicByteHigh = 0xab
	//MagicByteLow low byte of magic bytes
	MagicByteLow = 0x12
	//HeartBeat command byte of heart beat
	HeartBeat = 0x00
	//Request command byte of request
	Request = 0x01
	//Response command byte of response
	Response = 0x02
	//JSONSerialization byte of command serialization
	JSONSerialization = 0x01
)

//Command represent the request/response passed through server and client
type Command struct {
	SerialNumber      int64
	CommandType       byte
	SerializationType byte
	PayloadLength     int
	PayloadBytes      []byte
}
