package main

import (
	"github.com/ugol/gospan/serf"
	"flag"
)

func main() {

	nodeName := flag.String("node", "uno", "node name")
	serfBindAddr := flag.String("serf-bind", "0.0.0.0:7946", "serf bind addr")
	serfRpcAddr := flag.String("serf-rpc", "0.0.0.0:7373", "serf rpc bind addr")
	serfJoinAddr := flag.String("serf-join", "", "serf join bind addr")
	flag.Parse()
	serf.Start(*nodeName, *serfBindAddr, *serfRpcAddr, *serfJoinAddr)

	/*resultCh := serf.Start(*nodeName, *serfBindAddr, *serfRpcAddr, *serfJoinAddr)
	for {
		select {
		case <-resultCh:
			log.Println("Boh")
		}
		time.Sleep(10 * time.Second)
	}
	*/
}