package repository

import (
	"context"
	"testing"

	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/models"
)

var user = models.User{
	Id:       "5f5ea4b3-9ddb-4a20-7ad8-b62aa536a753",
	Username: "my_user",
	Password: "$2a$14$rc93aalD6OecMtluR8MM2uwI",
}

var userWrong = models.User{
	Id:       "5f5ea5b3-9ddb-4a20-7ad8-b63aa536a753",
	Username: "my_user!",
	Password: "$2a$14$rc93aalD6OecMtluR8MM2uwI",
}

var dbc = DBConnection()

var userRep = GetUsersRepository(DBConnection())

func delUser(dbConnect *bun.DB, uid string, t *testing.T) {
	ctx := context.Background()
	user := new(models.User)

	_, err := dbConnect.NewDelete().Model(user).Where("id = ?", uid).Exec(ctx)
	if err != nil {
		t.Fatalf("Error while deleting test user: %v", err)
	}
}

func TestSaveUser(t *testing.T) {
	delUser(dbc, user.Id, t)

	err := userRep.SaveUser(user)
	if err != nil {
		t.Fatalf("Saving user failed")
	}
}

func TestGetUserByFieldCorrect(t *testing.T) {
	tests := []struct {
		name  string
		field string
		value string
	}{
		{
			name:  "Get user by UserName field",
			field: "username",
			value: user.Username,
		},
		{
			name:  "Get user by Id field",
			field: "id",
			value: user.Id,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, e := userRep.GetUserByField(tt.field, tt.value)
			if e == false {
				t.Fatalf("TestGetUserByField failed, field: %v, value: %v ", tt.field, tt.value)
			}
		})
	}
}

func TestGetUserByFieldWrong(t *testing.T) {
	tests := []struct {
		name  string
		field string
		value string
	}{
		{
			name:  "Get user by UserName field",
			field: "username",
			value: userWrong.Username,
		},
		{
			name:  "Get user by Id field",
			field: "id",
			value: userWrong.Id,
		},
		{
			name:  "Get user by wrong field",
			field: "qwerty",
			value: "qwerty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, e := userRep.GetUserByField(tt.field, tt.value)
			if e != false {
				t.Fatalf("TestGetUserByFieldInCorrect failed, field: %v, value: %v ", tt.field, tt.value)
			}
		})
	}
}
