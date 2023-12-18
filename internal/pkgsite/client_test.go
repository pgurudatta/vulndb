// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgsite

import (
	"context"
	"flag"
	"testing"
)

var usePkgsite = flag.Bool("pkgsite", false, "use pkg.go.dev for tests")

func TestKnown(t *testing.T) {
	ctx := context.Background()

	for _, test := range []struct {
		name string
		in   string
		want bool
	}{
		{name: "valid", in: "golang.org/x/mod", want: true},
		{name: "invalid", in: "github.com/something/something", want: false},
	} {
		t.Run(test.name, func(t *testing.T) {
			cf, err := CacheFile(t)
			if err != nil {
				t.Fatal(err)
			}
			pc, err := TestClient(t, *usePkgsite, cf)
			if err != nil {
				t.Fatal(err)
			}

			got, err := pc.Known(ctx, test.in)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Errorf("%s: got %t, want %t", test.in, got, test.want)
			}
		})
	}
}
