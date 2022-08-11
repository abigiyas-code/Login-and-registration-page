CREATE TABLE "accounts" (
  "id" SERIAL PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "emailaddress" varchar NOT NULL,
  "username" varchar NOT NULL,
  "password" int NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now())
);
