import { DOMController } from "./dom";

const pad = (v: number) => ("00" + v).slice(-2);

export class TimerController {
  private startedAt = 0;
  private timer: number;
  private started = false;

  public constructor(private dom: DOMController) {}

  public start() {
    this.stop();
    this.dom.setActiveState(true);
    this.startedAt = Date.now();
    this.tick();
    this.timer = (setInterval(() => this.tick(), 10) as unknown) as number;
    this.started = true;
  }

  public hasStarted() {
    return this.started;
  }

  public reset() {
    this.stop();
    this.dom.setActiveState(false);
  }

  public stop() {
    this.started = false;
    this.timer && clearInterval(this.timer);
  }

  private tick() {
    this.updateLabel();
  }

  private updateLabel() {
    let s = Date.now() - this.startedAt;
    const ms = s % 1000;
    s = (s - ms) / 1000;
    const secs = s % 60;
    s = (s - secs) / 60;
    const mins = s % 60;

    this.dom.setActiveLabel(`${pad(mins)}:${pad(secs)}.${pad(Math.floor(ms / 10))}`);
  }
}
