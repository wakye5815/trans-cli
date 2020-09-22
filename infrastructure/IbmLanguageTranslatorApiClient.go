package infrastructure

import (
	"example.com.test/lib"
	"github.com/dghubble/sling"
)

type IbmLanguageTranslatorApiClient struct {
	sling   *sling.Sling
	version string
}

func NewIbmLanguageTranslatorApiClient(config *IbmApiConnectionConfig) *IbmLanguageTranslatorApiClient {
	return &IbmLanguageTranslatorApiClient{
		sling:   lib.NewCustomSling().Base(config.Url).SetBasicAuth("apikey", config.Key),
		version: config.Version,
	}
}

type IbmCommonParams struct {
	Version string `url:"version,omitempty"`
}

type GetLanguagesResponse struct {
	Languages []struct {
		Language           string `json:"language"`
		LanguageName       string `json:"language_name"`
		NativeLanguageName string `json:"native_language_name"`
		CountryCode        string `json:"country_code"`
		WordsSeparated     bool   `json:"words_separated"`
		Direction          string `json:"direction"`
		SupportedAsSource  bool   `json:"supported_as_source"`
		SupportedAsTarget  bool   `json:"supported_as_target"`
		Identifiable       bool   `json:"identifiable"`
	} `json:"languages"`
}

func (this *IbmLanguageTranslatorApiClient) GetLanguages() (*GetLanguagesResponse, error) {
	params := &IbmCommonParams{Version: this.version}
	response := new(GetLanguagesResponse)

	_, err := this.sling.Get("/v3/languages").QueryStruct(params).ReceiveSuccess(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type PostTranslateRequest struct {
	Text   []string `json:"text"`
	Source string   `json:"source"`
	Target string   `json:"target"`
}

type PostTranslateResponse struct {
	Translations []struct {
		Translation string `json:"translation"`
	} `json:"translations"`
	WordCount      int `json:"word_count"`
	CharacterCount int `json:"character_count"`
}

func (this *IbmLanguageTranslatorApiClient) PostTranslate(body *PostTranslateRequest) (*PostTranslateResponse, error) {
	params := &IbmCommonParams{Version: this.version}
	response := new(PostTranslateResponse)

	_, err := this.sling.Post("/v3/translate").QueryStruct(params).BodyJSON(body).ReceiveSuccess(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
