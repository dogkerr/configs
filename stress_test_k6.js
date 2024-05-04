import http from 'k6/http'
import {sleep, check } from 'k6';

// biar graph nginx di grafana naik 
// hati hati laptop anda bisa rusak wkwk
// jangan di run sampai selesai wkwkw
// cara install k6: https://k6.io/docs/get-started/installation/
// cara run: k6 run stress_testing_k6.js

export let options =  {
    insecureSkipTlsVerify: true,
    noConnectionReuse: false,
    stages: [
        { duration: '1m', target: 500},
        { duration: '2m', target: 600},
        { duration: '1m', target: 700},
        { duration: '2m', target: 700},

    ]
}

const nginx_url = "http://127.0.0.1:80"

export default () => {
    http.get(nginx_url)
    check(res, {'200': (r) => r.status === 200})
    sleep(1);
}
