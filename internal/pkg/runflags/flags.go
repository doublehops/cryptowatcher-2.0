package runflags

import "flag"

type FlagStruct struct {
	ConfigFile string
}

func GetFlags() FlagStruct {
	configFile := flag.String("config", "config.json", "Config file to use")
	flag.Parse()

	params := FlagStruct{
		ConfigFile: *configFile,
	}

	return params
}
