# drone-netapp-snapshot

[![Build Status](https://drone.seattleslow.com/api/badges/josmo/drone-netapp-snapshot/status.svg)](https://drone.seattleslow.com/josmo/drone-netapp-snapshot)

Drone plugin to snapshot a datastore before deployment. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build
```


## Docker

Build the container using `make`:

```
make deps docker
```

### Example

## Usage

Build and deploy from your current working directory:

```
docker run --rm                       \
  -e PLUGIN_URL=<source>              \
  -e PLUGIN_USERNAME=<user_name>      \
  -e PLUGIN_USER_NAME=<user_password> \
  -e PLUGIN_STORAGE_VM_KEY=<vm_key>   \
  -e PLUGIN_VOLUME_KEY=<volume_key>   \
  -e PLUGIN_NAME_PREFIX=<name_prefix>
  -v $(pwd):$(pwd)                    \
  -w $(pwd)                           \
  peloton/drone-netapp-snapshot
```
