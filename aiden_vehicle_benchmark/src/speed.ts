import { DOMController } from "./dom";
import { TimerController } from "./timer";

export class SpeedController {
  private currentSpeed = 0;
  private topSpeed = 0;

  public constructor(private dom: DOMController, private timer: TimerController) {}

  public reset() {
    this.currentSpeed = 0;
    this.topSpeed = 0;
    this.dom.setCurrentSpeed(this.currentSpeed);
    this.dom.setTopSpeed(this.topSpeed);
  }

  public updateSpeed(v: number) {
    const speed = Math.floor(v * 2.236936); // resolve to mph
    if (speed === this.currentSpeed) {
      return;
    }

    this.currentSpeed = speed;
    this.dom.setCurrentSpeed(speed);

    if (this.timer.hasStarted() && this.currentSpeed > this.topSpeed) {
      this.topSpeed = this.currentSpeed;
      this.dom.setTopSpeed(this.topSpeed);
    }
  }
}
