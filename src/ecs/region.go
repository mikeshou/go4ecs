package ecs

import ()

//查询可用数据中心
func DescribeRegions() (DescribeRegionsResult, error) {
	m := make(map[string]string)
	m["action"] = "DescribeRegions"
	br := DescribeRegionsResult{}
	return br, doInvoke(m, &br)
}

//查询 Zone 信息
func (region *ZoneParam) DescribeZones() (DescribeZonesResult, error) {
	m := make(map[string]string)
	m["RegionId"] = region.RegionId
	m["action"] = "DescribeZones"
	br := DescribeZonesResult{}
	return br, doInvoke(m, &br)
}
