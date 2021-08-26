import http from 'k6/http';

export let options = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: 'per-vu-iterations',
      vus: 3,
      iterations: 50,
      maxDuration: '1h30m',
    },
  },
};

export default function () {
  http.get('http://localhost:3000/api/collections');
}
