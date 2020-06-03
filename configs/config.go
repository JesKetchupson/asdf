package config

func ParceConfig() AsdfConfig {
	return AsdfConfig{
		ServerPort: ":8080", // todo get from somewhere
		DBUri:      "",
	}
}

type AsdfConfig struct {
	ServerPort string
	DBUri      string
}
