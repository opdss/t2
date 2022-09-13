package t2

import (
	"fmt"

	"github.com/opdss/t3"
)

func init() {

}

func SetT3(val int) int {
	fmt.Println("T2 SetT3(): t3 => %d", val)
	t3.T3 = val
	return t3.T3
}

func GetT3() int {

	fmt.Println("T2 GetT3(): t3 => %d", t3.T3)
	return t3.T3
}
