# Book : Extending Kubernetes pp213 - pp217 example

```
1. make main.go in cmd folder
2. make lister.go in pkg folder to the registration and discovery of the device plugin with the extend-k8s.io/example 
                    it returns a PluginInterface that will implement in plugin.go
3. make plugin.go in pkg folder                    
4. make go.mod in root folder
5. make Dockerfile to have docker image of Device-Plugin 
6. docker build & push with reference to docker-build-cmd.txt
7. make DaemonSet definition (.yaml) to deploy the Device_Plugin to Kubernetes cluster
   To check the status of custom devices from the node status data
   kubectl get node <node name> -w -o json | jq '.status.allocatable."extend-k8s.io/example'
8. Let's make a pod to use the custom device with Pod deployment definition
   When the pod runnit, execute the following command to check environment variables
   $ kubectl exec device-plugion-cosumer -- env
```

