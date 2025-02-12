// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	"sync"

	"github.com/cf-platform-eng/kibosh/k8s"
	api_v1 "k8s.io/api/core/v1"
	v1_beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type FakeCluster struct {
	GetClientStub        func() kubernetes.Interface
	getClientMutex       sync.RWMutex
	getClientArgsForCall []struct{}
	getClientReturns     struct {
		result1 kubernetes.Interface
	}
	getClientReturnsOnCall map[int]struct {
		result1 kubernetes.Interface
	}
	GetClientConfigStub        func() *rest.Config
	getClientConfigMutex       sync.RWMutex
	getClientConfigArgsForCall []struct{}
	getClientConfigReturns     struct {
		result1 *rest.Config
	}
	getClientConfigReturnsOnCall map[int]struct {
		result1 *rest.Config
	}
	CreateNamespaceStub        func(*api_v1.Namespace) (*api_v1.Namespace, error)
	createNamespaceMutex       sync.RWMutex
	createNamespaceArgsForCall []struct {
		arg1 *api_v1.Namespace
	}
	createNamespaceReturns struct {
		result1 *api_v1.Namespace
		result2 error
	}
	createNamespaceReturnsOnCall map[int]struct {
		result1 *api_v1.Namespace
		result2 error
	}
	DeleteNamespaceStub        func(name string, options *meta_v1.DeleteOptions) error
	deleteNamespaceMutex       sync.RWMutex
	deleteNamespaceArgsForCall []struct {
		name    string
		options *meta_v1.DeleteOptions
	}
	deleteNamespaceReturns struct {
		result1 error
	}
	deleteNamespaceReturnsOnCall map[int]struct {
		result1 error
	}
	ListPodsStub        func(nameSpace string, listOptions meta_v1.ListOptions) (*api_v1.PodList, error)
	listPodsMutex       sync.RWMutex
	listPodsArgsForCall []struct {
		nameSpace   string
		listOptions meta_v1.ListOptions
	}
	listPodsReturns struct {
		result1 *api_v1.PodList
		result2 error
	}
	listPodsReturnsOnCall map[int]struct {
		result1 *api_v1.PodList
		result2 error
	}
	GetDeploymentStub        func(string, string, meta_v1.GetOptions) (*v1_beta1.Deployment, error)
	getDeploymentMutex       sync.RWMutex
	getDeploymentArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 meta_v1.GetOptions
	}
	getDeploymentReturns struct {
		result1 *v1_beta1.Deployment
		result2 error
	}
	getDeploymentReturnsOnCall map[int]struct {
		result1 *v1_beta1.Deployment
		result2 error
	}
	ListServiceAccountsStub        func(string, meta_v1.ListOptions) (*api_v1.ServiceAccountList, error)
	listServiceAccountsMutex       sync.RWMutex
	listServiceAccountsArgsForCall []struct {
		arg1 string
		arg2 meta_v1.ListOptions
	}
	listServiceAccountsReturns struct {
		result1 *api_v1.ServiceAccountList
		result2 error
	}
	listServiceAccountsReturnsOnCall map[int]struct {
		result1 *api_v1.ServiceAccountList
		result2 error
	}
	CreateServiceAccountStub        func(string, *api_v1.ServiceAccount) (*api_v1.ServiceAccount, error)
	createServiceAccountMutex       sync.RWMutex
	createServiceAccountArgsForCall []struct {
		arg1 string
		arg2 *api_v1.ServiceAccount
	}
	createServiceAccountReturns struct {
		result1 *api_v1.ServiceAccount
		result2 error
	}
	createServiceAccountReturnsOnCall map[int]struct {
		result1 *api_v1.ServiceAccount
		result2 error
	}
	ListClusterRoleBindingsStub        func(meta_v1.ListOptions) (*rbacv1beta1.ClusterRoleBindingList, error)
	listClusterRoleBindingsMutex       sync.RWMutex
	listClusterRoleBindingsArgsForCall []struct {
		arg1 meta_v1.ListOptions
	}
	listClusterRoleBindingsReturns struct {
		result1 *rbacv1beta1.ClusterRoleBindingList
		result2 error
	}
	listClusterRoleBindingsReturnsOnCall map[int]struct {
		result1 *rbacv1beta1.ClusterRoleBindingList
		result2 error
	}
	CreateClusterRoleBindingStub        func(*rbacv1beta1.ClusterRoleBinding) (*rbacv1beta1.ClusterRoleBinding, error)
	createClusterRoleBindingMutex       sync.RWMutex
	createClusterRoleBindingArgsForCall []struct {
		arg1 *rbacv1beta1.ClusterRoleBinding
	}
	createClusterRoleBindingReturns struct {
		result1 *rbacv1beta1.ClusterRoleBinding
		result2 error
	}
	createClusterRoleBindingReturnsOnCall map[int]struct {
		result1 *rbacv1beta1.ClusterRoleBinding
		result2 error
	}
	CreateSecretStub        func(nameSpace string, secret *api_v1.Secret) (*api_v1.Secret, error)
	createSecretMutex       sync.RWMutex
	createSecretArgsForCall []struct {
		nameSpace string
		secret    *api_v1.Secret
	}
	createSecretReturns struct {
		result1 *api_v1.Secret
		result2 error
	}
	createSecretReturnsOnCall map[int]struct {
		result1 *api_v1.Secret
		result2 error
	}
	ListSecretsStub        func(nameSpace string, listOptions meta_v1.ListOptions) (*api_v1.SecretList, error)
	listSecretsMutex       sync.RWMutex
	listSecretsArgsForCall []struct {
		nameSpace   string
		listOptions meta_v1.ListOptions
	}
	listSecretsReturns struct {
		result1 *api_v1.SecretList
		result2 error
	}
	listSecretsReturnsOnCall map[int]struct {
		result1 *api_v1.SecretList
		result2 error
	}
	ListServicesStub        func(nameSpace string, listOptions meta_v1.ListOptions) (*api_v1.ServiceList, error)
	listServicesMutex       sync.RWMutex
	listServicesArgsForCall []struct {
		nameSpace   string
		listOptions meta_v1.ListOptions
	}
	listServicesReturns struct {
		result1 *api_v1.ServiceList
		result2 error
	}
	listServicesReturnsOnCall map[int]struct {
		result1 *api_v1.ServiceList
		result2 error
	}
	PatchStub        func(nameSpace string, name string, pt types.PatchType, data []byte, subresources ...string) (result *api_v1.ServiceAccount, err error)
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		nameSpace    string
		name         string
		pt           types.PatchType
		data         []byte
		subresources []string
	}
	patchReturns struct {
		result1 *api_v1.ServiceAccount
		result2 error
	}
	patchReturnsOnCall map[int]struct {
		result1 *api_v1.ServiceAccount
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCluster) GetClient() kubernetes.Interface {
	fake.getClientMutex.Lock()
	ret, specificReturn := fake.getClientReturnsOnCall[len(fake.getClientArgsForCall)]
	fake.getClientArgsForCall = append(fake.getClientArgsForCall, struct{}{})
	fake.recordInvocation("GetClient", []interface{}{})
	fake.getClientMutex.Unlock()
	if fake.GetClientStub != nil {
		return fake.GetClientStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getClientReturns.result1
}

func (fake *FakeCluster) GetClientCallCount() int {
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	return len(fake.getClientArgsForCall)
}

func (fake *FakeCluster) GetClientReturns(result1 kubernetes.Interface) {
	fake.GetClientStub = nil
	fake.getClientReturns = struct {
		result1 kubernetes.Interface
	}{result1}
}

func (fake *FakeCluster) GetClientReturnsOnCall(i int, result1 kubernetes.Interface) {
	fake.GetClientStub = nil
	if fake.getClientReturnsOnCall == nil {
		fake.getClientReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Interface
		})
	}
	fake.getClientReturnsOnCall[i] = struct {
		result1 kubernetes.Interface
	}{result1}
}

