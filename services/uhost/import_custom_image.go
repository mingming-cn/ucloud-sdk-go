//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost ImportCustomImage

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ImportCustomImageRequest is request schema for ImportCustomImage action
type ImportCustomImageRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 镜像名称
	ImageName *string `required:"true"`

	// UFile私有空间地址
	UFileUrl *string `required:"true"`

	// 操作系统平台，比如CentOS、Ubuntu、Windows、RedHat等，请参考控制台的镜像版本；若导入控制台上没有的操作系统，参数为Other
	OsType *string `required:"true"`

	// 操作系统详细版本，请参考控制台的镜像版本；OsType为Other时，输入参数为Other
	OsName *string `required:"true"`

	// 镜像格式，可选RAW、VHD、VMDK、qcow2
	Format *string `required:"true"`

	// 是否授权。必须填true
	Auth *bool `required:"true"`

	// 镜像描述
	ImageDescription *string `required:"false"`
}

// ImportCustomImageResponse is response schema for ImportCustomImage action
type ImportCustomImageResponse struct {
	response.CommonBase

	// 镜像Id
	ImageId string
}

// NewImportCustomImageRequest will create request of ImportCustomImage action.
func (c *UHostClient) NewImportCustomImageRequest() *ImportCustomImageRequest {
	req := &ImportCustomImageRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// ImportCustomImage - 把UFile的镜像文件导入到UHost，生成自定义镜像
func (c *UHostClient) ImportCustomImage(req *ImportCustomImageRequest) (*ImportCustomImageResponse, error) {
	var err error
	var res ImportCustomImageResponse

	err = c.client.InvokeAction("ImportCustomImage", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}