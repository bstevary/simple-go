CREATE TABLE
    "users" (
        user_id BIGSERIAL PRIMARY KEY,
        email VARCHAR(80) UNIQUE NOT NULL,
        hashed_password VARCHAR(255) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ
    );