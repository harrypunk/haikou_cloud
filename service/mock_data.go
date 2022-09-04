package service

import (
	"github.com/harrypunk/haikou_cloud/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddMockData(db *gorm.DB) error {
	return nil
}

func AddGrades(db *gorm.DB) error {
	var m1 = model.Grade{
		Model:       gorm.Model{ID: 101},
		Num:         7,
		DisplayText: "初一",
	}
	var m2 = model.Grade{
		Model:       gorm.Model{ID: 102},
		Num:         8,
		DisplayText: "初二",
	}
	var m3 = model.Grade{
		Model:       gorm.Model{ID: 103},
		Num:         9,
		DisplayText: "初三",
	}
	var h1 = model.Grade{
		Model:       gorm.Model{ID: 104},
		Num:         10,
		DisplayText: "高一",
	}
	var h2 = model.Grade{
		Model:       gorm.Model{ID: 105},
		Num:         11,
		DisplayText: "高二",
	}
	var h3 = model.Grade{
		Model:       gorm.Model{ID: 106},
		Num:         12,
		DisplayText: "高三",
	}
	var grades = []model.Grade{m1, m2, m3, h1, h2, h3}
	var result = db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&grades)
	return result.Error
}
