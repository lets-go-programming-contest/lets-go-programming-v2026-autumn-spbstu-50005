package cli

import (
	"flag"
	"fmt"
)

func MustGetConfigPath() string {
	var configPath *string = flag.String("config", "", "path to file does not find")
	flag.Parse()

	if *configPath == "" {
		panic("flag -config is necessary, but not set")
	}

	fmt.Println(*configPath)

	return *configPath
}
