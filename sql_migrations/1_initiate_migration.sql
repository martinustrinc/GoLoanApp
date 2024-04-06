-- +migrate Up
-- +migrate StatementBegin

-- SET SEARCH_PATH = 'loan_app';
CREATE TYPE locale_language AS ENUM ('en-US', 'id-ID');
CREATE TYPE record_status AS ENUM ('A', 'N', 'P');

CREATE SEQUENCE IF NOT EXISTS user_pkey_seq;
CREATE TABLE IF NOT EXISTS "user" (
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
    CONSTRAINT uq_user_userid UNIQUE (user_id)
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


-- +migrate StatementEnd