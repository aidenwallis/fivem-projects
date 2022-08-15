import "./styles/index.scss";

const container = document.getElementById("container")!;
const speed = document.getElementById("speed-meter")!;
const rpm = document.getElementById("rpm-meter")!;
const gear = document.getElementById("gear")!;
const health = document.getElementById("health-meter")!;
const showClass = "show";
const gearChangeClass = "gear-change";

const percentPath = (v: number) => `polygon(0 0, ${v}% 0, ${v}% 100%, 0 100%)`;

window.addEventListener("message", (event) => {
  const { data } = event;

  switch (data?.type) {
    case "entered-vehicle": {
      container?.classList.add(showClass);
      break;
    }

    case "left-vehicle": {
      container?.classList.remove(showClass);
      speed.textContent = "0";
      break;
    }

    case "update-stats": {
      data.speed !== undefined && (speed.textContent = data.speed);
      data.rpm !== undefined && (rpm.style.clipPath = percentPath(data.rpm));
      data.health !== undefined && (health.style.width = percentPath(data.health));

      if (data.gear !== undefined) {
        gear.classList.remove(gearChangeClass);
        gear.textContent = data.gear === 0 ? "R" : data.gear;
        void gear.offsetWidth; // force a repaint before reapplying class
        gear.classList.add(gearChangeClass);
      }
      break;
    }
  }
});
