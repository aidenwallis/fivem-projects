import { rpc } from "./rpc";

const showClass = "show";
const activeClass = "active";
const DEFAULT_DISTANCE = 20;

export class DOMController {
  private container = document.getElementById("container")!;
  private activeState = document.getElementById("active-state")!;
  private currentSpeed = document.getElementById("current-speed")!;
  private distanceModal = document.getElementById("distance-modal")!;
  private distanceSlider = document.getElementById("distance-slider")! as HTMLInputElement;
  private distanceReset = document.getElementById("distance-reset")!;
  private distanceSubmit = document.getElementById("distance-submit")!;
  private reminder = document.getElementById("reminder")!;
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
    if (!this.distanceModal) {
      throw new Error("distance modal node not found");
    }
    if (!this.topSpeed) {
      throw new Error("top speed node not found");
    }
    if (!this.reminder) {
      throw new Error("reminder node not found");
    }
    if (!this.distanceSubmit) {
      throw new Error("distance submit node not found");
    }

    this.registerSubmitHandler();
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
    if (active) {
      this.activeState.classList.add(activeClass);
      this.reminder.classList.remove(showClass);
    } else {
      this.activeState.classList.remove(activeClass);
      this.reminder.classList.add(showClass);
    }
  }

  public reset() {
    this.setActiveLabel("Inactive");
    this.setActiveState(false);
    this.topSpeed.classList.remove(activeClass);
  }

  public finish() {
    this.topSpeed.classList.add(activeClass);
  }

  public showDistanceModal(show: boolean) {
    show ? this.distanceModal.classList.add(showClass) : this.distanceModal.classList.remove(showClass);
  }

  public setDistanceSlider(value: number) {
    this.distanceSlider.value = value.toString();
    rpc('setDistance', { distance: this.distanceSlider.valueAsNumber });
  }

  private registerSubmitHandler() {
    this.distanceSlider.oninput = () => {
      this.setDistanceSlider(this.distanceSlider.valueAsNumber);
    };

    this.distanceReset.onclick = () => {
      this.setDistanceSlider(DEFAULT_DISTANCE);
    };

    this.distanceSubmit.onclick = (event) => {
      event.preventDefault();
      this.showDistanceModal(false);
      rpc('saveDistance', {});
    };
  }
}
