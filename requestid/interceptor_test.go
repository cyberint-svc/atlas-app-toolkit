package requestid

import (
	"context"
	"testing"

	mock_transport "github.com/infobloxopen/atlas-app-toolkit/mocks/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type testRequest struct{}

type testResponse struct{}

func TestUnaryServerInterceptorWithoutRequestId(t *testing.T) {
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		reqID, exists := FromContext(ctx)
		if exists && reqID == "" {
			t.Errorf("requestId must be generated by interceptor")
		}
		return &testResponse{}, nil
	}
	ctx := mock_transport.DummyContextWithServerTransportStream()
	_, err := UnaryServerInterceptor()(ctx, testRequest{}, nil, handler)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUnaryServerInterceptorWithDummyRequestId(t *testing.T) {
	dummyRequestID := newRequestID()
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		reqID, exists := FromContext(ctx)
		if !exists || reqID != dummyRequestID {
			t.Errorf("expected requestID: %q, returned requestId: %q", dummyRequestID, reqID)
		}
		return &testResponse{}, nil
	}
	ctx := mock_transport.DummyContextWithServerTransportStream()
	md := metadata.Pairs(DefaultRequestIDKey, dummyRequestID)
	newCtx := metadata.NewIncomingContext(ctx, md)
	_, err := UnaryServerInterceptor()(newCtx, testRequest{}, nil, handler)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestStreamServerInterceptorWithoutRequestId(t *testing.T) {
	handler := func(srv interface{}, stream grpc.ServerStream) error {
		reqID, exists := FromContext(stream.Context())
		if exists && reqID == "" {
			t.Errorf("requestId must be generated by interceptor")
		}
		return nil
	}
	ctx := mock_transport.DummyContextWithServerTransportStream()
	ms := mock_transport.NewMockServerStream(ctx)
	err := StreamServerInterceptor()(testRequest{}, ms, nil, handler)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestStreamServerInterceptorWithDummyRequestId(t *testing.T) {
	dummyRequestID := newRequestID()
	handler := func(srv interface{}, stream grpc.ServerStream) error {
		reqID, exists := FromContext(stream.Context())
		if !exists || reqID != dummyRequestID {
			t.Errorf("expected requestID: %q, returned requestId: %q", dummyRequestID, reqID)
		}
		return nil
	}
	ctx := mock_transport.DummyContextWithServerTransportStream()
	md := metadata.Pairs(DefaultRequestIDKey, dummyRequestID)
	newCtx := metadata.NewIncomingContext(ctx, md)
	streamInterceptor := StreamServerInterceptor()
	if err := streamInterceptor(testRequest{}, mock_transport.NewMockServerStream(newCtx), nil, handler); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
