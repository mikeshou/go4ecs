package ecs

import ()

//创建快照
func (snapshot *SnapshotCreateParam) CreateSnapshot() (CreateSnapshotResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = snapshot.InstanceId
	m["DiskId"] = snapshot.DiskId
	m["SnapshotName"] = snapshot.SnapshotName
	m["action"] = "CreateSnapshot"
	br := CreateSnapshotResult{}
	return br, doInvoke(m, &br)
}

//删除快照
func (snapshot *SnapshotDeleteParam) DeleteSnapshot() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = snapshot.InstanceId
	m["DiskId"] = snapshot.DiskId
	m["SnapshotId"] = snapshot.SnapshotId
	m["action"] = "DeleteSnapshot"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}

//查询磁盘设备的快照列表
func (snapshot *SnapshotDescribeParam) DescribeSnapshots() (DescribeSnapshotsResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = snapshot.InstanceId
	m["DiskId"] = snapshot.DiskId
	m["action"] = "DescribeSnapshots"
	br := DescribeSnapshotsResult{}
	return br, doInvoke(m, &br)
}

//查询快照详情
func (snapshot *SnapshotDescribeAttributeParam) DescribeSnapshotAttribute() (DescribeSnapshotAttributeResult, error) {
	m := make(map[string]string)
	m["RegionId"] = snapshot.RegionId
	m["SnapshotId"] = snapshot.SnapshotId
	m["action"] = "DescribeSnapshotAttribute"
	br := DescribeSnapshotAttributeResult{}
	return br, doInvoke(m, &br)
}

//回滚快照
func (snapshot *SnapshotRollbackParam) RollbackSnapshot() (BaseResult, error) {
	m := make(map[string]string)
	m["InstanceId"] = snapshot.InstanceId
	m["DiskId"] = snapshot.DiskId
	m["SnapshotId"] = snapshot.SnapshotId
	m["action"] = "RollbackSnapshot"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}
