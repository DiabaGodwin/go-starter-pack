ALTER TABLE users
ADD COLUMN password TEXT NOT NULL,
ADD COLUMN refresh_token TEXT,
ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();


CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER users_updated_at_trigger
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();