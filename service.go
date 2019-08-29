package dexter

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc/reflection"
	grpc "google.golang.org/grpc"
	pb "github.com/whiteblock/dexter/api/alerts"
	dataPb "github.com/whiteblock/dexter/api/data"
)

type dexterAlertsServer struct {
	db *gorm.DB
	dexterData dataPb.DataClient
}

func (s *dexterAlertsServer) CreateAlert(ctx context.Context, alert *pb.Alert) (*pb.Alert, error) {
	// Constraints
	// - Timeframe should be formatted correctly
	// - LineX
	a := Alert{
		ExternalID:	alert.ExternalId,
		Exchange:	alert.Exchange,
		Market:		alert.Market,
		Timeframe:	alert.Timeframe,
		LineA:          postgres.Jsonb{json.RawMessage(alert.LineA)},
		Condition:	AlertCondition(alert.Condition),
		LineB:          postgres.Jsonb{json.RawMessage(alert.LineB)},
		Frequency:	NotificationFrequency(alert.Frequency),
		MessageBody:	alert.MessageBody,
	}
	spew.Dump(a)
	s.db.Create(&a)
	newAlert := &pb.Alert{
		Id: uint64(a.ID),
	}
	// TODO - dexter.AddAlert to in-memory Chart
	chart := SetupChart(a, s.dexterData)
	chart.AddAlert(a)
	return newAlert, nil
}

func (s *dexterAlertsServer) ListAlerts(ctx context.Context, opts *pb.ListAlertsRequest) (*pb.ListAlertsResponse, error) {
	response := &pb.ListAlertsResponse{}
	log.Printf("ListAlerts %s", spew.Sdump(opts))
	var alerts []Alert
	if opts.ExternalId != 0 {
		s.db.Where("external_id = ?", opts.ExternalId).Find(&alerts)
	} else {
		s.db.Find(&alerts)
	}
	for _, alert := range alerts {
		pa := &pb.Alert{
			Id:             uint64(alert.ID),
			ExternalId:	alert.ExternalID,
			Exchange:	alert.Exchange,
			Market:		alert.Market,
			Timeframe:	alert.Timeframe,
			Condition:	pb.Condition(alert.Condition),
			Frequency:	pb.Frequency(alert.Frequency),
			MessageBody:	alert.MessageBody,
		}
		response.Alerts = append(response.Alerts, pa)
	}
	return response, nil
}

func (s *dexterAlertsServer) GetAlert(ctx context.Context, opts *pb.GetAlertRequest) (*pb.Alert, error) {
	var alert Alert
	log.Printf("GetAlert %s", spew.Sdump(opts))
	s.db.First(&alert, opts.AlertId)
	response := &pb.Alert{
		Id: uint64(alert.ID),
		ExternalId:	alert.ExternalID,
		Exchange:	alert.Exchange,
		Market:		alert.Market,
		Timeframe:	alert.Timeframe,
		Condition:	pb.Condition(alert.Condition),
		Frequency:	pb.Frequency(alert.Frequency),
		MessageBody:	alert.MessageBody,
	}
	return response, nil
}

func (s *dexterAlertsServer) UpdateAlert(ctx context.Context, alert *pb.Alert) (*pb.Alert, error) {
	response := &pb.Alert{}
	log.Printf("UpdateAlert %s", spew.Sdump(alert))
	var dbAlert Alert
	s.db.First(&dbAlert, alert.Id)
	dbAlert.Market = alert.Market
	dbAlert.Exchange = alert.Exchange
	dbAlert.Timeframe = alert.Timeframe
	dbAlert.Condition = AlertCondition(alert.Condition)
	dbAlert.Frequency = NotificationFrequency(alert.Frequency)
	dbAlert.MessageBody = alert.MessageBody
	s.db.Save(&dbAlert)
	chart := SetupChart(dbAlert, s.dexterData)
	chart.UpdateAlert(dbAlert)
	return response, nil
}

func (s *dexterAlertsServer) DeleteAlert(ctx context.Context, opts *pb.DeleteAlertRequest) (*pb.DeleteAlertResponse, error) {
	response := &pb.DeleteAlertResponse{}
	log.Printf("DeleteAlert %s", spew.Sdump(opts))
	alert := Alert{
		Model: gorm.Model{
			ID: uint(opts.AlertId),
		},
	}
	s.db.Delete(&alert)
	// TODO - Remove Alert from in-memory Chart
	//chart := SetupChart(alert, s.dexterData)
	//chart.RemoveAlert(alert)
	return response, nil
}

func (s *dexterAlertsServer) ListIndicators(ctx context.Context, opts *pb.ListIndicatorsRequest) (*pb.ListIndicatorsResponse, error) {
	response := &pb.ListIndicatorsResponse{}
	log.Printf("ListIndicators %s", spew.Sdump(opts))
	for _, i := range Indicators {
		indicatorSpec := &pb.Indicator{
			Name: i.Name,
			Inputs: i.Inputs,
			Outputs: i.Outputs,
		}
		response.Indicators = append(response.Indicators, indicatorSpec)
	}
	return response, nil
}

// StartServer starts the gRPC service for alert management
func StartServer(listen string, db *gorm.DB, conn *grpc.ClientConn) {
	var opts []grpc.ServerOption
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Listening on %s\n", listen)
	}
	client := dataPb.NewDataClient(conn)
	server := &dexterAlertsServer{
		db: db,
		dexterData: client,
	}
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	pb.RegisterAlertsServer(grpcServer, server)
	grpcServer.Serve(listener)
}
