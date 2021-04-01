package repositories

import (
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

// UserRepo user repository struct
type UserRepo struct {
	db interfaces.IDatabase
}

// NewUserRepository return new IUserRepository interface
func NewUserRepository(db interfaces.IDatabase) interfaces.IUserRepository {
	return &UserRepo{db: db}
}

// Login handle user login
func (r *UserRepo) Login(item *schema.LoginBodyParams) (*models.User, error) {
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

// Register new user
func (r *UserRepo) Register(item *schema.RegisterBodyParams) (*models.User, error) {
	var user models.User
	copier.Copy(&user, &item)
	if err := r.db.GetInstance().Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID get user by id
func (r *UserRepo) GetUserByID(id string) (*models.User, error) {
	user := models.User{}
	if r.db.GetInstance().Where("id = ? ", id).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// GetUserByToken get user by refresh token
func (r *UserRepo) GetUserByToken(token string) (*models.User, error) {
	var user models.User
	if r.db.GetInstance().Where("refresh_token = ? ", token).First(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// GetUsers get users by by query
func (r *UserRepo) GetUsers(queryParam *schema.UserQueryParam) (*[]models.User, error) {
	var query map[string]interface{}
	data, _ := json.Marshal(queryParam)
	json.Unmarshal(data, &query)

	var user []models.User
	if r.db.GetInstance().Where(query).Find(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// Update user
func (r *UserRepo) Update(userID string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error) {
	var body map[string]interface{}
	data, _ := json.Marshal(bodyParam)
	json.Unmarshal(data, &body)

	var change models.User
	if err := r.db.GetInstance().Model(&change).Where("id = ?", userID).Update(body).Error; err != nil {
		return nil, err
	}

	return &change, nil
}

// RemoveToken remove refresh token
func (r *UserRepo) RemoveToken(userID string) (*models.User, error) {
	var body = map[string]interface{}{"refresh_token": ""}
	var change models.User
	if err := r.db.GetInstance().Model(&change).Where("id = ?", userID).Update(body).Error; err != nil {
		return nil, err
	}

	return &change, nil
}

// Create new user
func (r *UserRepo) Create(user *models.User) error {
	if err := r.db.GetInstance().Create(&user).Error; err != nil {
		return errors.Wrap(err, "UserRepo.Create")
	}
	return nil
}

// GetByID get user by id
func (r *UserRepo) GetByID(id string) (*models.User, error) {
	user := models.User{}
	if err := r.db.GetInstance().Where("id = ? ", id).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "UserRepo.GetUserID")
	}

	return &user, nil
}

// List get user by UserQueryParam
func (r *UserRepo) List(param *schema.UserQueryParam) (*[]models.User, error) {
	var query map[string]interface{}
	data, _ := json.Marshal(param)
	json.Unmarshal(data, &query)

	var user []models.User
	if err := r.db.GetInstance().Where(query).Offset(param.Offset).Limit(param.Limit).Find(&user).Error; err != nil {
		return nil, errors.Wrap(err, "UserRepo.List")
	}

	return &user, nil
}
