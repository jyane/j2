package main

import (
	"fmt"
	"os"

	"j2/golang/util"
	configpb "j2/huer/config"
)

func main() {
	var config configpb.Config
	if err := util.ReadProtoTextFromLocalStorage("huer/config.pb.txt", &config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(config.DiscoveryEndpoint)
}
