package webhook

import (
	"testing"

	cmdManager "github.com/Dynatrace/dynatrace-operator/src/cmd/manager"
	"github.com/Dynatrace/dynatrace-operator/src/scheme"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
)

func TestCreateOptions(t *testing.T) {
	t.Run("implements interface", func(t *testing.T) {
		var provider cmdManager.Provider = NewProvider("certs-dir", "key-file", "cert-file")
		_, _ = provider.CreateManager("namespace", &rest.Config{})

		providerImpl := provider.(Provider)
		assert.Equal(t, "certs-dir", providerImpl.certificateDirectory)
		assert.Equal(t, "key-file", providerImpl.keyFileName)
		assert.Equal(t, "cert-file", providerImpl.certificateFileName)
	})
	t.Run("creates options", func(t *testing.T) {
		provider := Provider{}
		options := provider.createOptions("test-namespace")

		assert.NotNil(t, options)
		assert.Equal(t, scheme.Scheme, options.Scheme)
		assert.Equal(t, metricsBindAddress, options.MetricsBindAddress)
	})
}
