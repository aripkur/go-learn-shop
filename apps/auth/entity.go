package auth

import (
	"strings"
	"time"

	"github.com/aripkur/go-learn-shop/infra/response"
	"github.com/aripkur/go-learn-shop/utility"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	ROLE_Admin Role = "admin"
	ROLE_User  Role = "user"
)

type AuthEntity struct {
	Id        int       `db:"id"`
	PublicId  uuid.UUID `db:"public_id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      Role      `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func newFromRegisterRequest(req RegisterRequestPayload) AuthEntity {
	return AuthEntity{
		PublicId:  uuid.New(),
		Email:     req.Email,
		Password:  req.Password,
		Role:      ROLE_User,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func newFromLoginRequest(req LoginRequestPayload) AuthEntity {
	return AuthEntity{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (a AuthEntity) validate() (err error) {

	if err = a.validateEmail(); err != nil {
		return
	}

	if err = a.validatePassword(); err != nil {
		return
	}

	return
}

func (a AuthEntity) validateEmail() (err error) {
	if a.Email == "" {
		return response.ErrEmailRequired
	}

	emails := strings.Split(a.Email, "@")
	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}

	return
}

func (a AuthEntity) validatePassword() (err error) {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(a.Password) < 6 {
		return response.ErrPasswordInvalid
	}

	return
}

func (a AuthEntity) isExists() bool {
	return a.Id != 0
}

func (a *AuthEntity) encryptPassword(salt int) (err error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), salt)
	if err != nil {
		return
	}

	a.Password = string(encryptedPass)
	return
}

func (a AuthEntity) VerifyPasswordFromEncrypted(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
}

func (a AuthEntity) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(a.Password))
}

func (a AuthEntity) GenerateToken(secret string) (tokenString string, err error) {
	return utility.GenerateToken(a.PublicId.String(), string(a.Role), secret)
}
