# cluster-api-provider-metal

## Intro

The Cluster API Provider Metal (CAPM) is a project by [Talos Systems](https://www.talos-systems.com/) that provides a [Cluster API](https://github.com/kubernetes-sigs/cluster-api)(CAPI) infrastructure provider for bare metal use.
Given a reference to a bare metal server and some BMC info, this provider will reconcile the necessary custom resources and boot the nodes using IPMI.

## Corequisites

There are a few corequisites and assumptions that go into using this project:

- [Metal Controller Manager](https://github.com/talos-systems/metal-controller-manager)
- [Metal Metadata Server](https://github.com/talos-systems/metal-metadata-server)
- [Cluster API](https://github.com/kubernetes-sigs/cluster-api)
- [Cluster API Bootstrap Provider Talos](https://github.com/talos-systems/cluster-api-bootstrap-provider-talos) (optional)

## Building and Installing

This project can be built simply by running `make release` from the root directory.
Doing so will create a file called `_out/infrastructure-components.yaml`.
If you wish, you can tweak settings by editing the release yaml.
This file can then be installed into your management cluster with `kubectl apply -f _out/infrastructure-components.yaml`.

Note that CAPM should be deployed as part of a set of controllers for Cluster API.
You will need at least the upstream CAPI components and a bootstrap provider for v1alpha2 CAPI capabilities.

## Usage

CAPM supports two API types, MetalClusters and MetalMachines.
You can create YAML definitions of each and `kubectl apply` them as part of a larger CAPI cluster deployment.
Below are some example definitions of each.

MetalCluster:

```yaml
apiVersion: infrastructure.cluster.x-k8s.io/v1alpha2
kind: MetalCluster
metadata:
  name: talos
  namespace: default
spec:
  apiEndpoints:
    - host: 10.254.0.5
      port: 6443
```

Note the apiEndpoint specified above.
This field is blindly copied into the status for the MetalCluster resource, as CAPI upstream expects their to be an endpoint published.
In cloud environments, this is generally a loadbalancer DNS name.
In the case of bare metal, it should ideally be a loadbalanced IP or DNS name for all master machines.
In my above example, this is an IP attached to the loopback interface of all masters and exposed via BGP/ECMP.

MetalMachine:

```yaml
apiVersion: infrastructure.cluster.x-k8s.io/v1alpha2
kind: MetalMachine
metadata:
  name: talos-0
spec:
  serverRef:
    apiVersion: metal.arges.dev/v1alpha1
    kind: Server
    name: "00000000-0000-0000-0000-d05099d333e0"
    namespace: default
  bmc:
    endpoint: "192.168.1.222"
    user: "admin"
    pass: "******"

```

Note above that the MetalMachine requires a serverRef that corresponds to a server that has been discovered by the Metal Controller Manager that is a corequisite.
