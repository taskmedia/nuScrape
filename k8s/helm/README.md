# Kubernetes

Start local minikube cluster

```bash
minikube start --vm-driver=hyperkit
minikube addons enable ingress
```

To deploy nuScrape pleas use Helm.

# Helm

```bash
cd k8s/helm
helm install nuscrape ./
```

Access deployed application via URL:
<br />
http://k8s-local/nuscrape

Do not forget to add the `k8s-local` domain name to `/etc/hosts` for correct resolving.
You will get the correct IP to minikube with `minikube ip`.
