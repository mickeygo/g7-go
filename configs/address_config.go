package configs

// Item 数据配置项
type Item struct {
	Code string `json:"code"`
	Addr uint16 `json:"addr"`
	Len  uint16 `json:"len"`
	Type string `json:"type"`
}

// ReadArea 读区间，用于一次性读取
type ReadArea struct {
	// 起始地址
	Addr uint16
	// 地址长度
	Len uint16
}

// WriteArea 写区间
type WriteArea struct {
}

type area struct {
	start   uint16
	vlen    uint16
	address []span
}

// span 数据区域最小单元块
type span struct {
	start uint16
	vlen  uint8
	vtype string // int|real|string
}
