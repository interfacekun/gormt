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

func TestUserGameStatsUpdate(t *testing.T) {
	mgl := &md.Usergamestats{
		Area:   1,
		Stage:  1,
		Star:   1,
		Userid: 1,
	}
	db := OpenDb()
	// stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	// t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&mgl).Updates(
			map[string]interface{}{
				"games":    gorm.Expr("`games` + ?", 1),
				"feeCost":  gorm.Expr("`feeCost` + ?", 1),
				"goldCost": gorm.Expr("`goldCost` + ?", 0),
			},
		)
	})
	hlog.Infof("test user game stats update sql:\n%s\n", sql)

	// UPDATE `userGameStats` SET `feeCost`=`feeCost` + 1,`games`=`games` + 1,`goldCost`=`goldCost` + 0 WHERE `area` = 1 AND `stage` = 1 AND `star` = 1 AND `userID` = 1
}

func TestUserGameStatsFind(t *testing.T) {
	mgl := &md.Usergamestats{
		Area:   1,
		Stage:  1,
		Star:   1,
		Userid: 1,
	}
	db := OpenDb()
	// stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	// t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.First(mgl)
	})
	hlog.Infof("test user game stats find sql:\n%s\n", sql)

	// SELECT * FROM `userGameStats` WHERE `userGameStats`.`area` = 1 AND `userGameStats`.`stage` = 1 AND `userGameStats`.`star` = 1 AND `userGameStats`.`userID` = 1 ORDER BY `userGameStats`.`area` LIMIT 1
}

func TestStageGameStatsDayFind(t *testing.T) {
	mgl := &md.Stagegamestatsday{
		Time:  time.Now(),
		Area:  1,
		Stage: 1,
		Star:  1,
	}
	db := OpenDb()
	// stmt := db.Session(&gorm.Session{DryRun: true}).Create(mgl).Statement
	// t.Logf("test mini gameLog sql:\n%s\n", stmt.SQL.String())
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.First(mgl)
	})
	hlog.Infof("test stage game stats day find sql:\n%s\n", sql)

	// SELECT * FROM `stageGameStatsDay` WHERE `stageGameStatsDay`.`time` = '2024-05-21 15:48:23.449' AND `stageGameStatsDay`.`area` = 1 AND `stageGameStatsDay`.`stage` = 1 AND `stageGameStatsDay`.`star` = 1 ORDER BY `stageGameStatsDay`.`time` LIMIT 1
}

// db.Debug().Clauses(clause.OnConflict{
// 	Columns:   []clause.Column{{Name: "user_name"},{Name: "sex"}},
// 	DoUpdates: clause.AssignmentColumns([]string{"nick_name"}),
// }).Create(&users)
