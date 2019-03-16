### ninefs

`ninefs` is a multi-threaded, user space file server that speaks 9P2000.u protocol.

### Quick Start

Start the `ninefs` server, with protocol debugging to stderr.

```
sudo ninefs -d -p 5640
```

or

```
docker run --name=ninefs --rm --publish=0.0.0.0:5640:564 ninefs:latest
```

Mount it using the 9P kernel client, sometimes referred to as "v9fs".

```
sudo mount -t 9p -n 127.0.0.1 /mnt \
	-oaname=/tmp,version=9p2000.u,uname=root,access=user,port=5640
```
