package factory

import (
	dynatracev1alpha1 "github.com/Dynatrace/dynatrace-operator/pkg/apis/dynatrace/v1alpha1"
	_const "github.com/Dynatrace/dynatrace-operator/pkg/controller/const"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const KubeSystemUID = "a-very-unique-string"

func CreateFakeClient() client.Client {
	fakeClient := fake.NewFakeClientWithScheme(
		scheme.Scheme,
		&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      _const.ActivegateName,
				Namespace: _const.DynatraceNamespace,
			},
			Data: map[string][]byte{
				_const.DynatraceApiToken:  []byte("43"),
				_const.DynatracePaasToken: []byte("84"),
			},
		},
		&dynatracev1alpha1.DynaKube{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: _const.DynatraceNamespace,
				Name:      _const.ActivegateName,
			},
			Spec: dynatracev1alpha1.DynaKubeSpec{
				BaseDynaKubeSpec: dynatracev1alpha1.BaseDynaKubeSpec{
					APIURL: "https://ENVIRONMENTID.live.dynatrace.com/api",
				},
				OneAgent: dynatracev1alpha1.OneAgentSpec{
					Image: "dynatrace/oneagent:latest",
				},
				OneAgentAPMSpec: dynatracev1alpha1.OneAgentAPMSpec{
					Image: "dynatrace/oneagent:latest",
				},
				KubernetesMonitoringSpec: dynatracev1alpha1.KubernetesMonitoringSpec{
					Enabled: true,
				},
			},
		},
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: _const.KubeSystemNamespace,
				UID:  KubeSystemUID,
			},
		},
	)

	return fakeClient
}
