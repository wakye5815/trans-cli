package config

import "errors"

type ProviderType int

const (
	Invalid ProviderType = iota
	Ibm
	Weblio
)

func FindProviderType(value int) ProviderType {
	switch value {
	case int(Ibm):
		return Ibm
	case int(Weblio):
		return Weblio
	default:
		return Invalid
	}
}

type AppConfig struct {
	SourceLanguage     string
	TargetLanguage     string
	EnableProviderType ProviderType
}

func NewAppConfig(sourceLanguage string, targetLanguage string, enableProviderType ProviderType) (*AppConfig, error) {
	if len(sourceLanguage) == 0 {
		return nil, errors.New("sourceLanguage is empty")
	}
	if len(targetLanguage) == 0 {
		return nil, errors.New("targetLanguage is empty")
	}
	if enableProviderType == Invalid {
		return nil, errors.New("Invalid enableProviderType")
	}

	return &AppConfig{
		SourceLanguage:     sourceLanguage,
		TargetLanguage:     targetLanguage,
		EnableProviderType: enableProviderType,
	}, nil
}
