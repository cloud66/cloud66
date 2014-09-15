package cloud66

type Service struct {
	Name       string      `json:"name"`
	Containers []Container `json:"containers"`
}

func (c *Client) GetServices(stackUid string, serverUid *string) ([]Service, error) {
	var params interface{}
	if serverUid == nil {
		params = nil
	} else {
		params = struct {
			ServerUid string `json:"server_uid"`
		}{
			ServerUid: *serverUid,
		}
	}
	req, err := c.NewRequest("GET", "/stacks/"+stackUid+"/services.json", params)
	if err != nil {
		return nil, err
	}
	var serviceRes []Service
	return serviceRes, c.DoReq(req, &serviceRes)
}

func (c *Client) GetService(stackUid string, serviceName string, serverUid *string) (*Service, error) {
	var params interface{}
	if serverUid == nil {
		params = nil
	} else {
		params = struct {
			ServerUid string `json:"server_uid"`
		}{
			ServerUid: *serverUid,
		}
	}
	req, err := c.NewRequest("GET", "/stacks/"+stackUid+"/services/"+serviceName+".json", params)
	if err != nil {
		return nil, err
	}
	var servicesRes *Service
	return servicesRes, c.DoReq(req, &servicesRes)
}

func (c *Client) StopService(stackUid string, serviceName string, serverUid *string) (*AsyncResult, error) {
	var params interface{}
	if serverUid == nil {
		params = nil
	} else {
		params = struct {
			ServerUid string `json:"server_uid"`
		}{
			ServerUid: *serverUid,
		}
	}
	req, err := c.NewRequest("DELETE", "/stacks/"+stackUid+"/services/"+serviceName+".json", params)
	if err != nil {
		return nil, err
	}
	var asyncRes *AsyncResult
	return asyncRes, c.DoReq(req, &asyncRes)
}

func (s *Service) ServerContainerCountMap() map[string]int {
	var serverMap = make(map[string]int)
	for _, container := range s.Containers {
		if _, present := serverMap[container.ServerName]; !present {
			serverMap[container.ServerName] = 1
		} else {
			serverMap[container.ServerName] = serverMap[container.ServerName] + 1
		}
	}
	return serverMap
}
