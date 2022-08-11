package register

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"unicode"

	_ "github.com/lib/pq"

	"github.com/abigiya/internal/Internal/adaptor/framework/dbservice/sqlc"
	"github.com/abigiya/internal/Internal/port"
)

type ResponseAcceptor struct {
	Store  sqlc.Account
	Status string
}

type ConnectandCheck interface {
	CheckUserNameFormatAndExist(username string) (sqlc.Account, bool, error)
	CheckProperInput(fn, ln, email, un string, pswd int64) error
	ConfirmAccount(ctx context.Context, firstname, lastname, emailaddress, username string, password int64) (*ResponseAcceptor, error)
}

type DbPortCatch struct {
	DbPort port.DbPort
	rawun  *sqlc.Account
}

var sq *sqlc.SQLAdaptor

func errorDefine(err string) error {
	return errors.New(err)
}

func NewUser(dbDriver, dbSource string) ConnectandCheck {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("can not connect to database", err)
	}
	return &DbPortCatch{
		DbPort: sq.NewAdapter(conn),
		rawun:  &sqlc.Account{},
	}
}

func (dbcach *DbPortCatch) CheckUserNameFormatAndExist(username string) (sqlc.Account, bool, error) {
	query := dbcach.DbPort.Query()
	account, err := query.GetAccountByUsername(context.Background(), username)
	if err != nil {
		return account, false, err
	}
	return account, true, err

}

func ValidUserName(str string) bool {
	strings.Contains("@", str)
	for _, r := range str {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (portCatch *DbPortCatch) CheckProperInput(fn, ln, email, un string, pswd int64) error {
	if fn == " " {
		return errorDefine("empty First Name")
	}
	if ln == " " {
		return errorDefine("empty Last Name")
	}
	if un == " " {
		return errorDefine("empty User Name")
	}
	if !ValidUserName(un) {
		return errorDefine("invalid user name format, follow @username")
	}
	if email == " " {
		return errorDefine("empty email")
	}
	if pswd <= 4 && pswd >= 8 {
		return errorDefine("password should be in range of 4 and 8")
	}
	return nil
}

func (portCatch *DbPortCatch) ConfirmAccount(ctx context.Context, firstname, lastname, emailaddress, username string, password int64) (*ResponseAcceptor, error) {
	validAccount := portCatch.DbPort.Query()
	err := portCatch.CheckProperInput(firstname, lastname, emailaddress, username, password)
	if err != nil {
		return &ResponseAcceptor{Status: err.Error()}, err
	}

	compare, err := validAccount.GetAccountByUsername(context.Background(), username)
	if err != nil && compare.Username == portCatch.rawun.Username {
		notExistUserName, _, _ := portCatch.CheckUserNameFormatAndExist(username)
		return &ResponseAcceptor{Store: notExistUserName, Status: "user name exist"}, errorDefine("user name exist")
	}

	acc := portCatch.DbPort.CreateParams(firstname, lastname, emailaddress, username, password)
	if err != nil {
		return &ResponseAcceptor{Status: "unexpected error"}, err
	}
	return &ResponseAcceptor{Store: acc, Status: "Registered!"}, err
}
