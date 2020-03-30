// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/functions"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario4377(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "4377",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"Image_Id_ucloud": "#{u_get_image_resource($Region,$Zone)}",
				"saopaulo_image":  "uimage-1bkjka",
				"Region":          "cn-bj2",
				"Zone":            "cn-bj2-02",
			}
		},
		Owners: []string{"li.wei@ucloud.cn"},
		Title:  "内网-外网-ulb7自动化回归-基本操作-01",
		Steps: []*driver.Step{
			testStep4377DescribeImage01,
			testStep4377CreateUHostInstance02,
			testStep4377CreateULB03,
			testStep4377DescribeULB04,
			testStep4377CreateVServer05,
			testStep4377DescribeVServer06,
			testStep4377AllocateBackendBatch07,
			testStep4377DescribeVServer08,
			testStep4377UpdateBackendAttribute09,
			testStep4377UpdateBackendAttributeBatch10,
			testStep4377ReleaseBackend11,
			testStep4377DescribeVServer12,
			testStep4377AllocateBackend13,
			testStep4377DescribeVServer14,
			testStep4377UpdateULBAttribute15,
			testStep4377UpdateVServerAttribute16,
			testStep4377DescribeVServer17,
			testStep4377DeleteVServer18,
			testStep4377DescribeVServer19,
			testStep4377DeleteULB20,
			testStep4377DescribeULBSimple21,
			testStep4377PoweroffUHostInstance22,
			testStep4377TerminateUHostInstance23,
		},
	})
}

var testStep4377DescribeImage01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"Region":    step.Scenario.GetVar("Region"),
			"OsType":    "Linux",
			"ImageType": "Base",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("Image_Id_ucloud", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeImageResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取镜像列表",
	FastFail:      false,
}

