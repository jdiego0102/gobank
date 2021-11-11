ALTER TABLE IF EXISTS  "cuenta" DROP CONSTRAINT IF EXISTS "propietario_divisa_key";
 
ALTER TABLE IF EXISTS  "cuenta" DROP CONSTRAINT IF EXISTS "cuenta_propietario_fkey";

DROP TABLE IF EXISTS "users";