A basic set up of Prometheus x Grafana with Docker

Walkthrough of blog post which you can find [https://selfhost.esc.sh/prometheus-grafana/?ref=blog.esc.sh](here)



```sh
mkdir {grafana,prometheus}

docker compose up -d
```

Access to grafana at http://<ip>:3000
Access to prom UI at http://<ip>:9090

Sample prom query

```
{__name__=~"node_disk_write_time_seconds_total|go_threads|promhttp_metric_handler_errors_total"}
```