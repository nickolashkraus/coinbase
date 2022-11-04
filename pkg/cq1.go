package pkg

// Given an array of strings ("<clicks>,<subdomain>.<subdomain>.<tld>"),
// find the total number of clicks for each domain (e.g. <subdomain>,
// <subdomain>.<tld>, etc.).

import (
	"strconv"
	"strings"
)

// This problem was a warm-up and required simple string manipulation.
func getClicksByDomain(arr []string) map[string]int {
	ret := make(map[string]int)
	for _, elem := range arr {
		// Get substrings (ex. 900,google.com -> ["900", "google.com"])
		ss := strings.Split(elem, ",")
		clicks, _ := strconv.Atoi(ss[0])
		domain := ss[1]
		// Get domains (ex. google.com -> ["google", "com"])
		domains := strings.Split(domain, ".")
		// For each domain in domains (ex. ["google", "com"]), add the subdomain
		// and all variations (ex. "google.com") to the hash map with its
		// corresponding clicks.
		for i := 0; i < len(domains); i++ {
			d := strings.Join(domains[i:], ".")
			_, ok := ret[d]
			if !ok {
				ret[d] = clicks
			} else {
				ret[d] = ret[d] + clicks
			}
		}
	}
	return ret
}
