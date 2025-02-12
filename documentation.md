[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# phpfpm

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### phpfpm.accepted_connections

The number of requests accepted by the pool

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### phpfpm.active_processes

The number of active processes

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

### phpfpm.idle_processes

The The number of idle processes

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

### phpfpm.listen_queue

The number of requests in the queue of pending connections

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

### phpfpm.listen_queue_length

The size of the socket queue of pending connections

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

### phpfpm.max_active_processes

The maximum number of requests in the queue of pending connections since FPM has started

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### phpfpm.max_children_reached

The number of times, the process limit has been reached, when pm tries to start more children (works only for pm 'dynamic' and 'ondemand')

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### phpfpm.max_listen_queue

The maximum number of requests in the queue of pending connections since FPM has started

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### phpfpm.process.last_request_cpu

The %cpu the last request consumed

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| % | Gauge | Double |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| child | process id of the worker process | Any Int |

### phpfpm.process.last_request_memory

The max amount of memory the last request consumed

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| By | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| child | process id of the worker process | Any Int |

### phpfpm.process.request_duration

The duration in microseconds of the requests

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| us | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| child | process id of the worker process | Any Int |

### phpfpm.process.requests

The number of requests the process has served

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| child | process id of the worker process | Any Int |

### phpfpm.process.state

The state of the process (Idle, Running, ...)

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| child | process id of the worker process | Any Int |
| state | state of the process | Any Str |

### phpfpm.scrape_failures

The number of failures scraping from PHP-FPM

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### phpfpm.slow_requests

The number of requests that exceeded your 'request_slowlog_timeout' value

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| 1 | Sum | Int | Cumulative | true |

### phpfpm.start_since

The number of seconds since FPM has started

| Unit | Metric Type | Value Type | Aggregation Temporality | Monotonic |
| ---- | ----------- | ---------- | ----------------------- | --------- |
| s | Sum | Int | Cumulative | true |

### phpfpm.total_processes

The number of idle + active processes

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Int |

## Resource Attributes

| Name | Description | Values | Enabled |
| ---- | ----------- | ------ | ------- |
| phpfpm.pool_name | pool name | Any Str | true |
| phpfpm.scrape_uri | scrape URI | Any Str | true |
