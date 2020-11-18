package mytool

import (
	"bufio"
	"fmt"
	"os"
)

func GetMsg() string {
	reader := bufio.NewReader(os.Stdin)
	data, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
	}
	msg := string(data)

	return msg
}
