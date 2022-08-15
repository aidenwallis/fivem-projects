import { CreateSessionResponseBody, EventType } from "../types";

// MIN_DELTA is the minimum time a session is alive for
const MIN_DELTA = 300 * 1000;

const resolveDelta = (body: CreateSessionResponseBody) => {
  try {
    return Math.max(MIN_DELTA, new Date(body.expires_at).getTime() - Date.now() - MIN_DELTA);
  } catch (error) {
    return MIN_DELTA;
  }
};

export class TokenClient {
  private currentSession: CreateSessionResponseBody | null = null;
  private timer: NodeJS.Timeout;
  private failedAttempts = 0;

  public getToken() {
    return this.currentSession?.token || "";
  }

  public requestToken() {
    this.clearTimer();
    emitNet(EventType.TokenRequest);
    this.setNextRefresh(30 * 1000, true); // set a timeout, if the server doesn't respond in 30s, try again.
  }

  public receiveToken(body: CreateSessionResponseBody) {
    this.clearTimer(); // lets us cancel the timeout when we retry the request
    this.failedAttempts = 0;
    this.currentSession = body;
    this.setNextRefresh(resolveDelta(body));
  }

  public receiveError() {
    // We exponential backoff with a jitter for up to 10s to avoid outage stampedes
    this.failedAttempts++;
    this.setNextRefresh(Math.min((this.failedAttempts + Math.ceil(Math.random() * 10)) * 2000, 10 * 1000));
  }

  private clearTimer() {
    this.timer && clearTimeout(this.timer);
  }

  private setNextRefresh(duration: number, isRetry?: boolean) {
    this.clearTimer();
    console.log(`Scheduling next ${isRetry ? "retry" : "refresh"} in ${duration}ms`);
    this.timer = setTimeout(() => this.requestToken(), duration);
  }
}

export const tokenClient = new TokenClient();
