package model

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/maodou24/release-hub/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicId string `gorm:"column:public_id"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Login    string `gorm:"column:login;uniqueIndex"`
	TokenCipher string `gorm:"column:token_cipher" json:"-"`
	GitProviderType int `gorm:"column:git_provider_type;not null;default:1" json:"gitProviderType"`
}

func UpsertUser(ctx context.Context, user *User) error {
	db := database.GetDb().WithContext(ctx)
	if user.PublicId == "" {
		id := make([]byte, 16)
		if _, err := rand.Read(id); err != nil {
			return err
		}
		user.PublicId = hex.EncodeToString(id)
	}
	var existing User
	if err := db.Where("login = ?", user.Login).First(&existing).Error; err == nil {
		user.ID = existing.ID
		user.PublicId = existing.PublicId
		return db.Model(&existing).Updates(map[string]any{
			"name": user.Name, "email": user.Email, "token_cipher": user.TokenCipher,
			"git_provider_type": user.GitProviderType,
		}).Error
	} else if err != gorm.ErrRecordNotFound {
		return err
	}
	return db.Create(user).Error
}

func GetUserByPublicID(ctx context.Context, publicId string) (*User, error) {
	var u User
	err := database.GetDb().WithContext(ctx).Where("public_id = ?", publicId).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
