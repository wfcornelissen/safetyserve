package main

import (
	"fmt"

	"github.com/wfcornelissen/safetyserve/internal/dbHandling"
)

func main() {
	fmt.Println("Hello World")
}

func init() {
	dbHandling.DBInit()
}
