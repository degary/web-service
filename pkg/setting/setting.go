package setting

import "github.com/spf13/viper"

type setting struct {
	vp *viper.Viper
}

func NewSetting() (*setting,error)  {
	vp :=viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &setting{vp},nil
}
