package main

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Configuration string `short:"c" long:"configuration" description:"the configuration file" optional:"yes" default:"supervisord.conf"`
}

var (
	client *http.Client
	pool   *x509.CertPool
)

func init() {
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(pemCerts)
	client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}}
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	var options Options
	var parser = flags.NewParser(&options, flags.Default)
	parser.Parse()
	s := NewSupervisor(options.Configuration)
	s.Reload()

}
