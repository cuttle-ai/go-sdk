// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package discovery contains the utilities required to communicate with the cuttle platform's
//discovery service
package discovery

import (
	"github.com/cuttle-ai/brain/log"
	"github.com/hashicorp/consul/api"
)

//GetServices will return the services of the given name
func GetServices(config *api.Config, name string, l log.Log) ([]*api.AgentService, error) {
	/*
	 * We initialize the client
	 * Then we get the list of services
	 * Then will find the service with the given name
	 */
	//initializing the client
	client, err := api.NewClient(config)
	if err != nil {
		//error while initializing the client
		l.Error("error while initializing the client for finding the service", name)
		return nil, err
	}

	//getting all the services
	services, err := client.Agent().Services()
	if err != nil {
		//Error while getting all the services list
		l.Error("Error while getting the list of services registered while finding the service", name)
		return nil, err
	}

	//iterating through the services to find the service with the given name
	serviceList := []*api.AgentService{}
	for _, v := range services {
		if v.ID == name {
			serviceList = append(serviceList, v)
		}
	}
	return serviceList, nil
}
