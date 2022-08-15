import { CreateSessionResponse, EventType, isApiError } from "../types";
import { tokenClient } from "./client";

setImmediate(() => tokenClient.requestToken());

onNet(EventType.TokenResponse, (jsonifiedBody) => {
  try {
    const body: CreateSessionResponse = JSON.parse(jsonifiedBody);
    if (isApiError(body) || !body.token) {
      throw new Error(`Error while requesting token: ${jsonifiedBody}`);
    }

    console.log("Received session!");
    tokenClient.receiveToken(body);
  } catch (error) {
    console.error(`Failed to receive token: ${error.toString()}`);
    tokenClient.receiveError();
  }
});
