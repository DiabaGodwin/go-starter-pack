CREATE TABLE user_profiles(
      id UUID NOT NULL PRIMARY KEY,
      firstname varchar not null ,
      lastname varchar not null ,
      avatar_url varchar not null ,
      created_at TIMESTAMPTZ NOT NULL,
      updated_at TIMESTAMPTZ NOT NULL
)