package sql

import (
	"fmt"
	"testing"
)

func TestUser_Create(t *testing.T) {
	u := User{
		Name:    "maodou",
		Address: "every where",
		Age:     18,
	}

	err := u.Create()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUser_CreateReturnId(t *testing.T) {
	u := User{
		Name:    "maodou2",
		Address: "every where",
		Age:     18,
	}

	err := u.CreateReturnId()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(u)
}

func TestUser_QueryById(t *testing.T) {
	u := User{
		ID: 1,
	}

	err := u.QueryById()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(u)
}

func TestUser_UpdateById(t *testing.T) {
	u := User{
		ID:      1,
		Name:    "update",
		Address: "e",
		Age:     14,
	}

	err := u.UpdateById()
	if err != nil {
		t.Fatal(err)
	}

	updated := User{
		ID: 1,
	}
	err = updated.QueryById()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(updated)
}

func TestUser_DeleteById(t *testing.T) {
	u := User{
		ID: 4,
	}

	err := u.DeleteById()
	if err != nil {
		t.Fatal(err)
	}
}
