CREATE TABLE IF NOT EXISTS "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar(16) UNIQUE,
  "email" varchar(100),
  "password" varchar,
  "created_at" timestamp default(now()),
  "updated_at" timestamp null
);
