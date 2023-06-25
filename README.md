# daemonset-cronjob-hybrid
A k8s operator that manages a custom resource which is a hybrid of daemonset and cronjob.

## Custom resource
```
apiVersion: example.com/v1
kind: DaemonJob
metadata:
  name: hello
spec:
  selector:
    matchLabels:
      name: hello-app
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        metadata:
          labels:
            name: hello-app
        spec:
          containers:
          - name: hello
            image: busybox:1.28
            command:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
```

## ToDo
* Create a job based on custom resource.
* Create jobs for each node.
* Ensure jobs run on each node on schedule.

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

### References
* [Cronjob struct](https://pkg.go.dev/k8s.io/api/batch/v1#CronJob)
* [Daemonset struct](https://pkg.go.dev/k8s.io/api/apps/v1#DaemonSet)

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Setting up the cluster
```sh
kind create cluster --config kind-config.yaml
```

### Running on the cluster
1. Build and push your image to the location specified by `IMG`:

```sh
make docker-build kind-load IMG=<some-registry>/daemonset-cronjob-hybrid:tag
```

1. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/daemonset-cronjob-hybrid:tag
```

1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/demo_v1_daemonjob.yaml
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)
