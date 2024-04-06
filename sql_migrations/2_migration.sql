-- +migrate Up
-- +migrate StatementBegin


CREATE SEQUENCE IF NOT EXISTS account_loan_pkey_seq;
CREATE TABLE IF NOT EXISTS "account_loan" (
    id                      BIGINT NOT NULL             DEFAULT nextval('account_loan_pkey_seq'::regclass),
    month                   BIGINT,
    total_installment       VARCHAR(256),
    principal_installment   VARCHAR(256),
    interest_installment    VARCHAR(256),
    remaining_installment   VARCHAR(256),
    start_date              TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by              BIGINT,
    created_at              TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by              BIGINT,
    updated_at              TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted                 BOOLEAN                     DEFAULT FALSE,
    CONSTRAINT pk_account_loan_id PRIMARY KEY (id)
);


-- +migrate StatementEnd