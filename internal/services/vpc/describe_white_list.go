// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeWhiteListRequest is request schema for DescribeWhiteList action
type DescribeWhiteListRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// NATGateWay Id列表
	NATGWIds []string `required:"true"`
}

// DescribeWhiteListResponse is response schema for DescribeWhiteList action
type DescribeWhiteListResponse struct {
	response.CommonBase

	//
	DataSet []NatGWWhitelistDataSet

	// 满足条件的实例的总数
	TotalCount int
}

// NewDescribeWhiteListRequest will create request of DescribeWhiteList action.
func (c *VPCClient) NewDescribeWhiteListRequest() *DescribeWhiteListRequest {
	req := &DescribeWhiteListRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeWhiteList - 获取nat网关白名单列表
func (c *VPCClient) DescribeWhiteList(req *DescribeWhiteListRequest) (*DescribeWhiteListResponse, error) {
	var err error
	var res DescribeWhiteListResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeWhiteList", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
