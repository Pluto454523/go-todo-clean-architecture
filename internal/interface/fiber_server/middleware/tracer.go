package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer = otel.GetTracerProvider().Tracer("fiber-server")

func TracerMiddleware(c *fiber.Ctx) error {
	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	carrier := propagation.HeaderCarrier{}

	c.Request().Header.VisitAll(func(key, value []byte) {
		carrier.Set(string(key), string(value))
	})

	propagator.Inject(c.Context(), carrier)
	//parentCtx := propagator.Extract(c.Context(), carrier)

	spanOptions := []trace.SpanStartOption{
		trace.WithAttributes(semconv.HTTPMethodKey.String(c.Method())),
		trace.WithAttributes(semconv.HTTPTargetKey.String(string(c.Request().RequestURI()))),
		trace.WithAttributes(semconv.HTTPRouteKey.String(c.Route().Path)),
		trace.WithAttributes(semconv.HTTPURLKey.String(c.OriginalURL())),
		//trace.WithAttributes(semconv.NetHostIPKey.String(c.IP())),
		trace.WithAttributes(semconv.UserAgentOriginal(string(c.Request().Header.UserAgent()))),
		trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int(c.Request().Header.ContentLength())),
		trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Protocol())),
		trace.WithAttributes(semconv.NetTransportTCP),
		trace.WithSpanKind(trace.SpanKindServer),
	}

	ctx, span := Tracer.Start(c.Context(), fmt.Sprintf("%s %s", c.Method(), c.Path()), spanOptions...)
	defer span.End()

	//c.Locals(contextLocalKey, ctx)
	{
		propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
		carrier := propagation.HeaderCarrier{}
		propagator.Inject(ctx, carrier)

		for _, k := range carrier.Keys() {
			c.Response().Header.Set(k, carrier.Get(k))
		}
	}

	err := c.Next()

	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().StatusCode()))

	return err
}

func TracerMiddlewareV2(c *fiber.Ctx) error {

	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	carrier := propagation.HeaderCarrier{}

	c.Request().Header.VisitAll(func(key, value []byte) {
		carrier.Set(string(key), string(value))
	})

	ctx := propagator.Extract(c.Context(), carrier)

	spanOptions := []trace.SpanStartOption{
		trace.WithAttributes(semconv.HTTPMethodKey.String(c.Method())),
		trace.WithAttributes(semconv.HTTPTargetKey.String(string(c.Request().RequestURI()))),
		trace.WithAttributes(semconv.HTTPRouteKey.String(c.Route().Path)),
		trace.WithAttributes(semconv.HTTPURLKey.String(c.OriginalURL())),
		trace.WithAttributes(semconv.UserAgentOriginal(string(c.Request().Header.UserAgent()))),
		trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int(c.Request().Header.ContentLength())),
		trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Protocol())),
		trace.WithAttributes(semconv.NetTransportTCP),
		trace.WithSpanKind(trace.SpanKindServer),
	}

	ctx, span := Tracer.Start(ctx, fmt.Sprintf("%s %s", c.Method(), c.Path()), spanOptions...)
	defer span.End()

	c.Locals("otel_trace_context", ctx)

	err := c.Next()

	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().StatusCode()))

	return err
}

func TracerMiddlewareV3(c *fiber.Ctx) error {

	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	carrier := propagation.HeaderCarrier{}

	c.Request().Header.VisitAll(func(key, value []byte) {
		carrier.Set(string(key), string(value))
	})

	ctx := propagator.Extract(c.Context(), carrier)

	spanOptions := []trace.SpanStartOption{
		trace.WithAttributes(semconv.HTTPMethodKey.String(c.Method())),
		trace.WithAttributes(semconv.HTTPTargetKey.String(string(c.Request().RequestURI()))),
		trace.WithAttributes(semconv.HTTPRouteKey.String(c.Route().Path)),
		trace.WithAttributes(semconv.HTTPURLKey.String(c.OriginalURL())),
		trace.WithAttributes(semconv.UserAgentOriginal(string(c.Request().Header.UserAgent()))),
		trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int(c.Request().Header.ContentLength())),
		trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Protocol())),
		trace.WithAttributes(semconv.NetTransportTCP),
		trace.WithSpanKind(trace.SpanKindServer),
	}

	ctx, span := Tracer.Start(ctx, fmt.Sprintf("%s %s", c.Method(), c.Path()), spanOptions...)
	defer span.End()

	c.Locals("otel_trace_context", ctx)

	err := c.Next()

	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().StatusCode()))

	return err
}
