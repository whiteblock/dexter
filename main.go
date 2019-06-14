package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	dataPb "github.com/whiteblock/dexter/api/data"
)

var (
	help    = false
	verbose = false
	client  = ""
)

func supportedExchanges(client dataPb.DexterDataClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	exchanges, err := client.SupportedExchanges(ctx, &dataPb.ExchangesRequest{})
	if err != nil {
		log.Fatalln("Error", err)
	}
	log.Println(exchanges)
}

func streamCandles(client dataPb.DexterDataClient, request *dataPb.CandlesRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
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

func main() {
	flags := pflag.NewFlagSet("dexter", pflag.ExitOnError)
	flags.BoolVarP(&help, "help", "h", false, "Display help message")
	flags.BoolVarP(&verbose, "verbose", "v", false, "Be verbose")
	flags.StringVarP(&client, "client", "c", "0.0.0.0:50051", "Bind address of dexter-data service")
	conn, err := grpc.Dial(client, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to client", err)
	}
	defer conn.Close()

	client := dataPb.NewDexterDataClient(conn)
	supportedExchanges(client)
	streamCandles(client, &dataPb.CandlesRequest{Exchange: "binance", Market: "BTC/USDT", Timeframe: "5m"})
}
