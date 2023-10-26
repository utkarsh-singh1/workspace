1. Get details about running pods/ rc/ rs/ deployment.

```bash
kubectl get pods/rs(replicaset)/rc(replicacontrol)/deployment
```   

Note:- Use `-o wide` flag for wide information about k8s object.
Ex. 
```bash
kubectl get pods -o wide
```

2. Get brief description about pods/ rc/ rs/ deployment.

```bash
kubectl describe pods/rs/rc/deployment object_name
``` 

3. 


