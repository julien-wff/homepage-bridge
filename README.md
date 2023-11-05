# Homepage bridge

This is a custom request handler for my [Homepage](https://gethomepage.dev) dashboard

It currently supports:

- [Minio](https://min.io) through [prometheus](https://min.io/docs/minio/linux/operations/monitoring/collect-minio-metrics-using-prometheus.html)

## Environment variables

No default values are provided. If a required variable is not set, the application will panic.

| Variable         | Description                   | Required | Example                  |
|------------------|-------------------------------|----------|--------------------------|
| `MINIO_ENDPOINT` | Minio prometheus endpoint URL | x        | `http://prometheus:9090` |
| `MINIO_USERNAME` | Minio prometheus username     |          | `minio`                  |
| `MINIO_PASSWORD` | Minio prometheus password     |          | `minio`                  |
