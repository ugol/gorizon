package serf

import (
	"github.com/mitchellh/cli"
	"github.com/hashicorp/serf/command/agent"
	"os"
)

func Start(nodeName string, serfBindAddr string, serfRPCAddr string, serfJoinAddr string) int {

	shutdownCh := make(chan struct{})
	defer close(shutdownCh)

	ui := &cli.BasicUi{ErrorWriter:os.Stderr, Writer:os.Stdout}
	c := &agent.Command{
		ShutdownCh: shutdownCh,
		Ui:         ui,
	}

	args := []string{
		"-node", nodeName,
		"-bind", serfBindAddr,
		"-rpc-addr", serfRPCAddr,
		"-log-level", "info",
	}

	if serfJoinAddr != "" {
		args = append(args, "-join", serfJoinAddr)
	}

	resultCh := make(chan int)
	go func() {
		resultCh <- c.Run(args)
	}()

	select {
	case <-resultCh:
		return 0
	}

	return 0

}
