package port

import (
	"time"

	"github.com/MuhAndriansyah/grpc-bank-project/internal/application/domain/bank"
	"github.com/google/uuid"
)

type BankServicePort interface {
	FindCurrentBalance(acct string) float64
	FindExchangeRate(fromCur, toCur string, ts time.Time) float64
	CreateTransaction(acct string, t bank.Transaction) (uuid.UUID, error)
	CalculateTransactionSummary(tcur *bank.TransactionSummary, trans bank.Transaction) error
}
