import type { PaymentAccount } from './payment_account'
import type { RecurringPayment } from './recurring_payment'
import type { TransactionHistory } from './transaction_history'

export enum RecurringPaymentChargingMethod {
	Upfront,
	NonUpfront
}

export interface RecurringPaymentService {
	subscribe(signal: AbortSignal, account: PaymentAccount, billingID: string, callbackData: string): Promise<{ payment: RecurringPayment, history: TransactionHistory }>
	unsubscribe(signal: AbortSignal, payment: RecurringPayment): Promise<TransactionHistory>
	bill(signal: AbortSignal, payment: RecurringPayment): Promise<TransactionHistory>
}