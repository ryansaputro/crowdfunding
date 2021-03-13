package transaction

import (
	"crowdfunding/user"
)

type GetTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
