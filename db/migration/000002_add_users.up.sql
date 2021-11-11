CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "cuenta" ADD FOREIGN KEY ("propietario") REFERENCES "users" ("username");

-- CREATE UNIQUE INDEX ON "cuenta" ("propietario", "divisa"); 
ALTER TABLE "cuenta" ADD CONSTRAINT "propietario_divisa_key" UNIQUE ("propietario", "divisa");