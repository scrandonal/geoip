// Package ipset provides types and utilities for working with IP address sets,
// including both IPv4 and IPv6 CIDR ranges.
package ipset

import (
	"fmt"
	"net"
	"sort"
	"strings"
)

// IPSet represents a named collection of IP networks (CIDR ranges).
type IPSet struct {
	// Name is the identifier for this IP set (e.g., "cn", "us", "private")
	Name string

	// IPv4 holds the list of IPv4 CIDR ranges
	IPv4 []*net.IPNet

	// IPv6 holds the list of IPv6 CIDR ranges
	IPv6 []*net.IPNet
}

// New creates a new empty IPSet with the given name.
func New(name string) *IPSet {
	return &IPSet{
		Name: strings.ToLower(strings.TrimSpace(name)),
		IPv4: make([]*net.IPNet, 0),
		IPv6: make([]*net.IPNet, 0),
	}
}

// AddCIDR parses a CIDR string and adds it to the appropriate (IPv4 or IPv6) list.
// Returns an error if the CIDR string is invalid.
func (s *IPSet) AddCIDR(cidr string) error {
	cidr = strings.TrimSpace(cidr)
	if cidr == "" {
		return fmt.Errorf("empty CIDR string")
	}

	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("invalid CIDR %q: %w", cidr, err)
	}

	if network.IP.To4() != nil {
		s.IPv4 = append(s.IPv4, network)
	} else {
		s.IPv6 = append(s.IPv6, network)
	}

	return nil
}

// AddIPNet adds a pre-parsed *net.IPNet to the set.
func (s *IPSet) AddIPNet(network *net.IPNet) {
	if network == nil {
		return
	}
	if network.IP.To4() != nil {
		s.IPv4 = append(s.IPv4, network)
	} else {
		s.IPv6 = append(s.IPv6, network)
	}
}

// Contains reports whether the given IP address is contained within any
// of the networks in this set.
func (s *IPSet) Contains(ip net.IP) bool {
	if ip.To4() != nil {
		for _, network := range s.IPv4 {
			if network.Contains(ip) {
				return true
			}
		}
	} else {
		for _, network := range s.IPv6 {
			if network.Contains(ip) {
				return true
			}
		}
	}
	return false
}

// Len returns the total number of networks (IPv4 + IPv6) in the set.
func (s *IPSet) Len() int {
	return len(s.IPv4) + len(s.IPv6)
}

// Sort sorts the IPv4 and IPv6 network slices lexicographically by their
// string representation for deterministic output.
func (s *IPSet) Sort() {
	sort.Slice(s.IPv4, func(i, j int) bool {
		return s.IPv4[i].String() < s.IPv4[j].String()
	})
	sort.Slice(s.IPv6, func(i, j int) bool {
		return s.IPv6[i].String() < s.IPv6[j].String()
	})
}

// Merge combines another IPSet's networks into this set.
func (s *IPSet) Merge(other *IPSet) {
	if other == nil {
		return
	}
	s.IPv4 = append(s.IPv4, other.IPv4...)
	s.IPv6 = append(s.IPv6, other.IPv6...)
}

// String returns a human-readable summary of the IPSet.
func (s *IPSet) String() string {
	return fmt.Sprintf("IPSet{name: %q, ipv4: %d, ipv6: %d}", s.Name, len(s.IPv4), len(s.IPv6))
}
