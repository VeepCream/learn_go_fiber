CREATE TABLE users (
  "id" SERIAL,
  "username" VARCHAR(255) NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
)