BEGIN;
ALTER TABLE "payment_accounts" DROP CONSTRAINT IF EXISTS "unique_payment_accounts_user_uuid_foreign_id_service_id";
ALTER TABLE "payment_accounts" 
  ADD CONSTRAINT "unique_payment_accounts_user_uuid_foreign_id" 
  UNIQUE ("user_uuid", "foreign_id");
COMMIT;

