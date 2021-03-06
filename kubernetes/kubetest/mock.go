package kubetest

import (
	"github.com/stretchr/testify/mock"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/apps/v1beta2"
	auth_v1 "k8s.io/api/authorization/v1"
	batch_v1 "k8s.io/api/batch/v1"
	batch_v1beta1 "k8s.io/api/batch/v1beta1"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kiali/kiali/kubernetes"

	osappsv1 "github.com/openshift/api/apps/v1"
	osv1 "github.com/openshift/api/project/v1"
)

type K8SClientMock struct {
	mock.Mock
}

func NewK8SClientMock() *K8SClientMock {
	k8s := new(K8SClientMock)
	k8s.On("IsOpenShift").Return(true)
	return k8s
}

func (o *K8SClientMock) GetNamespace(namespace string) (*v1.Namespace, error) {
	args := o.Called(namespace)
	return args.Get(0).(*v1.Namespace), args.Error(1)
}

func (o *K8SClientMock) GetNamespaces() ([]v1.Namespace, error) {
	args := o.Called()
	return args.Get(0).([]v1.Namespace), args.Error(1)
}

func (o *K8SClientMock) GetProject(project string) (*osv1.Project, error) {
	args := o.Called(project)
	return args.Get(0).(*osv1.Project), args.Error(1)
}

func (o *K8SClientMock) GetProjects() ([]osv1.Project, error) {
	args := o.Called()
	return args.Get(0).([]osv1.Project), args.Error(1)
}

func (o *K8SClientMock) IsOpenShift() bool {
	args := o.Called()
	return args.Get(0).(bool)
}

func (o *K8SClientMock) GetDeployment(namespace string, deploymentName string) (*v1beta1.Deployment, error) {
	args := o.Called(namespace, deploymentName)
	return args.Get(0).(*v1beta1.Deployment), args.Error(1)
}

func (o *K8SClientMock) GetDeployments(namespace string) ([]v1beta1.Deployment, error) {
	args := o.Called(namespace)
	return args.Get(0).([]v1beta1.Deployment), args.Error(1)
}

func (o *K8SClientMock) GetDeploymentConfig(namespace string, deploymentName string) (*osappsv1.DeploymentConfig, error) {
	args := o.Called(namespace, deploymentName)
	return args.Get(0).(*osappsv1.DeploymentConfig), args.Error(1)
}

func (o *K8SClientMock) GetDeploymentConfigs(namespace string) ([]osappsv1.DeploymentConfig, error) {
	args := o.Called(namespace)
	return args.Get(0).([]osappsv1.DeploymentConfig), args.Error(1)
}

func (o *K8SClientMock) GetReplicaSet(namespace string, replicasetName string) (*v1beta2.ReplicaSet, error) {
	args := o.Called(namespace, replicasetName)
	return args.Get(0).(*v1beta2.ReplicaSet), args.Error(1)
}

func (o *K8SClientMock) GetReplicaSets(namespace string) ([]v1beta2.ReplicaSet, error) {
	args := o.Called(namespace)
	return args.Get(0).([]v1beta2.ReplicaSet), args.Error(1)
}

func (o *K8SClientMock) GetStatefulSet(namespace string, statefulsetName string) (*v1beta2.StatefulSet, error) {
	args := o.Called(namespace, statefulsetName)
	return args.Get(0).(*v1beta2.StatefulSet), args.Error(1)
}

func (o *K8SClientMock) GetStatefulSets(namespace string) ([]v1beta2.StatefulSet, error) {
	args := o.Called(namespace)
	return args.Get(0).([]v1beta2.StatefulSet), args.Error(1)
}

func (o *K8SClientMock) GetReplicationControllers(namespace string) ([]v1.ReplicationController, error) {
	args := o.Called(namespace)
	return args.Get(0).([]v1.ReplicationController), args.Error(1)
}

func (o *K8SClientMock) GetReplicationController(namespace string, replicationControllerName string) (*v1.ReplicationController, error) {
	args := o.Called(namespace, replicationControllerName)
	return args.Get(0).(*v1.ReplicationController), args.Error(1)
}

func (o *K8SClientMock) GetService(namespace string, serviceName string) (*v1.Service, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).(*v1.Service), args.Error(1)
}

func (o *K8SClientMock) GetServices(namespace string, selectorLabels map[string]string) ([]v1.Service, error) {
	args := o.Called(namespace, selectorLabels)
	return args.Get(0).([]v1.Service), args.Error(1)
}

func (o *K8SClientMock) GetPods(namespace, labelSelector string) ([]v1.Pod, error) {
	args := o.Called(namespace, labelSelector)
	return args.Get(0).([]v1.Pod), args.Error(1)
}

func (o *K8SClientMock) GetEndpoints(namespace string, serviceName string) (*v1.Endpoints, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).(*v1.Endpoints), args.Error(1)
}

func (o *K8SClientMock) GetServicePods(namespace string, serviceName string, serviceVersion string) (*v1.PodList, error) {
	args := o.Called(namespace, serviceName, serviceVersion)
	return args.Get(0).(*v1.PodList), args.Error(1)
}

