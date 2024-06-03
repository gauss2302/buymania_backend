package db

import (
	"github.com/gauss2302/buymania_backend/internal/models"
	_ "github.com/gauss2302/buymania_backend/internal/models"
	"github.com/jinzhu/gorm"
)

func (u *models.User) FindAllUsers(db *gorm.DB) (*[]models.User, error) {

}
