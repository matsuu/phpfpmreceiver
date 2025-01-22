# PHP-FPM Receiver

This receiver connects directly to PHP-FPM and fetch metrics.

## Configuration

### PHP-FPM

You must configure PHP-FPM to expose status information by editing PHP-FPM pool configuration.  Please see
[pm.status_path](https://www.php.net/manual/en/install.fpm.configuration.php#pm.status-path)
for a guide to configuring the PHP-FPM status page.

### Receiver Config

The following settings are required:

- `endpoint` (default: `[tcp://127.0.0.1:9000/status]`): The URLs of the PHP-FPM status endpoint

The following settings are optional:

- `collection_interval` (default = `10s`): This receiver collects metrics on an interval. This value must be a string readable by Golang's [time.ParseDuration](https://pkg.go.dev/time#ParseDuration). Valid time units are `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.

- `initial_delay` (default = `1s`): defines how long this receiver waits before starting.

Example:

```yaml
receivers:
  phpfpm:
    endpoint:
      - "tcp://127.0.0.1:9000/status"
    collection_interval: 10s
```

The full list of settings exposed for this receiver are documented [here](./config.go).
