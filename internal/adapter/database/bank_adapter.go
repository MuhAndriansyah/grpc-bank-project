package database

import (
	"log"
	"time"

	"github.com/MuhAndriansyah/grpc-bank-project/internal/application/domain/bank"
	"github.com/google/uuid"
)

func (a *DatabaseAdapter) GetBankAccountByAccountNumber(acct string) (BankAccountOrm, error) {
	var bankAccountOrm BankAccountOrm

	if err := a.db.First(&bankAccountOrm, "account_number = ?", acct).Error; err != nil {
		log.Printf("can't find bank account %v : %v\n", acct, err)
		return bankAccountOrm, err
	}

	return bankAccountOrm, nil
}

func (a *DatabaseAdapter) GetExchangeRateAtTimestamp(fromCur string, toCur string, ts time.Time) (BankExchangeRateOrm, error) {
	var bankExchangeRateOrm BankExchangeRateOrm

	err := a.db.Where("from_currency = ? AND to_currency = ? AND valid_from_timestamp <= ? AND valid_to_timestamp >= ?", fromCur, toCur, ts, ts).First(&bankExchangeRateOrm).Error

	if err != nil {
		log.Printf("can't find exchange rate from_currency %s and to_currency %s: %v", fromCur, toCur, err)
		return bankExchangeRateOrm, err
	}

	return bankExchangeRateOrm, nil
}

func (a *DatabaseAdapter) CreateTransaction(acct BankAccountOrm, t BankTransactionOrm) (uuid.UUID, error) {
	tx := a.db.Begin()

	if err := tx.Create(t).Error; err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	newAmount := t.Amount

	if t.TransactionType == bank.TransactionTypeOut {
		newAmount = -1 * t.Amount
	}

	newAccountBalance := acct.CurrentBalance + newAmount

	if err := tx.Model(&acct).Updates(
		map[string]interface{}{
			"current_balance": newAccountBalance,
			"updated_at":      time.Now(),
		},
	).Error; err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	tx.Commit()

	return t.TransactionUuid, nil
}
