module github.com/haproxytech/dataplaneapi

go 1.21

require (
	github.com/GehirnInc/crypt v0.0.0-20230320061759-8cc1b52080c5
	github.com/aws/aws-sdk-go-v2 v1.21.1
	github.com/aws/aws-sdk-go-v2/config v1.18.44
	github.com/aws/aws-sdk-go-v2/credentials v1.13.42
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.30.7
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.124.0
	github.com/docker/go-units v0.5.0
	github.com/dustinkirkland/golang-petname v0.0.0-20231002161417-6a283f1aaaf2
	github.com/fsnotify/fsnotify v1.6.0
	github.com/getkin/kin-openapi v0.120.0
	github.com/go-openapi/errors v0.20.4
	github.com/go-openapi/loads v0.21.2
	github.com/go-openapi/runtime v0.26.0
	github.com/go-openapi/spec v0.20.9
	github.com/go-openapi/strfmt v0.21.7
	github.com/go-openapi/swag v0.22.4
	github.com/go-openapi/validate v0.22.1
	github.com/google/renameio v1.0.1
	github.com/google/uuid v1.3.1
	github.com/haproxytech/client-native/v5 v5.0.1-0.20231011083306-519f1b08727d
	github.com/haproxytech/config-parser/v5 v5.0.1-0.20230913105857-585d83cb48ba
	github.com/jessevdk/go-flags v1.5.0
	github.com/json-iterator/go v1.1.12
	github.com/kr/pretty v0.3.1
	github.com/lestrrat-go/apache-logformat v0.0.0-20210106032603-24d066f940f8
	github.com/maruel/panicparse/v2 v2.3.1
	github.com/nathanaelle/syslog5424/v2 v2.0.5
	github.com/rs/cors v1.10.1
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.4
	golang.org/x/net v0.17.0
	golang.org/x/sys v0.13.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.42 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.36 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.44 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.36 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.15.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.17.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.23.1 // indirect
	github.com/aws/smithy-go v1.15.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/haproxytech/go-logger v1.1.0 // indirect
	github.com/invopop/yaml v0.2.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.mongodb.org/mongo-driver v1.12.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/haproxytech/client-native/v5 => /home/tom/Documents/client-native
