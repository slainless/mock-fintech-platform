BEGIN;

DROP INDEX IF EXISTS "idx_recurring_payments_scheduler_type";
DROP INDEX IF EXISTS "idx_recurring_payments_account_uuid";
DROP INDEX IF EXISTS "idx_recurring_payments_service_id";

DROP INDEX IF EXISTS "idx_transaction_histories_status";
DROP INDEX IF EXISTS "idx_transaction_histories_user_uuid";
DROP INDEX IF EXISTS "idx_transaction_histories_service_id";

DROP INDEX IF EXISTS "idx_monetary_accounts_user_uuid";
DROP INDEX IF EXISTS "idx_monetary_accounts_service_id";

DROP TABLE IF EXISTS "recurring_payments" CASCADE;
DROP TABLE IF EXISTS "transaction_histories" CASCADE;
DROP TABLE IF EXISTS "monetary_accounts" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;

COMMIT;