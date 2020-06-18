// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package datastores_test

import (
	"os"
	"testing"

	"github.com/cuttle-ai/brain/appctx"
	"github.com/cuttle-ai/brain/env"
	"github.com/cuttle-ai/go-sdk/services/datastores"
)

func TestListDatastores(t *testing.T) {
	appToken := os.Getenv("APP_TOKEN")
	discoveryURL := os.Getenv("DISCOVERY_URL")
	discoveryToken := os.Getenv("DISCOVERY_TOKEN")
	appCtx := appctx.NewAppCtx(appToken, discoveryToken, discoveryURL)
	env.LoadEnv(appCtx.Logger())
	_, err := datastores.ListDatastores(appCtx)
	if err != nil {
		t.Error("error while getting the list of datastores", err)
	}
}
