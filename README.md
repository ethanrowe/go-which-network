# Which Network?

This amazing go utility does an astounding thing:
* Given an argument list with at least one of a "CIDR-block=key" pair and an
  "IPaddr/interface-name" (can have multiple of either or both)...
* Outputs a tab-delimited sequence of the CIDR-block keys to which IPaddr/interface
  matches.

So, like, I want to know if I'm on subnet A or subnet B.  Suppose "eth0" is my
main interface (and let's suppose its address is `192.168.1.17`)

```bash
ethan@no-puppet ~$ which-network 192.168.100.0/24=subnet-A 192.168.1.0/24=subnet-B eth0
eth0	192.168.1.17	subnet-B
```

The output gives us tab-delimited columns of:
* The input interface or IP address
* The IP address used for its matching test
* The key given in association with the network (CIDR block) that matched the
  address.

If a named interface has multiple IP addresses, they'll all be considered
(independently).

All matches are listed.  The order of input controls the order of evaluation
(which is done network first).

Suppose I had both my `eth0` interface and a VPN tunnel `tun0` with IP address
`192.168.100.4` (subnet A in the earlier example).  And suppose we had a catch-all
network in there as well.

```bash
ethan@no-puppet ~$ which-network 192.168.100.0/24=subnet-A \
  192.168.1.0/24=subnet-B \
  0.0.0.0/0=default \
  eth0 \
  tun0
tun0	192.168.100.4	subnet-A
eth0	192.168.1.17	subnet-B
eth0	192.168.1.17	default
tun0	192.168.100.4	default
```

Building for Linux
==================

Assumes you've got docker handy.

```
./build
```

