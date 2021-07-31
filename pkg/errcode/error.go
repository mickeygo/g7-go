package errcode

import "fmt"

func Define(code uint32, module string, desc string) error {
	return fmt.Errorf("[%d] module:%s, desc:%s", code, module, desc)
}
