package modbus

import (
	"testing"
	"time"
)

func Test_modbusAdapter_ReadHoldingRegisters(t *testing.T) {
	conf := AdapterConfig{
		Addr:    "127.0.0.1:502",
		SlaveId: 0x01,
		Timeout: time.Second,
	}
	adapter := NewModbusTCPAdapter(conf)
	defer adapter.Close()

	buffer, err := adapter.ReadHoldingRegisters(1, 8)
	if err != nil {
		t.Errorf("ReadHoldingRegisters error, v% \n", err)
	}

	t.Logf("result: %v", string(buffer))
}
