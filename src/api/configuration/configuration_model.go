package configuration

type configuration struct {
	ConfigurationInstance ConfigurationInstance `yaml:"configuration_instance"`
}

type ConfigurationInstance struct {
	TpvClient  string `yaml:"tpv_client"`
	UserClient string `yaml:"user_client"`
	CurrencyClient string `yaml:"currency_client"`
}

var conf *configuration

func GetConfigurationInstance() *configuration {

	if conf == nil {
		conf = &configuration{}
	}
	return conf
}