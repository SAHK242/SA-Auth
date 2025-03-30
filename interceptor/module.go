package interceptor

import "go.uber.org/fx"

var Module = fx.Provide(
	fx.Annotate(
		NewChainInterceptor,
		fx.ResultTags(`name:"chainedGrpcInterceptor"`),
	),
)
