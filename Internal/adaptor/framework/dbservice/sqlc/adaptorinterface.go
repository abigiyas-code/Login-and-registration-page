package sqlc

import (
	"database/sql"
	"log"
)

type DbPortAdaptor interface {
	Query() Queries
	CreateParams(firstname, lastname, emailaddress, username string, password int64) Account
	AccountStructure() Account
	GetaAccountParams(username string, password int) GetAccountParams
	Close()
}

type SQLAdaptor struct {
	db      *sql.DB
	Queries *Queries
}

func (adapter *SQLAdaptor) NewAdapter(db *sql.DB) DbPortAdaptor {
	return &SQLAdaptor{
		db:      db,
		Queries: &Queries{},
	}
}

func (adapter *SQLAdaptor) Query() Queries {
	return *adapter.Queries
}
func (adapter *SQLAdaptor) CreateParams(firstname, lastname, emailaddress, username string, password int64) Account {
	return Account{Firstname: firstname, Lastname: lastname, Emailaddress: emailaddress, Username: username, Password: int64(password)}
}
func (adapter *SQLAdaptor) AccountStructure() Account {
	return Account{}
}
func (adapter *SQLAdaptor) GetaAccountParams(username string, password int) GetAccountParams {
	return GetAccountParams{Username: username, Password: int64(password)}
}
func (adpter *SQLAdaptor) Close() {
	err := adpter.Queries.Close()
	if err != nil {
		log.Println("closing query error ", err)
	}
	err = adpter.db.Close()
	if err != nil {
		log.Println("closing database error ", err)
	}
}
