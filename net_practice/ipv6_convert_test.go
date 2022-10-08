package net_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpv6Convert(t *testing.T) {
	ip := "2605:340:cdb1:123:7a17:0004:172f:55dd"
	assert.Equal(t, "340-pcdb1-123-7a17-4-172f-55dd", Ipv6ToHostName(ip), "The two words should be the same.")
}
