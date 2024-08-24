import type { ColumnType } from "kysely";
export type Generated<T> = T extends ColumnType<infer S, infer I, infer U>
  ? ColumnType<S, I | undefined, U>
  : ColumnType<T, T | undefined, T>;
export type Timestamp = ColumnType<Date, Date | string, Date | string>;

export type payment_accounts = {
    uuid: Generated<string>;
    user_uuid: string;
    service_id: string;
    name: string | null;
    foreign_id: string;
};
export type recurring_payments = {
    uuid: Generated<string>;
    service_id: string;
    account_uuid: string;
    scheduler_type: number;
    last_charge: Timestamp | null;
    foreign_id: string;
    charging_method: number;
};
export type schema_migrations = {
    version: string;
    dirty: boolean;
};
export type shared_account_access = {
    account_uuid: string;
    user_uuid: string;
    permission: number;
};
export type transaction_histories = {
    uuid: Generated<string>;
    account_uuid: string;
    dest_uuid: string | null;
    mutation: string;
    currency: string;
    status: number;
    address: string | null;
    transaction_note: string | null;
    transaction_date: Timestamp;
    transaction_type: number;
    issuer_uuid: string | null;
};
export type users = {
    uuid: Generated<string>;
    full_name: string | null;
    user_name: string | null;
    email: string;
    password_hash: string | null;
};
export type DB = {
    payment_accounts: payment_accounts;
    recurring_payments: recurring_payments;
    schema_migrations: schema_migrations;
    shared_account_access: shared_account_access;
    transaction_histories: transaction_histories;
    users: users;
};
