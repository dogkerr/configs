//  ini buat load testing
// biar graph network,cpu, ram nginx di grafananya naik
// cara install k6: https://k6.io/docs/get-started/installation/
// cara run: k6 run load_testing.js
import http from 'k6/http'
import {sleep, check } from 'k6';

export let options = {
    stages: [
        { duration: '8s', target: 100}, // dari 0 virutalUser/second ke  1000 virtualUser/second selama 30 detik
        { duration: '15s', target: 200},
        { duration: '8s', target: 200},
    ],
    thresholds: {
        http_req_duration: ['p(99)<1000'], // 99% request harus kurang dari 1s
    },
}


export default () => {
    const res =http.get("http://127.0.0.1:82")
    check(res, {'200': (r) => r.status === 200})
    sleep(1);
}
