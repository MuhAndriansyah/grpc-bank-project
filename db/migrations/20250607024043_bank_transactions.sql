-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bank_transactions (
    "transaction_uuid" UUID PRIMARY KEY,
    "account_uuid" UUID NOT NULL REFERENCES bank_accounts,
    "transaction_timestamp" TIMESTAMPTZ NOT NULL,
    "transaction_type" VARCHAR(25) NOT NULL,
    "amount" NUMERIC(15, 2) NOT NULL,
    "notes" TEXT,
    "created_at" TIMESTAMPTZ DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bank_transactions;
-- +goose StatementEnd
