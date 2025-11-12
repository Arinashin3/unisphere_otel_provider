# Disk Collector
Disk의 상태와 용량을 수집합니다.  
`metric`, `log`, `trace` 중 `metric`으로만 나타납니다.

## Metric Info
| Metric 명                 | 값                                                                                                                        | 유닛  | 추가 라벨 | 설명                             |
|--------------------------|--------------------------------------------------------------------------------------------------------------------------|-----|-------|--------------------------------|
| unisphere_disk_health    | 0: Unknown<br/>5: OK<br/>7: OK_BUT<br/>10: DEGRADED<br/>15: MINOR<br/>20: MAJOR<br/>25: CRITICAL<br/>30: NON_RECOVERABLE | -   |       | Disk의 현재 상태                    |
| unisphere_disk_info      | 1: emcPartNumber 감지<br/>0: emcPartNumber 미감지                                                                             | -   |       | Disk의 모델 및 파츠 정보               |
| unisphere_disk_size      | -                                                                                                                        | mb |       | Disk의 사이즈                      |
| unisphere_disk_is_in_use | 1: 사용중<br/>0: 미사용                                                                                                        | - |       | Disk가 유저에 의해 쓰여진 데이터가 존재하는지 여부 |

## Sample) Alert Rule
```yaml
  - name: UnisphereDiskHealthStatus
    labels:
      component: 'health_check'
    rules:
      - alert: 'UnisphereDiskHealthStatusWarning'
        expr: 'unisphere_disk_health == 7'
        for: 1m
        labels:
          severity: 'warning'
        annotations:
          summary: '{{ $labels.host_name }} is ping down.'
```


## Sample) Dashboard