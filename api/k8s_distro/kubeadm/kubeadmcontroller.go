package kubeadm

import "fmt"

type KubeadmDistro struct {
	IsHA    bool
	Version string
}

func (kubeadm *KubeadmDistro) DestroyControlPlane() {
	//TODO implement me
	panic("implement me")
}

func (kubeadm *KubeadmDistro) InstallApplication() {
	//TODO implement me
	panic("implement me")
}

func (kubeadm *KubeadmDistro) ConfigureControlPlane() {
	//TODO implement me
	fmt.Println("Kubeadm Config CP")
}

func (kubeadm *KubeadmDistro) ConfigureWorkerPlane() {
	// TODO implement me
}

func (kubeadm *KubeadmDistro) DestroyWorkerPlane() {
	// TODO implement me
}

func (kubeadm *KubeadmDistro) ConfigureDataStore() {
	// TODO implement me
}

func (kubeadm *KubeadmDistro) DestroyDataStore() {
	// TODO implement me
}

func (kubeadm *KubeadmDistro) ConfigureLoadbalancer() {}

func (kubeadm *KubeadmDistro) DestroyLoadbalancer() {}
