package flags

import "flag"

const (
	configName        = "config"
	configDefault     = "source/config.yaml"
	configDescription = "takes path to the config file"
)

func Parse() *string {
	configPath := flag.String(configName, configDefault, configDescription)
	flag.Parse()

	return configPath
}
