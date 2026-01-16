# Unisphere OTEL Provider

Unisphere OTEL Provider connects to the Unity console, scrape metric and logs.
and send them to Backends like Prometheus, Loki...

## Collectors
| Collector       | type     | Default Enabled | Description |
|-----------------|----------|-----------------|-------------|
| alert           | `log`    |             | |
| basicSystemInfo | `metric` |             | |
| capacity        | `metric` |             | |
| disk            | `metric` |             | |
| dpe             | `metric` |             | |
| ethernetPort    | `metric` |             | |
| event           | `log`    |             | |
| host            | `metric` |             | |
| lun             | `metric` |             | |



### Basic System Info
| Metric Name | unisphere_basic_system_info        |
|-------------|------------------------------------|
| Description | Information about unisphere system |
| Unit        | -                                  |
| Type        | gauge                              |
| Labels      | `model` `firmware.version`         |

----------------

### Capacity
| Metric Name | unisphere_capacity_total_capacity    |
|-------------|--------------------------------------|
| Description | Total capacity of unisphere capacity |
| Unit        | mb                                   |
| Type        | gauge                                |
| Labels      | -                                    |

| Metric Name | unisphere_capacity_used_capacity    |
|-------------|-------------------------------------|
| Description | Used capacity of unisphere capacity |
| Unit        | mb                                  |
| Type        | gauge                               |
| Labels      | -                                   |

| Metric Name | unisphere_capacity_free_capacity    |
|-------------|-------------------------------------|
| Description | Free capacity of unisphere capacity |
| Unit        | mb                                  |
| Type        | gauge                               |
| Labels      | -                                   |

| Metric Name | unisphere_capacity_preallocated_capacity     |
|-------------|----------------------------------------------|
| Description | pre-allocated capacity of unisphere capacity |
| Unit        | mb                                           |
| Type        | gauge                                        |
| Labels      | -                                            |

| Metric Name | unisphere_capacity_total_provision               |
|-------------|--------------------------------------------------|
| Description | Total provisioned capacity of unisphere capacity |
| Unit        | mb                                               |
| Type        | gauge                                            |
| Labels      | -                                                |

## Build
### Linux
1. Install golang on system


### Windows


### AIX
