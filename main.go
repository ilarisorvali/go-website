package main

import (
	"fmt"
	"os"

	"github.com/ilarisorvali/go-website/internal/server"
)

func main() {

	if err := server.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
