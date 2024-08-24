import type { MonetaryAmount } from './monetary'
import type { PaymentAccount } from './payment_account'
import type { TransactionHistory } from './transaction_history'
import type { User } from './user'

export const ErrInvalidAccountData = new Error('invalid account data');
export const ErrTransactionRejected = new Error('transaction rejected');

export enum TransactionType {
	Withdraw,
	Send
}

export interface PaymentService {
	send(signal: AbortSignal, user: User, source: PaymentAccount, dest: PaymentAccount, amount: number, callbackData: string): Promise<TransactionHistory>;
	withdraw(signal: AbortSignal, user: User, account: PaymentAccount, amount: number, callbackData: string): Promise<TransactionHistory>;
	getMatchingHistory(signal: AbortSignal, account: PaymentAccount, history: TransactionHistory): Promise<TransactionHistory>;
	balance(signal: AbortSignal, account: PaymentAccount): Promise<MonetaryAmount>;
	validate(signal: AbortSignal, user: User, accountForeignID: string, callbackData: string): Promise<void>;
}