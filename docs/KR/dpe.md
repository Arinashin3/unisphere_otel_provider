# DPE Collector
Disk Processor Enclosures의 상태와 최근 온도를 수집합니다.  
`metric`, `log`, `trace` 중 `metric`으로만 나타납니다.

## Metric Info
| Metric 명                          | 값                                                                                                                        | 유닛 | 설명         |
|-----------------------------------|--------------------------------------------------------------------------------------------------------------------------|----|------------|
| unisphere_dpe_health              | 0: Unknown<br/>5: OK<br/>7: OK_BUT<br/>10: DEGRADED<br/>15: MINOR<br/>20: MAJOR<br/>25: CRITICAL<br/>30: NON_RECOVERABLE | -  | DPE의 현재 상태 |
| unisphere_dpe_current_temperature | -                                                                                                                        | -  | DPE의 현재 온도 |

## Sample) Alert Rule
```yaml
```

## Sample) Dashboard