package reader

import (
	"bufio"
	"os"
	"strings"
)

func ReadData() string {
	reader := bufio.NewReader(os.Stdin)
	data := ""
	data, _ = reader.ReadString('\n')
	data = strings.TrimSuffix(data, "\n")
	return data
}
