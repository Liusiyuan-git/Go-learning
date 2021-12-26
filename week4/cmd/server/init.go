package main

import (
	"fmt"

	"database/sql"
	"github.com/Liusiyuan-git/Go-learning/week4/internal/repository"
	"github.com/Liusiyuan-git/Go-learning/week4/internal/service"
	"github.com/Liusiyuan-git/Go-learning/week4/internal/usecase"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var UserSet = wire.NewSet(
	service.NewUserService,
	repository.NewRepository,
	usecase.NewUserUsecase,
)

// 初始化mysql连接
func NewDB(v *viper.Viper) (*sql.DB, error) {
	client, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		v.Sub("mysql").GetString("role"),
		v.Sub("mysql").GetString("user"),
		v.Sub("mysql").GetString("ip"),
		v.Sub("mysql").GetString("port"),
		v.Sub("mysql").GetString("db")))
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 从config中取出mysql需要得配置参数
func InitConfig() (*viper.Viper, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return viper.GetViper(), nil
}
