# Rootchain helper command

Top level command for manipulating rootchain server.

## Start rootchain server

This command starts `ethereum/client-go` container which is basically geth node.

```bash
$ metachain rootchain server
```

## Fund initialized accounts

This command funds the initialized accounts via `metachain polybft-secrets` command.

```bash
$ metachain rootchain fund --data-dir data-dir- --num 2
```
Or
```bash
$ metachain rootchain fund --data-dir data-dir-1
```

## Deploy and initialize contracts

This command deploys and initializes rootchain contracts. Transactions are being sent to given `--json-rpc` endpoint and are signed by private key provided by `--adminKey` flag.

```bash
$ metachain rootchain init-contracts 
    --manifest <manifest_file_path> 
    --deployer-key <hex_encoded_rootchain_deployer_private_key>
    --json-rpc <json_rpc_endpoint> 
```

**Note:** In case `test` flag is provided, it engages test mode, which uses predefined test account private key to send transactions to the rootchain.