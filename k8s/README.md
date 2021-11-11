# Kubernetes

Start local minikube cluster

```bash
minikube start --vm-driver=hyperkit
minikube addons enable ingress
```

Add deployment, service and ingress to Kubernetes cluster

```bash
kubectl apply -f ./deployment.yml
kubectl apply -f ./service.yml
kubectl apply -f ./ingress.yml
```

Access application via URL:
<br />
http://k8s-local/nuscrape

Do not forget to add the `k8s-local` domain name to `/etc/hosts` for correct resolving.
You will get the correct IP to minikube with `minikube ip`.
