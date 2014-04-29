package ecs

import ()

//创建安全组
func (group *GroupCreateParam) CreateSecurityGroup() (CreateSecurityGroupResult, error) {
	m := make(map[string]string)
	m["RegionId"] = group.RegionId
	m["Description"] = group.Description
	m["action"] = "CreateSecurityGroup"
	br := CreateSecurityGroupResult{}
	result := doInvoke(m, &br)
	if result == nil {
		return br, nil
	} else {
		return CreateSecurityGroupResult{}, result
	}
}

//授权安全组权限
func (group *GroupAuthorizeParam) AuthorizeSecurityGroup() (BaseResult, error) {
	m := make(map[string]string)
	m["RegionId"] = group.RegionId
	m["SecurityGroupId"] = group.SecurityGroupId
	m["IpProtocol"] = group.IpProtocol
	m["PortRange"] = group.PortRange
	m["SourceGroupId"] = group.SourceGroupId
	m["SourceCidrIp"] = group.SourceCidrIp
	m["Policy"] = group.Policy
	m["NicType"] = group.NicType
	m["action"] = "AuthorizeSecurityGroup"
	br := BaseResult{}
	result := doInvoke(m, &br)
	if result == nil {
		return br, nil
	} else {
		return BaseResult{}, result
	}
}

//查询安全组规则
func (group *GroupDescribeParam) DescribeSecurityGroupAttribute() (DescribeSecurityGroupAttributeResult, error) {
	m := make(map[string]string)
	m["RegionId"] = group.RegionId
	m["SecurityGroupId"] = group.SecurityGroupId
	m["NicType"] = group.NicType
	m["action"] = "DescribeSecurityGroupAttribute"
	br := DescribeSecurityGroupAttributeResult{}
	result := doInvoke(m, &br)
	if result == nil {
		return br, nil
	} else {
		return DescribeSecurityGroupAttributeResult{}, result
	}
}

//查询安全组列表
func (group *GroupQueryParam) DescribeSecurityGroups() (DescribeSecurityGroupsResult, error) {
	m := make(map[string]string)
	m["RegionId"] = group.RegionId
	m["PageNumber"] = group.Page.PageNumber
	m["PageSize"] = group.Page.PageSize
	m["action"] = "DescribeSecurityGroups"
	br := DescribeSecurityGroupsResult{}
	result := doInvoke(m, &br)
	if result == nil {
		return br, nil
	} else {
		return DescribeSecurityGroupsResult{}, result
	}
}

//撤销安全组规则
func (group *GroupRevokeParam) RevokeSecurityGroup() (BaseResult, error) {
	m := make(map[string]string)
	m["RegionId"] = group.RegionId
	m["SecurityGroupId"] = group.SecurityGroupId
	m["IpProtocol"] = group.IpProtocol
	m["PortRange"] = group.PortRange
	m["SourceGroupId"] = group.SourceGroupId
	m["SourceCidrIp"] = group.SourceCidrIp
	m["Policy"] = group.Policy
	m["NicType"] = group.NicType
	m["action"] = "RevokeSecurityGroup"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//删除安全组
func (group *GroupDeleteParam) DeleteSecurityGroup() (BaseResult, error) {
	m := make(map[string]string)
	m["RegionId"] = group.RegionId
	m["SecurityGroupId"] = group.SecurityGroupId
	m["action"] = "DeleteSecurityGroup"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}
