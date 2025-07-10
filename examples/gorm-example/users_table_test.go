package gorm_example

import (
	"fmt"
	"testing"

	"github.com/maodou24/gorm-example/models"
	"github.com/maodou24/gorm-example/pg"
)

func TestCreateTable(t *testing.T) {
	if err := pg.GetDb().AutoMigrate(&models.User{}); err != nil {
		t.Error(err)
	}
}

func TestUserTableInsert(t *testing.T) {
	user := models.User{Name: "maodou", Age: 18}
	pg.GetDb().Create(&user)
}

func TestUserTableQueryFirst(t *testing.T) {
	// query
	var result models.User
	pg.GetDb().First(&result)

	fmt.Printf("%+v", result)
}

func TestUserTableQuery(t *testing.T) {
	// query
	user := models.User{Name: "maodou"}
	pg.GetDb().Find(&user)

	fmt.Printf("%+v", user)
}

func TestUserTableDelete(t *testing.T) {
	deleteUser := models.User{Name: "maodou"}
	if err := pg.GetDb().Where("name = ?", deleteUser.Name).Delete(&models.User{}).Error; err != nil {
		t.Error(err)
	}
}

func TestUserTableDeleteByAge(t *testing.T) {
	pg.GetDb().Where("age = ?", 18).Delete(&models.User{})
}
