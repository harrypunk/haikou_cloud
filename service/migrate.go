package service

import (
	m "github.com/harrypunk/haikou_cloud/model"
)

func Migrate() error {
	db, err := EnvConnection()
	if err != nil {
		return err
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&m.Family{},
			&m.Grade{},
			&m.School{},
			&m.Student{},
			&m.Parent{},
		)
	if err != nil {
		return err
	}
	return nil
}
