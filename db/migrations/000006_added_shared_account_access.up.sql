BEGIN;
CREATE TABLE "shared_account_access" (
  "account_uuid" uuid NOT NULL,
  "user_uuid" uuid NOT NULL,

  "permission" int NOT NULL,
  PRIMARY KEY ("account_uuid", "user_uuid")
);

ALTER TABLE "shared_account_access" ADD FOREIGN KEY ("account_uuid") REFERENCES "payment_accounts" ("uuid");
ALTER TABLE "shared_account_access" ADD FOREIGN KEY ("user_uuid") REFERENCES "users" ("uuid");

ALTER TABLE "transaction_histories" ADD COLUMN "issuer_uuid" uuid;
ALTER TABLE "transaction_histories" ADD FOREIGN KEY ("issuer_uuid") REFERENCES "users" ("uuid");
UPDATE "transaction_histories" SET "issuer_uuid" = (SELECT "user_uuid" FROM "payment_accounts" WHERE "payment_accounts"."uuid" = "transaction_histories"."account_uuid");
COMMIT;