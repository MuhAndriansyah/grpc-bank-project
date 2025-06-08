package application

import (
	"fmt"
	"log"
	"time"

	"github.com/MuhAndriansyah/grpc-bank-project/internal/adapter/database"
	"github.com/MuhAndriansyah/grpc-bank-project/internal/application/domain/bank"
	"github.com/MuhAndriansyah/grpc-bank-project/internal/port"
	"github.com/google/uuid"
)

type BankService struct {
	db port.BankDatabasePort
}

func NewBankService(dbPort port.BankDatabasePort) *BankService {
	return &BankService{
		db: dbPort,
	}
}

func (s *BankService) FindCurrentBalance(acct string) float64 {
	bankAccount, err := s.db.GetBankAccountByAccountNumber(acct)

	if err != nil {
		log.Println("Error on FindCurrentBalance :", err)
	}

	return bankAccount.CurrentBalance
}

func (s *BankService) FindExchangeRate(fromCur, toCur string, ts time.Time) float64 {
	exchangeRate, err := s.db.GetExchangeRateAtTimestamp(fromCur, toCur, ts)

	if err != nil {
		log.Println("Error on FindCurrentBalance :", err)
		return 0
	}

	return exchangeRate.Rate
}

func (s *BankService) CreateTransaction(acct string, t bank.Transaction) (uuid.UUID, error) {

	bankAccountOrm, err := s.db.GetBankAccountByAccountNumber(acct)

	if err != nil {
		return uuid.Nil, err
	}

	transacionOrm := database.BankTransactionOrm{
		TransactionUuid:      uuid.New(),
		AccountUuid:          bankAccountOrm.AccountUuid,
		TransactionTimestamp: time.Now(),
		Amount:               t.Amount,
		TransactionType:      t.TransactionType,
		Notes:                t.Notes,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	savedUuid, err := s.db.CreateTransaction(bankAccountOrm, transacionOrm)

	return savedUuid, err
}

func (s *BankService) CalculateTransactionSummary(tcur *bank.TransactionSummary, trans bank.Transaction) error {
	switch trans.TransactionType {
	case bank.TransactionTypeIn:
		tcur.SumIn += trans.Amount
	case bank.TransactionTypeOut:
		tcur.SumOut += trans.Amount
	default:
		return fmt.Errorf("unknown transaction type %v", trans.TransactionType)
	}

	tcur.SumTotal = tcur.SumIn - tcur.SumOut

	return nil
}
