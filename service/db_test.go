package service

import (
	"os"
	"testing"
)

func setLocalDB() {
	os.Setenv("DB_NAME", "haikou_test")
	os.Setenv("DB_USER", "onetwo")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_URL", "localhost")
}

func TestMigrate(t *testing.T) {
	setLocalDB()
	err := Migrate()
	if err != nil {
		t.Fail()
	}
}
