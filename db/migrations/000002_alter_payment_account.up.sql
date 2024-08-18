BEGIN;
ALTER TABLE "payment_accounts" ADD COLUMN "name" varchar(255);
ALTER TABLE "payment_accounts" ADD COLUMN "foreign_id" varchar(255) NOT NULL;
ALTER TABLE "payment_accounts" DROP COLUMN "balance";
ALTER TABLE "payment_accounts" DROP COLUMN "currency";

ALTER TABLE "payment_accounts" ADD CONSTRAINT "unique_payment_accounts_user_uuid_foreign_id" UNIQUE ("user_uuid", "foreign_id");
COMMIT;

