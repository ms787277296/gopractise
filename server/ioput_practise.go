package server

import (
	"bufio"
	"fmt"
	"os"

	"github.com/golang/glog"
)

func InputOutputPractise() {
	var buf []byte
	fmt.Println("123")
	n, err := os.Stdin.Read(buf)
	if err != nil {
		glog.Error(err)
	}
	glog.Info(string(buf[:n]))
}

func BufIO() {

	inputReader := bufio.NewReader(os.Stdin)

	input, err := inputReader.ReadString('\n')
	if err != nil {
		glog.Error(err)
	}

	glog.Info(input)

}

func WritePractise() (err error) {
	file, err := os.OpenFile("./hello.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		glog.Error(err)
		return
	}

	defer file.Close()

	file.WriteString("欢迎来到德莱联盟")

	return
}
