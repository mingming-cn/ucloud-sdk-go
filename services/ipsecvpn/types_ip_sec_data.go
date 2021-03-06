// Code is generated by ucloud-model, DO NOT EDIT IT.

package ipsecvpn

/*
IPSecData - IPSec参数

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type IPSecData struct {

	// IPSec通道中使用的认证算法
	IPSecAuthenticationAlgorithm string

	// IPSec通道中使用的加密算法
	IPSecEncryptionAlgorithm string

	// 指定VPN连接的本地子网，用逗号分隔
	IPSecLocalSubnetIds []string

	// 是否开启PFS功能,Disable表示关闭，数字表示DH组
	IPSecPFSDhGroup string

	// 使用的安全协议，ESP或AH
	IPSecProtocol string

	// 指定VPN连接的客户网段，用逗号分隔
	IPSecRemoteSubnets []string

	// IPSec中SA的生存时间
	IPSecSALifetime string

	// IPSec中SA的生存时间（以字节计）
	IPSecSALifetimeBytes string
}
