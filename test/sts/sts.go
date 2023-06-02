package sts

import (
	"encoding/json"
	goErrors "errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	faceid "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid/v20180301"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
)

type Sts struct {
	SecretId   string
	SecretKey  string
	credential common.CredentialIface
}
type Credential struct {
	Token        string `json:"token"`
	TmpSecretId  string `json:"tmpSecretId"`
	TmpSecretKey string `json:"tmpSecretKey"`
}
type IdCardVerifDetails struct {
	Result string `json:"result"`
	// 业务结果描述。
	Description string `json:"description"`
}

func NewSts(SecretId, SecretKey string) (*Sts, error) {
	s := &Sts{}
	if SecretId == "" || SecretKey == "" {
		provider := common.DefaultEnvProvider()
		credential, err := provider.GetCredential()
		if err != nil {
			return nil, err
		}
		s.credential = credential
	} else {
		credential := common.NewCredential(
			SecretId,
			SecretKey,
		)
		s.credential = credential
	}
	s.SecretId = s.credential.GetSecretId()
	s.SecretKey = s.credential.GetSecretKey()
	return s, nil
}

func (s *Sts) GetEFederationResponse(r R) (response *sts.GetFederationTokenResponse, err error) {
	m := make(map[string]interface{})
	m["Name"] = r.Name
	m["DurationSeconds"] = r.DurationSeconds
	policyBeta, _ := json.Marshal(r.Policy)
	m["Policy"] = string(policyBeta)
	client, err := sts.NewClient(s.credential, regions.Guangzhou, profile.NewClientProfile())
	if err != nil {
		return
	}
	request := sts.NewGetFederationTokenRequest()
	mBety, _ := json.Marshal(m)
	err = request.FromJsonString(string(mBety))
	if err != nil {
		return
	}
	// 通过client对象调用GetFederationToken接口，需要传入请求对象
	response, err = client.GetFederationToken(request)
	return
}
func (s *Sts) GetFederationCredentials(r R) (credential Credential) {
	response, err := s.GetEFederationResponse(r)
	if err != nil {
		return
	}
	if response != nil {
		if response.Response != nil {
			if response.Response.Credentials != nil {
				if response.Response.Credentials.Token != nil && response.Response.Credentials.TmpSecretId != nil && response.Response.Credentials.TmpSecretKey != nil {
					credential.TmpSecretId = *response.Response.Credentials.TmpSecretId
					credential.TmpSecretKey = *response.Response.Credentials.TmpSecretKey
					credential.Token = *response.Response.Credentials.Token
				}
			}
		}
	}
	return
}

func (s *Sts) IdCardVeri(idCard string, name string) (details IdCardVerifDetails, err error) {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	client, _ := faceid.NewClient(s.credential, "ap-guangzhou", cpf)

	request := faceid.NewIdCardVerificationRequest()

	request.IdCard = common.StringPtr(idCard)
	request.Name = common.StringPtr(name)

	response, err := client.IdCardVerification(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		err = goErrors.New("An API error has returned:" + err.Error())
		return
	}
	if err != nil {
		return
	}
	if response != nil {
		if response.Response != nil {
			if response.Response.Result != nil && response.Response.Description != nil {
				details.Result = *response.Response.Result
				details.Description = *response.Response.Description
			}
		}
	}
	return
}
