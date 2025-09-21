import http from "k6/http";

export let options = {
  scenarios: {
    high_load: {
      executor: "constant-arrival-rate",
      rate: 100000,       // requests per second
      timeUnit: "1s",
      duration: "30s",
      preAllocatedVUs: 50000, // pre-spawn VUs
      maxVUs: 100000,         // max allowed VUs
    },
  },
};

export default function () {
  http.get(":8080");
}
