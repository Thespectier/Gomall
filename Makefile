.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend --module github.com/qingz2/gomall/app/frontend -I ../../idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/user.proto --service user --module github.com/qingz2/gomall/rpc_gen
	@cd app/user && cwgo server --type RPC --I ../../idl --idl ../../idl/user.proto --service user --module github.com/qingz2/gomall/app/user --pass "-use github.com/qingz2/gomall/rpc_gen/kitex_gen"

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/cart.proto --service cart --module github.com/qingz2/gomall/rpc_gen
	@cd app/cart && cwgo server --type RPC --I ../../idl --idl ../../idl/cart.proto --service cart --module github.com/qingz2/gomall/app/cart --pass "-use github.com/qingz2/gomall/rpc_gen/kitex_gen"

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/payment.proto --service payment --module github.com/qingz2/gomall/rpc_gen
	@cd app/payment && cwgo server --type RPC --I ../../idl --idl ../../idl/payment.proto --service payment --module github.com/qingz2/gomall/app/payment --pass "-use github.com/qingz2/gomall/rpc_gen/kitex_gen"

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/checkout.proto --service checkout --module github.com/qingz2/gomall/rpc_gen
	@cd app/checkout && cwgo server --type RPC --I ../../idl --idl ../../idl/checkout.proto --service checkout --module github.com/qingz2/gomall/app/checkout --pass "-use github.com/qingz2/gomall/rpc_gen/kitex_gen"

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/product.proto --service product --module github.com/qingz2/gomall/rpc_gen
	@cd app/product && cwgo server --type RPC --I ../../idl --idl ../../idl/product.proto --service product --module github.com/qingz2/gomall/app/product --pass "-use github.com/qingz2/gomall/rpc_gen/kitex_gen"