package resources

import (
	"github.com/ksctl/ksctl/pkg/helpers/consts"
	"github.com/ksctl/ksctl/pkg/resources/controllers/cloud"
)

type KubernetesBootstrap interface {
	Setup(storage StorageFactory, operation consts.KsctlOperation) error

	ConfigureControlPlane(int, StorageFactory) error

	JoinWorkerplane(int, StorageFactory) error

	Version(string) KubernetesBootstrap

	CNI(string) (externalCNI bool)
}

type PreKubernetesBootstrap interface {
	Setup(cloud.CloudResourceState, StorageFactory, consts.KsctlOperation) error

	ConfigureDataStore(int, StorageFactory) error

	ConfigureLoadbalancer(StorageFactory) error
}
