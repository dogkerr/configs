# Dogkur config monitoring Prometheus, Grafana & config logging ELK stack

how to run monitoring:
docker network create dogker

1. nano /etc/docker/daemon.json, ini bikin docker metrics bisa di scrape sama prometheus

```
1.   isi daemon.json:
   {
   "metrics-addr": "0.0.0.0:9323",
   "experimental": true
   }



2.  sudo systemctl restart docker

```

2. allow dessire port fierewall:

```
1. https://www.inmotionhosting.com/support/security/how-to-install-firewalld-on-linux/
2. firewall-cmd --zone=public --add-port=9323/tcp
3. systemctl restart docker
```

3. docker compose up -d

4.buka dashboard grafana (localhost:3000) 5. buat service account + buat api key buat service acc tersebut 6. copy api key di grafana_config_writer.go

```
(398) r.Header.Add("Authorization", "Bearer <api_key>")

```

7. buat datasource prometheus di grafana dashboard

```
1. buka sidebar
2. klik datasource > add prometheus
3. server url : http://prometheus:9090
4. save & test
```

8. jalanin beberapa container dg label user_id yang samaa/beda

```


 <!-- 
 
 docker service create --name  ninggx  --publish 8080:80 --replicas 4 --container-label  user_id="18d2e020-538d-449a-8e9c-02e4e5cf41111"  nginx:latest
 -->

testing nya pake container go  aja:


 docker service create --name  go_container  --publish 8080:80 --replicas 3 --container-label  user_id="18d2e020-538d-449a-8e9c-02e4e5cf41111"  generate_user_dashboard_dan_perfomance_testing-go_container_log_user1:latest


```

9. masukin user_id ke grafana_config_writer

```
(410) 	uidDashboard := createNewDashboardPerUser("<user_id>")


```

10. jalanin golang config writer

```
go run grafana_config_writer.go
```

11. yang munculdi log golang itu link iframenya

```
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=8
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=9
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=1
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=34
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=10
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=37
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=5
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=31

```

12. jalanin load testing ke nginx pake k6 biar grafik grafananya naik

```
k6 run load_testing.js
```

## link iframe format (buat embed ke html)

- format:

```
http://127.0.0.1/d-solo/zrEtsX3d/zretsx3d?orgId=1&refresh=5s&from=now-5m&theme=light&to=now&panelId=1
```

- list from buat time rangenya (real time):

```
1. now-5m
2. now-15m
3. now-30m
4. now-1h
5. now-3h
6. now-6h
7. now-12h
8. now-24h
9. now-2d


```

- buat query param "to" nilainya "now" biar realtime

- list from buat time range (tapi pake dayFrom - dayTo) tidak real time:

```
format: &from=1709226000000&to=1709312399000
pake unix epoch (tapi kayake gak realtime kalo gini)
1709226000000:  Fri Mar 01 2024 00:00:00
1709312399000:     Friday, March 1, 2024 11:59:59
```

- list refresh dashboard period:

```
format: &refresh=5s
- 5s
- auto
- gak ngasih query param refresh (gak pernah direfresh )
- 10s
- 30s
- 1m
- 5m
- 15m
- 30m
- 1h
- 2h
- 1d

```

- tombol refresh kayaknya tinggal refresh iframenya deh ?

## Configuring Loki

1. ubah daemon.json nya docker biar log semua docker container dikirim ke loki

```
1. sudo nano /etc/docker/daemon.json

isi file:

{
   "metrics-addr": "0.0.0.0:9323",
   "experimental": true,
   "log-driver": "loki",
    "log-opts": {
        "loki-url": "http://localhost:3100/loki/api/v1/push",
        "loki-batch-size": "400"
    }

}

```

2. docker plugin install grafana/loki-docker-driver:2.9.4 --alias loki --grant-all-permissions

3. restart docker

```
 sudo systemctl restart docker

```

3. docker compose up -d

4. jalanin banyak request ke go_container_user1 dan go_container_user2 dengan k6

```
k6 run generate_log_golang.js

```

5. buka grafana (localhost:3000) & add datasource loki

```
1. buka sidebar
2. klik datasource >  add datasource
3. klik loki
4. masukin url: <your_ip_address>:3100

```

6. buka tab explore di grafana

```
1. buka sidebar > klik explore
2. pilih source nya loki
3. pilih label filer  "container_name"
4. pilih tanda "=~"
5. pilih go_container_api_user2, go_container_api_user1

6. coba juga query code:
- {container_name=~"go_container_api_user2|go_container_api_user1"} |= `401`

- {container_name=~"go_container_api_user2|go_container_api_user1"} |= `200`

- {container_name=~"go_container_api_user2|go_container_api_user1"} |= `500`

- {userId="18d2e020-538d-449a-8e9c-011212999"} |= `` | json | status_code="500"
- {userId="18d2e020-538d-449a-8e9c-011212999"} |= `` | json | status_code="400"
- {userId="18d2e020-538d-449a-8e9c-011212999"} |= `` | json | level="error"
```

6. bikins service account baru & buatin api key nya

7. import dashboard

```
1. dapetin id datasource loki dg command: curl --insecure  -H "Authorization: Bearer <api_key_service_account>"  http://localhost:3000/api/datasources
2. ubah datasourceId di file logs loki per-user-17... , jadi id datasource yang didapet tadi
3. buka tab dashboard
4. import file logs loki per user-17.....

```
