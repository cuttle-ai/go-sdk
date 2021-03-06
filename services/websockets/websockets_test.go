// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package websockets_test

import (
	"os"
	"testing"

	"github.com/cuttle-ai/brain/appctx"
	"github.com/cuttle-ai/brain/env"
	"github.com/cuttle-ai/brain/log"
	"github.com/cuttle-ai/brain/models"
	"github.com/cuttle-ai/go-sdk/services/websockets"
)

func TestSendInfoNotification(t *testing.T) {
	env.LoadEnv(log.NewLogger())
	appToken := os.Getenv("APP_TOKEN")
	discoveryURL := os.Getenv("DISCOVERY_URL")
	discoveryToken := os.Getenv("DISCOVERY_TOKEN")
	appCtx := appctx.NewAppCtx(appToken, discoveryToken, discoveryURL)
	err := websockets.SendInfoNotification(appCtx, models.Notification{Payload: "hi"})
	if err != nil {
		t.Error("error while sending notification through websockets server", err)
	}
}
