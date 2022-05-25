CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "first_name" varchar(256),
  "last_name" varchar(128),
  "full_name" varchar(512),
  "is_email_confirmed" boolean DEFAULT false,
  "email" varchar(128) NOT NULL,
  "user_name" varchar(512) NOT NULL,
  "role_id" int DEFAULT 1,
  "created_by" bigint,
  "updated_by" bigint,
  "created_timestamp" datetime DEFAULT (now()),
  "updated_timestamp" datetime DEFAULT (now()),
  "is_archived" boolean DEFAULT false
);

CREATE TABLE "login_details" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "password_hash" varchar(128) NOT NULL,
  "is_locked_out" boolean DEFAULT false,
  "failed_login_count" int DEFAULT 0,
  "created_by" bigint,
  "updated_by" bigint,
  "created_timestamp" datetime DEFAULT (now()),
  "updated_timestamp" datetime DEFAULT (now())
);

CREATE TABLE "events" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "name" varchar(128) NOT NULL,
  "ip_address" varchar(128),
  "user_agent" varchar(128),
  "created_by" bigint,
  "updated_by" bigint,
  "created_timestamp" datetime DEFAULT (now()),
  "updated_timestamp" datetime DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("id") REFERENCES "login_details" ("user_id");

ALTER TABLE "events" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
