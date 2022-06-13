package configuration

import (
	"github.com/go-resty/resty/v2"
	"github.com/jfrog/terraform-provider-shared/client"
)

func SendConfigurationPatch(content []byte, m interface{}) error {
	_, err := m.(*resty.Client).R().SetBody(content).
		SetHeader("Content-Type", "application/yaml").
		AddRetryCondition(client.RetryOnMergeError).
		Patch("artifactory/api/system/configuration")

	return err
}
