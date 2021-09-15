package rpc

import (
	"context"
	"encoding/base64"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"strings"
)

const (
	binHdrSuffix = "-bin"
)

var (
	grpcTag = opentracing.Tag{Key: string(ext.Component), Value: "gRPC"}
)

func ClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx, clientSpan := newClientSpanFromContext(ctx, tracer, method)
		clientSpan.SetTag("grpc.target", cc.Target())
		clientSpan.SetTag("grpc.method", method)
		err := invoker(newCtx, method, req, reply, cc, opts...)
		finishClientSpan(clientSpan, err)
		return err
	}
}

// metadataTextMap extends a metadata.MD to be an opentracing textmap
type metadataTextMap metadata.MD

// Set is a opentracing.TextMapReader interface that extracts values.
func (m metadataTextMap) Set(key, val string) {
	// gRPC allows for complex binary values to be written.
	encodedKey, encodedVal := encodeKeyValue(key, val)
	// The metadata object is a multiMap, and previous values may exist, but for opentracing headers, we do not append
	// we just override.
	m[encodedKey] = []string{encodedVal}
}

// ForeachKey is a opentracing.TextMapReader interface that extracts values.
func (m metadataTextMap) ForeachKey(callback func(key, val string) error) error {
	for k, vv := range m {
		for _, v := range vv {
			if err := callback(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// encodeKeyValue encodes key and value qualified for transmission via gRPC.
func encodeKeyValue(k, v string) (string, string) {
	k = strings.ToLower(k)
	if strings.HasSuffix(k, binHdrSuffix) {
		val := base64.StdEncoding.EncodeToString([]byte(v))
		v = val
	}
	return k, v
}

type clientSpanTagKey struct{}

func newClientSpanFromContext(ctx context.Context, tracer opentracing.Tracer, operateName string) (context.Context, opentracing.Span) {
	var parentSpanCtx opentracing.SpanContext
	// fetch possible exist parent context which withValue of span
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		parentSpanCtx = parent.Context()
	}
	opts := []opentracing.StartSpanOption{
		opentracing.ChildOf(parentSpanCtx),
		ext.SpanKindRPCClient,
		grpcTag,
	}
	if tag := ctx.Value(clientSpanTagKey{}); tag != nil {
		if opt, ok := tag.(opentracing.StartSpanOption); ok {
			opts = append(opts, opt)
		}
	}

	clientSpan := tracer.StartSpan(operateName, opts...)

	// create new textMap carrier
	md := metadataTextMap{}
	// inject data from current span to textMap carrier
	if err := tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, md); err != nil {
		// record inject error, can't panic, will break the span
		clientSpan.LogFields(log.String("inject-error", err.Error()))
	}

	ctxWithMetadata := metadata.NewOutgoingContext(ctx, metadata.MD(md))
	return opentracing.ContextWithSpan(ctxWithMetadata, clientSpan), clientSpan
}

func finishClientSpan(clientSpan opentracing.Span, err error) {
	if err != nil && err != io.EOF {
		ext.Error.Set(clientSpan, true)
		clientSpan.LogFields(log.String("invoke-error", err.Error()))
	}
	clientSpan.Finish()
}
