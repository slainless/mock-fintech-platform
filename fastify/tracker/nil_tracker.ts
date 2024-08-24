
import type { ErrorTracker } from '../platform/error_tracker'

export class NilTracker implements ErrorTracker {
  async report(signal: AbortSignal, error: Error): Promise<void> {}
}