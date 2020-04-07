// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package datastores has the sdk to interact with the datastores services to fetch/add/update/delete datastores in cuttle platform
package datastores

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/cuttle-ai/brain/log"
	"github.com/cuttle-ai/db-toolkit/datastores/services"
	"github.com/cuttle-ai/go-sdk/discovery"
	"github.com/cuttle-ai/go-sdk/httpclient"
	"github.com/hashicorp/consul/api"
	"github.com/jinzhu/gorm"
)

//ListDatastores returns the list of data stores available in the platform
//discoveryAddress is the address of the discovery service of cuttle platform
//discoveryToken is the token to be used to talk to the discovery service in the cuttle platform
//appToken is the authentication token of the user who is logged into the system
func ListDatastores(l log.Log, discoveryAddress, discoveryToken, appToken string) ([]services.Service, error) {
	/*
	 * First we will create the discovery config
	 * Then get the data integration services from discovery service
	 * Then will try to get the datastore list from each of them (whichever delivers first)
	 */
	//creating the discovery config
	dConfig := api.DefaultConfig()
	dConfig.Address = discoveryAddress
	dConfig.Token = discoveryToken

	//getting the data-integration services
	svs, err := discovery.GetServices(dConfig, "Brain-Data-Integeration-Service", l)
	if err != nil {
		//error while getting the services from the discovery
		l.Error("error while getting the list of Brain-Data-Integeration-Service from discovery service")
		return nil, err
	}

	//now we will try to get list of services
	result := []services.Service{}
	for _, v := range svs {
		targetURL := "http://" + v.Address + ":" + strconv.Itoa(v.Port) + "/services/datastore/list"
		l.Info("going to get the list of services from", targetURL)
		res, err := httpclient.Get(v.Address, targetURL, appToken, "auth-token")
		if err != nil {
			//error while making the request to get the list of services
			l.Error("error while getting the list of services from data-store-service at", targetURL, err)
			continue
		}
		defer res.Body.Close()

		//read the response
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			//error while making the reading the response from the api
			l.Error("error while reading the response of the list of services from data-store-service at", targetURL, err)
			continue
		}

		//parsing the response
		p := struct {
			Message string
			Data    []services.Service
		}{}
		err = json.Unmarshal(body, &p)
		if err != nil { //error while making the parsing the response from the api
			l.Error("error while parsing the response of the list of services from data-store-service at", targetURL, err)
			continue
		}

		//got the response
		l.Info("got the response message from the data-integration service", p.Message)
		result = p.Data
	}

	return result, nil
}

//GetDatastore returns the info of data store provided in the platform
//discoveryAddress is the address of the discovery service of cuttle platform
//discoveryToken is the token to be used to talk to the discovery service in the cuttle platform
//appToken is the authentication token of the user who is logged into the system
//serviceID is the id of the service
func GetDatastore(l log.Log, discoveryAddress, discoveryToken, appToken string, serviceID uint) (*services.Service, error) {
	/*
	 * First we will create the discovery config
	 * Then get the data integration services from discovery service
	 * Then will try to get the datastore list from each of them (whichever delivers first)
	 */
	//creating the discovery config
	dConfig := api.DefaultConfig()
	dConfig.Address = discoveryAddress
	dConfig.Token = discoveryToken

	//getting the data-integration services
	svs, err := discovery.GetServices(dConfig, "Brain-Data-Integeration-Service", l)
	if err != nil {
		//error while getting the services from the discovery
		l.Error("error while getting the list of Brain-Data-Integeration-Service from discovery service")
		return nil, err
	}

	//now we will try to get info of the service
	result := &services.Service{}
	payload, err := json.Marshal(services.Service{Model: gorm.Model{ID: serviceID}})
	if err != nil {
		//error while encoding the service
		l.Error("error while encoding the service")
		return nil, err
	}
	for _, v := range svs {
		targetURL := "http://" + v.Address + ":" + strconv.Itoa(v.Port) + "/services/datastore/get"
		l.Info("going to get the list of services from", targetURL)
		res, err := httpclient.Post(v.Address, targetURL, appToken, "auth-token", bytes.NewBuffer(payload))
		if err != nil {
			//error while making the request to get the info of the service
			l.Error("error while getting the info of service from data-store-service at", targetURL, err)
			continue
		}
		defer res.Body.Close()

		//read the response
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			//error while making the reading the response from the api
			l.Error("error while reading the response of the info of service from data-store-service at", targetURL, err)
			continue
		}

		//parsing the response
		p := struct {
			Message string
			Data    *services.Service
		}{}
		err = json.Unmarshal(body, &p)
		if err != nil { //error while making the parsing the response from the api
			l.Error("error while parsing the response of the info of service from data-store-service at", targetURL, err)
			continue
		}

		//got the response
		l.Info("got the response message from the data-integration service", p.Message)
		result = p.Data
	}

	return result, nil
}
