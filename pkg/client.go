package bidirectional

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	pb "grpc-bidirectional-stream/pkg/proto"

	"google.golang.org/grpc"
)

// SampleClient allows access to Bidirectional stream method and closing conn
type SampleClient interface {
	Bidirectional(ctx context.Context) error
	CloseConn() error
}

// Client struct has access to pb.SampleClient and grpc conn
type Client struct {
	route pb.SampleClient
	conn  *grpc.ClientConn
}

// NewSampleClient initializes a SampleClient
func NewSampleClient(conn *grpc.ClientConn) SampleClient {
	return &Client{conn: conn, route: pb.NewSampleClient(conn)}
}

// CloseConn closes conn
func (client *Client) CloseConn() error {
	return client.conn.Close()
}

// Bidirectional creates a bidirectional client stream
func (client *Client) Bidirectional(ctx context.Context) error {
	stream, err := client.route.SampleBidirectional(ctx)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		hello := []byte("hello world")
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
			err := stream.Send(&pb.SampleBidirectionalRequest{
				Data: hello,
			})
			if err != nil {
				// server returns with nil
				if err == io.EOF {
					break
				}
				log.Fatalf("stream Send fail: %s/n", err)
			}
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("stream send fail: %s\n", err)
		}
		wg.Done()
	}()

	wg.Add(1)

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("stream Recv fail: %s/n", err)
			}
			fmt.Println(string(res.GetData()))
		}
		wg.Done()
	}()
	wg.Wait()

	return nil
}
