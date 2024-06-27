package service

import (
	"context"
	pb "new/genproto/transportService"
	"new/storage/postgres"
)

type serverTransport struct {
	pb.UnimplementedTransportServiceServer
	Db *postgres.TransportRepo
}

func NewTransportService(db *postgres.TransportRepo) *serverTransport {
	return &serverTransport{Db: db}
}

func (s *serverTransport) GetBusSchedule(ctx context.Context, req *pb.BusNumber) (*pb.GetBusScheduleResponse, error) {
	return s.Db.GetBusSchedule(req)
}

func (s *serverTransport) TrackBusLocation(ctx context.Context, req *pb.BusNumber) (*pb.TrackBusLocationResponse, error) {
	return s.Db.GetBusLocation(req)
}

func (s *serverTransport) ReportTrafficJam(ctx context.Context, req *pb.ReportTrafficJamRequest) (*pb.ReportTrafficJamResponse, error) {
	err := s.Db.UpdateFeedback(req)
	if err != nil {
		return &pb.ReportTrafficJamResponse{Isaccepted: false}, err
	} else {
		return &pb.ReportTrafficJamResponse{Isaccepted: true}, err
	}
}
