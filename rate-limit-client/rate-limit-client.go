// Package main implements a client for rate-limiter service.
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sam09/rate-limiter/token-bucket"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	bucketName   = "test-bucket"
	maxAmount    = 1000
	refillTime   = 60 * 60
	refillAmount = 1000
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTokenBucketClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r1, err := c.CreateBucket(ctx, &pb.CreateBucketRequest{Name: bucketName, MaxAmount: maxAmount,
		RefillTime: refillAmount, RefillAmount: refillAmount})
	if err != nil {
		log.Fatalf("could not create bucket: %v", err)
	}
	log.Printf("Test Bucket created: %s", r1.GetBucketName())

	_, err = c.AddToken(ctx, &pb.AddTokenRequest{BucketName: bucketName})
	if err != nil {
		log.Fatalf("Could not add token: %v", err)
	}

	r2, err := c.ConsumeToken(ctx, &pb.ConsumeTokenRequest{BucketName: bucketName})
	if err != nil {
		log.Fatalf("Could not consume token: %v", err)
	}
	log.Printf("Greeting: %s", r2.GetToken())
}
