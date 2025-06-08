-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bank_exchange_rates (
    "exchange_rate_uuid" UUID PRIMARY KEY,
    "from_currency" VARCHAR(5) NOT NULL,
    "to_currency" VARCHAR(5) NOT NULL,
    "rate" NUMERIC(20, 10) NOT NULL,
    "valid_from_timestamp" TIMESTAMPTZ NOT NULL,
    "valid_to_timestamp" TIMESTAMPTZ NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bank_exchange_rates;
-- +goose StatementEnd
