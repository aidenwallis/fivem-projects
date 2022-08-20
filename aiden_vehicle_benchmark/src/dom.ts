const showClass = "show";
const activeClass = "active";

export class DOMController {
  private container = document.getElementById("container")!;
  private activeState = document.getElementById("active-state")!;
  private currentSpeed = document.getElementById("current-speed")!;
  private topSpeed = document.getElementById("top-speed")!;

  public constructor() {
    if (!this.activeState) {
      throw new Error("active state node not found");
    }
    if (!this.container) {
      throw new Error("container not found");
    }
    if (!this.currentSpeed) {
      throw new Error("current speed node not found");
    }
    if (!this.topSpeed) {
      throw new Error("top speed node not found");
    }
  }

  public setShow(show: boolean) {
    show ? this.container.classList.add(showClass) : this.container.classList.remove(showClass);
  }

  public setCurrentSpeed(speed: number) {
    this.currentSpeed.textContent = speed.toString() + " mph";
  }

  public setTopSpeed(speed: number) {
    this.topSpeed.textContent = speed.toString() + " mph";
  }

  public setActiveLabel(label: string) {
    this.activeState.textContent = label;
  }

  public setActiveState(active: boolean) {
    active ? this.activeState.classList.add(activeClass) : this.activeState.classList.remove(activeClass);
  }

  public reset() {
    this.setActiveLabel("Inactive");
    this.setActiveState(false);
    this.topSpeed.classList.remove(activeClass);
  }

  public finish() {
    this.topSpeed.classList.add(activeClass);
  }
}
