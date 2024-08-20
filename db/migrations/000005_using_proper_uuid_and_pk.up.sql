BEGIN;
SELECT gen_random_uuid();

ALTER TABLE "payment_accounts" DROP CONSTRAINT IF EXISTS "payment_accounts_user_uuid_fkey";
ALTER TABLE "transaction_histories" DROP CONSTRAINT IF EXISTS "transaction_histories_account_uuid_fkey";
ALTER TABLE "transaction_histories" DROP CONSTRAINT IF EXISTS "transaction_histories_dest_uuid_fkey";
ALTER TABLE "recurring_payments" DROP CONSTRAINT IF EXISTS "recurring_payments_account_uuid_fkey";

ALTER TABLE "users" DROP COLUMN "id";
ALTER TABLE "payment_accounts" DROP COLUMN "id";
ALTER TABLE "transaction_histories" DROP COLUMN "id";
ALTER TABLE "recurring_payments" DROP COLUMN "id";

ALTER TABLE "users" ALTER COLUMN "uuid" TYPE uuid USING "uuid"::uuid;
ALTER TABLE "users" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();  

ALTER TABLE "payment_accounts" ALTER COLUMN "uuid" TYPE uuid USING "uuid"::uuid;
ALTER TABLE "payment_accounts" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();;
ALTER TABLE "payment_accounts" ALTER COLUMN "user_uuid" TYPE uuid USING "user_uuid"::uuid;

ALTER TABLE "transaction_histories" ALTER COLUMN "uuid" TYPE uuid USING "uuid"::uuid;
ALTER TABLE "transaction_histories" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();  
ALTER TABLE "transaction_histories" ALTER COLUMN "account_uuid" TYPE uuid USING "account_uuid"::uuid;
ALTER TABLE "transaction_histories" ALTER COLUMN "dest_uuid" TYPE uuid USING "dest_uuid"::uuid; 

ALTER TABLE "recurring_payments" ALTER COLUMN "uuid" TYPE uuid USING "uuid"::uuid;
ALTER TABLE "recurring_payments" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();;
ALTER TABLE "recurring_payments" ALTER COLUMN "account_uuid" TYPE uuid USING "account_uuid"::uuid;

ALTER TABLE "users" ADD CONSTRAINT "users_primary_key" PRIMARY KEY ("uuid");
ALTER TABLE "payment_accounts" ADD CONSTRAINT "payment_accounts_primary_key" PRIMARY KEY ("uuid");
ALTER TABLE "transaction_histories" ADD CONSTRAINT "transaction_histories_primary_key" PRIMARY KEY ("uuid");
ALTER TABLE "recurring_payments" ADD CONSTRAINT "recurring_payments_primary_key" PRIMARY KEY ("uuid");

ALTER TABLE "payment_accounts" ADD CONSTRAINT "payment_accounts_user_uuid_fkey" FOREIGN KEY ("user_uuid") REFERENCES "users" ("uuid");
ALTER TABLE "transaction_histories" ADD CONSTRAINT "transaction_histories_account_uuid_fkey" FOREIGN KEY ("account_uuid") REFERENCES "payment_accounts" ("uuid");
ALTER TABLE "transaction_histories" ADD CONSTRAINT "transaction_histories_dest_uuid_fkey" FOREIGN KEY ("dest_uuid") REFERENCES "payment_accounts" ("uuid");
ALTER TABLE "recurring_payments" ADD CONSTRAINT "recurring_payments_account_uuid_fkey" FOREIGN KEY ("account_uuid") REFERENCES "payment_accounts" ("uuid");
COMMIT;