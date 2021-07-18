package configs

type Item struct {
	Code string `json:"code"`
	Addr uint16 `json:"addr"`
	Len  uint16 `json:"len"`
	typ  string `json:"type"`
}

type ReadArea struct {
}

type WriteArea struct {
}

type area struct {
	start   uint16
	vlen    uint16
	address []span
}

type span struct {
	start uint16
	vlen  uint8
	vtype string // int|real|string
}
