//  ini buat load testing
// biar graph network,cpu, ram nginx di grafananya naik
// cara install k6: https://k6.io/docs/get-started/installation/
// cara run: k6 run load_testing.js
import http from 'k6/http'
import {sleep, check } from 'k6';
import { Trend, Rate } from 'k6/metrics';

export let options = {
    stages: [
        { duration: '10s', target: 4}, // dari 0 virutalUser/second ke  4 virtualUser/second selama 30 detik
        { duration: '20s', target: 8},
        { duration: '10s', target: 4},
    ],
    thresholds: {
        http_req_duration: ['p(99)<1000'], // 99% request harus kurang dari 1s
    },
}

// trend
const user1AuthTrend = new Trend('List_user1Auth');
const user1BadTrend = new Trend('List_user1Bad');
const user1UnauthorizedTrend = new Trend('List_user1Unauthorized');
const user1ForbiddenTrend = new Trend('List_user1Forbidden');
const user1ServerErrorTrend = new Trend('List_user1ServerError');

const user2AuthTrend = new Trend('List_user2Auth');
const user2BadTrend = new Trend('List_user2Bad');
const user2UnauthorizedTrend = new Trend('List_user2Unauthorized');
const user2ForbiddenTrend = new Trend('List_user2Forbidden');
const user2ServerErrorTrend = new Trend('List_user2ServerError');


// Error rate
const user1AuthErrorRate = new Rate('List_user1Auth_Error_Rate');
const user1BadErrorRate = new Rate('List_user1Bad_Error_Rate');
const user1UnauthorizedErrorRate = new Rate('List_user1Unauthorized_Error_Rate');
const user1ForbiddenErrorRate = new Rate('List_user1Forbidden_Error_Rate');
const user1ServerErrorErrorRate = new Rate('List_user1ServerError_Error_Rate');

const user2AuthErrorRate = new Rate('List_user2Auth_Error_Rate');
const user2BadErrorRate = new Rate('List_user2Bad_Error_Rate');
const user2UnauthorizedErrorRate = new Rate('List_user2Unauthorized_Error_Rate');
const user2ForbiddenErrorRate = new Rate('List_user2Forbidden_Error_Rate');
const user2ServerErrorErrorRate = new Rate('List_user2ServerError_Error_Rate');


function getUrl(user_no) {
    let urls  = {
        auth: "http://localhost:823"+ user_no + "/auth",
        bad:  "http://localhost:823"+ user_no + "/bad",
        unauthorized:  "http://localhost:823"+ user_no + "/unauthorized",
        forbidden:  "http://localhost:823"+ user_no + "/forbidden",
        serverError:  "http://localhost:823"+ user_no + "/serverError",
    }
    return urls
}
const auth = "http://localhost:8231/auth"

export default () => {
    const userOneUrls = getUrl("1");
    const userTwoUrls = getUrl("2");

    const requests = {
        'user1_auth': {
          method: 'GET',
          url: userOneUrls.auth,
        },
        'user1_bad': {
          method: 'POST',
          url: userOneUrls.bad,
          body: {},
        },
        'user1_unauthorized': {
            method: 'GET',
            url: userOneUrls.unauthorized,
        },
        'user1_forbidden': {
            method: 'GET',
            url: userOneUrls.forbidden,
        },
        'user1_serverError': {
            method: 'GET',
            url: userOneUrls.serverError,
        },
        'user2_auth': {
            method: 'GET',
            url: userTwoUrls.auth,
        },
        'user2_bad': {
            method: 'POST',
            url: userTwoUrls.bad,
            body: {},
        },
        'user2_unauthorized': {
            method: 'GET',
            url: userTwoUrls.unauthorized,
        },
        'user2_forbidden': {
            method: 'GET',
            url: userTwoUrls.forbidden,
        },
        'user2_serverError': {
            method: 'GET',
            url: userTwoUrls.serverError,
        }  
      };


      const responses = http.batch(requests);

      // user 1 request all endpoints
      const user1AuthResp = responses['user1_auth'];
      check(user1AuthResp, {
        'status is 200': (r) => r.status === 200,
      }) || user1AuthErrorRate.add(1);

      user1AuthTrend.add(user1AuthResp.timings.duration)

      // bad request 
      const user1BadResp = responses['user1_bad'];
      check(user1BadResp, {
        'status is 200': (r) => r.status === 200,
      }) || user1BadErrorRate.add(1);

      user1BadTrend.add(user1BadResp.timings.duration)

      // unauthorized
      const user1UnauthorizedResp = responses['user1_unauthorized'];
      check(user1UnauthorizedResp, {
        'status is 200': (r) => r.status === 200,
      }) || user1UnauthorizedErrorRate.add(1);

      user1UnauthorizedTrend.add(user1UnauthorizedResp.timings.duration)


      // forbidden
      const user1ForbiddenResp = responses['user1_forbidden'];
      check(user1ForbiddenResp, {
        'status is 200': (r) => r.status === 200,
      }) || user1ForbiddenErrorRate.add(1);

      user1ForbiddenTrend.add(user1ForbiddenResp.timings.duration)

      // serverError
      const user1ServerErrorResp = responses['user1_serverError'];
      check(user1ServerErrorResp, {
        'status is 200': (r) => r.status === 200,
      }) || user1ServerErrorErrorRate.add(1);

      user1ServerErrorTrend.add(user1ServerErrorResp.timings.duration)






       // user 2 request all endpoints
       const user2AuthResp = responses['user2_auth'];
       check(user2AuthResp, {
         'status is 200': (r) => r.status === 200,
       }) || user2AuthErrorRate.add(1);
 
       user2AuthTrend.add(user2AuthResp.timings.duration)
 
       // bad request 
       const user2BadResp = responses['user2_bad'];
       check(user2BadResp, {
         'status is 200': (r) => r.status === 200,
       }) || user2BadErrorRate.add(1);
 
       user2BadTrend.add(user2BadResp.timings.duration)
 
       // unauthorized
       const user2UnauthorizedResp = responses['user2_unauthorized'];
       check(user2UnauthorizedResp, {
         'status is 200': (r) => r.status === 200,
       }) || user2UnauthorizedErrorRate.add(1);
 
       user2UnauthorizedTrend.add(user2UnauthorizedResp.timings.duration)
 
 
       // forbidden
       const user2ForbiddenResp = responses['user2_forbidden'];
       check(user2ForbiddenResp, {
         'status is 200': (r) => r.status === 200,
       }) || user2ForbiddenErrorRate.add(1);
 
       user2ForbiddenTrend.add(user2ForbiddenResp.timings.duration)
 
       // serverError
       const user2ServerErrorResp = responses['user2_serverError'];
       check(user2ServerErrorResp, {
         'status is 200': (r) => r.status === 200,
       }) || user2ServerErrorErrorRate.add(1);
 
       user2ServerErrorTrend.add(user2ServerErrorResp.timings.duration)
 
 
    sleep(1);
}
