package adapter

import "strings"

var (
	protocals = [...]string{"modbustcp", "opcua", "s7"}
)

type adapter struct {
	protocol string
}

func (apt *adapter) read() ([]byte, error) {
	if strings.EqualFold(apt.protocol, "modbustcp") {
		// TODO: handle the modbustcp.
	}

	return nil, nil
}

// must be the one of protocols
func checkProtocal(protocol string) bool {
	pass := false
	for _, v := range protocals {
		if strings.EqualFold(protocol, v) {
			pass = true
			break
		}
	}

	return pass
}
