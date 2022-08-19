import { createReactQueryHooks, TRPCClientErrorLike } from "@trpc/react";
import { inferProcedureInput, inferProcedureOutput } from "@trpc/server";
import * as React from "react";
import { QueryClient, QueryClientProvider } from "react-query";

import { AppRouter } from "../../../trpc";

declare const process: {
  env: {
    TRPC_HOST: string;
  };
};

export const trpc = createReactQueryHooks<AppRouter>();

let cachedToken: string;

export function setTRPCToken(token: string) {
  cachedToken = token;
}

export const TRPCProvider: React.FC<{ children: React.ReactNode }> = (props) => {
  const [queryClient] = React.useState(new QueryClient());
  const [trpcClient] = React.useState(
    trpc.createClient({
      url: process.env.TRPC_HOST,
      headers() {
        return {
          "x-fivem-auth": cachedToken,
        };
      },
    })
  );

  return (
    <trpc.Provider client={trpcClient} queryClient={queryClient}>
      <QueryClientProvider client={queryClient}>{props.children}</QueryClientProvider>
    </trpc.Provider>
  );
};

export type TClientError = TRPCClientErrorLike<AppRouter>;

/**
 * Enum containing all api query paths
 */
export type TQuery = keyof AppRouter["_def"]["queries"];

/**
 * Enum containing all api mutation paths
 */
export type TMutation = keyof AppRouter["_def"]["mutations"];

/**
 * This is a helper method to infer the output of a query resolver
 * @example type HelloOutput = InferQueryOutput<'hello'>
 */
export type InferQueryOutput<TRouteKey extends TQuery> = inferProcedureOutput<AppRouter["_def"]["queries"][TRouteKey]>;

/**
 * This is a helper method to infer the input of a query resolver
 * @example type HelloInput = InferQueryInput<'hello'>
 */
export type InferQueryInput<TRouteKey extends TQuery> = inferProcedureInput<AppRouter["_def"]["queries"][TRouteKey]>;

/**
 * This is a helper method to infer the output of a mutation resolver
 * @example type HelloOutput = InferMutationOutput<'hello'>
 */
export type InferMutationOutput<TRouteKey extends TMutation> = inferProcedureOutput<
  AppRouter["_def"]["mutations"][TRouteKey]
>;

/**
 * This is a helper method to infer the input of a mutation resolver
 * @example type HelloInput = InferMutationInput<'hello'>
 */
export type InferMutationInput<TRouteKey extends TMutation> = inferProcedureInput<
  AppRouter["_def"]["mutations"][TRouteKey]
>;
