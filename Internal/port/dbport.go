package port

import "github.com/abigiya/internal/Internal/adaptor/framework/dbservice/sqlc"

type DbPort interface {
	Close()
	Query() sqlc.Queries
	CreateParams(firstname, lastname, emailaddress, username string, password int64) sqlc.Account
	GetaAccountParams(username string, password int) sqlc.GetAccountParams
	AccountStructure() sqlc.Account
}
