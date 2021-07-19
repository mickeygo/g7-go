package modbus

import (
	"sync"
	"time"

	"github.com/goburrow/modbus"
)

type AdapterConfig struct {
	// Addr ip address, eg "127.0.0.1:502"
	Addr string
	// SlaveId eg: 0x01, or 0xFF
	SlaveId byte
	Timeout time.Duration
}

type modbusAdapter struct {
	handler *modbus.TCPClientHandler
	client  modbus.Client

	mu sync.Mutex
}

// NewModbusTCPAdapter creates an new TCP adapter
func NewModbusTCPAdapter(conf AdapterConfig) *modbusAdapter {
	handler := modbus.NewTCPClientHandler(conf.Addr)
	handler.SlaveId = conf.SlaveId
	handler.Timeout = conf.Timeout
	client := modbus.NewClient(handler)

	return &modbusAdapter{handler: handler, client: client}
}

// Connect connects the modbus server.
func (adapter *modbusAdapter) Connect() error {
	err := adapter.handler.Connect()
	return err
}

// ReadHoldingRegisters 
func (adapter *modbusAdapter) ReadHoldingRegisters(address uint16, number uint16) ([]byte, error) {
	buffer, err := adapter.client.ReadHoldingRegisters(address, number)
	return buffer, err
}

// Close closes the TCP connection.
func (adapter *modbusAdapter) Close() {
	adapter.handler.Close()
}
