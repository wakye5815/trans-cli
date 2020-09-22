package infrastructure

type IbmApiConnectionConfig struct {
	Key     string
	Url     string
	Version string
}

func NewIbmApiConnectionConfig(key string, url string, version string) *IbmApiConnectionConfig {
	return &IbmApiConnectionConfig{
		Key:     key,
		Url:     url,
		Version: version,
	}
}
