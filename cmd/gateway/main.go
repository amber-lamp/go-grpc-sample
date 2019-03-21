package main

import (
	"fmt"
	"os"

	"github.com/yuuyamad/go-grpc-sample/pkg/gateway"
)

func main() {
	if err := gateway.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
