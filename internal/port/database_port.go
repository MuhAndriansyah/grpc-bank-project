package port

import (
	"time"

	"github.com/MuhAndriansyah/grpc-bank-project/internal/adapter/database"
	"github.com/google/uuid"
)

type BankDatabasePort interface {
	GetBankAccountByAccountNumber(acct string) (database.BankAccountOrm, error)
	GetExchangeRateAtTimestamp(fromCur string, toCur string, ts time.Time) (database.BankExchangeRateOrm, error)
	CreateTransaction(acct database.BankAccountOrm, t database.BankTransactionOrm) (uuid.UUID, error)
}
