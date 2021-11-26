package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path"
	"sheetServerApi/global"
	model "sheetServerApi/internal/model/db"
	"sheetServerApi/internal/routers"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

//	init初始化函数
func init() {
	// 注入全局系统配置文件
	if err := setupSetting(); err != nil {
		logrus.Fatal(err)
		return
	}
	// 注入数据库连接池
	if err := setupDBEngine(); err != nil {
		logrus.Fatal(err)
		return
	}
	logrus.Info("init success...")
}

// 主函数
func main() {
	// 设置运行模式(release or debug ?)
	gin.SetMode(global.ServerSetting.RunMode)
	// 打印日志
	if global.ServerSetting.RunMode == "debug"{
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
			TimestampFormat:"2006-01-02 15:04:05",
		})
	}
	// 挂载路由
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Println("listen in " + global.ServerSetting.HttpPort)
	s.ListenAndServe()

}

//初始化全局服务器配置文件
func setupSetting() error {
	setting, err := NewSettings()
	if err != nil {
		return err
	}
	if err := setting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err := setting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err := setting.ReadSection("DatabaseOrm", &global.DatabaseOrmSetting); err != nil {
		return err
	}
	if err := setting.ReadSection("DatabaseSqlx", &global.DatabaseSqlxSetting); err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

//初始化全局数据库连接池
func setupDBEngine() error {
	var err error
	global.DBOrmEngine, err = model.NewDBOrmEngine(global.DatabaseOrmSetting)
	if err != nil {
		return err
	}
	global.DBSqlxEngine, err = model.NewDBSqlxEngine(global.DatabaseSqlxSetting)
	if err != nil {
		return err
	}
	return nil
}

// 读取配置文件
func NewSettings() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	// 载入全局配置文件
	dir, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
	vp.AddConfigPath(path.Join(dir, "configs"))
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
