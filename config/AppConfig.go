package config

import "errors"

type ProviderType int

const (
	Invalid ProviderType = iota
	Ibm
	Weblio
)

func FindProviderType(value int) (ProviderType, error) {
	switch value {
	case int(Ibm):
		return Ibm, nil
	case int(Weblio):
		return Weblio, nil
	default:
		return Invalid, errors.New("")
	}
}

type AppConfig struct {
	SourceLanguage     string
	TargetLanguage     string
	EnableProviderType ProviderType
}
