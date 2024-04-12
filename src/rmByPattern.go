package rmByPattern

import (
	"log"
)

func Run() {
	if err := newRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
