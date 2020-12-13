# Zookeeper-Go

A simple Zookeeper in Go

This is a part of the final project of Big Data System Foundation(B) in Autumn, 2020 by Group 2

## Getting Started

See https://github.com/DistAlchemist/Mongongo Getting Started

## Design

* Design mostly follow Zookeeper-V3.0.0

* Design is divided into two part
   * Quorum
   * Znode
   
* For Quorum, the base election algorithm is a raft-like election except ACK. Proposals are broadcasted but not need ACK as well. It can tolerate up to n failure nodes in a total of 2n+1 nodes. Reconnect and synchronization is supported. System reaches high availability and consistency.

* Znode support following operations:

```
   CREATE <dir>  ——create a dir
   DELETE <dir>  ——delete a dir
   DIR <dir>     ——show the content of a dir
   WATCH <dir> <port>   ——set a watcher on a dir, response will be on <port>
```

TODO：

* Send error back to client
* Persistent Znode
* ACK
* Add view and other consensus algorithm

## Example

You need to copy this project several times and change the settings in zoo.cfg and client.cfg to set address and port correctly.

Local-test.zip contains a basic example of this project that has been set correctly

By running:

```shell
cd test/zookeeper-go
go run cmd/server/main.go
```

```shell
cd test/zookeeper-go1
go run cmd/server/main.go
```

```shell
cd test/zookeeper-go2
go run cmd/server/main.go
```

in three terminal, you can see three servers start and begin to elect leader.

By running:

```shell
cd test/zookeeper-go
go run cmd/client/main.go
```

a client can be started and you can type like

```
please input command:
CREATE
please input create dir
ABC
```

to create  ABC under root.