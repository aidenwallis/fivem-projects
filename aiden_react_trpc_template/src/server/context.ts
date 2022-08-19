import { inferAsyncReturnType } from "@trpc/server";
import { CreateFastifyContextOptions } from "@trpc/server/adapters/fastify";
import { FastifyRequest } from "fastify";

import { getSession } from "./auth/sdk";

export type Context = inferAsyncReturnType<typeof createContext>;

const headerKey = "x-fivem-auth";

async function resolveSession(req: FastifyRequest) {
  const authHeader = (
    (Array.isArray(req.headers[headerKey]) ? req.headers[headerKey][0] : req.headers[headerKey]) || ""
  ).trim();
  if (!authHeader) {
    return null;
  }
  return getSession(authHeader);
}

export async function createContext({ req }: CreateFastifyContextOptions) {
  const session = await resolveSession(req);
  return { session };
}
