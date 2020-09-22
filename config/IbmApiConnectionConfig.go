package config

import "errors"

type IbmApiConnectionConfig struct {
	Key     string
	Url     string
	Version string
}

func NewIbmApiConnectionConfig(key string, url string, version string) (*IbmApiConnectionConfig, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	if len(url) == 0 {
		return nil, errors.New("url is empty")
	}
	if len(version) == 0 {
		return nil, errors.New("version is empty")
	}

	return &IbmApiConnectionConfig{
		Key:     key,
		Url:     url,
		Version: version,
	}, nil
}
