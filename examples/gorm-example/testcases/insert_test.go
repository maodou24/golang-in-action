package testcases

import (
	"fmt"
	"github.com/maodou24/gorm-example/models"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	user := models.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.Create(&user)

	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}

func TestInsertMoreRecords(t *testing.T) {
	users := []models.User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jack", Age: 19, Birthday: time.Now()},
	}

	result := db.Create(users)
	
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
