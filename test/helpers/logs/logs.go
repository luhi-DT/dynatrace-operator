//go:build e2e

package logs

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
)

func ReadLog(ctx context.Context, t *testing.T, environmentConfig *envconf.Config, namespace, podName, containerName string) string { //nolint:revive
	resources := environmentConfig.Client().Resources()

	var pod corev1.Pod
	require.NoError(t, resources.WithNamespace(namespace).Get(ctx, podName, namespace, &pod))

	clientset, err := kubernetes.NewForConfig(resources.GetConfig())
	require.NoError(t, err)

	logStream, err := clientset.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
		Container: containerName,
	}).Stream(ctx)
	require.NoError(t, err)

	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, logStream)
	require.NoError(t, err)

	return buffer.String()
}

func AssertContains(t *testing.T, logStream io.ReadCloser, contains string) {
	content := RequireContent(t, logStream)
	assert.Contains(t, content, contains)
}

func RequireContent(t *testing.T, logStream io.ReadCloser) string {
	buffer := new(bytes.Buffer)

	copied, err := io.Copy(buffer, logStream)
	require.NoError(t, err)
	require.Equal(t, int64(buffer.Len()), copied)

	return buffer.String()
}

func Contains(t *testing.T, logStream io.ReadCloser, contains string) bool {
	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, logStream)

	require.NoError(t, err)
	return strings.Contains(buffer.String(), contains)
}
