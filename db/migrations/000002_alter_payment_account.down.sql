BEGIN;
ALTER TABLE "payment_accounts" DROP COLUMN "name";
ALTER TABLE "payment_accounts" DROP COLUMN "foreign_id";
ALTER TABLE "payment_accounts" ADD COLUMN "balance" bigint NOT NULL DEFAULT 0;
ALTER TABLE "payment_accounts" ADD COLUMN "currency" varchar(24) NOT NULL;

ALTER TABLE "payment_accounts" DROP CONSTRAINT IF EXISTS "unique_payment_accounts_user_uuid_foreign_id";
COMMIT;

