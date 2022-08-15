// from internal/schema/codes/error_code.go
export enum ApiCode {
  Unknown = 0,
  InvalidBody = 1,
  ValidationError = 2,
  InvalidToken = 3,
}

export type ApiError = {
  code: ApiCode;
  message: string;
};

export type CreateSessionResponseBody = {
  identifiers: string[];
  metadata: unknown;
  created_at: string;
  expires_at: string;
  token: string;
};

export type CreateSessionResponse = ApiError | CreateSessionResponseBody;

export enum EventType {
  TokenRequest = "aiden:fivem_auth::token:request",
  TokenResponse = "aiden:fivem_auth::token:response",
}

export function isApiError(body: CreateSessionResponse): body is ApiError {
  return (body as ApiError)?.message !== undefined;
}
