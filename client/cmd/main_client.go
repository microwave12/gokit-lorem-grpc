package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/microwave12/gokit-lorem-grpc/client"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr = flag.String("grpcAddr", ":8081", "grpc address")
	)
	flag.Parse()

	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()

	loremService := client.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)

	var minStr, maxStr string

	minStr, args = pop(args)
	maxStr, args = pop(args)

	min, _ := strconv.Atoi(minStr)
	max, _ := strconv.Atoi(maxStr)

	switch cmd {
	case "word":
		msg, err := loremService.Word(ctx, min, max)
		if err != nil {
			fmt.Printf("%s", err)
			break
		}

		fmt.Println(msg)
	case "sentence":
		msg, err := loremService.Sentence(ctx, min, max)
		if err != nil {
			fmt.Printf("%s", err)
			break
		}

		fmt.Println(msg)
	case "paragraph":
		msg, err := loremService.Paragraph(ctx, min, max)
		if err != nil {
			fmt.Printf("%s", err)
			break
		}

		fmt.Println(msg)
	default:
		log.Fatalln("unknown command", cmd)
	}
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
