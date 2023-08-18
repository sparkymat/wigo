CREATE EXTENSION moddatetime;
CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name text NOT NULL,
  username text NOT NULL UNIQUE,
  encrypted_password text NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT current_timestamp,
  updated_at timestamp without time zone NOT NULL DEFAULT current_timestamp
);
CREATE TRIGGER users_updated_at
  BEFORE UPDATE
  ON users
  FOR EACH ROW
    EXECUTE FUNCTION moddatetime(updated_at);
