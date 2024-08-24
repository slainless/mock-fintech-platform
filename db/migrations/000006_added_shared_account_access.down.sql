BEGIN;
ALTER TABLE "transaction_histories" DROP CONSTRAINT IF EXISTS "shared_account_access_account_uuid_fkey";
ALTER TABLE "transaction_histories" DROP CONSTRAINT IF EXISTS "shared_account_access_user_uuid_fkey";
DROP TABLE "shared_account_access";

ALTER TABLE "transaction_histories" DROP CONSTRAINT IF EXISTS "transaction_histories_issuer_uuid_fkey";
ALTER TABLE "transaction_histories" DROP COLUMN "issuer_uuid";
COMMIT;