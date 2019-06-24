package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/whiteblock/dexter"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	dataPb "github.com/whiteblock/dexter/api/data"
)

var (
	help    = false
	verbose = false
	client  = ""
)

func supportedExchanges(client dataPb.DataClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	exchanges, err := client.SupportedExchanges(ctx, &dataPb.ExchangesRequest{})
	if err != nil {
		log.Fatalln("Error", err)
	}
	log.Println(exchanges)
}

func streamCandles(client dataPb.DataClient, request *dataPb.CandlesRequest) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.StreamCandles(ctx, request)
	if err != nil {
		log.Fatalln("Could not stream", err)
	}
	for {
		candle, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Streaming error", err)
		}
		log.Println(candle)
	}
}

func demo(conn *grpc.ClientConn) {
	client := dataPb.NewDataClient(conn)
	supportedExchanges(client)
	streamCandles(client, &dataPb.CandlesRequest{Exchange: "binance", Market: "BTC/USDT", Timeframe: "5m"})
}

// dexter [OPTION]

func main() {
	flags := pflag.NewFlagSet("dexter", pflag.ExitOnError)
	flags.BoolVarP(&help, "help", "h", false, "Display help message")
	flags.BoolVarP(&verbose, "verbose", "v", false, "Be verbose")
	flags.StringVarP(&client, "client", "c", "0.0.0.0:50051", "Bind address of dexter-data service")
	flags.SortFlags = false
	flags.Parse(os.Args)
	if help {
		fmt.Print("DEXter Technical Analysis Service\n\n")
		fmt.Print("Usage: dexter [OPTION]...\n\n")
		flags.PrintDefaults()
		os.Exit(0)
	}
	conn, err := grpc.Dial(client, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to client", err)
	}
	defer conn.Close()
	demo(conn)

	// TODO Start gRPC Service
}
