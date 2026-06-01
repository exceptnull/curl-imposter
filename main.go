package main

import (
	"fmt"
	"os"

	"github.com/exceptnull/curl-imposter/cmd" // Update this import if your module name is different!
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}