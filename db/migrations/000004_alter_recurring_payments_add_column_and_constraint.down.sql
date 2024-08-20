BEGIN;
ALTER TABLE "recurring_payments" DROP COLUMN "foreign_id";
ALTER TABLE "recurring_payments" DROP COLUMN "charging_method";
ALTER TABLE "recurring_payments" DROP CONSTRAINT IF EXISTS "recurring_payments_account_uuid_foreign_id_service_id";
COMMIT;