CREATE TABLE feeds (
  id uuid PRIMARY KEY,
  user_id uuid NOT NULL REFERENCES users(id),
  public_id text NOT NULL UNIQUE,
  title text NOT NULL DEFAULT '',
  description text NOT NULL DEFAULT '',
  url text NOT NULL,
  last_fetched_at timestamp without time zone NOT NULL DEFAULT current_timestamp,
  created_at timestamp without time zone NOT NULL DEFAULT current_timestamp,
  updated_at timestamp without time zone NOT NULL DEFAULT current_timestamp
);
