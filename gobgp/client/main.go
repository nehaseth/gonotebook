package main

import (
	"fmt"
	
	"github.com/osrg/gobgp/client"
	"github.com/osrg/gobgp/config"
)

func main()  {
	client, err := client.New("localhost:50051")
	err = client.StartServer(&config.Global{
		Config: config.GlobalConfig{
			As:       1,
			RouterId: "1.1.1.1",
			Port:     1790,
		},
	})
	
	fmt.Println("err ", err)
	err = client.AddNeighbor(&config.Neighbor{
		Config: config.NeighborConfig{
			NeighborAddress: "10.0.0.1",
			PeerAs:          2,
		},
	})
	_, err = client.GetNeighbor("10.0.0.1")
	_, err = client.GetNeighbor("10.0.0.2")
	
}