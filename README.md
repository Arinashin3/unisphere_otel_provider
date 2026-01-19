# Unisphere OTEL Provider

Unisphere OTEL Provider connects to the Unity console, scrape metric and logs.
and send them to Backends like Prometheus, Loki...



## Index

- [Set up](#set-up)
- [Collector List](#collector-list)
- [Metric List](#metric-list)
  - [Basic System Info](#basic-system-info)


## Set up
### Case 1. Send to Backend directly

#### Unisphere OTEL Provider Config 
(unisphere_otel_provider.yml)
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
    auth: <authKey1>
    insecure: true

auths:
  - name: <authKey1>
    username: <unisphere-username>
    password: <unisphere-password>

collectors:
  basicSystemInfo:
    enabled: true
  systemCapacty:
    enabled: true
  metric:
    enabled: true
    paths:
      # CPU
      - "sp.*.physical.coreCount"
      - "sp.*.cpu.summary.busyTicks"
      - "sp.*.cpu.summary.idleTicks"
      - "sp.*.cpu.summary.waitTicks"

      # Memory
      - "sp.*.memory.summary.totalBytes"
      - "sp.*.memory.summary.totalUsedBytes"
      - "sp.*.memory.summary.cachedBytes"
      - "sp.*.memory.summary.freeBytes"
      - "sp.*.memory.summary.buffersBytes"

      # Cache Memory
      - "sp.*.memory.bufferCache.lookups"
      - "sp.*.memory.bufferCache.hits"
      - "sp.*.blockCache.global.summary.cleanPages"
      - "sp.*.blockCache.global.summary.dirtyBytes"
      - "sp.*.blockCache.global.summary.dirtyPages"
      - "sp.*.blockCache.global.summary.flushedBlocks"
      - "sp.*.blockCache.global.summary.flushes"

      # Physical Disk
      - "sp.*.physical.disk.*.readBlocks"
      - "sp.*.physical.disk.*.writeBlocks"
      - "sp.*.physical.disk.*.reads"
      - "sp.*.physical.disk.*.writes"

      # FibreChannel
      - "sp.*.fibreChannel.fePort.*.readBlocks"
      - "sp.*.fibreChannel.fePort.*.writeBlocks"
      - "sp.*.fibreChannel.fePort.*.reads"
      - "sp.*.fibreChannel.fePort.*.writes"

      # Iscsi
      - "sp.*.iscsi.fePort.*.readBlocks"
      - "sp.*.iscsi.fePort.*.writeBlocks"
      - "sp.*.iscsi.fePort.*.reads"
      - "sp.*.iscsi.fePort.*.writes"

      # Ethernet Port
      - "sp.*.net.device.*.bytesIn"
      - "sp.*.net.device.*.bytesOut"
      - "sp.*.net.device.*.pktsIn"
      - "sp.*.net.device.*.pktsOut"
```

#### Prometheus
1. Enable `--web.enable-otlp-receiver` feature in prometheus.
2. In prometheus's config file `prometheus.yml` , set as below
```yaml
otlp:
  promote_all_resource_attributes: true

storage:
  tsdb:
    out_of_order_time_window: 30m
```

### Case 2. Use Opentelemetry-Collector Gateway



## Collector List
| Collector       | type     | Default Enabled | Description                                        |
|-----------------|----------|-----------------|----------------------------------------------------|
| alert           | `log`    |                 | Scrape alerts Log                                  |
| basicSystemInfo | `metric` |                 | Scrape system information(model, firmware version) |
| capacity        | `metric` |                 | Scrape system's capacity                           |
| disk            | `metric` |                 | Scrape Disk Health and size                        |
| dpe             | `metric` |                 | Scrape DPE Health and Temperate                    |
| ethernetPort    | `metric` |                 | Scrape Ethernet Port Health                        |
| event           | `log`    |                 | Scrape Event Log                                   |
| host            | `metric` |                 | Scrape Host's Information and Health               |
| lun             | `metric` |                 | Scrape Lun's Information and Size                  |
| metric          | `metric` |                 | query metric instant (using RealTimeQuery API)     |


## Metric List
### Basic System Info
> Metric Name:: **unisphere_basic_system_info**  
> Description:: Information about unisphere system  
> > Unit:: `N/A`  
> > Type:: `gauge`  
> > Attributes:: `model` `firmware.version`  
> > Value:: 1

----------------

### Capacity
> Metric Name:: **unisphere_capacity_total_capacity**  
> Description:: Total capacity of unisphere capacity  
> > Unit:: `mb`  
> > Type:: `gauge`  
> > Attributes:: `N/A`  
> > Value:: `float64`

> Metric Name:: **unisphere_capacity_used_capacity**  
> Description:: Used capacity of unisphere capacity
> > Unit:: `mb`  
> > Type:: `gauge`  
> > Attributes:: `N/A`  
> > Value:: `float64`

> Metric Name:: **unisphere_capacity_free_capacity**  
> Description:: Free capacity of unisphere capacity
> > Unit:: `mb`  
> > Type:: `gauge`  
> > Attributes:: `N/A`  
> > Value:: `float64`

> Metric Name:: **unisphere_capacity_preallocated_capacity**  
> Description:: pre-allocated capacity of unisphere capacity
> > Unit:: `mb`  
> > Type:: `gauge`  
> > Attributes:: `N/A`  
> > Value:: `float64`

> Metric Name:: **unisphere_capacity_total_provision**  
> Description:: Total provisioned capacity of unisphere capacity
> > Unit:: `mb`  
> > Type:: `gauge`  
> > Attributes:: `N/A`  
> > Value:: `float64`
---

### Disk
Scrape Disk Health and Size  
- API: `/api/types/disk/instances`

#### Configuration Example

```yaml
collector:
  disk:
    enabled: true
```


#### Metric List
> Metric Name:: **unisphere_disk_info**  
> Description:: Information of the associated resource  
> > Unit:: `N/A`  
> > Type:: `gauge`  
> > Attributes:: `disk.id` `slot.id` `disk.model` `disk.part`  
> > Value:: 1

> Metric Name:: **unisphere_disk_health**  
> Description:: Health of the associated resource
> > Unit:: `N/A`  
> > Type:: `gauge`  
> > Attributes:: `disk.id` `slot.id`  
> > Value:: `enum`   
> 0 = UNKNOWN  
> 5 = OK  
> 7 = OK_BUT  
> 10 = DEGRADED  
> 15 = MINOR  
> 20 = MAJOR  
> 25 = CRITICAL  
> 30 = NON-RECOVERABLE

> Metric Name:: **unisphere_disk_size**  
> Description:: Usable capacity  
> > Unit:: `mb`  
> > Type:: `gauge`  
> > Attributes:: `disk.id` `slot.id`  
> > Value:: `float64`


> Metric Name:: **unisphere_disk_is_in_use**  
> Description:: Indicates whether the drive contains user-written data  
> > Unit:: `N/A`  
> > Type:: `gauge`  
> > Attributes:: `disk.id` `slot.id`  
> > Value:: `float64`

> Metric Name::
> Description::
> > Unit::
> > Type::
> > Attributes::
> > Value::
---
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
