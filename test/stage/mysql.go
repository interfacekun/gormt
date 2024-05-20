package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "embed"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	MySQLConf struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Database     string `yaml:"database"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		MaxOpenConns int    `yaml:"maxOpenConns"` // 最大连接数
		MaxIdleConns int    `yaml:"maxIdleConns"` // 最大空闲连接
		MaxLifetime  int    `yaml:"maxLifetime"`  // mysql timeout 单位分钟，要小于mysql设置的time_out
		ReplicaSet   int    `yaml:"replicaSet"`   // 副本数，一个db开多少个service
		Parallel     int    `yaml:"parallel"`     // 每个service同时执行多少个sql
		SqlTimeout   int    `yaml:"sqlTimeout"`   // sql执行超时，单位秒
	}
)

//go:embed config.yaml
var configData []byte

func loadConfig() (*MySQLConf, error) {

	config := &MySQLConf{}
	err := yaml.Unmarshal(configData, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func Open() (db *gorm.DB, sqlDb *sql.DB, err error) {
	conf, err := loadConfig()
	if err != nil {
		return
	}
	dialector := mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?"+
			"charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password,
		conf.Host, conf.Port, conf.Database,
	))
	db, err = gorm.Open(
		dialector,
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		},
	)
	if err != nil {
		hlog.Fatalf("mysql init failed, err: ", err)
		return
	}
	sqlDb, err = db.DB()
	if err != nil {
		hlog.Fatalf("mysql get db failed, err: ", err)
		return
	}
	sqlDb.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDb.SetMaxIdleConns(conf.MaxIdleConns)
	// mysql default conn timeout=8h, should < mysql_timeout
	sqlDb.SetConnMaxLifetime(
		time.Minute * time.Duration(conf.MaxLifetime),
	)
	err = sqlDb.Ping()
	if err != nil {
		hlog.Fatalf("mysql ping failed, err: ", err)
	}
	hlog.Infof("mysql conn pool has initiated.")
	return
}
