# ODH AI Edge Use Cases
Artifacts in support of ODH Edge use cases

| Components                           | Version |
|--------------------------------------|---------|
| OpenShift                            | 4.13    |
| Open Data Hub                        | 2.x     |
| Red Hat Advanced Cluster Management  | 2.8     |
| Tekton Pipelines                     | 2.8     |
| Quay Registry                        | 2.8     |


### Proof of Concept Edge use case with ACM

The main objective is to showcase that a user can take a trained model, use a pipeline to package it with all the dependencies and deploy it at the near edge location(s) in a centralized way.

1. Provision OpenShift Cluster
1. Configure the default Identity Provider
1. Install Red Hat Advanced Cluster Management
1. Register the clusters
   * Core (local-cluster)
   * Near Edge
1. Deploy Open Data Hub to the Core cluster
   1. Develop the model in a notebook
   1. Build the model using a pipeline
   1. Push the model to the image registry accessible by the near edge cluster
   1. Update the GitOps config for the near edge
