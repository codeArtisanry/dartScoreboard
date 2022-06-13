package controllers_test

import (
	models "dartscoreboard/models/database"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("start")
	os.Chdir("..")
	models.Migration("test.db")
	dbtest := models.Database("test.db")

	if dbtest == nil {
		fmt.Println("database not connected ")
	}

	ret := m.Run()
	fmt.Println("end")
	teardown()
	os.Exit(ret)
}

func teardown() {
	os.Remove("test.db")
}
