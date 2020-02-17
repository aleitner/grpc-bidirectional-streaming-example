package client

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	bidirectional "grpc-bidirectional-stream/pkg"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var (
	port int
)

func initializeFlags() {
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()
}

func main() {
	app := cli.NewApp()

	// Set up connection with rpc server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc Dial fail: %s/n", err)
	}

	client := bidirectional.NewSampleClient(conn)
	defer client.CloseConn()

	app.Commands = []cli.Command{
		{
			Name:    "bidirectional",
			Aliases: []string{"u"},
			Usage:   "bidirectional data",
			Action: func(c *cli.Context) error {
				err := client.Bidirectional(context.Background())
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("app Run fail: %s/n", err)
	}
}
