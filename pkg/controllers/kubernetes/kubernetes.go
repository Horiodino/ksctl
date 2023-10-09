package kubernetes

import (
	"fmt"

	k3sPkg "github.com/kubesimplify/ksctl/internal/k8sdistros/k3s"
	kubeadmPkg "github.com/kubesimplify/ksctl/internal/k8sdistros/kubeadm"
	"github.com/kubesimplify/ksctl/internal/k8sdistros/universal"
	"github.com/kubesimplify/ksctl/pkg/resources"
	. "github.com/kubesimplify/ksctl/pkg/utils/consts"
)

func HydrateK8sDistro(client *resources.KsctlClient) error {
	switch client.Metadata.K8sDistro {
	case K8S_K3S:
		client.Distro = k3sPkg.ReturnK3sStruct()
	case K8S_KUBEADM:
		client.Distro = kubeadmPkg.ReturnKubeadmStruct()
	default:
		return fmt.Errorf("[kubernetes] Invalid k8s provider")
	}
	return nil
}

func ConfigureCluster(client *resources.KsctlClient) error {
	err := client.Distro.ConfigureLoadbalancer(client.Storage)
	if err != nil {
		return err
	}

	for no := 0; no < client.Metadata.NoDS; no++ {
		err := client.Distro.ConfigureDataStore(no, client.Storage)
		if err != nil {
			return err
		}
	}
	for no := 0; no < client.Metadata.NoCP; no++ {
		err := client.Distro.Version(client.Metadata.K8sVersion).ConfigureControlPlane(no, client.Storage)
		if err != nil {
			return err
		}
	}

	for no := 0; no < int(client.Metadata.NoWP); no++ {
		err := client.Distro.Version(client.Metadata.K8sVersion).JoinWorkerplane(no, client.Storage)
		if err != nil {
			return err
		}
	}
	return nil
}

func JoinMoreWorkerPlanes(client *resources.KsctlClient, start, end int) error {
	for no := start; no < end; no++ {
		err := client.Distro.Version(client.Metadata.K8sVersion).JoinWorkerplane(no, client.Storage)
		if err != nil {
			return err
		}
	}
	return nil
}

func DelWorkerPlanes(client *resources.KsctlClient, hostnames []string) error {

	kubeconfigPath, _, err := client.Distro.GetKubeConfig(client.Storage)
	if err != nil {
		return err
	}

	kubernetesClient := universal.Kubernetes{
		Metadata:      client.Metadata,
		StorageDriver: client.Storage,
	}
	if err := kubernetesClient.ClientInit(kubeconfigPath); err != nil {
		return err
	}

	for _, hostname := range hostnames {
		if err := kubernetesClient.DeleteWorkerNodes(hostname); err != nil {
			return err
		}
	}
	return nil
}