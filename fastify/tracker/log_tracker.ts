import type { ErrorTracker } from '../platform/error_tracker'

export class LogTracker implements ErrorTracker {
	async report(signal: AbortSignal, error: Error): Promise<void> {
		console.error(error)
	}
}