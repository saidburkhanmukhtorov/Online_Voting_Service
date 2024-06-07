CREATE TABLE parties(
    id UUID PRIMARY KEY,
    name VARCHAR(256),
    slogan VARCHAR,
    opened_date DATE,
    description VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TYPE gender AS ENUM ('m', 'f');
CREATE TABLE public(
    id UUID PRIMARY KEY,
    name VARCHAR(32),
    last_name VARCHAR(32),
    phone VARCHAR(13) UNIQUE,
    email VARCHAR(64) UNIQUE,
    birthday DATE,
    gender gender,
    parties_id UUID REFERENCES parties(id) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);