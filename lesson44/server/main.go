package main

import (
	"context"
	"log"
	"net"
	pb "translater/protos/translate"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTranslaterServer
	Translations map[string]string
}

func main() {
	translations := map[string]string{
		"Hello":        "Salom",
		"World":        "Dunyo",
		"Go":           "Bor",
		"Thank you":    "Rahmat",
		"Goodbye":      "Xayr",
		"Please":       "Iltimos",
		"Yes":          "Ha",
		"No":           "Yo'q",
		"Friend":       "Do'st",
		"Family":       "Oila",
		"Love":         "Sevgi",
		"Happy":        "Baxtli",
		"Sad":          "Qayg'uli",
		"Morning":      "Ertalab",
		"Night":        "Kechasi",
		"Food":         "Ovqat",
		"Water":        "Suv",
		"House":        "Uy",
		"School":       "Maktab",
		"Work":         "Ish",
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	s := grpc.NewServer()

	pb.RegisterTranslaterServer(s, &Server{Translations: translations})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *Server) GetTranslateWord(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	translatedWord := make(map[string]string)

	for _, word := range req.Word {
		if translation, ok := s.Translations[word]; ok {
			translatedWord[word] = translation
		} else {
			translatedWord[word] = "Not found"
		}
	}
	return &pb.TranslateResponse{TranslatedWord: translatedWord}, nil
}
