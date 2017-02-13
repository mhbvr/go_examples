# go_examples
Base examples of using golang for system programming and administration tasks. Implementation of base versions of good known tools with go.

## Simple tools

* *ls* - list of files and directories. Add special option for recursive calculation of actual directory size (sum of content sizes).
* *find* tool with limited functionality. Find by name regexp, modification time, size, MIME type.
* *ps* tool with threads support.
* *lsof* like tool to find open files and processes who used it.
* Log of starting/stopping of all procceses in system.
* *vmtouch* like tool to find what part of file(s) are placed in page cache. Is it posible to know what part in active list and what part in inactive list?

## Example apps

* *udp/tcp ping pong*. Simple network benchmark with multicast support.
* *rdma ping pong*. Is it possible to work with Infiniband/RoCE with Go? The one possibility is https://github.com/jsgilmore/ib
* ebpf with golang
* RSS reader. XML parsing.
