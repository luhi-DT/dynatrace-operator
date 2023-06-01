package webhook

import (
	"crypto/tls"

	"github.com/Dynatrace/dynatrace-operator/src/scheme"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

const (
	metricsBindAddress = ":8383"
	port               = 8443
)

type Provider struct {
	certificateDirectory string
	certificateFileName  string
	keyFileName          string
}

func NewProvider(certificateDirectory string, keyFileName string, certificateFileName string) Provider {
	return Provider{
		certificateDirectory: certificateDirectory,
		certificateFileName:  certificateFileName,
		keyFileName:          keyFileName,
	}
}

func (provider Provider) CreateManager(namespace string, config *rest.Config) (manager.Manager, error) {
	mgr, err := ctrl.NewManager(config, provider.createOptions(namespace))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return mgr, nil
}

func (provider Provider) createOptions(namespace string) ctrl.Options {
	return ctrl.Options{
		Scheme:             scheme.Scheme,
		MetricsBindAddress: metricsBindAddress,
		WebhookServer:      provider.setupWebhookServer(),
		Cache: cache.Options{
			Namespaces: []string{
				namespace,
			},
		},
	}
}

func (provider Provider) setupWebhookServer() webhook.Server {
	tlsConfig := func(config *tls.Config) {
		config.MinVersion = tls.VersionTLS13
	}

	options := webhook.Options{
		Port:     port,
		CertDir:  provider.certificateDirectory,
		CertName: provider.certificateFileName,
		KeyName:  provider.keyFileName,
		TLSOpts:  []func(*tls.Config){tlsConfig},
	}

	return webhook.NewServer(options)
}
