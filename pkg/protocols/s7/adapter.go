package s7

import (
	"log"
	"time"

	"github.com/robinson/gos7"
)

type AdapterConfig struct {
	// Addr ip address, eg "127.0.0.1"
	Addr    string
	Rack    int
	Slot    int
	Timeout time.Duration
}

type s7Adapter struct {
	handler *gos7.TCPClientHandler
	client  gos7.Client
}

// NewModbusTCPAdapter creates an new TCP adapter
func NewS7Adapter(conf AdapterConfig) *s7Adapter {
	handler := gos7.NewTCPClientHandler(conf.Addr, conf.Rack, conf.Slot)
	handler.Timeout = conf.Timeout
	handler.Logger = log.Default()

	client := gos7.NewClient(handler)
	return &s7Adapter{handler: handler, client: client}
}

// Connect connects the modbus server.
func (apt *s7Adapter) Connect() error {
	err := apt.handler.Connect()
	return err
}

// AGReadDB reads data blocks
func (apt *s7Adapter) AGReadDB(dbNumber, start, size int, buffer []byte) (err error) {
	err = apt.client.AGReadDB(dbNumber, start, size, buffer)
	return
}

func (apt *s7Adapter) Close() {
	apt.handler.Close()
}
