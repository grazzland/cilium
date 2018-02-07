<!-- This file was autogenerated via cilium cmdref, do not edit manually-->

## cilium kvstore set

Set a key and value

### Synopsis


Set a key and value

```
cilium kvstore set [options] <key>
```

### Examples

```
cilium kvstore set foo=bar
```

### Options

```
      --key string     Key
      --value string   Value
```

### Options inherited from parent commands

```
      --config string     config file (default is $HOME/.cilium.yaml)
  -D, --debug             Enable debug messages
  -H, --host string       URI to server-side API
      --kvstore string    kvstore type
      --kvstore-opt map   kvstore options (default map[])
```

### SEE ALSO
* [cilium kvstore](cilium_kvstore.html)	 - Direct access to the kvstore
