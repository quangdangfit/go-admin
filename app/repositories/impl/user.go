package impl

import (
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"go-admin/app/models"
	"go-admin/app/repositories"
	"go-admin/app/schema"
	"go-admin/dbs"
	"go-admin/pkg/utils"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository() repositories.IUserRepository {
	return &UserRepo{db: dbs.Database}
}

func (u *UserRepo) Login(item *schema.LoginBodyParam) (*models.User, error) {
	user := &models.User{}
	if dbs.Database.Where("username = ? ", item.Username).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(item.Password))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func (u *UserRepo) Register(item *schema.RegisterBodyParam) (*models.User, error) {
	var user models.User
	copier.Copy(&user, &item)
	hashedPassword := utils.HashAndSalt([]byte(item.Password))
	user.Password = hashedPassword

	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetUserByID(id string) (*models.User, error) {
	user := models.User{}
	if dbs.Database.Where("id = ? ", id).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *UserRepo) GetUserByToken(token string) (*models.User, error) {
	var user models.User
	if dbs.Database.Where("refresh_token = ? ", token).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *UserRepo) GetUsers(queryParam *schema.UserQueryParam) (*[]models.User, error) {
	var query map[string]interface{}
	data, _ := json.Marshal(queryParam)
	json.Unmarshal(data, &query)

	var user []models.User
	if dbs.Database.Where(query).Find(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *UserRepo) Update(userId string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error) {
	var body map[string]interface{}
	data, _ := json.Marshal(bodyParam)
	json.Unmarshal(data, &body)

	var change models.User
	if err := dbs.Database.Model(&change).Where("id = ?", userId).Update(body).Error; err != nil {
		return nil, err
	}

	return &change, nil
}
