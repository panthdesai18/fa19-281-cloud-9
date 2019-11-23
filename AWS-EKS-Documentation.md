## Amazon  Elastic Kubernetes Service(EKS) Documentation:

* I have used Amazon's Elastic Kubernetes Service to deploy my order backend.

### Refrence Link: https://docs.aws.amazon.com/eks/latest/userguide/getting-started.html

### Below are the steps for EKS deployment:
      Amazon EKS using eksctl

* #### Install the latest AWS CLI:
      
      pip install awscli --upgrade --user

* #### Configure Your AWS CLI Credentials
      
      $ aws configure
      AWS Access Key ID [None]: <!your access key>
      AWS Secret Access Key [None]: <!your secret key>
      Default region name [None]: us-west-2
      Default output format [None]: json

    Install the Weaveworks Homebrew taps
      brew tap weaveworks/tap
    Install or upgrade eksctl
      brew install weaveworks/tap/eksctl
      brew upgrade eksctl && brew link --overwrite eksctl

* #### Create Your Amazon EKS Cluster and Worker Nodes
      
      eksctl create cluster \
      --name prod \
      --version 1.14 \
      --region us-west-2 \
      --nodegroup-name standard-workers \
      --node-type t3.medium \
      --nodes 3 \
      --nodes-min 1 \
      --nodes-max 4 \
      --managed

    It takes about 10-15 mins to create the cluster. Test the cluster after creation
      kubectl get svc

    Output:
      NAME             TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
      svc/kubernetes   ClusterIP   10.100.0.1   <none>        443/TCP   1m     
 
* #### for cluster

* Create .yaml file an add the following code for cluster and node groups


    apiVersion: eksctl.io/v1alpha5
        kind: ClusterConfig
        metadata:
    name: <name of your cluster>
        region: us-east-1
    vpc:
    id: "vpc-06c0e6a617c62ccf6"
    cidr: "10.0.0.0/16"
    subnets:
    private:
      us-east-1a:
        id: "subnet-05601ffdcf9fd8b8f"
        cidr: "10.0.4.0/24"
      us-east-1b:
        id: "subnet-059b98ddc9832f9fd"
        cidr: "10.0.1.0/24"
    public:
      us-east-1a:
        id: "subnet-04721aae9616a5b1a"
        cidr: "10.0.2.0/24"
      us-east-1b:
        id: "subnet-0b3365b32ec3dfbeb"
        cidr: "10.0.0.0/24"

* #### for nodeGroups

    name: <your worker node names>
    labels: { role: workers}
    instanceType: t2.micro
    desiredCapacity: 1
    minSize: 1
    maxSize: 3
    volumeSize: 50
    volumeType: gp2
    privateNetworking: true
    availabilityZones: ["us-east-1a"]
    ssh:
      allow: true
      publicKeyName: <your-public-key>
    tags:
      'service': 'order'

eksctl create cluster -f <! your yaml file name>

