-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE admin (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    login_id VARCHAR(128),
    password VARCHAR(128)
);

CREATE TABLE guild (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    status SMALLINT NOT NULL,   -- 1:仮登録、2:本登録

    create_user_id BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_user_id BIGINT,
    updated_at TIMESTAMP WITH TIME ZONE,
    delete_user_id BIGINT,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE guest_token (
    id BIGSERIAL PRIMARY KEY,
    guild_id BIGINT NOT NULL REFERENCES guild (id),
    mail VARCHAR(256) NOT NULL,
    token VARCHAR(256) NOT NULL,
    expiration_date TIMESTAMP WITH TIME ZONE NOT NULL,
    accepted_number VARCHAR(64) NOT NULL
);

CREATE TABLE owner (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(256),
    mail VARCHAR(256) NOT NULL,
    login_id VARCHAR(128),
    password VARCHAR(128),

    create_user_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_user_id BIGINT,
    updated_at TIMESTAMP WITH TIME ZONE,
    delete_user_id BIGINT,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE guild_owner_relation (
    id BIGSERIAL PRIMARY KEY,
    guild_id BIGINT NOT NULL REFERENCES guild (id),
    owner_id BIGINT NOT NULL REFERENCES owner (id)
);

CREATE TABLE participant (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,

    create_user_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_user_id BIGINT,
    updated_at TIMESTAMP WITH TIME ZONE,
    delete_user_id BIGINT,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE task (
    id BIGSERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    status SMALLINT NOT NULL,
    continuity SMALLINT NOT NULL,
    due_datetime TIMESTAMP,

    create_user_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_user_id BIGINT,
    updated_at TIMESTAMP WITH TIME ZONE,
    delete_user_id BIGINT,
    deleted_at TIMESTAMP WITH TIME ZONE
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE task;
DROP TABLE participant;
DROP TABLE guild_owner_relation;
DROP TABLE owner;
DROP TABLE guest_token;
DROP TABLE guild;
DROP TABLE admin;
-- +migrate StatementEnd
