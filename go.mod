module cortify

go 1.15

require (
	github.com/buildpacks/pack v0.14.0
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-delve/delve v1.6.0 // indirect
	github.com/markbates/pkger v0.17.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/ory/viper v1.7.4
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.1-0.20201006035406-b97b5ead31f7
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.19.1
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/client v0.19.1
	knative.dev/eventing v0.19.0
	knative.dev/pkg v0.0.0-20201103163404-5514ab0c1fdf
	knative.dev/serving v0.19.0
)

replace (
	// Nail down k8 deps to align with transisitive deps
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
)
