package main

import (
	"alertagator/connectors"
	"os"
)

func main() {
	args := os.Args[1:]
	connectors.DatadogEvents(args)
}
