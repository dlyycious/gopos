package authmodule

import (
	"time"

	"github.com/dlyycious/gopos/database"
	"github.com/dlyycious/gopos/model"
)

func Logout(token string) (model.TokenLog, error) {
	db := database.DB

	query := model.TokenLog{Token: token}
	found := model.TokenLog{}
	if err := db.Model(query).Where("token = ?", query.Token).First(&found).Error; err != nil {
		return model.TokenLog{}, err
	}

	found.ExpiresAt = time.Now()
	db.Save(&found)

	return found, nil
}
