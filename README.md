# drone-netapp-snapshot

[![Build Status](https://drone.seattleslow.com/api/badges/josmo/drone-netapp-snapshot/status.svg)](https://drone.seattleslow.com/josmo/drone-netapp-snapshot)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)
[![Go Doc](https://godoc.org/github.com/josmo/drone-netapp-snapshot?status.svg)](http://godoc.org/github.com/josmo/drone-netapp-snapshot)
[![Go Report](https://goreportcard.com/badge/github.com/josmo/drone-netapp-snapshot)](https://goreportcard.com/report/github.com/josmo/drone-netapp-snapshot)
[![](https://images.microbadger.com/badges/image/peloton/drone-netapp-snapshot.svg)](https://microbadger.com/images/peloton/drone-netapp-snapshot "Get your own image badge on microbadger.com")

Drone plugin to snapshot a datastore before deployemtn. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `drone cli`:

```
drone exec
```

### Example

## Usage

Build and deploy from your current working directory:

```
docker run --rm                          \
  -e PLUGIN_URL=<source>                 \
  -e PLUGIN_USERNAME=<user_name>     \
  -e PLUGIN_USER_NAME=<user_password>  \
  -e PLUGIN_STORAGE_VM_KEY=<vm_key>            \  
  -e PLUGIN_VOLUME_KEY=<volume_key>         \
  -e PLUGIN_NAME_PREFIX=<name_prefix>
  -v $(pwd):$(pwd)                       \
  -w $(pwd)                              \
  peloton/drone-netapp-snapshot
```
