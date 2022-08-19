import fetch from "node-fetch";
import { z } from "zod";

import { appConfig } from "../config";

const jsonField = z.union([z.string(), z.number(), z.boolean(), z.null()]);
type JsonField = z.infer<typeof jsonField>;
type JsonMap = JsonField | { [key: string]: JsonMap } | JsonMap[];
const jsonMap: z.ZodType<JsonMap> = z.lazy(() => z.union([jsonField, z.array(jsonMap), z.record(jsonMap)]));

const responseSchema = z.object({
  identifiers: z.array(z.string()).min(1),
  metadata: jsonMap,
  created_at: z.string(),
  expires_at: z.string(),
});

export type Session = z.infer<typeof responseSchema>;

// getSession fetches the session from aiden_auth
export async function getSession(token: string): Promise<Session | null> {
  try {
    const resp = await fetch(appConfig.auth.host + "/v1/sessions", {
      headers: {
        "x-fivem-auth": token,
      },
    });
    if (resp.status === 401) {
      // unauthorized session
      return null;
    }

    const data = (await resp.json()) as Session;

    // validate response from api
    await responseSchema.parseAsync(data);

    return data;
  } catch (error) {
    console.error(`Failed to resolve session`, error);
    return null;
  }
}
