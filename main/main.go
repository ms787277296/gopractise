package main

import (
	"flag"
	"ms/gopractise/server"

	_ "github.com/golang/glog"
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

	server.PrintNum()
}
