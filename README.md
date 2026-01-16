# Unisphere OTEL Provider

Unisphere OTEL Provider connects to the Unity console, scrape metric and logs.
and send them to Backends like Prometheus, Loki...


-----------------
## Index

- [Set up](#set-up)
- [Collector List](#collector-list)
- [Metric List](#metric-list)
  - [Basic System Info](#basic-system-info)

---------------
## Set up
### Case 1. Send to Backend directly

unisphere_otel_provider.yml
```yaml
global:
  client:
    insecure: true
    interval: 1m
    
server:
  metrics:
    endpoint: http://<prometheus-address>:9090
    api_path: /api/v1/otlp/v1/metrics
    insecure: true
    enable: true
    
  logs:
    endpoint: http://<loki-address>:3100
    api_path: /otlp/v1/logs
    insecure: true
    enable: true

clients:
  - endpoint: https://<unisphere-address>
    auth: authKey
    insecure: true

collectors:
  basicSystemInfo:
    enabled: true
  systemCapacty:
    enabled: true
  metric:
    enabled: true
    paths:
      - "sp.*.physical.coreCount"
```

prometheus.yml
```yaml
scrape:


```

### Case 2. Use Opentelemetry-Collector Gateway

-------------

## Collector List
| Collector       | type     | Default Enabled | Description                                        |
|-----------------|----------|-----------------|----------------------------------------------------|
| alert           | `log`    |                 | Scrape alerts and send to loki                     |
| basicSystemInfo | `metric` |                 | Scrape system information(model, firmware version) |
| capacity        | `metric` |                 | Scrape system's capacity                           |
| disk            | `metric` |                 | Scrape Disk Health and size                        |
| dpe             | `metric` |                 |                                                    |
| ethernetPort    | `metric` |                 |                                                    |
| event           | `log`    |                 |                                                    |
| host            | `metric` |                 |                                                    |
| lun             | `metric` |                 |                                                    |
| metric          | `metric`  |                 | query metric instant                               |


## Metric List
### Basic System Info
> **unisphere_basic_system_info**  
> Information about unisphere system

| Metric Name | unisphere_basic_system_info        |
|-------------|------------------------------------|
| Description | Information about unisphere system |
| Unit        | -                                  |
| Type        | gauge                              |
| Labels      | `model` `firmware.version`         |
| Value       | 1                                  |



----------------

### Capacity
| Metric Name | unisphere_capacity_total_capacity    |
|-------------|--------------------------------------|
| Description | Total capacity of unisphere capacity |
| Unit        | mb                                   |
| Type        | gauge                                |
| Labels      | -                                    |
| Value       | -                                  |

| Metric Name | unisphere_capacity_used_capacity    |
|-------------|-------------------------------------|
| Description | Used capacity of unisphere capacity |
| Unit        | mb                                  |
| Type        | gauge                               |
| Labels      | -                                   |
| Value       | -                                  |

| Metric Name | unisphere_capacity_free_capacity    |
|-------------|-------------------------------------|
| Description | Free capacity of unisphere capacity |
| Unit        | mb                                  |
| Type        | gauge                               |
| Labels      | -                                   |
| Value       | -                                  |

| Metric Name | unisphere_capacity_preallocated_capacity     |
|-------------|----------------------------------------------|
| Description | pre-allocated capacity of unisphere capacity |
| Unit        | mb                                           |
| Type        | gauge                                        |
| Labels      | -                                            |
| Value       | -                                  |

| Metric Name | unisphere_capacity_total_provision               |
|-------------|--------------------------------------------------|
| Description | Total provisioned capacity of unisphere capacity |
| Unit        | mb                                               |
| Type        | gauge                                            |
| Labels      | -                                                |
| Value       | -                                  |


### Disk
| Metric Name | unisphere_disk_info                              |
|-------------|--------------------------------------------------|
| Description | Information of the associated resource           |
| Unit        | -                                                |
| Type        | gauge                                            |
| Labels      | `disk.id`  `slot.id` `disk.model` `disk.part`    |
| Value       | 1                                                |

| Metric Name | unisphere_disk_health                                                                                                    |
|-------------|--------------------------------------------------------------------------------------------------------------------------|
| Description | Health of the associated resource                                                                                        |
| Unit        | -                                                                                                                        |
| Type        | gauge                                                                                                                    |
| Labels      | `disk.id` `slot.id`                                                                                                      |
| Value       | 0: Unknown<br/>5: OK<br/>7: OK_BUT<br/>10: DEGRADED<br/>15: MINOR<br/>20: MAJOR<br/>25: CRITICAL<br/>30: NON_RECOVERABLE |

| Metric Name | unisphere_disk_size |
|-------------|---------------------|
| Description | Usable capacity     |
| Unit        | mb                  |
| Type        | gauge               |
| Labels      | `disk.id` `slot.id` |
| Value       | -                   |

| Metric Name | unisphere_disk_is_in_use                               |
|-------------|--------------------------------------------------------|
| Description | Indicates whether the drive contains user-written data |
| Unit        | -                                                      |
| Type        | gauge                                                  |
| Labels      | `disk.id` `slot.id`                                    |
| Value       | -                                                      |


### DPE
| Metric Name | unisphere_dpe_health                                                                                                     |
|-------------|--------------------------------------------------------------------------------------------------------------------------|
| Description | health about DPE of system                                                                                               |
| Unit        | -                                                                                                                        |
| Type        | gauge                                                                                                                    |
| Labels      | `dpe.id`                                                                                                                 |
| Value       | 0: Unknown<br/>5: OK<br/>7: OK_BUT<br/>10: DEGRADED<br/>15: MINOR<br/>20: MAJOR<br/>25: CRITICAL<br/>30: NON_RECOVERABLE |

| Metric Name | unisphere_dpe_current_temperature |
|-------------|-----------------------------------|
| Description | current temperature of the DPE    |
| Unit        | -                                 |
| Type        | gauge                             |
| Labels      | `dpe.id`                          |
| Value       | -                                 |

## Build
### Linux
1. Install golang on system


### Windows


### AIX
