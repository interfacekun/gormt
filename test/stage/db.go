package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

func OpenDb() (db *gorm.DB) {
	db, _, err := Open()
	db.Logger.LogMode(0)
	if err != nil {
		hlog.Fatalf("connect db fail %s", err.Error())
		return nil
	}
	return
}
