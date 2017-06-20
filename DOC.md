Use the netapp plugin to snapshot volumes in [netapp](http://netapp.com).

The following parameters are used to configure this plugin:

- `url` - url to your rancher server, including protocol and port
- `user_name` - api user name
- `user_password` - api password
- `storage_vm_key` - vm key for the vm to have the snapshot
- `volume_key` - volume to have the snapshot
- `prefix_name` - prefix of the volume name

The following is a sample Rancher configuration in your `.drone.yml` file:

```yaml
backup:
  snapshot:
    image: peloton/drone-netapp-snapshot
    url: https://apiturl/whatever/snapshot
    user_name: 1234567abcdefg
    user_password: abcdefg1234567
    storage_vm_key: afsldkfjljq;lkjg;lkjasdf
    volume_key: asldkfj;alksjdf
    prefix_name: my-prefix
```

if you want to add secrets for the user_name and user_password it's NETAPP_USER_NAME and NETAPP_USER_PASSWORD

