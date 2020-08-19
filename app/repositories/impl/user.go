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

func (u *UserRepo) Login(item *schema.Login) (*models.User, error) {
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

func (u *UserRepo) Register(item *schema.Register) (*models.User, error) {
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

func (u *UserRepo) GetUsers(queryParam *schema.UserQueryParam) (*[]models.User, error) {
	var query map[string]interface{}
	data, _ := json.Marshal(queryParam)
	json.Unmarshal(data, &queryParam)

	var user []models.User
	if dbs.Database.Where(query).Find(&user).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
