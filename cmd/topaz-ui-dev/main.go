// Command topaz-ui-dev runs the UI as a standalone server for local development.
package main

import (
	"log"

	"github.com/aserto-dev/go-topaz-ui/internal/web"
)

func main() {
	if err := web.ServeWeb("localhost:8086"); err != nil {
		log.Fatal(err)
	}
}
