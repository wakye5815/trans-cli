//設定値読み込み処理の隠蔽

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func BuildAppConfig() (*AppConfig, error) {
	enableProviderType := FindProviderType(viper.GetInt("ENABLE_PROVIDER"))
	config, err := NewAppConfig(viper.GetString("SOURCE_LANGUAGE"), viper.GetString("TARGET_LANGUAGE"), enableProviderType)
	if err != nil {
		return nil, fmt.Errorf("Failed building AppConfig: %w", err)
	}
	return config, nil
}

func BuildIbmApiConnnectionConfig() (*IbmApiConnectionConfig, error) {
	config, err := NewIbmApiConnectionConfig(
		viper.GetString("IBM_API_KEY"),
		viper.GetString("IBM_API_URL"),
		"2018-05-01",
	)
	if err != nil {
		return nil, fmt.Errorf("Failed building IbmApiConnectionConfig: %w", err)
	}
	return config, nil
}
