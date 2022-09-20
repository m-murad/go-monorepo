package main

import (
	"fmt"

	"github.com/m-murad/go-monorepo/go-common/aws"
	"github.com/m-murad/go-monorepo/go-common/logger"
)

func main() {
	_ = aws.NewAWS("213")
	_ = logger.NewLogger()
	fmt.Println("--")
}
