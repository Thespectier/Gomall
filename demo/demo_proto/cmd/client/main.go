package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/qingz2/Gomall/demo/demo_proto/biz/middleware"
	"github.com/qingz2/Gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/qingz2/Gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware),
	)
	if err != nil {
		panic(err)
	}
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	res, err := c.Echo(ctx, &pbapi.Request{Message: "hello"})
	var bizErr *kerrors.GRPCBizStatusError
	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			log.Printf("%#v", bizErr)
		}
		klog.Fatal(err)
	}
	fmt.Printf("%v", res)
}
