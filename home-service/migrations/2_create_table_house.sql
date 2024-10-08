-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE houses (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    year INT CHECK (year > 0),
    developer VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE flats (
    id SERIAL PRIMARY KEY,
    house_id INT REFERENCES houses(id) ON DELETE CASCADE,
    flat_num INT CHECK (flat_num > 0),
    price INT CHECK (price >= 0),
    rooms INT CHECK (rooms > 0),
    status VARCHAR(20) CHECK (status IN ('created', 'approved', 'declined', 'on moderation')),
    moderator_id UUID REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    house_id INT REFERENCES houses(id) ON DELETE CASCADE,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE moderations (
    id SERIAL PRIMARY KEY,
    flat_id INT REFERENCES flats(id) ON DELETE CASCADE,
    moderator_id UUID REFERENCES users(id)
);

CREATE INDEX idx_moderations_flat_id ON moderations (flat_id);
CREATE INDEX idx_moderations_moderator_id ON moderations (moderator_id);
CREATE INDEX idx_flats_house_id ON flats (house_id);
CREATE INDEX idx_flats_status ON flats (status);
CREATE INDEX idx_flats_moderator_id ON flats (moderator_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS houses;
DROP TABLE IF EXISTS flats;
DROP TABLE IF EXISTS subscriptions;

DROP INDEX IF EXISTS idx_flats_moderator_id;
DROP INDEX IF EXISTS idx_flats_house_id;
DROP INDEX IF EXISTS idx_flats_status;
-- +goose StatementEnd
