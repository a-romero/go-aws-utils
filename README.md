# go-aws-utils
Go module to handle AWS services.

## Secrets
This modules allows both to easily retrieve a secret from AWS Secrets Manager, as well as to store new ones. It assumes
permissions are appropriately set through IAM policies.

### Example
```
import (

  "github.com/a-romero/go-aws-utils/secrets"

)

type SvcConfig struct {
	Address    string `json:"address"`
	SecretPath string `json:"secretPath"`
	Project    string `json:"project"`
	Table      string `json:"table"`
}

type SvcClientIface interface {
	Query(string, time.Time, time.Time) ([][]string, error)
}

var _ SvcClientIface = (*SvcClient)(nil)

type SvcClient struct {
	httpClient    	*http.Client
	svcURL      	*url.URL
	project       	string
	key           	string
	table         	string
	queryTemplate 	*template.Template
}

func NewSvcClient(cfg SvcConfig) (SvcClientIface, error) {
	parsed, err := url.Parse(cfg.Address)
	if err != nil {
		return nil, err
	}

	apiKey, err := secrets.GetSecret(cfg.SecretPath)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	temp, err := template.New("query").Parse(queryTemplate)
	if err != nil {
		return nil, err
	}
	return &SvcClient{
		svcURL:      	parsed,
		httpClient:    	c,
		project:       	cfg.Project,
		key:           	apiKey,
		table:         	cfg.Table,
		queryTemplate: 	temp,
	}, nil
}

func (kc *SvcClient) Query(id string, startTimestamp time.Time, endTimestamp time.Time) ([][]string, error) {
	...
}
```
