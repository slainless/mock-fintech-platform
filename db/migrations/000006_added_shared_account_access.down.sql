BEGIN;
DROP TABLE "shared_account_access";
ALTER TABLE "transaction_histories" DROP CONSTRAINT "transaction_histories_issuer_uuid_fkey";
ALTER TABLE "transaction_histories" DROP COLUMN "issuer_uuid";
COMMIT;