package ecs

import ()

//增加磁盘设备
func (dev *DeviceAddParam) AddDisk() (AddDiskResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = dev.InstanceId
	m["Size"] = dev.Size
	m["SnapshotId"] = dev.SnapshotId
	m["action"] = "AddDisk"
	br := AddDiskResult{}
	return br, doInvoke(m, &br)
}

//删除磁盘
func (dev *DeviceDeleteParam) DeleteDisk() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = dev.InstanceId
	m["DiskId"] = dev.DiskId
	m["action"] = "DeleteDisk"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//查询实例磁盘列表
func (dev *DeviceDescribeParam) DescribeInstanceDisks() (DescribeInstanceDisksResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = dev.InstanceId
	m["action"] = "DescribeInstanceDisks"
	br := DescribeInstanceDisksResult{}
	return br, doInvoke(m, &br)
}
