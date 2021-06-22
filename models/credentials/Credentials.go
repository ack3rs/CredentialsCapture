package credentials

import (
	"errors"
	"fmt"

	"github.com/acky666/CredentialsCapture/db"
	l "github.com/acky666/ackyLog"
	"github.com/go-playground/validator"
)

// swagger:model User
type User struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Country   string `json:"country" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

var validate *validator.Validate

func (u *User) Save() error {
	defer l.TIMED("Saved Record")()

	validate = validator.New()
	err := validate.Struct(u)
	if err != nil {

		ReturnedErr := ""
		for _, err := range err.(validator.ValidationErrors) {

			ReturnedErr = fmt.Sprintf("Validation error '%s' failed on '%s'", err.Field(), err.Tag())
			l.ERROR("%s", ReturnedErr)
		}

		return errors.New(ReturnedErr)
	}

	l.DEBUG("Saving: Firstname: [F-CYAN]%s[F-NORMAL] Lastname: [F-CYAN]%s[F-NORMAL] Country: [F-CYAN]%s[F-NORMAL] Email: [F-CYAN]%s[F-NORMAL]", u.Firstname, u.Lastname, u.Country, u.Email)

	sql := "INSERT INTO credentials.users (Firstname,Surname,Country,Email) VALUES (?,?,?,?)"
	id, rowEffected, err := db.ExecutePrepared(sql, u.Firstname, u.Lastname, u.Country, u.Email)

	l.DEBUG("DB Result - NewID:%d Rows:%v Err:%v", id, rowEffected, err)
	if err != nil {
		return err
	}

	if rowEffected == 0 {
		return errors.New("Failed to write to database")
	}

	return nil
}
