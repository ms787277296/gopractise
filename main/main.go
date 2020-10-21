package main

import (
	"flag"
	"ms/gopractise/server"
	_ "ms/gopractise/server"

	"github.com/golang/glog"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	// err := server.JsonTransform()
	// if err != nil {
	// 	glog.Fatal(err)
	// }

	// err := server.ProtobufPractise()
	// if err != nil {
	// 	glog.Fatal(err)
	// }

	// server.PrintNum()

	// err :=server.TCPServer()
	// if err != nil {
	// 	glog.Fatal(err)
	// }
	//var buf []byte

	err := server.WritePractise()
	if err != nil {
		glog.Fatal(err)
	}

}
