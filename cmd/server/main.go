package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/MelihcanSrky/Go-gRPC/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct {
	pb.UnimplementedTranslatorServer
}

func (s *server) Translate(ctx context.Context, req *pb.TranslationRequest) (*pb.TranslationResponse, error) {
	translatedText := reverseString(req.Text)
	fmt.Printf("Translated Text: %s", translatedText)
	return &pb.TranslationResponse{TranslatedText: translatedText}, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
