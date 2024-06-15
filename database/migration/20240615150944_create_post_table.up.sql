CREATE TABLE IF NOT EXISTS "posts" (
  "id" varchar PRIMARY KEY,
  "title" varchar(60),
  "content" TEXT,
  "user_id" varchar,
  "created_at" timestamp default(now()),
  "updated_at" timestamp null
);

ALTER TABLE posts
ADD CONSTRAINT fk_post_user FOREIGN KEY (user_id) REFERENCES users(id);