func (fake *FakeCluster) GetClientConfig() *rest.Config {
	fake.getClientConfigMutex.Lock()
	ret, specificReturn := fake.getClientConfigReturnsOnCall[len(fake.getClientConfigArgsForCall)]
	fake.getClientConfigArgsForCall = append(fake.getClientConfigArgsForCall, struct{}{})
	fake.recordInvocation("GetClientConfig", []interface{}{})
	fake.getClientConfigMutex.Unlock()
	if fake.GetClientConfigStub != nil {
		return fake.GetClientConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getClientConfigReturns.result1
}

func (fake *FakeCluster) GetClientConfigCallCount() int {
	fake.getClientConfigMutex.RLock()
	defer fake.getClientConfigMutex.RUnlock()
	return len(fake.getClientConfigArgsForCall)
}

func (fake *FakeCluster) GetClientConfigReturns(result1 *rest.Config) {
	fake.GetClientConfigStub = nil
	fake.getClientConfigReturns = struct {
		result1 *rest.Config
	}{result1}
}

func (fake *FakeCluster) GetClientConfigReturnsOnCall(i int, result1 *rest.Config) {
	fake.GetClientConfigStub = nil
	if fake.getClientConfigReturnsOnCall == nil {
		fake.getClientConfigReturnsOnCall = make(map[int]struct {
			result1 *rest.Config
		})
	}
	fake.getClientConfigReturnsOnCall[i] = struct {
		result1 *rest.Config
	}{result1}
}

func (fake *FakeCluster) CreateNamespace(arg1 *api_v1.Namespace) (*api_v1.Namespace, error) {
	fake.createNamespaceMutex.Lock()
	ret, specificReturn := fake.createNamespaceReturnsOnCall[len(fake.createNamespaceArgsForCall)]
	fake.createNamespaceArgsForCall = append(fake.createNamespaceArgsForCall, struct {
		arg1 *api_v1.Namespace
	}{arg1})
	fake.recordInvocation("CreateNamespace", []interface{}{arg1})
	fake.createNamespaceMutex.Unlock()
	if fake.CreateNamespaceStub != nil {
		return fake.CreateNamespaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createNamespaceReturns.result1, fake.createNamespaceReturns.result2
}

func (fake *FakeCluster) CreateNamespaceCallCount() int {
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	return len(fake.createNamespaceArgsForCall)
}

func (fake *FakeCluster) CreateNamespaceArgsForCall(i int) *api_v1.Namespace {
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	return fake.createNamespaceArgsForCall[i].arg1
}

func (fake *FakeCluster) CreateNamespaceReturns(result1 *api_v1.Namespace, result2 error) {
	fake.CreateNamespaceStub = nil
	fake.createNamespaceReturns = struct {
		result1 *api_v1.Namespace
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) CreateNamespaceReturnsOnCall(i int, result1 *api_v1.Namespace, result2 error) {
	fake.CreateNamespaceStub = nil
	if fake.createNamespaceReturnsOnCall == nil {
		fake.createNamespaceReturnsOnCall = make(map[int]struct {
			result1 *api_v1.Namespace
			result2 error
		})
	}
	fake.createNamespaceReturnsOnCall[i] = struct {
		result1 *api_v1.Namespace
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) DeleteNamespace(name string, options *meta_v1.DeleteOptions) error {
	fake.deleteNamespaceMutex.Lock()
	ret, specificReturn := fake.deleteNamespaceReturnsOnCall[len(fake.deleteNamespaceArgsForCall)]
	fake.deleteNamespaceArgsForCall = append(fake.deleteNamespaceArgsForCall, struct {
		name    string
		options *meta_v1.DeleteOptions
	}{name, options})
	fake.recordInvocation("DeleteNamespace", []interface{}{name, options})
	fake.deleteNamespaceMutex.Unlock()
	if fake.DeleteNamespaceStub != nil {
		return fake.DeleteNamespaceStub(name, options)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteNamespaceReturns.result1
}

func (fake *FakeCluster) DeleteNamespaceCallCount() int {
	fake.deleteNamespaceMutex.RLock()
	defer fake.deleteNamespaceMutex.RUnlock()
	return len(fake.deleteNamespaceArgsForCall)
}

func (fake *FakeCluster) DeleteNamespaceArgsForCall(i int) (string, *meta_v1.DeleteOptions) {
	fake.deleteNamespaceMutex.RLock()
	defer fake.deleteNamespaceMutex.RUnlock()
	return fake.deleteNamespaceArgsForCall[i].name, fake.deleteNamespaceArgsForCall[i].options
}

func (fake *FakeCluster) DeleteNamespaceReturns(result1 error) {
	fake.DeleteNamespaceStub = nil
	fake.deleteNamespaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCluster) DeleteNamespaceReturnsOnCall(i int, result1 error) {
	fake.DeleteNamespaceStub = nil
	if fake.deleteNamespaceReturnsOnCall == nil {
		fake.deleteNamespaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteNamespaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCluster) ListPods(nameSpace string, listOptions meta_v1.ListOptions) (*api_v1.PodList, error) {
	fake.listPodsMutex.Lock()
	ret, specificReturn := fake.listPodsReturnsOnCall[len(fake.listPodsArgsForCall)]
	fake.listPodsArgsForCall = append(fake.listPodsArgsForCall, struct {
		nameSpace   string
		listOptions meta_v1.ListOptions
	}{nameSpace, listOptions})
	fake.recordInvocation("ListPods", []interface{}{nameSpace, listOptions})
	fake.listPodsMutex.Unlock()
	if fake.ListPodsStub != nil {
		return fake.ListPodsStub(nameSpace, listOptions)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listPodsReturns.result1, fake.listPodsReturns.result2
}

func (fake *FakeCluster) ListPodsCallCount() int {
	fake.listPodsMutex.RLock()
	defer fake.listPodsMutex.RUnlock()
	return len(fake.listPodsArgsForCall)
}

func (fake *FakeCluster) ListPodsArgsForCall(i int) (string, meta_v1.ListOptions) {
	fake.listPodsMutex.RLock()
	defer fake.listPodsMutex.RUnlock()
	return fake.listPodsArgsForCall[i].nameSpace, fake.listPodsArgsForCall[i].listOptions
}

func (fake *FakeCluster) ListPodsReturns(result1 *api_v1.PodList, result2 error) {
	fake.ListPodsStub = nil
	fake.listPodsReturns = struct {
		result1 *api_v1.PodList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListPodsReturnsOnCall(i int, result1 *api_v1.PodList, result2 error) {
	fake.ListPodsStub = nil
	if fake.listPodsReturnsOnCall == nil {
		fake.listPodsReturnsOnCall = make(map[int]struct {
			result1 *api_v1.PodList
			result2 error
		})
	}
	fake.listPodsReturnsOnCall[i] = struct {
		result1 *api_v1.PodList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) GetDeployment(arg1 string, arg2 string, arg3 meta_v1.GetOptions) (*v1_beta1.Deployment, error) {
	fake.getDeploymentMutex.Lock()
	ret, specificReturn := fake.getDeploymentReturnsOnCall[len(fake.getDeploymentArgsForCall)]
	fake.getDeploymentArgsForCall = append(fake.getDeploymentArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 meta_v1.GetOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetDeployment", []interface{}{arg1, arg2, arg3})
	fake.getDeploymentMutex.Unlock()
	if fake.GetDeploymentStub != nil {
		return fake.GetDeploymentStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeploymentReturns.result1, fake.getDeploymentReturns.result2
}

func (fake *FakeCluster) GetDeploymentCallCount() int {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return len(fake.getDeploymentArgsForCall)
}

func (fake *FakeCluster) GetDeploymentArgsForCall(i int) (string, string, meta_v1.GetOptions) {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return fake.getDeploymentArgsForCall[i].arg1, fake.getDeploymentArgsForCall[i].arg2, fake.getDeploymentArgsForCall[i].arg3
}

func (fake *FakeCluster) GetDeploymentReturns(result1 *v1_beta1.Deployment, result2 error) {
	fake.GetDeploymentStub = nil
	fake.getDeploymentReturns = struct {
		result1 *v1_beta1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) GetDeploymentReturnsOnCall(i int, result1 *v1_beta1.Deployment, result2 error) {
	fake.GetDeploymentStub = nil
	if fake.getDeploymentReturnsOnCall == nil {
		fake.getDeploymentReturnsOnCall = make(map[int]struct {
			result1 *v1_beta1.Deployment
			result2 error
		})
	}
	fake.getDeploymentReturnsOnCall[i] = struct {
		result1 *v1_beta1.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListServiceAccounts(arg1 string, arg2 meta_v1.ListOptions) (*api_v1.ServiceAccountList, error) {
	fake.listServiceAccountsMutex.Lock()
	ret, specificReturn := fake.listServiceAccountsReturnsOnCall[len(fake.listServiceAccountsArgsForCall)]
	fake.listServiceAccountsArgsForCall = append(fake.listServiceAccountsArgsForCall, struct {
		arg1 string
		arg2 meta_v1.ListOptions
	}{arg1, arg2})
	fake.recordInvocation("ListServiceAccounts", []interface{}{arg1, arg2})
	fake.listServiceAccountsMutex.Unlock()
	if fake.ListServiceAccountsStub != nil {
		return fake.ListServiceAccountsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listServiceAccountsReturns.result1, fake.listServiceAccountsReturns.result2
}

func (fake *FakeCluster) ListServiceAccountsCallCount() int {
	fake.listServiceAccountsMutex.RLock()
	defer fake.listServiceAccountsMutex.RUnlock()
	return len(fake.listServiceAccountsArgsForCall)
}

func (fake *FakeCluster) ListServiceAccountsArgsForCall(i int) (string, meta_v1.ListOptions) {
	fake.listServiceAccountsMutex.RLock()
	defer fake.listServiceAccountsMutex.RUnlock()
	return fake.listServiceAccountsArgsForCall[i].arg1, fake.listServiceAccountsArgsForCall[i].arg2
}

func (fake *FakeCluster) ListServiceAccountsReturns(result1 *api_v1.ServiceAccountList, result2 error) {
	fake.ListServiceAccountsStub = nil
	fake.listServiceAccountsReturns = struct {
		result1 *api_v1.ServiceAccountList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListServiceAccountsReturnsOnCall(i int, result1 *api_v1.ServiceAccountList, result2 error) {
	fake.ListServiceAccountsStub = nil
	if fake.listServiceAccountsReturnsOnCall == nil {
		fake.listServiceAccountsReturnsOnCall = make(map[int]struct {
			result1 *api_v1.ServiceAccountList
			result2 error
		})
	}
	fake.listServiceAccountsReturnsOnCall[i] = struct {
		result1 *api_v1.ServiceAccountList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) CreateServiceAccount(arg1 string, arg2 *api_v1.ServiceAccount) (*api_v1.ServiceAccount, error) {
	fake.createServiceAccountMutex.Lock()
	ret, specificReturn := fake.createServiceAccountReturnsOnCall[len(fake.createServiceAccountArgsForCall)]
	fake.createServiceAccountArgsForCall = append(fake.createServiceAccountArgsForCall, struct {
		arg1 string
		arg2 *api_v1.ServiceAccount
	}{arg1, arg2})
	fake.recordInvocation("CreateServiceAccount", []interface{}{arg1, arg2})
	fake.createServiceAccountMutex.Unlock()
	if fake.CreateServiceAccountStub != nil {
		return fake.CreateServiceAccountStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createServiceAccountReturns.result1, fake.createServiceAccountReturns.result2
}

func (fake *FakeCluster) CreateServiceAccountCallCount() int {
	fake.createServiceAccountMutex.RLock()
	defer fake.createServiceAccountMutex.RUnlock()
	return len(fake.createServiceAccountArgsForCall)
}

func (fake *FakeCluster) CreateServiceAccountArgsForCall(i int) (string, *api_v1.ServiceAccount) {
	fake.createServiceAccountMutex.RLock()
	defer fake.createServiceAccountMutex.RUnlock()
	return fake.createServiceAccountArgsForCall[i].arg1, fake.createServiceAccountArgsForCall[i].arg2
}

func (fake *FakeCluster) CreateServiceAccountReturns(result1 *api_v1.ServiceAccount, result2 error) {
	fake.CreateServiceAccountStub = nil
	fake.createServiceAccountReturns = struct {
		result1 *api_v1.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) CreateServiceAccountReturnsOnCall(i int, result1 *api_v1.ServiceAccount, result2 error) {
	fake.CreateServiceAccountStub = nil
	if fake.createServiceAccountReturnsOnCall == nil {
		fake.createServiceAccountReturnsOnCall = make(map[int]struct {
			result1 *api_v1.ServiceAccount
			result2 error
		})
	}
	fake.createServiceAccountReturnsOnCall[i] = struct {
		result1 *api_v1.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListClusterRoleBindings(arg1 meta_v1.ListOptions) (*rbacv1beta1.ClusterRoleBindingList, error) {
	fake.listClusterRoleBindingsMutex.Lock()
	ret, specificReturn := fake.listClusterRoleBindingsReturnsOnCall[len(fake.listClusterRoleBindingsArgsForCall)]
	fake.listClusterRoleBindingsArgsForCall = append(fake.listClusterRoleBindingsArgsForCall, struct {
		arg1 meta_v1.ListOptions
	}{arg1})
	fake.recordInvocation("ListClusterRoleBindings", []interface{}{arg1})
	fake.listClusterRoleBindingsMutex.Unlock()
	if fake.ListClusterRoleBindingsStub != nil {
		return fake.ListClusterRoleBindingsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listClusterRoleBindingsReturns.result1, fake.listClusterRoleBindingsReturns.result2
}

func (fake *FakeCluster) ListClusterRoleBindingsCallCount() int {
	fake.listClusterRoleBindingsMutex.RLock()
	defer fake.listClusterRoleBindingsMutex.RUnlock()
	return len(fake.listClusterRoleBindingsArgsForCall)
}

func (fake *FakeCluster) ListClusterRoleBindingsArgsForCall(i int) meta_v1.ListOptions {
	fake.listClusterRoleBindingsMutex.RLock()
	defer fake.listClusterRoleBindingsMutex.RUnlock()
	return fake.listClusterRoleBindingsArgsForCall[i].arg1
}

func (fake *FakeCluster) ListClusterRoleBindingsReturns(result1 *rbacv1beta1.ClusterRoleBindingList, result2 error) {
	fake.ListClusterRoleBindingsStub = nil
	fake.listClusterRoleBindingsReturns = struct {
		result1 *rbacv1beta1.ClusterRoleBindingList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListClusterRoleBindingsReturnsOnCall(i int, result1 *rbacv1beta1.ClusterRoleBindingList, result2 error) {
	fake.ListClusterRoleBindingsStub = nil
	if fake.listClusterRoleBindingsReturnsOnCall == nil {
		fake.listClusterRoleBindingsReturnsOnCall = make(map[int]struct {
			result1 *rbacv1beta1.ClusterRoleBindingList
			result2 error
		})
	}
	fake.listClusterRoleBindingsReturnsOnCall[i] = struct {
		result1 *rbacv1beta1.ClusterRoleBindingList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) CreateClusterRoleBinding(arg1 *rbacv1beta1.ClusterRoleBinding) (*rbacv1beta1.ClusterRoleBinding, error) {
	fake.createClusterRoleBindingMutex.Lock()
	ret, specificReturn := fake.createClusterRoleBindingReturnsOnCall[len(fake.createClusterRoleBindingArgsForCall)]
	fake.createClusterRoleBindingArgsForCall = append(fake.createClusterRoleBindingArgsForCall, struct {
		arg1 *rbacv1beta1.ClusterRoleBinding
	}{arg1})
	fake.recordInvocation("CreateClusterRoleBinding", []interface{}{arg1})
	fake.createClusterRoleBindingMutex.Unlock()
	if fake.CreateClusterRoleBindingStub != nil {
		return fake.CreateClusterRoleBindingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createClusterRoleBindingReturns.result1, fake.createClusterRoleBindingReturns.result2
}

func (fake *FakeCluster) CreateClusterRoleBindingCallCount() int {
	fake.createClusterRoleBindingMutex.RLock()
	defer fake.createClusterRoleBindingMutex.RUnlock()
	return len(fake.createClusterRoleBindingArgsForCall)
}

func (fake *FakeCluster) CreateClusterRoleBindingArgsForCall(i int) *rbacv1beta1.ClusterRoleBinding {
	fake.createClusterRoleBindingMutex.RLock()
	defer fake.createClusterRoleBindingMutex.RUnlock()
	return fake.createClusterRoleBindingArgsForCall[i].arg1
}

func (fake *FakeCluster) CreateClusterRoleBindingReturns(result1 *rbacv1beta1.ClusterRoleBinding, result2 error) {
	fake.CreateClusterRoleBindingStub = nil
	fake.createClusterRoleBindingReturns = struct {
		result1 *rbacv1beta1.ClusterRoleBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) CreateClusterRoleBindingReturnsOnCall(i int, result1 *rbacv1beta1.ClusterRoleBinding, result2 error) {
	fake.CreateClusterRoleBindingStub = nil
	if fake.createClusterRoleBindingReturnsOnCall == nil {
		fake.createClusterRoleBindingReturnsOnCall = make(map[int]struct {
			result1 *rbacv1beta1.ClusterRoleBinding
			result2 error
		})
	}
	fake.createClusterRoleBindingReturnsOnCall[i] = struct {
		result1 *rbacv1beta1.ClusterRoleBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) CreateSecret(nameSpace string, secret *api_v1.Secret) (*api_v1.Secret, error) {
	fake.createSecretMutex.Lock()
	ret, specificReturn := fake.createSecretReturnsOnCall[len(fake.createSecretArgsForCall)]
	fake.createSecretArgsForCall = append(fake.createSecretArgsForCall, struct {
		nameSpace string
		secret    *api_v1.Secret
	}{nameSpace, secret})
	fake.recordInvocation("CreateSecret", []interface{}{nameSpace, secret})
	fake.createSecretMutex.Unlock()
	if fake.CreateSecretStub != nil {
		return fake.CreateSecretStub(nameSpace, secret)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createSecretReturns.result1, fake.createSecretReturns.result2
}

func (fake *FakeCluster) CreateSecretCallCount() int {
	fake.createSecretMutex.RLock()
	defer fake.createSecretMutex.RUnlock()
	return len(fake.createSecretArgsForCall)
}

func (fake *FakeCluster) CreateSecretArgsForCall(i int) (string, *api_v1.Secret) {
	fake.createSecretMutex.RLock()
	defer fake.createSecretMutex.RUnlock()
	return fake.createSecretArgsForCall[i].nameSpace, fake.createSecretArgsForCall[i].secret
}

func (fake *FakeCluster) CreateSecretReturns(result1 *api_v1.Secret, result2 error) {
	fake.CreateSecretStub = nil
	fake.createSecretReturns = struct {
		result1 *api_v1.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) CreateSecretReturnsOnCall(i int, result1 *api_v1.Secret, result2 error) {
	fake.CreateSecretStub = nil
	if fake.createSecretReturnsOnCall == nil {
		fake.createSecretReturnsOnCall = make(map[int]struct {
			result1 *api_v1.Secret
			result2 error
		})
	}
	fake.createSecretReturnsOnCall[i] = struct {
		result1 *api_v1.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListSecrets(nameSpace string, listOptions meta_v1.ListOptions) (*api_v1.SecretList, error) {
	fake.listSecretsMutex.Lock()
	ret, specificReturn := fake.listSecretsReturnsOnCall[len(fake.listSecretsArgsForCall)]
	fake.listSecretsArgsForCall = append(fake.listSecretsArgsForCall, struct {
		nameSpace   string
		listOptions meta_v1.ListOptions
	}{nameSpace, listOptions})
	fake.recordInvocation("ListSecrets", []interface{}{nameSpace, listOptions})
	fake.listSecretsMutex.Unlock()
	if fake.ListSecretsStub != nil {
		return fake.ListSecretsStub(nameSpace, listOptions)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listSecretsReturns.result1, fake.listSecretsReturns.result2
}

func (fake *FakeCluster) ListSecretsCallCount() int {
	fake.listSecretsMutex.RLock()
	defer fake.listSecretsMutex.RUnlock()
	return len(fake.listSecretsArgsForCall)
}

func (fake *FakeCluster) ListSecretsArgsForCall(i int) (string, meta_v1.ListOptions) {
	fake.listSecretsMutex.RLock()
	defer fake.listSecretsMutex.RUnlock()
	return fake.listSecretsArgsForCall[i].nameSpace, fake.listSecretsArgsForCall[i].listOptions
}

func (fake *FakeCluster) ListSecretsReturns(result1 *api_v1.SecretList, result2 error) {
	fake.ListSecretsStub = nil
	fake.listSecretsReturns = struct {
		result1 *api_v1.SecretList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListSecretsReturnsOnCall(i int, result1 *api_v1.SecretList, result2 error) {
	fake.ListSecretsStub = nil
	if fake.listSecretsReturnsOnCall == nil {
		fake.listSecretsReturnsOnCall = make(map[int]struct {
			result1 *api_v1.SecretList
			result2 error
		})
	}
	fake.listSecretsReturnsOnCall[i] = struct {
		result1 *api_v1.SecretList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListServices(nameSpace string, listOptions meta_v1.ListOptions) (*api_v1.ServiceList, error) {
	fake.listServicesMutex.Lock()
	ret, specificReturn := fake.listServicesReturnsOnCall[len(fake.listServicesArgsForCall)]
	fake.listServicesArgsForCall = append(fake.listServicesArgsForCall, struct {
		nameSpace   string
		listOptions meta_v1.ListOptions
	}{nameSpace, listOptions})
	fake.recordInvocation("ListServices", []interface{}{nameSpace, listOptions})
	fake.listServicesMutex.Unlock()
	if fake.ListServicesStub != nil {
		return fake.ListServicesStub(nameSpace, listOptions)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listServicesReturns.result1, fake.listServicesReturns.result2
}

func (fake *FakeCluster) ListServicesCallCount() int {
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	return len(fake.listServicesArgsForCall)
}

func (fake *FakeCluster) ListServicesArgsForCall(i int) (string, meta_v1.ListOptions) {
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	return fake.listServicesArgsForCall[i].nameSpace, fake.listServicesArgsForCall[i].listOptions
}

func (fake *FakeCluster) ListServicesReturns(result1 *api_v1.ServiceList, result2 error) {
	fake.ListServicesStub = nil
	fake.listServicesReturns = struct {
		result1 *api_v1.ServiceList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) ListServicesReturnsOnCall(i int, result1 *api_v1.ServiceList, result2 error) {
	fake.ListServicesStub = nil
	if fake.listServicesReturnsOnCall == nil {
		fake.listServicesReturnsOnCall = make(map[int]struct {
			result1 *api_v1.ServiceList
			result2 error
		})
	}
	fake.listServicesReturnsOnCall[i] = struct {
		result1 *api_v1.ServiceList
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) Patch(nameSpace string, name string, pt types.PatchType, data []byte, subresources ...string) (result *api_v1.ServiceAccount, err error) {
	var dataCopy []byte
	if data != nil {
		dataCopy = make([]byte, len(data))
		copy(dataCopy, data)
	}
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		nameSpace    string
		name         string
		pt           types.PatchType
		data         []byte
		subresources []string
	}{nameSpace, name, pt, dataCopy, subresources})
	fake.recordInvocation("Patch", []interface{}{nameSpace, name, pt, dataCopy, subresources})
	fake.patchMutex.Unlock()
	if fake.PatchStub != nil {
		return fake.PatchStub(nameSpace, name, pt, data, subresources...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.patchReturns.result1, fake.patchReturns.result2
}

func (fake *FakeCluster) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *FakeCluster) PatchArgsForCall(i int) (string, string, types.PatchType, []byte, []string) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return fake.patchArgsForCall[i].nameSpace, fake.patchArgsForCall[i].name, fake.patchArgsForCall[i].pt, fake.patchArgsForCall[i].data, fake.patchArgsForCall[i].subresources
}

func (fake *FakeCluster) PatchReturns(result1 *api_v1.ServiceAccount, result2 error) {
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 *api_v1.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) PatchReturnsOnCall(i int, result1 *api_v1.ServiceAccount, result2 error) {
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 *api_v1.ServiceAccount
			result2 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 *api_v1.ServiceAccount
		result2 error
	}{result1, result2}
}

func (fake *FakeCluster) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getClientMutex.RLock()
	defer fake.getClientMutex.RUnlock()
	fake.getClientConfigMutex.RLock()
	defer fake.getClientConfigMutex.RUnlock()
	fake.createNamespaceMutex.RLock()
	defer fake.createNamespaceMutex.RUnlock()
	fake.deleteNamespaceMutex.RLock()
	defer fake.deleteNamespaceMutex.RUnlock()
	fake.listPodsMutex.RLock()
	defer fake.listPodsMutex.RUnlock()
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	fake.listServiceAccountsMutex.RLock()
	defer fake.listServiceAccountsMutex.RUnlock()
	fake.createServiceAccountMutex.RLock()
	defer fake.createServiceAccountMutex.RUnlock()
	fake.listClusterRoleBindingsMutex.RLock()
	defer fake.listClusterRoleBindingsMutex.RUnlock()
	fake.createClusterRoleBindingMutex.RLock()
	defer fake.createClusterRoleBindingMutex.RUnlock()
	fake.createSecretMutex.RLock()
	defer fake.createSecretMutex.RUnlock()
	fake.listSecretsMutex.RLock()
	defer fake.listSecretsMutex.RUnlock()
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCluster) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ k8s.Cluster = new(FakeCluster)
