package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/cactus/go-statsd-client/statsd"
	flags "github.com/jessevdk/go-flags"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"code.cloudfoundry.org/lager"

	"crypto/tls"
	"crypto/x509"

	"strconv"

	"time"

	"code.cloudfoundry.org/perm/cmd"
	cmdflags "code.cloudfoundry.org/perm/cmd/flags"
	"code.cloudfoundry.org/perm/pkg/ioutilx"
	"code.cloudfoundry.org/perm/pkg/monitor"
	"code.cloudfoundry.org/perm/pkg/monitor/recording"
	"code.cloudfoundry.org/perm/pkg/perm"
)

type options struct {
	Perm permOptions `group:"Perm" namespace:"perm"`

	StatsD statsDOptions `group:"StatsD" namespace:"statsd"`

	Logger cmdflags.LagerFlag

	Frequency       time.Duration `long:"frequency" description:"Frequency with which the probe is issued" default:"5s"`
	RequestDuration time.Duration `long:"request-duration" description:"Time after which a request is considered to have failed" default:"100ms"`
	Timeout         time.Duration `long:"timeout" description:"Time after which the probe will cancel a run" default:"10s"`
}

type permOptions struct {
	Hostname      string                 `long:"hostname" description:"Hostname used to resolve the address of Perm" required:"true"`
	Port          int                    `long:"port" description:"Port used to connect to Perm" required:"true"`
	CACertificate []ioutilx.FileOrString `long:"ca-certificate" description:"File path(s) of Perm's CA certificate (and UAA's CA if --require-auth)"`
	RequireAuth   bool                   `long:"require-auth" description:"Enable the monitor to talk to perm using oauth"`
	TokenURL      string                 `long:"token-url" description:"URL to uaa's token endpoint (only required if '--require-auth' is provided)"`
	ClientID      string                 `long:"client-id" description:"UAA Client ID used to fetch token (only required if '--require-auth' is provided)"`
	ClientSecret  string                 `long:"client-secret" description:"UAA Client Secret used to fetch token (only required if '--require-auth' is provided)"`
}

type statsDOptions struct {
	Hostname string `long:"hostname" description:"Hostname used to connect to StatsD server" required:"true"`
	Port     int    `long:"port" description:"Port used to connect to StatsD server" required:"true"`
}

type probeOptions struct {
}

func main() {
	parserOpts := &options{}
	parser := flags.NewParser(parserOpts, flags.Default)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	logger, _ := parserOpts.Logger.Logger("perm-monitor")

	logger.Debug(starting)
	defer logger.Debug(finished)

	//////////////////////
	// Setup StatsD Client
	statsDAddr := net.JoinHostPort(parserOpts.StatsD.Hostname, strconv.Itoa(parserOpts.StatsD.Port))
	statsDClient, err := statsd.NewBufferedClient(statsDAddr, "", 0, 0)
	if err != nil {
		logger.Error(failedToConnectToStatsD, err, lager.Data{
			"addr": statsDAddr,
		})
		os.Exit(1)
	}
	defer statsDClient.Close()
	//////////////////////

	//////////////////////
	// Setup Perm GRPC Client
	//
	//// Setup TLS Credentials
	pool := x509.NewCertPool()

	for _, certPath := range parserOpts.Perm.CACertificate {
		caPem, e := certPath.Bytes(ioutilx.InjectableOS{}, ioutilx.InjectableIOReader{})
		if e != nil {
			logger.Error(failedToReadCertificate, e, lager.Data{
				"location": certPath,
			})
			os.Exit(1)
		}

		if ok := pool.AppendCertsFromPEM(caPem); !ok {
			logger.Error(failedToAppendCertToPool, e, lager.Data{
				"location": certPath,
			})
			os.Exit(1)
		}
	}

	addr := net.JoinHostPort(parserOpts.Perm.Hostname, strconv.Itoa(parserOpts.Perm.Port))
	opts := []perm.DialOption{perm.WithTLSConfig(&tls.Config{RootCAs: pool})}

	if parserOpts.Perm.RequireAuth {
		tsConfig := clientcredentials.Config{
			ClientID:     parserOpts.Perm.ClientID,
			ClientSecret: parserOpts.Perm.ClientSecret,
			TokenURL:     parserOpts.Perm.TokenURL,
		}

		uaaClient := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: pool,
				},
			},
		}
		ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &uaaClient)

		tokenSource := tsConfig.TokenSource(ctx)

		opts = append(opts, perm.WithTokenSource(tokenSource))
	}

	client, err := perm.Dial(addr, opts...)
	recordedClient := recording.NewClient(client, &durationRecorder{})

	if err != nil {
		logger.Error(failedToCreatePermClient, err)
		os.Exit(1)
	}
	defer client.Close()

	probe := monitor.NewProbe(recordedClient)

	statter := monitor.NewStatter(statsDClient)

	ticker := time.NewTicker(parserOpts.Frequency)
	for range ticker.C {
		cmd.RecordProbeResults(logger, probe, statter, parserOpts.RequestDuration, parserOpts.Timeout)
		statter.SendStats(logger)
	}
}

type durationRecorder struct{}

func (r *durationRecorder) Observe(duration time.Duration) error {
	return nil
}
