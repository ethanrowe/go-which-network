# Which Network?

This amazing go utility does an astounding thing:
* Given an argument list with at least one of a "CIDR-block=key" pair and an
  "IPaddr/interface-name/hostname" (can have multiple of either or both)...
* Outputs a tab-delimited sequence of the CIDR-block keys to which
  IPaddr/interface/hostname matches.

So, like, I want to know if I'm on subnet A or subnet B.  Suppose "eth0" is my
main interface (and let's suppose its address is `192.168.1.17`)

```bash
ethan@no-puppet ~$ which-network 192.168.100.0/24=subnet-A 192.168.1.0/24=subnet-B eth0
eth0	192.168.1.17	subnet-B
```

The output gives us tab-delimited columns of:
* The input interface or hostname or IP address
* The IP address used for its matching test
* The key given in association with the network (CIDR block) that matched the
  address.

If a named interface/host has multiple IP addresses, they'll all be considered
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

Here's an example that includes hostnames.  These networks are mostly stupid,
but they illustrate the usage.  (The "nopuppet" name is my local hostname,
named for a moment of soaring oratory during the 2016 US Presidential debates.)

```bash
ethan@nopuppet ~$ which-network \
  127.0.0.0/24=local/X \
  127.0.1.0/24=local/Y \
  172.16.24.0/24=/my/special/vpn \
  0.0.0.0/0=catch-all \
  localhost \
  nopuppet \
  127.0.0.1 \
  127.0.1.1 \
  tun0 google.com \
  amazonaws.com
localhost	127.0.0.1	local/X
127.0.0.1	127.0.0.1	local/X
nopuppet	127.0.1.1	local/Y
127.0.1.1	127.0.1.1	local/Y
tun0	172.16.24.18	/my/special/vpn
localhost	127.0.0.1	catch-all
nopuppet	127.0.1.1	catch-all
127.0.0.1	127.0.0.1	catch-all
127.0.1.1	127.0.1.1	catch-all
tun0	172.16.24.18	catch-all
google.com	172.217.164.110	catch-all
amazonaws.com	207.171.166.22	catch-all
amazonaws.com	72.21.206.80	catch-all
amazonaws.com	72.21.210.29	catch-all
```

Building for Linux
==================

Assumes you've got docker handy.

```
./build
```

