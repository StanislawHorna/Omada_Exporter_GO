# Omada Exporter

Prometheus Exporter written in Go for TP-Link Omada SDN.

This exporter is built with native Go HTTP and [prometheus/client_golang](https://github.com/prometheus/client_golang)
libraries to expose metrics from Omada SDN to Prometheus.
Metrics are queried live from the controller when the `/metrics` endpoint is accessed,
avoiding unnecessary polling when not in use.

Supports both Omada OpenAPI and Web API to accommodate current platform limitations.

## Grafana Dashboards

### [Site Overview](/Dashboards/Site_Overview.json)

![image](/Pictures/Site_Overview.png)

### [Router](/Dashboards/Router.json)

![image](/Pictures/Router_1.png)

### [Switch](/Dashboards/Switch.json)

![image](/Pictures/Switch_1.png)

### [Access Point](/Dashboards/AccessPoint.json)

![image](/Pictures/AccessPoint_1.png)

## 🚀 Getting Started

### Prerequisites

- Omada Controller with configured OpenAPI Client and Service user account, details [here](#omada-authentication-setup)
-

### Setup

1. Clone the Repository

```shell
git clone https://github.com/StanislawHorna/Omada_Exporter_Go.git
cd Omada_Exporter_Go
```

2. Create `.env` File with variables from [Configuration Parameters](#configuration-parameters)

```
LOG_LEVEL=info

OMADA_URL=https://<omada_ip_address>:<omada_port>
OMADA_SITE_NAME=<your_omada_site_name>
OMADA_CLIENT_ID=<open_api_client_id>
OMADA_CLIENT_SECRET=<open_api_client_secret>
OMADA_USERNAME=<web_ui_username>
OMADA_PASSWORD=<web_ui_password>

LOKI_URL=http://<loki_ip_address>:<loki_port>
LOKI_ENV=prod
LOKI_APP_VERSION=1.0.0
```

3. Build the Docker image

```shell
docker build -t omada_exporter .
```

4. Run the container

```shell
docker run -d \
  --env-file .env \
  -p 8080:8080 \
  --name omada_exporter \
  omada_exporter
```

5. Configure Prometheus to scrape the data:

```YAML
scrape_configs:
  - job_name: "Omada"
    static_configs:
      - targets: ["<docker_host_ip>:8080"]
        labels:
          device_name: "<omada_controller_friendly_name>"
    metrics_path: /metrics
```

6. Import dashboards from [`Dashboards`](/Dashboards/) directory to your Grafana instance.

> [!TIP]
> After starting up the exporter navigate to `http://"<docker_host_ip>:8080/metrics`
> and check `omada_http_client_requests_total` metrics.
> There should be **no** metrics with code different different than 200

### Omada Authentication Setup

- OpenAPI Client – Created via: `Settings -> Platform Integration`.
  Assign admin role for full API access.
- Service User – Create under: `Account section` at `Global level`.
  Assign viewer role for read-only access.

### Configuration Parameters

All configuration values are read from environment variables. These can be provided via a .env file, environment variables in your system, or directly in Docker Compose.

| Variable              | Description                                                            | Default |
| --------------------- | ---------------------------------------------------------------------- | ------- |
| `LOG_LEVEL`           | Logging verbosity level (_debug_, _info_, _warn_, _error_)             | `error` |
| `OMADA_URL`           | Full base URL to the Omada controller (e.g., https://192.168.1.1:8043) | -       |
| `OMADA_SITE_NAME`     | Name of the site you wish to monitor                                   | -       |
| `OMADA_CLIENT_ID`     | OpenAPI Client ID created in Omada Platform Integration                | -       |
| `OMADA_CLIENT_SECRET` | OpenAPI Client Secret                                                  | -       |
| `OMADA_USERNAME`      | Username for Web API access                                            | -       |
| `OMADA_PASSWORD`      | Password for Web API access                                            | -       |
