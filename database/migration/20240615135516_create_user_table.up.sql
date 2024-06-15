CREATE TABLE IF NOT EXISTS "users" (
  "id" integer PRIMARY KEY,
  "username" varchar(16) UNIQUE,
  "email" varchar(100),
  "password" varchar,
  "created_at" timestamp default(now()),
  "updated_at" timestamp null
);
