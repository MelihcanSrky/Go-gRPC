package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/MelihcanSrky/Go-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultText = "You can go"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
	text = flag.String("text", defaultText, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTranslatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Translate(ctx, &pb.TranslationRequest{Text: *text, SourceLanguage: "en-US", TargetLanguage: "tr"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Translated Text: %s", r.GetTranslatedText())
}
