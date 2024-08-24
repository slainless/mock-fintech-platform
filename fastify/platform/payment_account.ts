import type { payment_accounts, shared_account_access } from '../prisma/generated/types'

export interface SharedAccountAccess extends shared_account_access {}
export interface PaymentAccount extends payment_accounts {
	permission: number
}
export interface PaymentAccountDetail extends PaymentAccount {
	permissions: SharedAccountAccess[]
}
