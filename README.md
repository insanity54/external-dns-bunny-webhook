# ExternalDNS - bunny.net provider

[ExternalDNS](https://github.com/kubernetes-sigs/external-dns) synchronises
exposed Kubernetes Services and Ingresses with DNS providers.

This repository contains an ExternalDNS provider for [bunny.net](https://bunny.net).


## Important

This is a fork of <https://github.com/contaimlabs/external-dns-bunny-webhook>.
As of writing, the upstream appears to be unmaintained.

This provider is not officially supported by [bunny.net](https://bunny.net).


## Deployment

An example of deploying the provider with Flux can be seen at
<https://nossa.ee/~talya/vyx/blob/main/-/flux/cassax/external-dns.yaml>.

Configuration options are available below and may be set using environment
variables on the webhook container.


## Configuration

The provider can be configured using the following environment variables:

| Environment Variable    | Required | Description                                                                  | Default     |
|-------------------------|----------|------------------------------------------------------------------------------|-------------|
| `BUNNY_API_KEY`         | Yes      | The API key used to authenticate with the Bunny.net API.                     |             |
| `BUNNY_DRY_RUN`         | No       | If set to `true`, the provider will not make any changes to the DNS records. | `false`     |
| `WEBHOOK_HOST`          | No       | The host to use for the webhook endpoint.                                    | `localhost` |
| `WEBHOOK_PORT`          | No       | The port to use for the webhook endpoint.                                    | `8888`      |
| `WEBHOOK_READ_TIMEOUT`  | No       | The read timeout for the webhook endpoint.                                   | `60s`       |
| `WEBHOOK_WRITE_TIMEOUT` | No       | The write timeout for the webhook endpoint.                                  | `60s`       |
| `HEALTH_HOST`           | No       | The host to use for the health endpoint.                                     | `0.0.0.0`   |
| `HEALTH_PORT`           | No       | The port to use for the health endpoint.                                     | `8080`      |
| `HEALTH_READ_TIMEOUT`   | No       | The read timeout for the health endpoint.                                    | `60s`       |
| `HEALTH_WRITE_TIMEOUT`  | No       | The write timeout for the health endpoint.                                   | `60s`       |


## Provider-Specific Annotations

The following annotations may be added to sources to control behavior of the DNS
records created by this provider:

### `external-dns.alpha.kubernetes.io/webhook-bunny-disabled`

If set to `true`, the DNS record will be managed but set to disabled in the
Bunny API. This annotation is optional and will default to `false` if not
provided. Disabling a record will cause it to not respond to DNS queries, but
will still be managed by the provider and visible in the Bunny.net dashboard.


### `external-dns.alpha.kubernetes.io/webhook-bunny-monitor-type`

The monitor type to use for the DNS record. Valid values are `none` (default),
`http`, and `ping`. This annotation is optional and will default to `none` if
not provided, which will create a standard DNS record without any monitoring.


### `external-dns.alpha.kubernetes.io/webhook-bunny-weight`

The weight to use for the DNS record. Valid values are between 1 and 100. This
annotation is optional and will default to `100` if not provided. Any value
outside of the valid range will be set to the nearest valid value, and any
non-integer value will result in the default value being used.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
