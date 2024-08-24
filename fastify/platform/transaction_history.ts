import type { transaction_histories } from '../prisma/generated/types'

export interface TransactionHistory extends transaction_histories {
	serviceID: string
	userUUID: string
}