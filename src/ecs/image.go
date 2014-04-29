package ecs

import ()

//查询可用镜像
func (img *ImageDescribeParam) DescribeImages() (DescribeImagesResult, error) {
	m := make(map[string]string)
	m["RegionId"] = img.RegionId
	m["PageNumber"] = img.Page.PageNumber
	m["PageSize"] = img.Page.PageSize
	m["action"] = "DescribeImages"
	br := DescribeImagesResult{}
	return br, doInvoke(m, &br)
}

//创建自定义镜像
func (img *ImageCreateParam) CreateImage() (CreateImageResult, error) {
	m := make(map[string]string)
	m["RegionId"] = img.RegionId
	m["SnapshotId"] = img.SnapshotId
	m["ImageVersion"] = img.ImageVersion
	m["Description"] = img.Description
	m["Visibility"] = img.Visibility
	m["action"] = "CreateImage"
	br := CreateImageResult{}
	return br, doInvoke(m, &br)
}

//删除自定义镜像
func (img *ImageDeleteParam) DeleteImage() (BaseResult, error) {
	m := make(map[string]string)
	m["RegionId"] = img.RegionId
	m["ImageId"] = img.ImageId
	m["action"] = "DeleteImage"
	br := BaseResult{}
	return br, doInvoke(m, &br)
}
