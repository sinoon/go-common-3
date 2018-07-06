package grpcx

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/opentracing/opentracing-go"
	"github.com/qeelyn/go-common/grpcx/registry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
)

type serverOptions struct {
	tracer                   opentracing.Tracer
	logger                   *zap.Logger
	unaryServerInterceptors  []grpc.UnaryServerInterceptor
	streamServerInterceptors []grpc.StreamServerInterceptor
	authFunc                 grpc_auth.AuthFunc
	prometheus               bool
	prometheusListen         string
	register                 registry.Registry
	registryServiceName      string
}

func (t *serverOptions) applyOption(opts ...Option) *serverOptions {
	for _, v := range opts {
		v(t)
	}
	return t
}

type Option func(*serverOptions)

func WithLogger(logger *zap.Logger) Option {
	return func(options *serverOptions) {
		options.logger = logger
	}
}

func WithTracer(tracer opentracing.Tracer) Option {
	return func(options *serverOptions) {
		options.tracer = tracer
	}
}

func WithUnaryServerInterceptor(intercoptors ...grpc.UnaryServerInterceptor) Option {
	return func(options *serverOptions) {
		options.unaryServerInterceptors = append(options.unaryServerInterceptors, intercoptors...)
	}
}

func WithStreamServerInterceptor(intercoptors ...grpc.StreamServerInterceptor) Option {
	return func(options *serverOptions) {
		options.streamServerInterceptors = append(options.streamServerInterceptors, intercoptors...)
	}
}

func WithAuthFunc(authFunc grpc_auth.AuthFunc) Option {
	return func(options *serverOptions) {
		options.authFunc = authFunc
	}
}

func WithPrometheus(listen string) Option {
	return func(options *serverOptions) {
		options.prometheus = true
		if listen != "" {
			options.prometheusListen = listen
		} else {
			log.Printf("use prometheus but no standalone partten!")
		}
	}
}

func WithRegistry(register registry.Registry, serviceName string) Option {
	return func(options *serverOptions) {
		options.register = register
		options.registryServiceName = serviceName
	}
}
