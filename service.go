package dexter

import (
	"context"
	"log"
	"net"
	"github.com/jinzhu/gorm"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc/reflection"
	grpc "google.golang.org/grpc"
	pb "github.com/whiteblock/dexter/api/alerts"
)

type dexterAlertsServer struct {
	db *gorm.DB
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
		Condition:	AlertCondition(alert.Condition),
		Frequency:	NotificationFrequency(alert.Frequency),
		MessageBody:	alert.MessageBody,
	}
	s.db.Create(&a)
	newAlert := &pb.Alert{
		Id: uint64(a.ID),
	}
	return newAlert, nil
}

func (s *dexterAlertsServer) ListAlerts(ctx context.Context, opts *pb.ListAlertsRequest) (*pb.ListAlertsResponse, error) {
	response := &pb.ListAlertsResponse{}
	var alerts []Alert
	s.db.Where("external_id = ?", opts.ExternalId).Find(&alerts)
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
	var dbAlert Alert
	s.db.First(&dbAlert, alert.Id)
	dbAlert.Market = alert.Market
	dbAlert.Exchange = alert.Exchange
	dbAlert.Timeframe = alert.Timeframe
	dbAlert.Condition = AlertCondition(alert.Condition)
	dbAlert.Frequency = NotificationFrequency(alert.Frequency)
	dbAlert.MessageBody = alert.MessageBody
	s.db.Save(&dbAlert)
	return response, nil
}

func (s *dexterAlertsServer) DeleteAlert(ctx context.Context, opts *pb.DeleteAlertRequest) (*pb.DeleteAlertResponse, error) {
	response := &pb.DeleteAlertResponse{}
	alert := Alert{
		Model: gorm.Model{
			ID: uint(opts.AlertId),
		},
	}
	s.db.Delete(&alert)
	return response, nil
}

func (s *dexterAlertsServer) ListIndicators(ctx context.Context, opts *pb.ListIndicatorsRequest) (*pb.ListIndicatorsResponse, error) {
	response := &pb.ListIndicatorsResponse{}
	log.Printf("ListIndicators")
	var indicators []Indicator
	s.db.Find(&indicators)
	spew.Dump(&indicators)
	for _, v := range indicators {
		indicatorSpec := &pb.Indicator{
			Name: v.Name,
			Implementation: v.Implementation,
			Source: v.Source,
			// TODO - Figure out what I should do with the Jsonb fields.
		}
		response.Indicators = append(response.Indicators, indicatorSpec)
	}
	return response, nil
}

// StartServer starts the gRPC service for alert management
func StartServer(listen string, db *gorm.DB) {
	var opts []grpc.ServerOption
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Listening on %s\n", listen)
	}
	server := &dexterAlertsServer{
		db: db,
	}
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	pb.RegisterAlertsServer(grpcServer, server)
	grpcServer.Serve(listener)
}
