package command

import (
	"fmt"

	"example.com.test/config"
	"example.com.test/infrastructure"
	"example.com.test/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCommand = &cobra.Command{
	Use:   "trans",
	Short: "Translation words",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile("./.env")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("file not found")
			} else {
				panic(err)
			}
		}

		enableProviderType, err := config.FindProviderType(viper.GetInt("ENABLE_PROVIDER"))
		appConfig := &config.AppConfig{
			SourceLanguage:     viper.GetString("SOURCE_LANGUAGE"),
			TargetLanguage:     viper.GetString("TARGET_LANGUAGE"),
			EnableProviderType: enableProviderType,
		}

		ibmApiConnectionConfig := infrastructure.NewIbmApiConnectionConfig(
			viper.GetString("IBM_API_KEY"),
			viper.GetString("IBM_API_URL"),
			"2018-05-01",
		)

		client := infrastructure.NewIbmLanguageTranslatorApiClient(ibmApiConnectionConfig)
		var service service.TranslateService = service.NewIbmTranslateService(appConfig, client)

		availableLanguages, err := service.GetAvailableLanguages()
		if err != nil {
			panic(err)
		}

		translatedTexts, err := service.Translate("リンゴ")
		if err != nil {
			panic(err)
		}

		fmt.Println(availableLanguages)
		fmt.Println(translatedTexts)
	},
}
