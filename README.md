## DRKChain

Official `Go` implementation of the **DRKChain** protocol.

## Building the source

For prerequisites and detailed build instructions please read the `Installation Instructions` below.

### Installation Instructions for Linux/Unix
(You can use Linux/Unix servers of DigitalOcean, AWS, etc.)

Update package manager and install `git`

```
sudo apt-get update && sudo apt-get -y upgrade
sudo apt-get install git
```

Clone the repository to a directory of your choosing:

```
git clone https://github.com/DRKtechnology/DRKChain.git
```

Install latest distribution of [Go](https://golang.org/) if you don't have it already. You can refer to below document guide to install Go on Ubuntu 18.04

https://github.com/golang/go/wiki/Ubuntu

Building `drkchain` requires both a Go (version 1.10 or later) and a C compiler. You can install
them using your favourite package manager.

```
sudo apt-get install -y build-essential
```

Once the dependencies are installed. Finally, build the `drkchain` program using the following command.

```shell
cd DRKChain
make drkchain
```

or, to build the full suite of utilities:

```shell
make all
```

You can now run `build/bin/drkchain` to start your node.

## Running `drkchain`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `drkchain` instance.

### Validator node on the DRKChain network

Create an drkchain account to sign the block for the validator

```
./build/bin/drkchain account new
```

It will promtly you to input password for the new account. And take node that `Your new account is locked with a password. Please give a password. Do not forget this password.`. For example, we input `123456` as password.

After you input matched password and repeat password of `123456`, a new account was generated something like following

```
Your new key was generated

Public address of the key:   e94d5a92aeacaff99fe1ab6988ddb3e66c9fd4eb
Path of the secret key file: /root/.drkchain/keystore/UTC--2019-11-05T15-46-35.685344950Z--e94d5a92aeacaff99fe1ab6988ddb3e66c9fd4eb

- You can share your public address with anyone. Others need it to interact with you.
- You must NEVER share the secret key with anyone! The key controls access to your funds!
- You must BACKUP your key file! Without the key, it's impossible to access account funds!
- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!
```

To unlock account in the cli for signing, we will firstly write out the account password above into a file named `passwd` by running `echo 123456 > passwd` then passing the `passwd` file name into `--password` option to run validator node as below.

Run node as validator with `--mine` cli option as below, we also need to unlock the etherbase account with password. For instance, we run validator node with account `e94d5a92aeacaff99fe1ab6988ddb3e66c9fd4eb`

```
nohup ./build/bin/drkchain --mine --unlock e94d5a92aeacaff99fe1ab6988ddb3e66c9fd4eb --password passwd --etherbase e94d5a92aeacaff99fe1ab6988ddb3e66c9fd4eb
```

Then come to `drkchain` **validator dapp** to deposit and register validator with coinbase of `e94d5a92aeacaff99fe1ab6988ddb3e66c9fd4eb` that we unlock on above.

### Full node on the DRKChain network

By far the most common scenario is people wanting to simply interact with the DRKChain
network: create accounts; transfer funds; deploy and interact with contracts. For this
particular use-case the user doesn't care about years-old historical data, so we can
fast-sync quickly to the current state of the network. To do so:

```shell
$ drkchain console
```

This command will:
 * Start `drkchain` in fast sync mode (default, can be changed with the `--syncmode` flag),
   causing it to download more data in exchange for avoiding processing the entire history
   of the DRKChain network, which is very CPU intensive.
 * Start up `drkchain`'s built-in interactive `JavaScript console`,
   (via the trailing `console` subcommand) through which you can invoke all official `web3` methods
   as well as `drkchain`'s own `management APIs`.
   This tool is optional and if you leave it out you can always attach to an already running
   `drkchain` instance with `drkchain attach`.

### Programmatically interfacing `drkchain` nodes

As a developer, sooner rather than later you'll want to start interacting with `drkchain` and the
DRKChain network via your own programs and not manually through the console. To aid
this, `drkchain` has built-in support for a JSON-RPC based APIs.
These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `drkchain`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

  * `--rpc` Enable the HTTP-RPC server
  * `--rpcaddr` HTTP-RPC server listening interface (default: `localhost`)
  * `--rpcport` HTTP-RPC server listening port (default: `8545`)
  * `--rpcapi` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
  * `--rpccorsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--wsaddr` WS-RPC server listening interface (default: `localhost`)
  * `--wsport` WS-RPC server listening port (default: `8546`)
  * `--wsapi` API's offered over the WS-RPC interface (default: `eth,net,web3`)
  * `--wsorigins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: `admin,debug,eth,miner,net,personal,shh,txpool,web3`)
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `drkchain` node configured with the above flags and you'll
need to speak `JSON-RPC` on all transports. You can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
drkchain nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!**
