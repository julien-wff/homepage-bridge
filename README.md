# Homepage bridge

This is a custom request handler for my [Homepage](https://gethomepage.dev) dashboard

It currently supports:

- Minio through prometheus
- Docspell

## Environment variables

| Variable         | Description                   | Required |
|------------------|-------------------------------|----------|
| `MINIO_ENDPOINT` | Minio prometheus endpoint URL | Yes      |
| `MINIO_USERNAME` | Minio prometheus username     | No       |
| `MINIO_PASSWORD` | Minio prometheus password     | No       |
