package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wakye5815/trans-cli/config"
	"github.com/wakye5815/trans-cli/infrastructure"
	"github.com/wakye5815/trans-cli/service"
)

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "trans",
		Short: "Translation words",
		Run: func(cmd *cobra.Command, args []string) {
			appConfig, err := config.BuildAppConfig()
			if err != nil {
				panic(err)
			}

			ibmApiConnectionConfig, err := config.BuildIbmApiConnnectionConfig()
			if err != nil {
				panic(err)
			}

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
}
