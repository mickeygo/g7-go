package adapter

// Reader the interface for read data of protocols.
type Reader interface {
	Read(addr, number uint16) ([]byte, error)
}
