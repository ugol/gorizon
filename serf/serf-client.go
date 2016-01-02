package serf

import (
	"github.com/hashicorp/serf/client"
	"log"
)


func TestConnection() {

	config := client.Config {
		Addr:"127.0.0.1:7373",
	}

	rpc, err := client.ClientFromConfig(&config)
	if (err != nil) {
		log.Fatal(err)
	}

	members, err := rpc.Members()
	if (err != nil) {
		log.Fatal(err)
	}
	for _, member := range members {
		log.Printf("Member: %s\n", member)
	}

	list := []string{"0.0.0.0:7946"}

	i, err := rpc.Join(list, true)
	if (err != nil) {
		log.Fatal(err)
	}
	log.Printf("Member: %d\n", i)
	members, err = rpc.Members()
	if (err != nil) {
		log.Fatal(err)
	}
	for _, member := range members {
		log.Printf("Member: %s\n", member)
	}

	ch := make(chan map[string]interface{})
	stream, err := rpc.Stream("", ch)

	for {
		x := <-ch
		log.Printf("%s\n", x)
	}

	defer rpc.Stop(stream)

}