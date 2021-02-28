# Cortify
Container-as-a-service solution Built on top of Kubernetes and Knative



Roadmap :

* Manage Knative Control-Plane on different Tenants.
* Manage Multi-Tenancy with MongoDb
* Implement Encryption-as-a-Service for different runtimes.

* Add Built-in  observability and Tracing for Containers within a given env.

... 
Features :
    # Service Settings :
        
        - Deployment Platform :
            * Fully Managed
            * Sperate Tenant
        - ServiceName:


    # Configure Revision :
        * Deploy revision from container
        * Build and Deploy revision from Git.
        - Advanced Settings : 

            1-Container:    
            + initContainers + MultiContainers
                * General
                    ContainerPort: 8080
                    ContainerCommand: 
                    ContainerArgs:
                    ContainerServiceAccount:
                * Capacity
                    Memory Allocated: 256MiB
                    CPU Allocated: 1
                    RequestTimeout: 300
                    MaximumRequestsperCntainer: 80
                * Auto-Scaling:
                    Min: 0
                    Max: 100
            2-Variables:
                Environment variables:
                    Name: Environment
                    Value: Production
                    + Variable as  File
                Secrets: 
                    Secret Manager
                    Secret as File ( TLS certificat..)

            3- Connections:   
                Database Connections
                OAuth As a service

    # Configure how this service is triggered:
        Ingress:
            Internal
            PubliclyAccessible
        Authentication:
            Allow unauthenticated invocations
            RequireAuthentification (ex OAuth)
        Triggers:
            Eventing.
