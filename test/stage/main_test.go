package main

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	md "github.com/xxjwxc/gormt/model"
	"gorm.io/gorm"
)

func TestMiniGameLog(t *testing.T) {
	mgl := &md.Minigamelog{Nickname: "ikun"}
	db := OpenDb()
	// stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	// t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(mgl)
	})

	hlog.Infof("test mini gameLog sql:\n%s\n", sql)

}
