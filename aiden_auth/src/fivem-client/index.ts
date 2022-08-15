import "./events";

import { tokenClient } from "./client";

// getSessionToken lets other scripts pull the active session token
global.exports("getSessionToken", () => tokenClient.getToken());

// forceSessionRefresh lets other scripts force a session reload, the call is asynchronous and will
// instantly return, clients should poll getSessionToken and wait for a changed response.
global.exports("forceSessionRefresh", () => tokenClient.requestToken());
