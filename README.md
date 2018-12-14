# healthy [![Build Status](https://travis-ci.org/localghost/healthy.svg?branch=master)](https://travis-ci.org/localghost/healthy)

_healthy_ is a service that via HTTP REST API exposes current status of configured health checks.

The main use case for this service is to have single endpoint for checking current status of an entire infrastructure.
For example, in CI, after deploying or modifying your infrastructure the test can easily check whether the infrastructure
is fully operational before executing its scenario.
The same in product, having a single check is very convenient for the Ops team to quickly verify product status.

## configuration

Configuration can be either JSON or YAML or actually anything that is supported by the [spf13/viper](https://github.com/spf13/viper) library.
Configuration options can also be changed via environment variables, the pattern is `HEALTHY_OPTION_PATH`, e.g.:
```
HEALTHY_SERVER_LISTEN_ON=localhost:8888 ./healthy --config ./healthy.yml 
```

There are two main configuration sections: `server` and `checks`, for details please see [here](https://github.com/localghost/healthy/wiki/Configuration-schema).

### example

```
server:
  listen_on: 127.0.0.1:8199

checks:
  google:
    type: http
    url: http://google.com
  rabbit:
    type: dial
    address: 127.0.0.1:5672
  echo:
    command: echo "hello world!"
```

Other examples can be found in `examples/`.

## docker images   

_healthy_ is also released as docker image available for example as `zkostrzewa/healthy:0.1.0`.

Until [#20](https://github.com/localghost/healthy/issues/20) is fixed docker image requires mounting custom configuration file to `/etc/healthy/healthy.yml` or `/healthy.yml`.
