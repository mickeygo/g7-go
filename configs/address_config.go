package configs

// Address 地址数据
type Address struct {
	// 唯一编号
	Code string `json:"code"`
	// 地址
	Addr uint16 `json:"addr"`
	// 长度 bool->1byte, int->2byte, dint->4byte, real->4byte, string->specified
	Len uint16 `json:"len"`
	// 数据类型 bool/int/dint|real|string
	Type string `json:"type"`
	// 描述
	Desc string `json:"desc"`
}
