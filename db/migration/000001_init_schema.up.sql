CREATE TABLE "cuenta" (
  "id" bigserial PRIMARY KEY,
  "propietario" varchar NOT NULL,
  "tope" bigint NOT NULL,
  "divisa" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "ingreso" (
  "id" bigserial PRIMARY KEY,
  "cuenta_id" bigint NOT NULL,
  "monto" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "transferencia" (
  "id" bigserial PRIMARY KEY,
  "from_cuenta_id" bigint NOT NULL,
  "to_cuenta_id" bigint NOT NULL,
  "monto" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "ingreso" ADD FOREIGN KEY ("cuenta_id") REFERENCES "cuenta" ("id");

ALTER TABLE "transferencia" ADD FOREIGN KEY ("from_cuenta_id") REFERENCES "cuenta" ("id");

ALTER TABLE "transferencia" ADD FOREIGN KEY ("to_cuenta_id") REFERENCES "cuenta" ("id");

CREATE INDEX ON "cuenta" ("propietario");

CREATE INDEX ON "ingreso" ("cuenta_id");

CREATE INDEX ON "transferencia" ("from_cuenta_id");

CREATE INDEX ON "transferencia" ("to_cuenta_id");

CREATE INDEX ON "transferencia" ("from_cuenta_id", "to_cuenta_id");

COMMENT ON COLUMN "ingreso"."monto" IS 'Puede ser negativo o positivo';

COMMENT ON COLUMN "transferencia"."monto" IS 'Debe ser positivo';
