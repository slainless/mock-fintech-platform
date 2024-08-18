BEGIN;
ALTER TABLE "payment_accounts" ADD COLUMN "name" varchar(255);
ALTER TABLE "payment_accounts" ADD COLUMN "foreign_id" varchar(255);
ALTER TABLE "payment_accounts" DROP COLUMN "balance";
ALTER TABLE "payment_accounts" DROP COLUMN "currency";
COMMIT;

