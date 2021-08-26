import http from 'k6/http';

export let options = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: 'per-vu-iterations',
      vus: 3,
      iterations: 100,
      maxDuration: '30m',
    },
  },
};

export default function () {
  http.get('https://api-golang-restful.herokuapp.com/api/collections');
}
