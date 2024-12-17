// Package alicloud
/*
 * Copyright 2024-2025 Bilibili Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pvtz"
	"github.com/bilibili/HCP/cloudrepo"
)

// CloudRegion 云地域
type CloudRegion struct {
	cloudrepo.CloudProductCommon
}

// GetCID 获取云地域CID
func (c *CloudRegion) GetCID() string {
	return c.CID
}

// GetName 获取云地域名称
func (c *CloudRegion) GetName() string {
	return c.Name
}

// GetStatus 获取云地域状态
func (c *CloudRegion) GetStatus() string {
	return c.Status
}

// ListRegion 获取地域列表
func (c *Client) ListRegion(req *cloudrepo.GetCloudProductReq) ([]cloudrepo.CloudRegion, error) {
	cli, err := c.clientPvtz()
	if err != nil {
		return nil, err
	}
	request := pvtz.CreateDescribeRegionsRequest()
	request.RegionId = req.ResourceID
	request.AcceptLanguage = "zh-CN"
	resp, err := cli.DescribeRegions(request)
	if err != nil {
		return nil, err
	}
	return list2DoRegion(resp)
}

func list2DoRegion(resp *pvtz.DescribeRegionsResponse) (list []cloudrepo.CloudRegion, err error) {
	for _, v := range resp.Regions.Region {
		_ = v
		list = append(list, &CloudRegion{
			CloudProductCommon: cloudrepo.CloudProductCommon{
				CID:    v.RegionId,
				Name:   v.RegionName,
				Status: "available",
			},
		})
	}
	return
}