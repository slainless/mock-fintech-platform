BEGIN;
ALTER TABLE "recurring_payments" DROP CONSTRAINT "recurring_payments_account_uuid_fkey";
ALTER TABLE "transaction_histories" DROP CONSTRAINT "transaction_histories_dest_uuid_fkey";
ALTER TABLE "transaction_histories" DROP CONSTRAINT "transaction_histories_account_uuid_fkey";
ALTER TABLE "payment_accounts" DROP CONSTRAINT "payment_accounts_user_uuid_fkey";

DROP TABLE IF EXISTS "recurring_payments";
DROP TABLE IF EXISTS "transaction_histories";
DROP TABLE IF EXISTS "payment_accounts";
DROP TABLE IF EXISTS "users";
COMMIT;