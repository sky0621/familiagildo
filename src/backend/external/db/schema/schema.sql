CREATE TABLE guild (
    id BIGSERIAL PRIMARY KEY,
    node_id UUID NOT NULL UNIQUE,
    name TEXT NOT NULL,

    create_user_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE,
    update_user_id BIGINT,
    updated_at TIMESTAMP WITH TIME ZONE,
    delete_user_id BIGINT,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX ON guild (node_id);

CREATE TABLE admin (
    id BIGSERIAL PRIMARY KEY,
    node_id UUID NOT NULL UNIQUE,
    name TEXT NOT NULL,
    login_id VARCHAR(128),
    password VARCHAR(128)
);
CREATE INDEX ON admin (node_id);

CREATE TABLE owner (
    id BIGSERIAL PRIMARY KEY,
    node_id UUID NOT NULL UNIQUE,
    name TEXT NOT NULL,

    create_user_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE,
    update_user_id BIGINT,
    updated_at TIMESTAMP WITH TIME ZONE,
    delete_user_id BIGINT,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX ON owner (node_id);

CREATE TABLE participant (
    id BIGSERIAL PRIMARY KEY,
    node_id UUID NOT NULL UNIQUE,
    name TEXT NOT NULL,

    create_user_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE,
    update_user_id BIGINT,
    updated_at TIMESTAMP WITH TIME ZONE,
    delete_user_id BIGINT,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX ON participant (node_id);

CREATE TABLE task (
    id BIGSERIAL PRIMARY KEY,
    node_id UUID NOT NULL UNIQUE,
    content TEXT NOT NULL,
    status SMALLINT NOT NULL,
    continuity SMALLINT NOT NULL,
    due_datetime TIMESTAMP,

    create_user_id BIGINT,
    created_at TIMESTAMP WITH TIME ZONE,
    update_user_id BIGINT,
    updated_at TIMESTAMP WITH TIME ZONE,
    delete_user_id BIGINT,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX ON task (node_id);
