// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package octopus has the sdk to interact with the octopus services
package octopus

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/cuttle-ai/brain/log"
	"github.com/cuttle-ai/db-toolkit/datastores/services"
	"github.com/cuttle-ai/go-sdk/discovery"
	"github.com/cuttle-ai/go-sdk/httpclient"
	"github.com/hashicorp/consul/api"
)

//RemoveDict will remove the dict corresponding to a user from the cache
func RemoveDict(l log.Log, discoveryAddress, discoveryToken, appToken string) error {
	/*
	 * First we will create the discovery config
	 * Then get the data integration services from discovery service
	 * Then will try to remove from each of them
	 */
	//creating the discovery config
	dConfig := api.DefaultConfig()
	dConfig.Address = discoveryAddress
	dConfig.Token = discoveryToken

	//getting the octopus services
	svs, err := discovery.GetServices(dConfig, "Brain-Octopus-Service", l)
	if err != nil {
		//error while getting the services from the discovery
		l.Error("error while getting the list of Brain-Octopus-Service from discovery service")
		return err
	}

	//now we will try to remove the dict
	for _, v := range svs {
		targetURL := "http://" + v.Address + ":" + strconv.Itoa(v.Port) + "/dict/remove"
		l.Info("going to remove the dict from", targetURL)
		res, err := httpclient.Get(v.Address, targetURL, appToken, "auth-token")
		if err != nil {
			//error while making the request to remove the dict
			l.Error("error while removing the dict from octopus service at", targetURL, err)
			continue
		}
		defer res.Body.Close()

		//read the response
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			//error while reading the response from the api
			l.Error("error while reading the response of the remove dict from octopus service at", targetURL, err)
			continue
		}

		//parsing the response
		p := struct {
			Message string
			Data    []services.Service
		}{}
		err = json.Unmarshal(body, &p)
		if err != nil {
			//error while the parsing the response from the api
			l.Error("error while parsing the response of the remove dict from octopus service at", targetURL, err)
			continue
		}

		//got the response
		l.Info("got the response message from the doctopus service", p.Message)
	}

	return nil
}

//UpdateDict will update the dict corresponding to a user in cache with updated datasets
func UpdateDict(l log.Log, discoveryAddress, discoveryToken, appToken string) error {
	/*
	 * First we will create the discovery config
	 * Then get the data integration services from discovery service
	 * Then will try to update in each of them
	 */
	//creating the discovery config
	dConfig := api.DefaultConfig()
	dConfig.Address = discoveryAddress
	dConfig.Token = discoveryToken

	//getting the octopus services
	svs, err := discovery.GetServices(dConfig, "Brain-Octopus-Service", l)
	if err != nil {
		//error while getting the services from the discovery
		l.Error("error while getting the list of Brain-Octopus-Service from discovery service")
		return err
	}

	//now we will try to remove the dict
	for _, v := range svs {
		targetURL := "http://" + v.Address + ":" + strconv.Itoa(v.Port) + "/dict/update"
		l.Info("going to remove the dict from", targetURL)
		res, err := httpclient.Get(v.Address, targetURL, appToken, "auth-token")
		if err != nil {
			//error while making the request to update the dict
			l.Error("error while updating the dict from octopus service at", targetURL, err)
			continue
		}
		defer res.Body.Close()

		//read the response
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			//error while reading the response from the api
			l.Error("error while reading the response of the update dict from octopus service at", targetURL, err)
			continue
		}

		//parsing the response
		p := struct {
			Message string
			Data    []services.Service
		}{}
		err = json.Unmarshal(body, &p)
		if err != nil {
			//error while the parsing the response from the api
			l.Error("error while parsing the response of the update dict from octopus service at", targetURL, err)
			continue
		}

		//got the response
		l.Info("got the response message from the doctopus service", p.Message)
	}

	return nil
}
