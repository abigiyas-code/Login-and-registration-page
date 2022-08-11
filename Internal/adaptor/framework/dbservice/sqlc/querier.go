// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	GetAccount(ctx context.Context, arg GetAccountParams) (Account, error)
	GetAccountByUsername(ctx context.Context, username string) (Account, error)
}

var _ Querier = (*Queries)(nil)
