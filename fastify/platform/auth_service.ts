import type { FastifyRequest } from 'fastify'

export const ErrUnsupportedCredential = new Error("unsupported credential")
export const ErrEmptyCredential = new Error("empty credential")
export const ErrInvalidCredential = new Error("invalid credential")
export const ErrUnsupportedHeader = new Error("unsupported header")

export interface AuthService {
	serviceID(): string
	validate(signal: AbortSignal, credential: any): Promise<string>
	credential(signal: AbortSignal, request: FastifyRequest): Promise<any>
}