package checktcp

import "bufio"
import "io"
import "fmt"
import "strings"
import "os"

func ChecktcpM(targets_file_path string) {
	var targetslist []string
	file, err := os.Open(targets_file_path)
	if err != nil {
		fmt.Println(targets_file_path, " open err.")
	}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
                break
			}
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			targetslist = append(targetslist, line)
		}
	}
    for _, targets := range targetslist {
        Checktcp(targets)
    }
}
