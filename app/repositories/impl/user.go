package impl

import (
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/quangdangfit/go-admin/app/dbs"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/pkg/utils"
)

type UserRepo struct {
	db dbs.IDatabase
}

func NewUserRepository(db dbs.IDatabase) repositories.IUserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Login(item *schema.LoginBodyParam) (*models.User, error) {
	user := &models.User{}
	if r.db.GetInstance().Where("username = ? ", item.Username).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(item.Password))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func (r *UserRepo) Register(item *schema.RegisterBodyParam) (*models.User, error) {
	var user models.User
	copier.Copy(&user, &item)
	hashedPassword := utils.HashAndSalt([]byte(item.Password))
	user.Password = hashedPassword

	if err := r.db.GetInstance().Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByID(id string) (*models.User, error) {
	user := models.User{}
	if dbs.Database.Where("id = ? ", id).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *UserRepo) GetUserByToken(token string) (*models.User, error) {
	var user models.User
	if dbs.Database.Where("refresh_token = ? ", token).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *UserRepo) GetUsers(queryParam *schema.UserQueryParam) (*[]models.User, error) {
	var query map[string]interface{}
	data, _ := json.Marshal(queryParam)
	json.Unmarshal(data, &query)

	var user []models.User
	if dbs.Database.Where(query).Find(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *UserRepo) Update(userId string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error) {
	var body map[string]interface{}
	data, _ := json.Marshal(bodyParam)
	json.Unmarshal(data, &body)

	var change models.User
	if err := dbs.Database.Model(&change).Where("id = ?", userId).Update(body).Error; err != nil {
		return nil, err
	}

	return &change, nil
}

func (r *UserRepo) RemoveToken(userId string) (*models.User, error) {
	var body = map[string]interface{}{"refresh_token": ""}
	var change models.User
	if err := dbs.Database.Model(&change).Where("id = ?", userId).Update(body).Error; err != nil {
		return nil, err
	}

	return &change, nil
}
