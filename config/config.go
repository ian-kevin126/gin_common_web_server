package config

type Configuration struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	MySQL MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql PGSQL `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}
