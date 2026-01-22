# open-telemetry
simple grafana logs using Go

<img width="345" height="471" alt="image" src="https://github.com/user-attachments/assets/2bc90b71-1cd0-448e-94de-7106f0c4ae21" />



1st things 1st :
A.
- put file docker folder to C:
- PS from C:\docker
- docker compose up -d
- docker should be run
 <img width="1459" height="252" alt="image" src="https://github.com/user-attachments/assets/9662dc23-6394-43ea-a744-39f03c58a222" />
<img width="1448" height="122" alt="image" src="https://github.com/user-attachments/assets/1a1ea0f2-3626-44e5-91e9-1c9b6e9d7f46" />

- run main.go
- go to http://localhost:3000/
- find Connections -> Data Sources
- Add new data source
  <img width="1498" height="300" alt="image" src="https://github.com/user-attachments/assets/13632cde-049b-40f8-87ea-3426fe0dc075" />
loki : http://loki:3100
prometheus : http://host.docker.internal:9090
tempo : http://host.docker.internal:3200
- environment created

B.
- trigger curl api
    curl http://localhost:8080/user/list
    curl http://localhost:8080/order/list
- go to http://localhost:3000/
- find Logs
- or go to http://localhost:3000/, find Data Sources
- find Loki, click explore
- on Label filters put container
- pick value /grafana
- logs will appear

notes :
Loki query example : 
{container="otel-demo-app"} | json | msg="api_called"

{container="otel-demo-app"} | json | trace_id="abcd1234"

{container="/grafana"} |= `f8f765c0574462259b4cc1573defdfff` |= `/user/list` |= `statusCode=200`
