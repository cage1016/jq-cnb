# JQ Buildpack

![Version](https://img.shields.io/badge/dynamic/json?url=https://cnb-registry-api.herokuapp.com/api/v1/buildpacks/cage1016/jq-cnb&label=Version&query=$.latest.version)

A [Cloud Native Buildpack](https://buildpacks.io) that include jq


## Buildpack registry

https://registry.buildpacks.io/buildpacks/cage1016/jq-cnb

## Usage

you could assign specific jq version by `echo 1.5 > .jq-version` at build time, default version is  1.6

```
pack build myapp --buildpack cage1016/jq-cnb@1.1.0
```

### URI

```
urn:cnb:registry:cage1016/jq-cnb
```

### Supported Stacks

- google
- io.buildpacks.stacks.bionic
- io.buildpacks.samples.stacks.bionic
- heroku-18
- heroku-20