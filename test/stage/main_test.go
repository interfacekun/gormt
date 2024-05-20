package main

import (
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	md "github.com/xxjwxc/gormt/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func TestStageGameStatsDay(t *testing.T) {
	mgl := &md.Stagegamestatsday{Time: time.Now()}
	db := OpenDb()
	// stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	// t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(mgl)
	})
	hlog.Infof("test stage game stats day sql:\n%s\n", sql)

	// INSERT INTO `stageGameStatsDay` (`time`,`area`,`stage`,`star`,`playersAll`,`players`,`passPlayers`,`passGameAvg`,`feeCostAvg`,`passGameMax`,`passGameMin`,`goldCost`) VALUES ('2024-05-20 10:52:57.276',0,0,0,0,0,0,0,0,0,0,0) ON DUPLICATE KEY UPDATE `playersAll`=VALUES(`playersAll`),`players`=VALUES(`players`),`passPlayers`=VALUES(`passPlayers`),`passGameAvg`=VALUES(`passGameAvg`),`feeCostAvg`=VALUES(`feeCostAvg`),`passGameMax`=VALUES(`passGameMax`),`passGameMin`=VALUES(`passGameMin`),`goldCost`=VALUES(`goldCost`)

}

func TestStageUser(t *testing.T) {
	mgl := &md.Stageuser{
		Area: 1,
	}
	db := OpenDb()
	// stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	// t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(mgl)
	})
	hlog.Infof("test stage user sql:\n%s\n", sql)

	// INSERT INTO `stageUser` (`area`,`stage`,`star`,`userID`,`createdAt`) VALUES (1,0,0,0,'0000-00-00 00:00:00') ON DUPLICATE KEY UPDATE `createdAt`=VALUES(`createdAt`)
}

func TestUserGameStatDay(t *testing.T) {
	mgl := &md.Usergamestatsday{
		Area: 1,
	}
	db := OpenDb()
	// stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	// t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(mgl)
	})
	hlog.Infof("test user game stats day sql:\n%s\n", sql)

	// INSERT INTO `userGameStatsDay` (`time`,`area`,`stage`,`star`,`userID`,`games`,`state`,`passTime`,`feeCost`,`goldCost`) VALUES ('0000-00-00 00:00:00',1,0,0,0,0,0,'0000-00-00 00:00:00',0,0) ON DUPLICATE KEY UPDATE `games`=VALUES(`games`),`state`=VALUES(`state`),`passTime`=VALUES(`passTime`),`feeCost`=VALUES(`feeCost`),`goldCost`=VALUES(`goldCost`)
}

func TestUserGameStats(t *testing.T) {
	mgl := &md.Usergamestats{
		Area: 1,
	}
	db := OpenDb()
	// stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	// t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(mgl)
	})
	hlog.Infof("test user game stats sql:\n%s\n", sql)

	// INSERT INTO `userGameStats` (`area`,`stage`,`star`,`userID`,`games`,`state`,`passTime`,`feeCost`,`goldCost`,`startGold`,`leftGold`) VALUES (1,0,0,0,0,0,'0000-00-00 00:00:00',0,0,0,0) ON DUPLICATE KEY UPDATE `games`=VALUES(`games`),`state`=VALUES(`state`),`passTime`=VALUES(`passTime`),`feeCost`=VALUES(`feeCost`),`goldCost`=VALUES(`goldCost`),`startGold`=VALUES(`startGold`),`leftGold`=VALUES(`leftGold`)
}

// db.Debug().Clauses(clause.OnConflict{
// 	Columns:   []clause.Column{{Name: "user_name"},{Name: "sex"}},
// 	DoUpdates: clause.AssignmentColumns([]string{"nick_name"}),
// }).Create(&users)
