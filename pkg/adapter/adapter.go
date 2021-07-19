package adapter

import "strings"

var (
	protocals = [...]string{"modbustcp", "opcua", "s7"}
)

func checkProtocal(protocal string) bool {
	pass := false
	for _, v := range protocals {
		if strings.EqualFold(protocal, v) {
			pass = true

			break
		}
	}

	return pass
}
