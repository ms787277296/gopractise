package server

import (
	"github.com/golang/glog"
)

func SumProfit() (sum float64) {
	a:= 6.0
	b:= 1.08
	sum = 6.0
	for i:=0; i<29;i++{
		sum = sum*b+a
	}
	
	glog.Info(sum)
	return
}
