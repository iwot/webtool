package webtool

import (
	"database/sql"
	"net/http"
)

type TransactionInterface interface {
	Writer() http.ResponseWriter
	Request() *http.Request
	DefaultParams() map[string]string
}

type TransactionWithDbInterface interface {
	TransactionInterface
	SetDb(*sql.DB)
	Db() *sql.DB
}

type Transaction struct {
	w        http.ResponseWriter
	r        *http.Request
	defaults map[string]string
}

// データベースを使用する場合のトランザクション
type TransactionWithDb struct {
	*Transaction
}
