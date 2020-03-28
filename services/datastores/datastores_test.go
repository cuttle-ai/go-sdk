// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package datastores_test

import (
	"os"
	"testing"

	"github.com/cuttle-ai/go-sdk/env"
	"github.com/cuttle-ai/go-sdk/log"
	"github.com/cuttle-ai/go-sdk/services/datastores"
)

func TestListDatastores(t *testing.T) {
	l := log.NewLogger()
	env.LoadEnv(l)
	appToken := os.Getenv("APP_TOKEN")
	discoveryURL := os.Getenv("DISCOVERY_URL")
	discoveryToken := os.Getenv("DISCOVERY_TOKEN")
	_, err := datastores.ListDatastores(l, discoveryURL, discoveryToken, appToken)
	if err != nil {
		t.Error("error while getting the list of datastores", err)
	}
}
