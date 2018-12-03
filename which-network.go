package main

import (
  "fmt"
  "net"
  "os"
  "strings"
)

type Network struct {
  block *net.IPNet
  key string
}

type Candidate struct {
  addr net.IP
  key string
}

type Match struct {
  network Network
  candidate Candidate
}

func ParseCandidate(candidate string) []Candidate {
  // We can pretty safely default to capacity of 1, since
  // even an interface name is likely to only have 1 address.
  candidates := make([]Candidate, 0, 1)
  ip := net.ParseIP(candidate)
  if ip != nil {
    candidates = append(candidates, Candidate{ip, candidate})
  } else {
    iface, err := net.InterfaceByName(candidate)
    if err == nil {
      addrs, err := iface.Addrs()
      if err == nil {
        for _, addr := range addrs {
          ip, _, _ := net.ParseCIDR(addr.String())
          candidates = append(candidates, Candidate{ip, candidate})
        }
      }
    }
  }
  return candidates
}

func ParseNetworkSpec(spec string) (n *Network, err error) {
  parts := strings.SplitN(spec, "=", 2)
  _, network, err := net.ParseCIDR(parts[0])
  if err == nil {
    n = &Network{network, parts[1]}
  } else {
    n = nil
  }
  return n, err
}

func IsNetworkSpec(spec string) bool {
  return strings.Contains(spec, "=")
}

func ParseInputs(args []string) (nets []Network, candidates []Candidate) {
  nets = make([]Network, 0, len(args))
  candidates = make([]Candidate, 0, len(args))
  for _, arg := range args {
    if IsNetworkSpec(arg) {
      net, err := ParseNetworkSpec(arg)
      if err == nil {
        nets = append(nets, *net)
      }
    } else {
      candidates = append(candidates, ParseCandidate(arg)...)
    }
  }
  return nets, candidates
}

func ExtractMatches(nets []Network, candidates []Candidate) []Match {
  matches := make([]Match, 0, len(nets) * len(candidates))
  for _, network := range nets {
    for _, candidate := range candidates {
      if network.block.Contains(candidate.addr) {
        matches = append(matches, Match{network, candidate})
      }
    }
  }
  return matches
}

func main() {
  nets, candidates := ParseInputs(os.Args[1:])
  for _, match := range ExtractMatches(nets, candidates) {
    fmt.Printf("%s\t%s\t%s\n", match.candidate.key, match.candidate.addr.String(), match.network.key)
  }
}

