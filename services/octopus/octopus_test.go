// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package octopus_test

import (
	"os"
	"testing"

	"github.com/cuttle-ai/brain/env"
	"github.com/cuttle-ai/brain/log"
	"github.com/cuttle-ai/go-sdk/services/octopus"
)

func TestRemoveDict(t *testing.T) {
	l := log.NewLogger()
	env.LoadEnv(l)
	appToken := os.Getenv("APP_TOKEN")
	discoveryURL := os.Getenv("DISCOVERY_URL")
	discoveryToken := os.Getenv("DISCOVERY_TOKEN")
	err := octopus.RemoveDict(l, discoveryURL, discoveryToken, appToken)
	if err != nil {
		t.Error("error while removing the dict from octopus service", err)
	}
}
