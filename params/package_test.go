// Copyright 2014 Canonical Ltd.

package params_test

import (
	"testing"

	jujutesting "github.com/juju/testing"
)

func TestPackage(t *testing.T) {
	jujutesting.MgoTestPackage(t, nil)
}