var testStep4377CreateUHostInstance02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":        step.Scenario.GetVar("Zone"),
			"Tag":         "Default",
			"Region":      step.Scenario.GetVar("Region"),
			"Password":    "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":        "ulb-host",
			"Memory":      1024,
			"MachineType": "N",
			"LoginMode":   "Password",
			"ImageId":     step.Scenario.GetVar("Image_Id_ucloud"),
			"Disks": []map[string]interface{}{
				{
					"IsBoot": "True",
					"Size":   20,
					"Type":   "LOCAL_NORMAL",
				},
			},
			"CPU": 1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("UHostId_01", step.Must(utils.GetValue(resp, "UHostIds.0")))
		step.Scenario.SetVar("IP_01", step.Must(utils.GetValue(resp, "IPs.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建云主机",
	FastFail:      false,
}

var testStep4377CreateULB03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreateULBRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBName":   "测试",
			"Tag":       "Default",
			"Region":    step.Scenario.GetVar("Region"),
			"InnerMode": "No",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateULB(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("ULBId", step.Must(utils.GetValue(resp, "ULBId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建负载均衡",
	FastFail:      false,
}

var testStep4377DescribeULB04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeULBRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
			"Offset": 0,
			"Limit":  60,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeULB(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ULBId", step.Scenario.GetVar("ULBId"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(3) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取负载均衡信息",
	FastFail:      false,
}

var testStep4377CreateVServer05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewCreateVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerName":     "vserver-test",
			"ULBId":           step.Scenario.GetVar("ULBId"),
			"Region":          step.Scenario.GetVar("Region"),
			"Protocol":        "HTTP",
			"PersistenceType": "UserDefined",
			"PersistenceInfo": "huangchao",
			"Method":          "Roundrobin",
			"ListenType":      "RequestProxy",
			"FrontendPort":    80,
			"ClientTimeout":   60,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "创建VServer",
	FastFail:      false,
}

var testStep4377DescribeVServer06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("VServerId", step.Must(utils.GetValue(resp, "DataSet.0.VServerId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ListenType", "RequestProxy", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VServerName", "vserver-test", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Protocol", "HTTP", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.FrontendPort", 80, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Method", "Roundrobin", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ClientTimeout", "60", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4377AllocateBackendBatch07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewAllocateBackendBatchRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId": step.Scenario.GetVar("VServerId"),
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Region":    step.Scenario.GetVar("Region"),
			"Backends": []interface{}{
				step.Must(functions.Concat(step.Scenario.GetVar("UHostId_01"), "|UHost|80|1|", step.Scenario.GetVar("IP_01"), "|0")),
			},
			"ApiVersion": 3,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AllocateBackendBatch(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "批量添加VServer后端节点",
	FastFail:      false,
}

var testStep4377DescribeVServer08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("BackendId", step.Must(utils.GetValue(resp, "DataSet.0.BackendSet.0.BackendId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ListenType", "RequestProxy", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.VServerName", "vserver-test", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Protocol", "HTTP", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.FrontendPort", 80, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.Method", "Roundrobin", "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.ClientTimeout", "60", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4377UpdateBackendAttribute09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewUpdateBackendAttributeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Weight":    0,
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Region":    step.Scenario.GetVar("Region"),
			"Port":      80,
			"Enabled":   0,
			"BackendId": step.Scenario.GetVar("BackendId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateBackendAttribute(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "更新后端实例属性",
	FastFail:      false,
}

var testStep4377UpdateBackendAttributeBatch10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("UpdateBackendAttributeBatch")
		err = req.SetPayload(map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
			"Attribute": []interface{}{
				step.Must(functions.Concat(step.Scenario.GetVar("BackendId"), "|80|0|0")),
			},
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "批量更新后端实例属性",
	FastFail:      false,
}

var testStep4377ReleaseBackend11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewReleaseBackendRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Region":    step.Scenario.GetVar("Region"),
			"BackendId": step.Scenario.GetVar("BackendId"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.ReleaseBackend(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "释放后端实例",
	FastFail:      false,
}

var testStep4377DescribeVServer12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.BackendSet", 0, "len_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4377AllocateBackend13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewAllocateBackendRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Weight":       0,
			"VServerId":    step.Scenario.GetVar("VServerId"),
			"ULBId":        step.Scenario.GetVar("ULBId"),
			"ResourceType": "UHost",
			"ResourceId":   step.Scenario.GetVar("UHostId_01"),
			"Region":       step.Scenario.GetVar("Region"),
			"Port":         80,
			"Enabled":      1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.AllocateBackend(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "添加后端实例",
	FastFail:      false,
}

var testStep4377DescribeVServer14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("DataSet.0.BackendSet", 1, "len_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4377UpdateULBAttribute15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewUpdateULBAttributeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
			"Name":   "测试-改",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateULBAttribute(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "更新负载均衡属性",
	FastFail:      false,
}

var testStep4377UpdateVServerAttribute16 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewUpdateVServerAttributeRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerName":     "vserver-gai",
			"VServerId":       step.Scenario.GetVar("VServerId"),
			"ULBId":           step.Scenario.GetVar("ULBId"),
			"Region":          step.Scenario.GetVar("Region"),
			"PersistenceType": "ServerInsert",
			"Method":          "Source",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateVServerAttribute(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "更新VServer属性",
	FastFail:      false,
}

var testStep4377DescribeVServer17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId": step.Scenario.GetVar("VServerId"),
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Region":    step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("TotalCount", 1, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4377DeleteVServer18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDeleteVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId": step.Scenario.GetVar("VServerId"),
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Region":    step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    20,
	RetryInterval: 10 * time.Second,
	Title:         "删除VServer",
	FastFail:      false,
}

var testStep4377DescribeVServer19 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDescribeVServerRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VServerId": step.Scenario.GetVar("VServerId"),
			"ULBId":     step.Scenario.GetVar("ULBId"),
			"Region":    step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeVServer(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取VServer信息",
	FastFail:      false,
}

var testStep4377DeleteULB20 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("ULB")
		if err != nil {
			return nil, err
		}
		client := c.(*ulb.ULBClient)

		req := client.NewDeleteULBRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteULB(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "删除负载均衡",
	FastFail:      false,
}

var testStep4377DescribeULBSimple21 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("")
		if err != nil {
			return nil, err
		}
		client := c.(*ucloud.Client)

		req := client.NewGenericRequest()
		_ = req.SetAction("DescribeULBSimple")
		err = req.SetPayload(map[string]interface{}{
			"ULBId":  step.Scenario.GetVar("ULBId"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}
		resp, err := client.GenericInvoke(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    10,
	RetryInterval: 10 * time.Second,
	Title:         "获取负载均衡信息",
	FastFail:      false,
}

var testStep4377PoweroffUHostInstance22 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("UHostId_01"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.PoweroffUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      false,
}

var testStep4377TerminateUHostInstance23 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("UHostId_01"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.TerminateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(60) * time.Second,
	MaxRetries:    30,
	RetryInterval: 10 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}
