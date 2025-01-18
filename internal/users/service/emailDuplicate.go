package service

import (
	"fmt"
	"syncal/database"
	"syncal/internal/users/models"
)

func EmailDuplicate(email string) bool {
	var user models.User
	result := database.Database.Where("mail = ?", email).First(&user)
	fmt.Println(result.RowsAffected)
	return result.RowsAffected > 0
}
