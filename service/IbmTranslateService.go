package service

import (
	"example.com.test/config"
	"example.com.test/infrastructure"
)

type IbmTranslateService struct {
	appConfig                      *config.AppConfig
	ibmLanguageTranslatorApiClient *infrastructure.IbmLanguageTranslatorApiClient
}

func NewIbmTranslateService(
	appConfig *config.AppConfig,
	ibmLanguageTranslatorApiClient *infrastructure.IbmLanguageTranslatorApiClient,
) *IbmTranslateService {
	return &IbmTranslateService{
		appConfig:                      appConfig,
		ibmLanguageTranslatorApiClient: ibmLanguageTranslatorApiClient,
	}
}

func (this *IbmTranslateService) Translate(text string) (translatedText []string, err error) {
	requestBody := &infrastructure.PostTranslateRequest{
		Text:   []string{text},
		Source: this.appConfig.SourceLanguage,
		Target: this.appConfig.TargetLanguage,
	}

	response, err := this.ibmLanguageTranslatorApiClient.PostTranslate(requestBody)
	if err != nil {
		return nil, err
	}

	for _, translation := range response.Translations {
		translatedText = append(translatedText, translation.Translation)
	}

	return
}

func (this *IbmTranslateService) GetAvailableLanguages() (languages []string, err error) {
	response, err := this.ibmLanguageTranslatorApiClient.GetLanguages()
	if err != nil {
		return nil, err
	}

	for _, language := range response.Languages {
		languages = append(languages, language.Language)
	}

	return
}
