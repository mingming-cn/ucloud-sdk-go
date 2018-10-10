//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api VPC DeleteSubnet

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteSubnetRequest is request schema for DeleteSubnet action
type DeleteSubnetRequest struct {
	request.CommonBase

	// 子网ID
	SubnetId *string `required:"true"`
}

// DeleteSubnetResponse is response schema for DeleteSubnet action
type DeleteSubnetResponse struct {
	response.CommonBase
}

// NewDeleteSubnetRequest will create request of DeleteSubnet action.
func (c *VPCClient) NewDeleteSubnetRequest() *DeleteSubnetRequest {
	req := &DeleteSubnetRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteSubnet - 删除子网
func (c *VPCClient) DeleteSubnet(req *DeleteSubnetRequest) (*DeleteSubnetResponse, error) {
	var err error
	var res DeleteSubnetResponse

	err = c.client.InvokeAction("DeleteSubnet", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}