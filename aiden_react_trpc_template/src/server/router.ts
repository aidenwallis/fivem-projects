import * as trpc from "@trpc/server";

import { Context } from "./context";

export const r = trpc.router<Context>().query("currentUser.get-session", {
  resolve({ ctx }) {
    if (!ctx.session) {
      return null;
    }
    return {
      identifiers: ctx.session.identifiers,
      createdAt: new Date(ctx.session.created_at),
      expiresAt: new Date(ctx.session.expires_at),
    };
  },
});

// This is exported again in ../trpc.ts, this is to try and discourage people from using something potentially
// incorrectly and dumping a bunch of typescript into their client bundles
export type _appRouter = typeof r;
