// This package provides immutable UUID structs and the functions
// NewV3, NewV4, NewV5 and Parse() for generating versions 3, 4
// and 5 UUIDs as specified in RFC 4122.
//
// Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>
// Copyright (C) 2012-2019 by Philip Schlump - MIT License

package uuid

import "regexp"

// -------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------
const hexUUIDPattern = "^([a-z0-9]{8})-([a-z0-9]{4})-([1-5][a-z0-9]{3})-([a-z0-9]{4})-([a-z0-9]{12})$"

var uuidPattern *regexp.Regexp

func init() {
	uuidPattern = regexp.MustCompile(hexUUIDPattern)
}

func IsUUID(s string) bool {
	md := uuidPattern.FindStringSubmatch(s)
	if md == nil {
		return false
	}
	return true
}
