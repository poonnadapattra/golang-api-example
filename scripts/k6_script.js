import http from 'k6/http';
import { check } from 'k6';

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

var url = 'https://api-golang-restful.herokuapp.com/api/collections'

export default function () {
  let res = http.get(url);
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
}
