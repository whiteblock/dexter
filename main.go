package main

import (
	"context"
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
}
