API KEY grafana= glsa_2lJAFLqG9KWOwA2lfpybWL5vikEqnwnc_ece2bad7

curl -X POST --insecure -H "Authorization: Bearer glsa_2lJAFLqG9KWOwA2lfpybWL5vikEqnwnc_ece2bad7" -H "Content-Type: application/json" -d "{\"dashboard\":$(cat docker-quest-prometheus.json)}" http://localhost:3000/api/dashboards/db


curl -X POST --insecure -H "Authorization: Bearer glsa_2lJAFLqG9KWOwA2lfpybWL5vikEqnwnc_ece2bad7" -H "Content-Type: application/json" -d @docker-quest-prometheus.json http://localhost:3000/api/dashboards/db


curl --insecure  -H "Authorization: Bearer glsa_2lJAFLqG9KWOwA2lfpybWL5vikEqnwnc_ece2bad7"  http://localhost:3000/api/datasources 


curl -X POST --insecure -H"Authorization: Bearer glsa_2lJAFLqG9KWOwA2lfpybWL5vikEqnwnc_ece2bad7" -d "{\"datasource\":$(cat data_source.json)}"  http://localhost:3000/api/datasources