func (o *K8SClientMock) GetIstioDetails(namespace string, serviceName string) (*kubernetes.IstioDetails, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).(*kubernetes.IstioDetails), args.Error(1)
}

func (o *K8SClientMock) GetGateways(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetGateway(namespace string, gateway string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, gateway)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetVirtualServices(namespace string, serviceName string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetVirtualService(namespace string, virtualservice string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, virtualservice)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetDestinationRules(namespace string, serviceName string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetDestinationRule(namespace string, destinationrule string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, destinationrule)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceEntries(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceEntry(namespace string, serviceEntryName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, serviceEntryName)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetIstioRules(namespace string) (*kubernetes.IstioRules, error) {
	args := o.Called(namespace)
	return args.Get(0).(*kubernetes.IstioRules), args.Error(1)
}

func (o *K8SClientMock) GetIstioRuleDetails(namespace string, istiorule string) (*kubernetes.IstioRuleDetails, error) {
	args := o.Called(namespace, istiorule)
	return args.Get(0).(*kubernetes.IstioRuleDetails), args.Error(1)
}

func (o *K8SClientMock) GetQuotaSpecs(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetQuotaSpec(namespace string, quotaSpecName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, quotaSpecName)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetQuotaSpecBindings(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetQuotaSpecBinding(namespace string, quotaSpecBindingName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, quotaSpecBindingName)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetCronJob(namespace, cronjobName string) (*batch_v1beta1.CronJob, error) {
	args := o.Called(namespace, cronjobName)
	return args.Get(0).(*batch_v1beta1.CronJob), args.Error(1)
}

func (o *K8SClientMock) GetCronJobs(namespace string) ([]batch_v1beta1.CronJob, error) {
	args := o.Called(namespace)
	return args.Get(0).([]batch_v1beta1.CronJob), args.Error(1)
}

func (o *K8SClientMock) GetJob(namespace, jobName string) (*batch_v1.Job, error) {
	args := o.Called(namespace, jobName)
	return args.Get(0).(*batch_v1.Job), args.Error(1)
}

func (o *K8SClientMock) GetJobs(namespace string) ([]batch_v1.Job, error) {
	args := o.Called(namespace)
	return args.Get(0).([]batch_v1.Job), args.Error(1)
}

func (o *K8SClientMock) GetSelfSubjectAccessReview(namespace, api, resourceType string, verbs []string) ([]*auth_v1.SelfSubjectAccessReview, error) {
	args := o.Called(namespace, api, resourceType, verbs)
	return args.Get(0).([]*auth_v1.SelfSubjectAccessReview), args.Error(1)
}

func (o *K8SClientMock) DeleteIstioObject(api, namespace, objectType, objectName string) error {
	args := o.Called(api, namespace, objectType, objectName)
	return args.Error(0)
}

func (o *K8SClientMock) FakeService() *v1.Service {
	return &v1.Service{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "httpbin",
			Namespace: "tutorial",
			Labels: map[string]string{
				"app":     "httpbin",
				"version": "v1"}},
		Spec: v1.ServiceSpec{
			ClusterIP: "fromservice",
			Type:      "ClusterIP",
			Selector:  map[string]string{"app": "httpbin"},
			Ports: []v1.ServicePort{
				{
					Name:     "http",
					Protocol: "TCP",
					Port:     3001},
				{
					Name:     "http",
					Protocol: "TCP",
					Port:     3000}}}}
}

func (o *K8SClientMock) FakeServiceList() []v1.Service {
	return []v1.Service{
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:      "reviews",
				Namespace: "tutorial",
				Labels: map[string]string{
					"app":     "reviews",
					"version": "v1"}},
			Spec: v1.ServiceSpec{
				ClusterIP: "fromservice",
				Type:      "ClusterIP",
				Selector:  map[string]string{"app": "reviews"},
				Ports: []v1.ServicePort{
					{
						Name:     "http",
						Protocol: "TCP",
						Port:     3001},
					{
						Name:     "http",
						Protocol: "TCP",
						Port:     3000}}}},
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:      "httpbin",
				Namespace: "tutorial",
				Labels: map[string]string{
					"app":     "httpbin",
					"version": "v1"}},
			Spec: v1.ServiceSpec{
				ClusterIP: "fromservice",
				Type:      "ClusterIP",
				Selector:  map[string]string{"app": "httpbin"},
				Ports: []v1.ServicePort{
					{
						Name:     "http",
						Protocol: "TCP",
						Port:     3001},
					{
						Name:     "http",
						Protocol: "TCP",
						Port:     3000}}}},
	}
}

func (o *K8SClientMock) FakePodList() []v1.Pod {
	return []v1.Pod{
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:   "reviews-v1",
				Labels: map[string]string{"app": "reviews", "version": "v1"}}},
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:   "reviews-v2",
				Labels: map[string]string{"app": "reviews", "version": "v2"}}},
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:   "httpbin-v1",
				Labels: map[string]string{"app": "httpbin", "version": "v1"}}},
	}
}
