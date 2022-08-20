import "./styles/index.scss";

import { DOMController } from "./dom";
import { SpeedController } from "./speed";
import { TimerController } from "./timer";

const dom = new DOMController();
const timer = new TimerController(dom);
const speed = new SpeedController(dom, timer);

window.addEventListener("message", (event) => {
  const t = event?.data?.type || "";

  switch (t) {
    case "ui-show": {
      const show = !!event?.data?.show;
      dom.reset();
      dom.setShow(show);
      timer.reset();
      show && speed.reset();
      break;
    }

    case "update-speed": {
      speed.updateSpeed(event?.data?.speed || 0);
      break;
    }

    case "recording": {
      timer.start();
      break;
    }

    case "finished": {
      timer.stop();
      dom.finish();
      break;
    }

    case "adjust-distance": {
      dom.setDistanceSlider(event?.data?.distance || 20);
      dom.showDistanceModal(true);
      break;
    }
  }
});
