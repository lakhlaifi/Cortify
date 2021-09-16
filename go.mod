module cortify

go 1.15

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.1-0.20201006035406-b97b5ead31f7 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.19.1
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/client v0.19.1
	knative.dev/serving v0.19.0
)

replace (
	// Nail down k8 deps to align with transisitive deps
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
)
