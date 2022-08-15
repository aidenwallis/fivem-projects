import { ApiCode, CreateSessionResponse, EventType } from "../types";
import { clearSessions, createSession, dropSession } from "./utils";

// handle token request events
onNet(EventType.TokenRequest, () => {
  const source = global.source;
  const respond = (body: CreateSessionResponse) => emitNet(EventType.TokenResponse, source, JSON.stringify(body));

  createSession(getPlayerIdentifiers(source))
    .then((body) => {
      setImmediate(() => respond(body));
    })
    .catch((error) => {
      setImmediate(() => respond({ code: ApiCode.Unknown, message: error.toString() }));
    });
});

// handle when someone disconnects
on("playerDropped", () => {
  const source = global.source;
  setImmediate(() => dropSession(getPlayerIdentifiers(source)));
});

setImmediate(() => {
  clearSessions();
});
