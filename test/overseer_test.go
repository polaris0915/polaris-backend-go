package test

import (
	"fmt"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestOverseer(t *testing.T) {
	overseer.Run(overseer.Config{
		Program: func(state overseer.State) {
			log.Printf("app (%s) listening...", state.ID)
			http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "app (%s) says hello!\n", state.ID)
			}))
			http.Serve(state.Listener, nil)
		},
		Address: ":8080",
		Fetcher: &fetcher.HTTP{
			URL:      "http://localhost:4000/binaries/myapp",
			Interval: 1 * time.Second,
		},
	})
}
