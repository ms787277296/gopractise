package main

import (
	"flag"
	"gopractise/server"

	"github.com/golang/glog"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	answer := server.Add(1, 2)
	glog.Infof("result is %d", answer)

	answer = server.Sub(1, 2)
	glog.Infof("result is %d", answer)

}
