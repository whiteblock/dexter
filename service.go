package dexter

import (
	"context"
	"log"
	"net"
	grpc "google.golang.org/grpc"
	pb "github.com/whiteblock/dexter/api/alerts"
)

type dexterAlertsServer struct {
}

func (s *dexterAlertsServer) CreateAlert(ctx context.Context, alert *pb.Alert) (*pb.Alert, error) {
	newAlert := &pb.Alert{}
	return newAlert, nil
}

func (s *dexterAlertsServer) ListAlerts(ctx context.Context, opts *pb.ListAlertsRequest) (*pb.ListAlertsResponse, error) {
	response := &pb.ListAlertsResponse{}
	return response, nil
}

func (s *dexterAlertsServer) GetAlert(ctx context.Context, opts *pb.GetAlertRequest) (*pb.Alert, error) {
	alert := &pb.Alert{}
	return alert, nil
}

func (s *dexterAlertsServer) UpdateAlert(ctx context.Context, alert *pb.Alert) (*pb.Alert, error) {
	updatedAlert := &pb.Alert{}
	return updatedAlert, nil
}

func (s *dexterAlertsServer) DeleteAlert(ctx context.Context, opts *pb.DeleteAlertRequest) (*pb.DeleteAlertResponse, error) {
	response := &pb.DeleteAlertResponse{}
	return response, nil
}

func (s *dexterAlertsServer) ListIndicators(ctx context.Context, opts *pb.ListIndicatorsRequest) (*pb.ListIndicatorsResponse, error) {
	response := &pb.ListIndicatorsResponse{}
	return response, nil
}

func newServer() *dexterAlertsServer {
	s := &dexterAlertsServer{}
	return s
}

// StartServer starts the gRPC service for alert management
func StartServer(listen string) {
	var opts []grpc.ServerOption
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Listening on %s\n", listen)
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAlertsServer(grpcServer, newServer())
	grpcServer.Serve(listener)
}
