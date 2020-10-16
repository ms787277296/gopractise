package server

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	contact_pb "ms/testproto/gen/go/contact"
)

func ProtobufPractise() (err error) {
	var contact1 contact_pb.ContactBook

	for i:=0;i<60;i++{
		persion := contact_pb.Persion{
			Id:int32(i),
			Name:"ms",
		}

		phone := &contact_pb.Phone{
			Num:"123",
			Type:contact_pb.PhoneType_android,
		}
		persion.Phones = append(persion.Phones, phone)
		contact1.Persions = append(contact1.Persions, &persion)
	}

	data,err := proto.Marshal(&contact1)
	if err!= nil{
		glog.Error(err)
		return
	}

	// fmt.Print(data)

	var contact2 contact_pb.ContactBook

	err = proto.Unmarshal(data, &contact2)

	fmt.Printf("proto is %#v\n", &contact2)

	return
}
