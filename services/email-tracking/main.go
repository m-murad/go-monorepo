package main

import (
	"github.com/m-murad/go-monorepo/go-common/aws"
	"github.com/m-murad/go-monorepo/go-common/logger"
)

func main() {
	_ = aws.NewAWS()
	_ = logger.NewLogger()
}
