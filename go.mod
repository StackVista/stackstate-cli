module gitlab.com/stackvista/stackstate-cli2

go 1.15

replace (
	k8s.io/api => k8s.io/api v0.19.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.2
	k8s.io/apiserver => k8s.io/apiserver v0.19.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.2
	k8s.io/client-go => k8s.io/client-go v0.19.2
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.2
	k8s.io/code-generator => k8s.io/code-generator v0.19.2
	k8s.io/component-base => k8s.io/component-base v0.19.2
	k8s.io/cri-api => k8s.io/cri-api v0.19.2
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.2
	k8s.io/kops => k8s.io/kops v1.19.1
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.2
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.2
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.2
	k8s.io/kubectl => k8s.io/kubectl v0.19.2
	k8s.io/kubelet => k8s.io/kubelet v0.19.2
	k8s.io/kubernetes => k8s.io/kubernetes v1.19.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.2
	k8s.io/metrics => k8s.io/metrics v0.19.2
	k8s.io/mount-utils => k8s.io/mount-utils v0.19.2
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.2
)

require (
	github.com/aws/aws-sdk-go-v2/config v1.8.3
	github.com/aws/aws-sdk-go-v2/service/s3 v1.16.1
	github.com/docker/docker v1.4.2-0.20200309214505-aa6a9891b09c // indirect
	github.com/gosuri/uilive v0.0.4 // indirect
	github.com/hierynomus/taipan v1.2.0
	github.com/logrusorgru/aurora/v3 v3.0.0
	github.com/manifoldco/promptui v0.8.0
	github.com/markbates/pkger v0.17.1
	github.com/mcuadros/go-defaults v1.2.0
	github.com/mitchellh/copystructure v1.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rs/zerolog v1.25.0
	github.com/spf13/cobra v1.2.1
	github.com/stackvista/sandbox-operator v0.5.0
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	k8s.io/kops v1.22.0
	k8s.io/kubectl v0.22.2 // indirect
	sigs.k8s.io/controller-runtime v0.10.2
)
