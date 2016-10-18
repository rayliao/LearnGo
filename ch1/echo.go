package ch1

import (
	"fmt"
	"os"
	"strings"
)

// Echo func 输出参数
func Echo() {
	fmt.Println(strings.Join(os.Args[:], "***"))
}
