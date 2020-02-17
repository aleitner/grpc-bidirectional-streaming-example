package bidirectional

import (
	"fmt"
	"io"

	pb "grpc-bidirectional-stream/pkg/proto"
)

// SampleServer
type SampleServer struct {
}

// SampleBidirectional receives data from the client and responds with a SampleBidirectionalResponse
func (s *SampleServer) SampleBidirectional(stream pb.Sample_SampleBidirectionalServer) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		fmt.Println("received", res.GetData())
		err = stream.Send(&pb.SampleBidirectionalResponse{
			Data: res.GetData(),
		})
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	return nil
}
