package main

import (
	"flag"
	"fmt"

	pb "github.com/morix1500/cloudendpoint/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type credential struct {
	key     string
	referer string
	jwt     string
}

func (c credential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"x-api-key":     c.key,
		"referer":       c.referer,
		"authorization": "Bearer " + c.jwt,
	}, nil
}

func (credential) RequireTransportSecurity() bool {
	return false
}

func main() {
	var addr, msg, key, referer, jwt string
	flag.StringVar(&addr, "addr", "127.0.0.1:50051", "server address")
	flag.StringVar(&msg, "msg", "Hello", "message")
	flag.StringVar(&key, "key", "invalid", "API Key")
	flag.StringVar(&referer, "referer", "invalid", "referer")
	flag.StringVar(&jwt, "jwt", "invalid", "JSON Web Token")
	flag.Parse()

	cred := credential{
		key:     key,
		referer: referer,
		jwt:     jwt,
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithPerRPCCredentials(cred))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewEchoServiceClient(conn)

	ctx := context.Background()
	req := &pb.Request{Message: &pb.Msg{Message: msg}}

	res, err := c.Echo1(ctx, req)
	if err == nil {
		fmt.Println("Echo1: succeeded: ", res.Message)
	} else {
		fmt.Println("Echo1: failed: ", err)
	}
	res, err = c.Echo2(ctx, req)
	if err == nil {
		fmt.Println("Echo2: succeeded: ", res.Message)
	} else {
		fmt.Println("Echo2: failed: ", err)
	}
	res, err = c.Echo3(ctx, req)
	if err == nil {
		fmt.Println("Echo3: succeeded: ", res.Message)
	} else {
		fmt.Println("Echo3: failed: ", err)
	}
	res, err = c.Echo4(ctx, req)
	if err == nil {
		fmt.Println("Echo4: succeeded: ", res.Message)
	} else {
		fmt.Println("Echo4: failed: ", err)
	}
}
