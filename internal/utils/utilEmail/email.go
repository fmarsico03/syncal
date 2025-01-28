package utilEmail

import (
	"errors"
	"gorm.io/gorm"
	"syncal/database"
	modelsUser "syncal/internal/users/models"
)

func FindUserByEmail(email string) (modelsUser.User, error) {
	var user modelsUser.User
	err := database.Database.Where("mail = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errors.New("user not found")
	}
	return user, err
}
