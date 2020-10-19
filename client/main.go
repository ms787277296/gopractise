package main

import(
	"net"
	"github.com/golang/glog"
	"os"
	"strings"
	"bufio"
)

func main(){
	conn,err := net.Dial("tcp", "localhost:50000")
	if err != nil{
		glog.Error(err)
		return
	}

	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for{
		input,err:= inputReader.ReadString('\n')
		if err !=nil{
			glog.Error(err)
		}

		trimmedinput := strings.TrimSpace(input)
		_,err =conn.Write([]byte(trimmedinput))
		if err != nil{
			glog.Error(err)
		}

	}
}

