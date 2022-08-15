import fetch from "node-fetch";

import { CreateSessionResponse } from "../types";

export const sleep = (duration: number) => new Promise((resolve) => setTimeout(resolve, duration));

const serviceHost_ = () => GetConvar("aiden_fivem_auth_internal_addr", "");
const fetch_ = <T>(options: { path: string; body?: T }) =>
  fetch(options.path, {
    headers: {
      "Content-Type": "application/json; charset=utf-8",
    },
    method: "POST",
    body: options.body ? JSON.stringify(options.body) : undefined,
  });

export const createSession = (identifiers: string[]) => {
  const serviceHost = serviceHost_();
  if (!serviceHost) {
    throw new Error("No service host set.");
  }

  return fetch_({
    path: serviceHost + "/v1/sessions",
    body: {
      identifiers: identifiers,
      metadata: {
        // You can really put any metadata you want here.
        currentSource: source,
      },
    },
  }).then((resp) => resp.json() as unknown as CreateSessionResponse);
};

export const clearSessions = () => {
  const serviceHost = serviceHost_();
  if (!serviceHost) {
    return;
  }

  return fetch_({ path: serviceHost + "/v1/clear-sessions" })
    .then((resp) => {
      if (resp.status !== 200) {
        throw new Error(`Non 200 http status code returned: ${resp.status}`);
      }

      console.log("Cleared sessions!");
    })
    .catch((error) => {
      console.error(`Failed to clear sessions: ${error.toString()}`);
    });
};

export const dropSession = async (identifiers: string[]) => {
  const serviceHost = GetConvar("aiden_fivem_auth_internal_addr", "");
  if (!serviceHost) {
    return;
  }

  const attempt = () =>
    fetch_({
      path: serviceHost + "/v1/drop-session",
      body: identifiers,
    })
      .then((resp) => {
        const success = resp.status === 200;

        !success &&
          console.error("Failed to drop session for identifiers", {
            status: resp.status,
            identifiers,
          });

        return success;
      })
      .catch((error) => {
        console.error("Failed to drop session for identifiers", {
          error: error.toString(),
          identifiers,
        });
        return false;
      });

  for (let i = 0; i < 10; i++) {
    if (await attempt()) {
      return;
    }

    await sleep(1000);
  }
};
