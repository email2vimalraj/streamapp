CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "streams" (
  "id" bigserial PRIMARY KEY,
  "stream_name" varchar UNIQUE NOT NULL,
  "stream_link" varchar NOT NULL,
  "username" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "invitees" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar,
  "mobile_number" varchar,
  "inviter" varchar NOT NULL,
  "stream_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "streams"
ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
ALTER TABLE "invitees"
ADD FOREIGN KEY ("inviter") REFERENCES "users" ("username");
ALTER TABLE "invitees"
ADD FOREIGN KEY ("stream_id") REFERENCES "streams" ("id");
CREATE INDEX ON "streams" ("username");
CREATE INDEX ON "streams" ("username", "stream_name");
CREATE INDEX ON "invitees" ("inviter");