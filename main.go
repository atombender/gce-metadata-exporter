package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	goflags "github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Flags struct {
	ListenAddress string `long:"listen-address" description:"Address to listen to." value-name:"[HOST][:PORT]" required:"true" default:":9100"`
}

func (f *Flags) parse() error {
	parser := goflags.NewParser(f, goflags.HelpFlag|goflags.PassDoubleDash)
	args, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*goflags.Error); ok && e.Type == goflags.ErrHelp {
			parser.WriteHelp(os.Stdout)
			os.Exit(0)
		}
		return err
	}
	if len(args) > 0 {
		return errors.New("Too many command line arguments")
	}
	return nil
}

func main() {
	var flags Flags
	if err := flags.parse(); err != nil {
		log.Fatal(err)
	}

	NewGCEMetadataCollector().Start()

	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Serving on %s", flags.ListenAddress)
	log.Fatal(http.ListenAndServe(flags.ListenAddress, nil))
}
