package configs

// Address 地址数据
type Address struct {
	// 唯一编号
	Code string `json:"code"`
	// 地址
	Addr uint16 `json:"addr"`
	// 长度
	Len uint16 `json:"len"`
	// 数据类型 int|real|string
	Type string `json:"type"`
	// 描述
	Desc string `json:"desc"`
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
	addr  uint16
	vlen  uint8
	vtype string // int|real|string
}
