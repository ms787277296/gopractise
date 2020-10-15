module ms/gopractise

go 1.14

replace ms/testproto/gen/go/contact => ../testproto/gen/go/contact

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/stretchr/testify v1.6.1
	ms/testproto/gen/go/contact v0.0.0-00010101000000-000000000000
)
