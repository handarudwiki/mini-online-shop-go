package user

import (
	"handarudwiki/mini-online-shop-go/infra/response"
	"handarudwiki/mini-online-shop-go/utility"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	ROLE_ADMIN Role = "admin"
	ROLE_USER  Role = "user"
)

type UserEntity struct {
	ID        int       `db:"id"`
	PublicID  uuid.UUID `db:"public_id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      Role      `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewFormRegisterRequest(req RegisterRequestPayload) UserEntity {
	return UserEntity{
		Email:     req.Email,
		Password:  req.Password,
		PublicID:  uuid.New(),
		Role:      ROLE_USER,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewFormLoginRequest(req LoginRequestPayload) UserEntity {
	return UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (u UserEntity) Validate() (err error) {
	err = u.ValidateEmail()
	if err != nil {
		return
	}
	err = u.ValidatePassword()
	if err != nil {
		return
	}

	return
}

func (u UserEntity) ValidateEmail() (err error) {
	if u.Email == "" {
		return response.ErrEmailRequired
	}
	emails := strings.Split(u.Email, "@")
	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}
	return
}

func (u UserEntity) ValidatePassword() (err error) {
	if u.Password == "" {
		return response.ErrPasswordRequired
	}
	if len(u.Password) <= 6 {
		return response.ErrPasswordInvalidLength
	}
	return
}

func (u UserEntity) IsExists() bool {
	return u.ID != 0
}

func (u UserEntity) VerifyPassword(plain string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	if err != nil {
		return
	}
	return
}

func (u UserEntity) GenerateToken(secret string) (tokenString string, err error) {
	tokenString, err = utility.GenerateToken(u.PublicID.String(), string(u.Role), secret)
	if err != nil {
		return
	}
	return
}

func (u *UserEntity) HashPassword(salt int) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)
	if err != nil {
		return
	}
	u.Password = string(hashedPassword)
	return
}
