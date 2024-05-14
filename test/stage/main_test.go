package main

import (
	"testing"

	md "github.com/xxjwxc/gormt/model"
	"gorm.io/gorm"
)

func TestMiniGameLog(t *testing.T) {
	mgl := &md.Minigamelog{Nickname: "ikun"}
	db := OpenDb()
	stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
}
