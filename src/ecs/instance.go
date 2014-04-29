package ecs

import (
	"go-uuid/uuid"
)

//创建
func (ins *InstanceCreateParam) CreateInstance() (CreateInstanceResult, error) {
	m := make(map[string]string)
	m["RegionId"] = ins.RegionId
	m["ImageId"] = ins.ImageId
	m["InstanceType"] = ins.InstanceType
	m["SecurityGroupId"] = ins.SecurityGroupId
	m["InternetMaxBandwidthIn"] = ins.InternetMaxBandwidthIn
	m["InternetMaxBandwidthOut"] = ins.InternetMaxBandwidthOut
	m["HostName"] = ins.HostName
	m["Password"] = ins.Password
	m["ZoneId"] = ins.ZoneId
	m["ClientToken"] = uuid.NewRandom().String()
	m["action"] = "CreateInstance"
	br := CreateInstanceResult{}
	return br, doInvoke(m, &br)
}

//启动
func (ins *InstanceStartParam) StartInstance() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = ins.InstanceId
	m["action"] = "StartInstance"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//停止
func (ins *InstanceStopParam) StopInstance() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = ins.InstanceId
	m["ForceStop"] = ins.ForceStop
	m["action"] = "StopInstance"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//重启
func (ins *InstanceRebootParam) RebootInstance() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = ins.InstanceId
	m["ForceStop"] = ins.ForceStop
	m["action"] = "RebootInstance"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//重置
func (ins *InstanceResetParam) ResetInstance() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = ins.InstanceId
	m["ImageId"] = ins.ImageId
	m["DiskType"] = ins.Disk.Type
	m["action"] = "ResetInstance"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//修改实例属性
func (ins *InstanceModifyParam) ModifyInstanceAttribute() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = ins.InstanceId
	m["Password"] = ins.Password
	m["HostName"] = ins.HostName
	m["SecurityGroupId"] = ins.SecurityGroupId
	m["action"] = "ModifyInstanceAttribute"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//查询实例状态
func (ins *InstanceQueryParam) DescribeInstanceStatus() (DescribeInstanceStatusResult, error) {
	m := make(map[string]string)
	m["RegionId"] = ins.RegionId
	m["ZoneId"] = ins.ZoneId
	m["PageNumber"] = ins.Page.PageNumber
	m["PageSize"] = ins.Page.PageSize
	m["action"] = "DescribeInstanceStatus"
	br := DescribeInstanceStatusResult{}
	return br, doInvoke(m, &br)
}

//查询实例信息
func (ins *InstanceDescribeParam) DescribeInstanceAttribute() (DescribeInstanceAttributeResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = ins.InstanceId
	m["action"] = "DescribeInstanceAttribute"
	br := DescribeInstanceAttributeResult{}
	return br, doInvoke(m, &br)
}

//删除
func (ins *InstanceDeleteParam) DeleteInstance() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = ins.InstanceId
	m["action"] = "DeleteInstance"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//查询实例资源规格列表
func DescribeInstanceTypes() (DescribeInstanceTypesResult, error) {
	m := make(map[string]string)
	m["action"] = "DescribeInstanceTypes"
	br := DescribeInstanceTypesResult{}
	return br, doInvoke(m, &br)
}
