DROP TRIGGER IF EXISTS users_updated_at_trigger ON users;
DROP FUNCTION IF EXISTS update_timestamp();

ALTER TABLE users
DROP COLUMN IF EXISTS password_hash,
DROP COLUMN IF EXISTS refresh_token,
DROP COLUMN IF EXISTS updated_at;