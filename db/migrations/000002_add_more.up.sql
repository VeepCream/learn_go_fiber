CREATE TABLE more (
  "id" SERIAL,
  "name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now()
)