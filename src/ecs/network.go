package ecs

import ()

//分配公网 IP 地址
func (network *NetworkAllocateParam) AllocatePublicIpAddress() (AllocatePublicIpAddressResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = network.InstanceId
	m["action"] = "AllocatePublicIpAddress"
	br := AllocatePublicIpAddressResult{}
	return br, doInvoke(m, &br)
}

//释放公网 IP 地址
func (network *NetworkReleaseParam) ReleasePublicIpAddress() (BaseResult, error) {
	m := make(map[string]string)
	m["PublicIpAddress"] = network.PublicIpAddress
	m["action"] = "ReleasePublicIpAddress"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}
