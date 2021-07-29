package s7

import (
	"testing"
	"time"

	"github.com/robinson/gos7"
)

func Test_s7Adapter_AGReadDB(t *testing.T) {
	conf := AdapterConfig{
		Addr:    "192.168.0.1",
		Rack:    0,
		Slot:    1,
		Timeout: time.Second,
	}
	adapter := NewS7Adapter(conf)
	adapter.Connect()
	defer adapter.Close()

	buf := make([]byte, 255)
	err := adapter.AGReadDB(2710, 8, 2, buf)
	if err != nil {
		t.Errorf("AGReadDB error, v% \n", err)
	}

	var s7 gos7.Helper
	var result uint16
	s7.GetValueAt(buf, 0, &result)

	t.Logf("result: %d", result)
}
