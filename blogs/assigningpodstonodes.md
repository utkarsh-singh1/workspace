
Practical - Report:- 
<h1> Methods to Assign Pods to Nodes </h1>

```
Some of the terms to keep in remember if you are reading this guide - 

- Topology - Set of a worker nodes in an zone, rezions or any user defined terms. You can can also create in your local environment using any minikube, kubeadm or kind. topologykey is a key of label that is assigned to a node, Ex. If we are talkikg about zones a set of nodes gonna have similiar labels like zone: zone1, another set of nodes in other zone gonna have label zone2.

- maxSkew - Maximum differnce between no. of pods between 2  topology domain ( remember topology domain depends on level means on the level of node, zone or regions) , Ex - On level of zone suppose there are 3 zones z1, z2, z3 and nodes on z1 are 3, z2 are 2, z3 are 1 then maximum differnce = 2, which is between zone1 (max no. of nodes) and zone3 (min no. of nodes), so maxskew will be 2. 
```


<h2> Inter-pod-affinity / pod-affinity - </h2>

podAffinity depends on scheduled pods labels present on the node and also on labels on node or set of nodes also known as _topologykey_, these constraints will determine if current pod is gonna scheduled on the node or not. If the current pod already scheduled on the node has a similiar label as one on upcoming pod using podAffinity , the upcoming pod will schedule in the nodes with topologyKey mention in the upcoming Pods podspec. Now on the basis of if kube-scheduler wanna shedule or not these two methods will be most important -

1- requiredDuringSchedulingIgnoredDuringExecution - It is necessary to new pod to follow the rule of podAffinity or kube-scheduler will not schedule the new pod.

2- preferredDuringSchedulingIgnoredDuringExecution - kube-scheduler will try to follow the podAffinity rule but if required it will try to schedule pod.

Remember even if you can schedule the pod , it will always be schedule on a one topologydomain , New pod will always schedule on the any node among set of nodes on a single topologydomain. Conversely, if there are no Pods with similiar labels as pod scheduled in Zone , the scheduler will not assign the example Pod to any node in that zone.


<h2> Inter-pod-anti-affinity / pod-AntiAffinity - </h2>

Same as podaffinity, podAntiAffinity depends on scheduled pods labels present on the node and also on labels on node or set of nodes also known as _topologykey_, these constraints will determine if current pod is gonna scheduled on the node or not.But unlike podAffinity, podAntiAffinity will not schedule on that topologyDomain where these pods with similiar label as mention under labelSelector. Similarly like podAffinity ,Now on the basis of if kube-scheduler wanna shedule or not these two terms will be most important -

1- requiredDuringSchedulingIgnoredDuringExecution - It is necessary to new pod to follow the rule of podAffinity or kube-scheduler will not schedule the new pod.

2- preferredDuringSchedulingIgnoredDuringExecution - kube-scheduler will try to follow the podAffinity rule but if required it will try to schedule pod.

Conversely, the anti-affinity rule does not impact scheduling into Zone if there are no Pods with with similar labels are available.



> Note - If no node with label key mentioned under both podaffinity and podAntiAffinity is avialable in that no pod can scheduled, and pod will be in pending state.


<h3> podTopologySpread -  </h3>

Control how pods/ workloads are spread across cluster among topologydomain such as regions, zones, nodes, and other user-defined topology domains.

It is somewhat differnt from both podAffinity and podAntiAffinity, it can shedule on an empty nodes (topologykey mention under pods podspec should match the nodes) even if there is is labelSelector is mentioned under podspec, Unlike podaffinity where no matching pods can not a schedule a new pod, here labelSelector is used to find matching Pods. Pods that match this label selector are counted to determine the number of Pods in their corresponding topology domain. Unlike podAffinity all pods are not spread in a single topologyDomain, the pods are distributed across different topologyDomain, because here pods label are only used for counting number of pods in their topologyDomain and spread is moredue to maxSkew mentioned under podspec. 

maxSkew will determine the pod spread by - 

- if you select `whenUnsatisfiable: DoNotSchedule`, then maxSkew defines the maximum permitted difference between the number of matching pods in the target topology and the global minimum (the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains). For example, if you have 3 zones with 2, 2 and 1 matching pods respectively, MaxSkew is set to 1 then the global minimum is 1.

- if you select `whenUnsatisfiable: ScheduleAnyway`, the scheduler gives higher precedence to topologies that would help reduce the skew.

> whenUnsatisfiable: DoNotSchedule and whenUnsatisfiable: ScheduleAnyway are constraints of podTopologySpread.
