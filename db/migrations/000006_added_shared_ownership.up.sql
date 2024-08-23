CREATE TABLE "shared_ownership" (
  "account_uuid" uuid NOT NULL,
  "user_uuid" uuid NOT NULL,

  status smallint NOT NULL,
  PRIMARY KEY ("account_uuid", "user_uuid")
)