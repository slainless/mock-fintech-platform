export interface ErrorTracker {
	report(signal: AbortSignal, error: Error): Promise<void>
}