package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/whiteblock/dexter"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	dataPb "github.com/whiteblock/dexter/api/data"
)

var (
	help    = false
	verbose = false
	client  = ""
	listen  = ""
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

// load alerts from database and start analysis on them
func loadAlerts(conn *grpc.ClientConn, db *gorm.DB) {
	var alerts []dexter.Alert
	client := dataPb.NewDataClient(conn)

	// for every alert create a chart if needed
	db.Find(&alerts)
	for _, alert := range alerts {
		dexter.SetupChart(alert, client)
	}
}

// dexter [OPTION]

func main() {
	flags := pflag.NewFlagSet("dexter", pflag.ExitOnError)
	flags.BoolVarP(&help, "help", "h", false, "Display help message")
	flags.BoolVarP(&verbose, "verbose", "v", false, "Be verbose")
	flags.StringVarP(&client, "client", "c", "0.0.0.0:50051", "Bind address of dexter-data service")
	flags.StringVarP(&listen, "listen", "l", "0.0.0.0:50052", "Bind address of dexter-alerts service")
	flags.SortFlags = false
	flags.Parse(os.Args)
	if help {
		fmt.Print("DEXter Technical Analysis Service\n\n")
		fmt.Print("Usage: dexter [OPTION]...\n\n")
		flags.PrintDefaults()
		os.Exit(0)
	}
	// connect to database
	err := godotenv.Load()
	maxTries := 5
	tries := 0
	connect := os.Getenv("PG_URL")
	var wg sync.WaitGroup
	var db *gorm.DB
	wg.Add(1)

	// Try a few times to connect to the database.
	func() {
		for {
			db, err = gorm.Open("postgres", connect)
			if err != nil {
				log.Println("Could not connect to database", err)
				tries = tries + 1
				if tries >= maxTries {
					log.Fatalln("Exceeded max tries")
				}
				time.Sleep(3 * time.Second)
				continue
			} else {
				//defer db.Close()
				wg.Done()
				break
			}
		}
	}()

	wg.Wait()

	// connect to dexter-data gRPC service
	conn, err := grpc.Dial(client, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to client", err)
	}
	defer conn.Close()

	//demo(conn)

	// load alerts from database
	loadAlerts(conn, db)

	// Start dexter gRPC Service
	dexter.StartServer(listen, db, conn)
}
