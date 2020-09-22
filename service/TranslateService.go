package service

type TranslateService interface {
	Translate(text string) (translatedText []string, err error)
	GetAvailableLanguages() (languages []string, err error)
}
