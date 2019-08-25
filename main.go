package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jimmysawczuk/try"
	"github.com/pkg/errors"
)

var timeout = 60 * time.Second
var interval = 1 * time.Second
var host string

func init() {
	flag.DurationVar(&timeout, "timeout", 60*time.Second, "total amount of time to try")
	flag.DurationVar(&interval, "interval", 1*time.Second, "amount of time to wait between tries")
	flag.StringVar(&host, "host", "http://localhost:5601", "host of kibana server")

	flag.Parse()
}

func main() {
	start := time.Now()

	log.Printf("attempting to connect to kibana (trying for %s, %s between attempts)", timeout, interval)

	if err := try.Try(connectToKibana(host), timeout, interval); err != nil {
		log.Fatal(err)
	}

	log.Printf("connected in %s", time.Now().Sub(start).Truncate(time.Millisecond))
}

// connectToKibana returns a function which attempts to connect to a kibana server and ping it using the endpoint provided.
// It will ping the api status endpoint and wait for a 200 to be returned on it.
func connectToKibana(endpoint string) func() error {
	return func() error {
		client := &http.Client{
			Timeout: time.Second * 1,
		}

		resp, err := client.Get(fmt.Sprintf("%v/api/status", endpoint))
		if err != nil {
			return errors.Wrap(err, "get on kibana status")
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("service returning %v", resp.StatusCode)
		}

		return nil
	}
}
