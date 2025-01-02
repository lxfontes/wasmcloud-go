package wasmbus

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.wasmcloud.dev/wasmbus/wasmbustest"
)

func TestServerRegisterHandler(t *testing.T) {
	defer wasmbustest.MustStartNats(t)()
	nc, err := NatsConnect(NatsDefaultURL)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()

	bus := NewNatsBus(nc)
	server := NewServer(bus, "test")
	server.RegisterHandler("test", ServerHandlerFunc(func(ctx context.Context, msg *Message) error {
		reply := NewMessage(msg.Reply)
		reply.Data = []byte("hello")
		return server.Publish(reply)
	}))

	resp, err := server.Request(context.TODO(), NewMessage("test"))
	if err != nil {
		t.Fatal(err)
	}

	if want, got := "hello", string(resp.Data); want != got {
		t.Fatalf("want %q, got %q", want, got)
	}

	resp, err = server.Request(context.TODO(), NewMessage("error"))
	if err == nil {
		t.Fatal("expected NO_RESPONDERS")
	}
}

func TestServerDrain(t *testing.T) {
	defer wasmbustest.MustStartNats(t)()
	nc, err := NatsConnect(NatsDefaultURL)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()

	bus := NewNatsBus(nc)
	server := NewServer(bus, "test")
	checkpoint := make(chan bool)
	server.RegisterHandler("slow", ServerHandlerFunc(func(ctx context.Context, msg *Message) error {
		close(checkpoint)
		<-time.After(200 * time.Millisecond)
		return server.Publish(NewMessage(msg.Reply))
	}))

	errCh := make(chan error, 1)
	go func() {
		_, err = server.Request(context.TODO(), NewMessage("slow"))
		errCh <- err
	}()
	<-checkpoint

	// NOTE(lxf): waits for all in-flight messages to be processed
	if err := server.Drain(); err != nil {
		t.Fatal(err)
	}

	reqErr := <-errCh
	if reqErr != nil {
		t.Fatal(reqErr)
	}
}

func TestServerErrorStream(t *testing.T) {
	defer wasmbustest.MustStartNats(t)()
	nc, err := NatsConnect(NatsDefaultURL)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()

	bus := NewNatsBus(nc)
	server := NewServer(bus, "test")
	bomb := errors.New("bomb")
	bombCh := make(chan error, 1)
	server.RegisterHandler("bomb", ServerHandlerFunc(func(ctx context.Context, msg *Message) error {
		reply := NewMessage(msg.Reply)
		server.Publish(reply)
		return bomb
	}))

	go func() {
		for err := range server.ErrorStream() {
			bombCh <- err.Err
		}
	}()

	_, err = server.Request(context.TODO(), NewMessage("bomb"))
	if err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-bombCh:
		if want, got := bomb, err; want != got {
			t.Fatalf("want %v, got %v", want, got)
		}
	default:
		t.Fatal("expected error")
	}
}

type testRequest struct{}
type testResponse struct {
	Hello string `json:"hello"`
}

func TestRequestHandler(t *testing.T) {
	defer wasmbustest.MustStartNats(t)()
	nc, err := NatsConnect(NatsDefaultURL)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()

	bus := NewNatsBus(nc)
	server := NewServer(bus, "test")
	handler := NewRequestHandler(testRequest{}, testResponse{}, func(ctx context.Context, req *testRequest) (*testResponse, error) {
		return &testResponse{
			Hello: "world",
		}, nil
	})
	server.RegisterHandler("test", handler)

	rawResp, err := server.Request(context.TODO(), NewMessage("test"))
	if err != nil {
		t.Fatal(err)
	}

	resp := &testResponse{}
	if err := Decode(rawResp, resp); err != nil {
		t.Fatal(err)
	}

	if want, got := "world", resp.Hello; want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}