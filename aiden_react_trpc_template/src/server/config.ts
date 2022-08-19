import * as fs from "fs";
import * as path from "path";
import { z } from "zod";

const schema = z.object({
  environment: z.enum(["development", "production"]),
  port: z.number(),

  // how to connect to our db
  database: z.object({
    url: z.string(),
  }),

  // aiden_auth options
  auth: z.object({
    host: z.string(),
  }),

  trpc: z.object({
    host: z.string(),
  }),
});

export const appConfig = JSON.parse(fs.readFileSync(path.join(__dirname, "../../config.json"), "utf-8")) as z.infer<
  typeof schema
>;

try {
  schema.parse(appConfig);
} catch (error) {
  console.error(`Invalid config.json!`);
  console.error(error);
}
