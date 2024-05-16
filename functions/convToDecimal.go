package reload

import (
	"fmt"
	"strconv"
)

func ConvToDecimal(str string, base int) string {
	decimal_num, err := strconv.ParseInt(str, base, 64)
	if err != nil {
		fmt.Println(err)
	}
	str = strconv.Itoa(int(decimal_num))
	return (str)
}
