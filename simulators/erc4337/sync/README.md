# ERC4337 in Hive

This is *experimental*.

## Configuration

Bundler configuration is provided via environment vars, for the clients to build configs in their format of preference.
The configuration (incl. formatting of values) matches the [Bundler p2p spec](https://github.com/eth-infinitism/bundler-spec/blob/main/p2p-specs/p2p-interface.md)).

**Each configuration var is prefixed with `HIVE_ERC4337_CONFIG_`**.

And configuration variable is runtime-configurable, i.e. client docker images can adopt the Hive config without rebuild. 

## Container preparation

Note `{}` is used for variable substitution, not part of content.

### Bundler Client

#### Files

A bundler configuration file
```
/hive/input/config.yaml
```


#### Env vars

Every standard bundler config var prefixed with `HIVE_ERC4337_CONFIG_`, and additionally:

```yaml
# Comma separated list of Eth1 nodes to communicate with.
# Clients should strip off everything after first comma if they do not load-balancing between them.
# If it is left empty, the beacon node should use a "dummy eth1" mode, where it fills Eth1 votes with mock data.
HIVE_ERC4337_ETH1_RPC_ADDRS: ""
```
