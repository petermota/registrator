package eureka

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gliderlabs/registrator/bridge"
)

type EurekaInstance struct {
	Instance EurekaInfo `json:"instance"`
}

type EurekaInfo struct {
	InstanceID         string         `json:"instanceId"`
	HostName           string         `json:"hostName"`
	App                string         `json:"app"`
	IPAddr             string         `json:"ipAddr"`
	Status             string         `json:"status"`
	Port               PortInfo       `json:"port"`
	DataCenter         DataCenterInfo `json:"dataCenterInfo"`
	Lease              LeaseInfo      `json:"leaseInfo"`
	LastDirtyTimestamp int64          `json:"lastDirtyTimestamp"`
}

type PortInfo struct {
	Number  int    `json:"$"`
	Enabled string `json:"@enabled"`
}

type DataCenterInfo struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
}

type LeaseInfo struct {
	RenewalIntervalInSecs int `json:"renewalIntervalInSecs"`
}

func (r *EurekaAdapter) getAllApps() error {
	_, error := http.Get("http://" + r.host + r.path + "/apps")
	return error
}

func (r *EurekaAdapter) registerApp(svc *bridge.Service) error {
	host := strings.Split(svc.ID, ":")[0]
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	instance := EurekaInstance{EurekaInfo{svc.ID, host, svc.Name, svc.IP, "UP",
		PortInfo{svc.Port, "true"},
		DataCenterInfo{"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo", "MyOwn"},
		LeaseInfo{30}, millis}}

	jsonBody, _ := json.Marshal(instance)
	_, error := http.Post("http://"+r.host+r.path+"/apps/"+svc.Name,
		"application/json",
		bytes.NewBuffer(jsonBody))

	return error
}

func (r *EurekaAdapter) deregisterApp(svc *bridge.Service) error {
	client := &http.Client{}
	req, error := http.NewRequest(http.MethodDelete,
		"http://"+r.host+r.path+"/apps/"+svc.Name+"/"+svc.ID,
		nil)
	if error != nil {
		return error
	}

	_, error = client.Do(req)
	return error
}
