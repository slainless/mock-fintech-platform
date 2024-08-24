import type { FastifyRequest } from 'fastify'
import { ErrInvalidCredential, ErrUnsupportedCredential, type AuthService } from '../platform/auth_service'
import { verify } from 'jsonwebtoken'

export class EmailJWTAuthService implements AuthService {
	constructor(protected secretKey: string) {}

	serviceID(): string {
		return "supabase_jwt"
	}

	async credential(signal: AbortSignal, request: FastifyRequest): Promise<any> {
		const token = request.headers.authorization
		if(token == null) throw ErrInvalidCredential
		if(token.startsWith('Bearer '))	return token.substring(7)
		throw ErrUnsupportedCredential
	}
	
	async validate(signal: AbortSignal, credential: any): Promise<string> {
		if(typeof credential != 'string') throw ErrInvalidCredential
		
		signal.throwIfAborted()
		const decoded = verify(credential, this.secretKey)
		
		if(typeof decoded == 'string') throw ErrInvalidCredential
		if(!decoded.email) throw ErrInvalidCredential
		return decoded.email
	}
}