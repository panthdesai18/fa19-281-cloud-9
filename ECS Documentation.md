## Amazon  Elastic Container Service(ECS) Documentation:

* I have used Elastic Container Service to deploy my menu backend.Amazon Elastic Container Service (Amazon ECS) is a highly scalable, fast, container management service that makes it easy to run, stop, and manage Docker containers on a cluster.

### Refrence Link: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/Welcome.html

### Below are the steps for ECS deployment:

* #### Login to AWS account:
      select the ECS service

* #### Select the ECS Deployment Types:

      I have selected the Rolling Update which is managed by AWS.
      
* #### Creating a Service:

      Configure Basic Parameters
      select the region that your cluster is in
      On the Task Definition name page, select the revision of the task definition from which to create your service.
      Review the task definition, and choose Actions, Create Service.
      On the Configure service page, fill out the following parameters accordingly:
        Launch Type: Amazon EC2 container
        Platform version: any version of the service
        Cluster:
        Service name:
        Service type: REPLICA
        Number of Tasks: 1
        
* #### Configure a Network:

      Cluster VPC: CMPE281
      Subnets: public subnets
      Security Groups: one with all ports opened
      Auto-assign public IP: yes

      
* #### (Optional) Configuring Your Service to Use Service Discovery:

            
* #### (Optional) Configuring Your Service to Use Service Auto Scaling:

      
* #### Review and Create your service:


      
