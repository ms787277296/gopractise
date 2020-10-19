package server

import(
	"net"
	"github.com/golang/glog"
)

func TCPServer()(err error){
	glog.Info("start TCP server:")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil{
		glog.Error(err)
		return err
	}

	for {
		conn,err:= listen.Accept()
		if err != nil{
			glog.Error(err)
			continue
		}

		go TCPProcess(conn)
	}

}

func TCPProcess(conn net.Conn){
	defer conn.Close()
	for{
		buf := make([]byte,512)
		n,err := conn.Read(buf)
		if err != nil{
			glog.Error(err)
			break
		}
		str := string(buf[:n])
		glog.Info(str)
	}
}

