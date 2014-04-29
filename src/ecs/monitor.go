package ecs

import (

)

//查看云服务器监控信息
func (monitor *MonitorParam) GetMonitorData() (GetMonitorDataResult, error) {
	m := make(map[string]string)
	m["RegionId"] = monitor.RegionId
	m["InstanceId"] = monitor.InstanceId 
	m["action"] = "GetMonitorData"
	br := GetMonitorDataResult{}
	return br, doInvoke(m, &br)
}