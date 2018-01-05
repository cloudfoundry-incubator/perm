package main

import (
	"context"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm/monitor"
	"github.com/satori/go.uuid"
)

const (
	AdminProbeTickDuration = 30 * time.Second
	AdminProbeTimeout      = 3 * time.Second
)

func RunAdminProbe(ctx context.Context, logger lager.Logger, wg *sync.WaitGroup, probe *monitor.AdminProbe, statter *monitor.Statter) {
	defer wg.Done()

	var err error

	for range time.NewTicker(AdminProbeTickDuration).C {
		err = runAdminProbe(ctx, logger, probe)
		if err != nil {
			statter.SendFailedAdminProbe(logger.Session("metrics"))
		} else {
			statter.SendSuccessfulAdminProbe(logger.Session("metrics"))
		}
	}
}

func runAdminProbe(ctx context.Context, logger lager.Logger, probe *monitor.AdminProbe) error {
	u := uuid.NewV4()

	defer probe.Cleanup(ctx, logger.Session("cleanup"), u.String())

	cctx, cancel := context.WithTimeout(ctx, AdminProbeTimeout)
	defer cancel()

	return probe.Run(cctx, logger.Session("run"), u.String())
}
