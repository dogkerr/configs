//  ini buat load testing
// biar graph network,cpu, ram nginx di grafananya naik
// cara install k6: https://k6.io/docs/get-started/installation/
// cara run: k6 run load_testing.js
import http from 'k6/http'
import {sleep, check } from 'k6';

export let options = {
    stages: [
        { duration: '30s', target: 1000}, // dari 0 virutalUser/second ke  1000 virtualUser/second selama 30 detik
        { duration: '2m', target: 3000},
        { duration: '30s', target: 1000},
    ],
    thresholds: {
        http_req_duration: ['p(99)<1000'], // 99% request harus kurang dari 1s
    },
}


export default () => {
    const res =http.get("http://localhost:80")
    check(res, {'200': (r) => r.status === 200})
    sleep(1);
}