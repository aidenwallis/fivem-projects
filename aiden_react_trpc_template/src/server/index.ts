import cors from "@fastify/cors";
import { fastifyTRPCPlugin } from "@trpc/server/adapters/fastify";
import fastify from "fastify";

import { appConfig } from "./config";
import { createContext } from "./context";
import { r } from "./router";

const server = fastify({
  maxParamLength: 5000,
});

(async () => {
  try {
    await server.register(cors);

    await server.register(fastifyTRPCPlugin, {
      prefix: "/trpc",
      trpcOptions: {
        router: r,
        createContext: createContext,
      },
    });

    await server.listen({ port: appConfig.port });
    console.log(`Server running on port ${appConfig.port}`);
  } catch (err) {
    server.log.error(err);
    process.exit(1);
  }
})();
