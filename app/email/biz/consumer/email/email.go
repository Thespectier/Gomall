package email

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"github.com/qingz2/gomall/app/email/infra/mq"
	"github.com/qingz2/gomall/app/email/infra/notify"
	"github.com/qingz2/gomall/rpc_gen/kitex_gen/email"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	// Connect to a server

	// tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}
		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)

		// consumer otel
		// ctx := context.Background()
		// ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(m.Header))
		// _, span := tracer.Start(ctx, "shop-email-consumer")
		// defer span.End()
		// consumer otel
	})
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe() //nolint:errcheck
		mq.Nc.Close()
	})
}
