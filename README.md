# Rivine Chain Template

A template repository for all your Go-Rivine-based blockchains.

**TODO**:

- replace all `<UNDEFINED>`
- replace the `UNDEFNED_CLIENT_NAME` dir
- replace the `UNDEFINED_DAEMON_NAME` dir

## devnet

You can run the chain easily on your computer to play with the software already.

First step is to launch a daemon from your console:

```
<cli client name> --network devnet --no-bootstrap -Mgctwbe
```

the above launches a <chain name> daemon on devnet, using no bootstrap
(meaning it doesn't try to connect to bootstrap nodes or wait for such nodes to know if you're sync or not),
enabling also the explorer module.

Once you have that you can recover the genesis wallet so you can start creating blocks and have money to spend:

```
<cli client name> wallet recover --plain \
    --seed "carbon boss inject cover mountain fetch fiber fit tornado cloth wing dinosaur proof joy intact fabric thumb rebel borrow poet chair network expire else"
```

As this wallet is recovered as a plain wallet it does not have to be unlocked and is ready for use:

```
$ <cli client name> wallet
Wallet status:
Encrypted, Unlocked
Confirmed Balance:   100006530 GFT
Locked Balance:      0 GFT
Unconfirmed Delta:   + 0 GFT
BlockStakes:         3000 BS
```

Please consult the `--help` menus of the `<cli client name>` command and all its subcommands for more information on how to use the CLI.

### Using multiple wallets on the same machine

A single `<cli daemon name>` daemon doesn't allow multiple wallets for the time being.
In order to have multiple wallets running on the same machine you therefore need
to run multiple `<cli daemon name>` daemons, with each daemon:
  - using a unique persistent directory (either by starting each daemon from a different directory or
    by explicitly setting it using the `--persistent-dir` flag);
  - exposing itself using a unique port.
These different can manually be connected to one another using the `goldchainc gateway connect localhost:[port]` command.

