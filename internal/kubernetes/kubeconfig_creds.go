package kubernetes

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ksctl/ksctl/pkg/helpers/consts"
	"github.com/ksctl/ksctl/pkg/types"
	"k8s.io/client-go/tools/clientcmd"
)

func httpClient(caCert, clientCert, clientKey []byte) (*tls.Config, error) {

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.X509KeyPair(clientCert, clientKey)
	if err != nil {
		return nil, log.NewError(kubernetesCtx, "Error loading client certificate and key", "Reason", err)
	}

	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}
	return tlsConfig, nil
}

func ExtractURLAndTLSCerts(kubeconfig, clusterContextName string) (url string, tlsConf *tls.Config, err error) {

	config, err := clientcmd.Load([]byte(kubeconfig))
	if err != nil {
		return "", nil, log.NewError(kubernetesCtx, "failed deserializes the contents into Config object", "Reason", err)
	}

	cluster := config.Clusters[clusterContextName]
	usr := config.AuthInfos[clusterContextName]

	kubeapiURL := cluster.Server
	caCert := cluster.CertificateAuthorityData
	clientCert := usr.ClientCertificateData
	clientKey := usr.ClientKeyData

	tlsConf, _err := httpClient(caCert, clientCert, clientKey)
	if _err != nil {
		return "", nil, _err
	}
	return kubeapiURL, tlsConf, nil
}

func transferData(kubeconfig,
	clusterContextName,
	podName,
	podNs string,
	podPort int,
	v *types.StorageStateExportImport) error {

	url, tlsConf, err := ExtractURLAndTLSCerts(kubeconfig, clusterContextName)
	if err != nil {
		return err
	}

	out, err := json.Marshal(v)
	if err != nil {
		return log.NewError(kubernetesCtx, "failed to marshal the exported stateDocuments", "Reason", err)
	}

	url = fmt.Sprintf("%s/api/v1/namespaces/%s/pods/%s:%d/proxy/import", url, podNs, podName, podPort)

	log.Debug(kubernetesCtx, "full url for state transfer", "url", url)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(out))
	if err != nil {
		return log.NewError(kubernetesCtx, "failed, client could not create request", "Reason", err)
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConf,
	}

	client := &http.Client{Transport: tr, Timeout: 1 * time.Minute}

	for counter := 0; counter <= int(consts.CounterMaxRetryCount); counter++ {

		res, err := client.Do(req)
		if err != nil {
			return log.NewError(kubernetesCtx, "failed, client error making http request", "Reason", err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return log.NewError(kubernetesCtx, "failed, to read response", "Reason", err)
		}

		if res.StatusCode < 300 {
			log.Success(kubernetesCtx, "Response of successful state transfer", "StatusCode", res.StatusCode, "Response", string(body))
			break
		}
		if counter == int(consts.CounterMaxRetryCount) {
			return log.NewError(kubernetesCtx, "failed, to send data", "Headers", res.Header, "StatusCode", res.StatusCode, "Response", string(body))
		}

		log.Warn(kubernetesCtx, "Error from state transfer",
			"failed", counter,
			"maxRetries", int(consts.CounterMaxRetryCount),
			"StatusCode", res.StatusCode,
			"Response", string(body))

		time.Sleep(10 * time.Second)
	}

	return nil
}
