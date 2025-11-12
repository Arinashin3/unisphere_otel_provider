# Unisphere OTEL

unisphere_otel은 OpenTelemetry를 이용하여 Unisphere REST API로 부터 Dell EMC Unity Storage 장비의 성능 및 정보를 수집합니다.  
수집된 데이터는 opentelemetry-collector 또는 OpenTelemetry를 지원하는 백엔드에 직접 전송할 수 있습니다.  
시스템 및 호스트 및 용량 정보의 경우, 각각에 관련된 API에서 데이터를 수집하며,  
성능 정보의 경우, MetricRealTimeQuery API를 통해 최대 48개의 메트릭 데이터를 요청합니다.



## 구조

## 수집 가능 데이터
 - CPU 사용량
 - 버퍼 캐시 히트
 - 시스템 상태(하드웨어, 인터페이스, Lun, 호스트, Disk 등)
 - 서버실 온도
 - Fibre Channel 별 성능 지표(BPS, IOPS)
 - iSCSI 별 성능 지표(BPS, IOPS)
 - 물리 디스크 별 성능 지표(처리 시간, 대기 시간, BPS, IOPS, Queue Langth, 용량)
 - LUN 별 성능 지표(응답시간, BPS, IOPS, 용량)
 - 시스템 정보(firmware version, 모델 명)
 - 호스트 이니시에이터 상태
 - 호스트-LUN 매핑 정보
 - 이벤트 및 Alert 정보(로그)
 - 기타 등등...

## 분석 가능 항목
이 도구를 통해 다음과 같은 중요한 성능에 대한 질문에 답할 수 있습니다.

### 시스템
 - 시스템에서 경고 및 중요 이벤트가 생성되었습니까?
 - 디스크 격납장치의 현재 상태는 어떻습니까?
 - 각 스토리지 프로세서에서 사용 가능한 Cache Clean Page는 얼마나 있습니까?
 - 각 스토리지 프로세서에서 읽기/쓰기 요청이 캐시에서 얼마나 잘 처리되었습니까? (캐시 히트율)

### 디스크
 - 각 디스크의 현재 상태는 어떻습니까?
 - 각 디스크에서 IO 작업이 얼마나 잘 수행되고 있습니까?
 - 각 디스크에서의 초당 처리량이 얼마나 됩니까?

### Fibre Channel
 - 각 Fibre Channel의 현재 상태는 어떻습니까?
 - 각 Fibre Channel에서 초당 IO 요청량은 얼마나 됩니까?
 - 각 Fibre Channel에서 초당 처리량은 얼마나 됩니까?

### iSCSI
- 각 iSCSI의 현재 상태는 어떻습니까?
- 각 iSCSI에서 초당 IO 요청량은 얼마나 됩니까?
- 각 iSCSI에서 초당 처리량은 얼마나 됩니까?

### 호스트
 - 각 호스트의 현재 상태는 어떻습니까?
 - 각 호스트의 이니시에이터 상태는 어떻습니까?
 - 각 호스트와 매핑된 Lun은 무엇입니까?

### LUN
 - 각 LUN의 현재 상태는 어떻습니까?
 - 각 LUN의 용량은 얼마나 됩니까?