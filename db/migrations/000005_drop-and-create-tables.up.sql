-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS citext;

-- Drop in dependency order
DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS users;

-- =========================
-- users
-- =========================
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       email CITEXT NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,

                       role VARCHAR(20) NOT NULL DEFAULT 'user',
                       status VARCHAR(20) NOT NULL DEFAULT 'pending',

                       email_verified_at TIMESTAMPTZ NULL,
                       last_login_at TIMESTAMPTZ NULL,

                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMPTZ NULL,

                       CONSTRAINT users_role_check
                           CHECK (role IN ('user', 'admin')),

                       CONSTRAINT users_status_check
                           CHECK (status IN ('pending', 'active', 'suspended', 'deleted'))
);

-- =========================
-- user_profiles
-- =========================
CREATE TABLE user_profiles (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,

                               first_name VARCHAR(100) NOT NULL,
                               last_name VARCHAR(100) NOT NULL,
                               display_name VARCHAR(150) NOT NULL,

                               phone VARCHAR(30) NULL,
                               avatar_url TEXT NULL,
                               bio TEXT NULL,
                               date_of_birth DATE NULL,
                               gender VARCHAR(30) NULL,

                               country VARCHAR(100) NULL,
                               city VARCHAR(100) NULL,
                               address TEXT NULL,

                               created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                               updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                               CONSTRAINT user_profiles_display_name_check
                                   CHECK (char_length(display_name) > 0)
);

-- =========================
-- indexes
-- =========================
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
CREATE INDEX idx_users_created_at ON users(created_at);

CREATE INDEX idx_user_profiles_user_id ON user_profiles(user_id);
CREATE INDEX idx_user_profiles_display_name ON user_profiles(display_name);
CREATE INDEX idx_user_profiles_phone ON user_profiles(phone);