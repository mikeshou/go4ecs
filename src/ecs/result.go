package ecs

import (
	"fmt"
)

type BaseResult struct {
	RequestId string
}

type ErrorResult struct {
	BaseResult
	HostId  string
	Code    string
	Message string
}

func (e ErrorResult) Error() string {
	return fmt.Sprintf("code:%v; msg:%v", e.Code, e.Message)
}

type DescribeInstanceStatusResult struct {
	BaseResult
	TotalCount       int
	PageNumber       int
	PageSize         int
	InstanceStatuses instanceStatus
}

type instanceStatus struct {
	InstanceStatus []instanceStatusResult
}

type instanceStatusResult struct {
	InstanceId string
	Status     string
}

type DescribeInstanceAttributeResult struct {
	BaseResult
	InstanceId              string
	ImageId                 string
	RegionId                string
	ZoneId                  string
	InstanceType            string
	HostName                string
	Status                  string
	PublicIpAddress         ipAddress
	InnerIpAddress          ipAddress
	InternetMaxBandwidthOut int
	InternetMaxBandwidthIn  int
	SecurityGroups          securityGroupId
}

type ipAddress struct {
	IpAddress []string
}

type securityGroupId struct {
	SecurityGroupId []string
}

type CreateInstanceResult struct {
	BaseResult
	InstanceId string
	ZoneId     string
}

type AddDiskResult struct {
	BaseResult
	DiskId string
}

type DescribeInstanceDisksResult struct {
	BaseResult
	Disks []disksResult
}

type disksResult struct {
	Disk []diskResult
}

type diskResult struct {
	DiskId string
	Size   int
	Type   disk
}

type CreateSnapshotResult struct {
	BaseResult
	SnapshotId string
}

type DescribeSnapshotsResult struct {
	BaseResult
	Snapshots snapshotsResult
}

type snapshotsResult struct {
	Snapshot []snapshotResult
}

type snapshotResult struct {
	SnapshotId   string
	Progress     string
	CreationTime string
}

type DescribeSnapshotAttributeResult struct {
	BaseResult
	SnapshotId   string
	Progress     string
	CreationTime string
}

type DescribeImagesResult struct {
	BaseResult
	RegionId   string
	TotalCount int
	PageNumber int
	PageSize   int
	Images     imagesResult
}

type imagesResult struct {
	Images []imageResult
}

type imageResult struct {
	ImageId      string
	ImageVersion string
	Platform     string
	Description  string
	Size         int
}

type CreateImageResult struct {
	BaseResult
	ImageId string
}

type AllocatePublicIpAddressResult struct {
	BaseResult
	IpAddress string
}

type CreateSecurityGroupResult struct {
	BaseResult
	SecurityGroupId string
}

type DescribeSecurityGroupAttributeResult struct {
	BaseResult
	SecurityGroupId string
	RegionId        string
	Description     string
	Permissions     permissionsResult
}
type permissionsResult struct {
	Permission []permissionResult
}

type permissionResult struct {
	IpProtocol    string
	PortRange     string
	SourceGroupId string
	Policy        string
	NicType       string
}

type DescribeSecurityGroupsResult struct {
	BaseResult
	TotalCount     int
	PageSize       int
	RegionId       string
	PageNumber     int
	SecurityGroups securityGroupsResult
}

type securityGroupsResult struct {
	SecurityGroup []securityGroupResult
}

type securityGroupResult struct {
	SecurityGroupId string
	Description     string
}

type DescribeRegionsResult struct {
	BaseResult
	Regions regionsResult
}

type regionsResult struct {
	Region []regionResult
}

type regionResult struct {
	RegionId string
}

type DescribeZonesResult struct {
	BaseResult
	Zones zonesResult
}

type zonesResult struct {
	Zone []zoneResult
}

type zoneResult struct {
	ZoneId string
}

type GetMonitorDataResult struct {
	BaseResult
	MonitorData instanceMonitorDatasResult
}

type instanceMonitorDatasResult struct {
	InstanceMonitorData []instanceMonitorDataResult
}

type instanceMonitorDataResult struct {
	InstanceId        string
	CPU               int
	Memory            int
	IntranetRX        int
	IntranetTX        int
	IntranetFlow      int
	IntranetBandwidth int
	InternetRX        int
	InternetTX        int
	InternetFlow      int
	InternetBandwidth int
	IOPSRead          int
	IOPSWrite         int
	BPSRead           int
	BPSWrite          int
	TimeStamp         string
}

type DescribeInstanceTypesResult struct {
	InstanceTypes instanceTypesResult
}

type instanceTypesResult struct {
	InstanceType []instanceTypeResult
}

type instanceTypeResult struct {
	InstanceTypeId string
	CpuCoreCount   int
	MemorySize     float32
}
