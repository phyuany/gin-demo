package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"jkdev.cn/api/common"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
