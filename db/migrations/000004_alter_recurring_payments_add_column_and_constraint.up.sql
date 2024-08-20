BEGIN;
ALTER TABLE "recurring_payments" ADD COLUMN "foreign_id" varchar(255) NOT NULL;
ALTER TABLE "recurring_payments" 
  ADD CONSTRAINT "recurring_payments_account_uuid_foreign_id_service_id" 
  UNIQUE ("account_uuid", "foreign_id", "service_id");
COMMIT;

