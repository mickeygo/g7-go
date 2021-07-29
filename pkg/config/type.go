package config

// 数据类型
type DataType int

const (
	_        DataType = iota
	BOOL              // 位
	BYTE              // 字节（8位无符号）
	CHAR              // 字节（ASCII字符）
	WORD              // 字（16位无符号）
	DWORD             // 双字（32位无符号）
	INT               // 整数（16位有符号）
	DINT              // 长整数（32位有符号）
	REAL              // 32位浮点数
	STRING            // 字符串
	DATETIME          // 日期
)

type area struct {
	start   uint16
	vlen    uint16
	address []span
}

// span 数据区域最小单元块
type span struct {
	addr  uint16
	vlen  uint8
	vtype DataType
}
