-- +migrate Up
-- +migrate StatementBegin

CREATE TYPE locale_language AS ENUM ('en-US', 'id-ID');
CREATE TYPE record_status AS ENUM ('A', 'N', 'P');

CREATE SEQUENCE IF NOT EXISTS user_pkey_seq;
CREATE TABLE "user" (
    id              BIGINT NOT NULL             DEFAULT nextval('user_pkey_seq'::regclass),
    user_id         BIGINT,
    locale          locale_language,
    first_name      VARCHAR(50),
    last_name       VARCHAR(50),
    username        VARCHAR(20),
    password        VARCHAR(20),
    email           VARCHAR(50),
    phone           VARCHAR(100),
    status          record_status               DEFAULT 'A',
    created_by      BIGINT,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by      BIGINT,
    updated_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN                     DEFAULT FALSE,
    CONSTRAINT pk_user_id PRIMARY KEY (id),
    CONSTRAINT uq_user_userid UNIQUE (client_id)
);

CREATE SEQUENCE IF NOT EXISTS role_pkey_seq;
CREATE TABLE "role"
(
    id             BIGINT NOT NULL             DEFAULT nextval('role_pkey_seq'::regclass),
    role_name      VARCHAR(256),
    description    TEXT,
    permission     TEXT,
    level          SMALLINT,
    created_by     BIGINT,
    created_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by     BIGINT,
    updated_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted        BOOLEAN                     DEFAULT FALSE,
    CONSTRAINT pk_role_id PRIMARY KEY (id),
    CONSTRAINT uq_role_role_name UNIQUE (role_name)
);

CREATE SEQUENCE IF NOT EXISTS account_loan_pkey_seq;
CREATE TABLE "account_loan" (
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