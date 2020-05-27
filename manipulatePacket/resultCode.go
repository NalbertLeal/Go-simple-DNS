package manipulatePacket

const (
	NOERROR  = 0
	FORMERR  = 1
	SERVFAIL = 2
	NXDOMAIN = 3
	NOTIMP   = 4
	REFUSED  = 5
)

func isResultCodeAnError(code uint8) bool {
	return NOERROR == code
}
