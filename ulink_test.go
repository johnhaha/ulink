package ulink_test

import (
	"testing"

	"github.com/johnhaha/ulink"
)

func TestGenerate(t *testing.T) {
	res := ulink.NewSimpleUlink("xxx", "yyy.zzz")
	if res.Applinks.Details[0].AppID != "xxx.yyy.zzz" {
		t.Fatal(res.Applinks.Details[0].AppID)
	}
}
