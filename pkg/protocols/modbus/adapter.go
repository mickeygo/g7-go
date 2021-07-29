package modbus

import (
	"log"
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
}

// NewModbusTCPAdapter creates an new TCP adapter
func NewModbusTCPAdapter(conf AdapterConfig) *modbusAdapter {
	handler := modbus.NewTCPClientHandler(conf.Addr)
	handler.SlaveId = conf.SlaveId
	handler.Timeout = conf.Timeout
	handler.Logger = log.Default()

	client := modbus.NewClient(handler)
	return &modbusAdapter{handler: handler, client: client}
}

// Connect connects the modbus server.
func (apt *modbusAdapter) Connect() error {
	err := apt.handler.Connect()
	return err
}

// ReadHoldingRegisters
func (apt *modbusAdapter) ReadHoldingRegisters(address uint16, number uint16) ([]byte, error) {
	buffer, err := apt.client.ReadHoldingRegisters(address, number)
	return buffer, err
}

// Close closes the TCP connection.
func (apt *modbusAdapter) Close() {
	apt.handler.Close()
}
