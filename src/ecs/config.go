package ecs

type Config struct {
	AccessKeyId  string
	AccessSecret string
}

const (
	Url              string = "http://ecs.aliyuncs.com/?"
	Version          string = "2013-01-10"
	SignatureMethod  string = "HMAC-SHA1"
	Timestamp        string = "YYYY-MM-DDThh:mm:ssZ"
	SignatureVersion string = "1.0"
)

type Result struct {
	success bool
	code    int
	msg     string
	data    string
}

type InstanceBaseParam struct {
	InstanceId string
}

type InstanceStartParam struct {
	InstanceBaseParam
}

type InstanceStopParam struct {
	InstanceBaseParam
	ForceStop string
}

type InstanceRebootParam struct {
	InstanceStopParam
}

type InstanceResetParam struct {
	InstanceBaseParam
	ImageId string
	Disk    disk
}

type InstanceModifyParam struct {
	InstanceBaseParam
	Password        string
	HostName        string
	SecurityGroupId string
}

type InstanceQueryParam struct {
	RegionId string
	ZoneId   string
	Page     page
}

type InstanceDescribeParam struct {
	InstanceBaseParam
}

type InstanceDeleteParam struct {
	InstanceBaseParam
}

type InstanceCreateParam struct {
	RegionId                string
	ImageId                 string
	InstanceType            string
	SecurityGroupId         string
	InternetMaxBandwidthIn  string
	InternetMaxBandwidthOut string
	HostName                string
	Password                string
	ZoneId                  string
}

type DeviceAddParam struct {
	InstanceId string
	Size       string
	SnapshotId string
}

type DeviceDeleteParam struct {
	InstanceId string
	DiskId     string
}

type DeviceDescribeParam struct {
	InstanceId string
}

type GroupCreateParam struct {
	RegionId    string
	Description string
}

type GroupAuthorizeParam struct {
	RegionId        string
	SecurityGroupId string
	IpProtocol      string
	PortRange       string
	SourceGroupId   string
	SourceCidrIp    string
	Policy          string
	NicType         string
}

type GroupDescribeParam struct {
	RegionId        string
	SecurityGroupId string
	NicType         string
}

type GroupQueryParam struct {
	RegionId string
	Page     page
}

type GroupRevokeParam struct {
	GroupAuthorizeParam
}

type GroupDeleteParam struct {
	RegionId        string
	SecurityGroupId string
}

type ImageDescribeParam struct {
	RegionId     string
	Page         page
}

type ImageCreateParam struct {
	ImageId      string
	RegionId     string
	SnapshotId   string
	ImageVersion string
	Description  string
	Visibility   string
}

type ImageDeleteParam struct {
	ImageId      string
	RegionId     string
}

type NetworkAllocateParam struct {
	InstanceId      string
}

type NetworkReleaseParam struct {
	PublicIpAddress string
}

type SnapshotCreateParam struct {
	InstanceId   string
	DiskId       string
	SnapshotName string
}

type SnapshotDeleteParam struct {
	InstanceId   string
	DiskId       string
	SnapshotId   string
}

type SnapshotDescribeParam struct {
	InstanceId   string
	DiskId       string
}

type SnapshotDescribeAttributeParam struct {
	SnapshotId   string
	RegionId     string
}

type SnapshotRollbackParam struct {
	InstanceId   string
	DiskId       string
	SnapshotId   string
}

type ZoneParam struct {
	RegionId string
}

type MonitorParam struct {
	InstanceId string
	RegionId   string
}

type disk struct {
	Type string `system|data`
}

type page struct {
	PageNumber string
	PageSize   string
}
