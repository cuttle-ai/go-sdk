// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package websockets has the sdk to interact with the websockets services
package websockets

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/cuttle-ai/brain/appctx"
	"github.com/cuttle-ai/brain/models"
	"github.com/cuttle-ai/db-toolkit/datastores/services"
	"github.com/cuttle-ai/go-sdk/discovery"
	"github.com/cuttle-ai/go-sdk/httpclient"
	"github.com/hashicorp/consul/api"
)

func sendNotification(appCtx appctx.AppContext, n models.Notification) error {
	/*
	 * First we will create the discovery config
	 * Then get the websockets servers from discovery service
	 * Then will try to send notification to websockets services from each of them (whichever delivers first)
	 */
	//creating the discovery config
	dConfig := api.DefaultConfig()
	dConfig.Address = appCtx.DiscoveryAddress()
	dConfig.Token = appCtx.DiscoveryToken()
	l := appCtx.Logger()

	//getting the web sockets servers
	svs, err := discovery.GetServices(dConfig, "Brain-Websockets-Server", l)
	if err != nil {
		//error while getting the services from the discovery
		l.Error("error while getting the list of Brain-Websockets-Server from discovery service")
		return err
	}

	//trying to send to any of the web sockets server
	payload, err := json.Marshal(n)
	if err != nil {
		//error while encoding the notification payload
		l.Error("error while encoding the notification payload")
		return err
	}
	for _, v := range svs {
		targetURL := "http://" + v.Address + ":" + strconv.Itoa(v.Port) + "/notifications/send"
		l.Info("going to send notification to websockets server at", targetURL)
		res, err := httpclient.Post(v.Address, targetURL, appCtx.AccessToken(), "auth-token", bytes.NewBuffer(payload))
		if err != nil {
			//error while sending notification to websockets server
			l.Error("error while sending notification to websockets server at", targetURL, err)
			continue
		}
		defer res.Body.Close()

		//read the response
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			//error while making the reading the response from the api
			l.Error("error while reading the response from websockets server at", targetURL, err)
			continue
		}

		//parsing the response
		p := struct {
			Message string
			Data    []services.Service
		}{}
		err = json.Unmarshal(body, &p)
		if err != nil { //error while making the parsing the response from the api
			l.Error("error while parsing the response from websockets server at", targetURL, err)
			continue
		}

		//got the response
		l.Info("got the response message from the websockets server", p.Message)
		break
	}
	return nil
}

//SendInfoNotification will send a info notification to the user's websocket clients
func SendInfoNotification(appCtx appctx.AppContext, n models.Notification) error {
	n.Event = models.InfoNotification
	return sendNotification(appCtx, n)
}
