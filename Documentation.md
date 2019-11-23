# CMPE - 281 Team Hackathon Project
# Burger Point
## Team Members:
* ### Panth Desai
* ### Udit Marolia
* ### Nivali Reddy
* ### Ramya Bandi
* ### Ayushi Gupta
<br/>

## Architecture Diagram:

## Description:

### BurgerPoint:  
BurgerPoint is an interactive platform for ordering burgers. It has a login user who will create an order and then he can review the order by looking at the cart and then place the order. The user can also view available outlets based on their zipcode. The admin can add and delete the items from the backend. No user iterface has been provided for it.
The BurgerPoint will be very helpful to users to place orders.


### Frontend - User/Admin:  
User will interact with appliction from frontend. Request from here will be directed to Amazon API Gateway. API gateway will
interact with microservices.
**Technology**: React, HTML, CSS. 

<br/>

## AFK Scale Cube:

#### Figure:

### X-axis:
X-axis scaling consists of running multiple copies of an application behind a load balancer. We have 3 docker images for User
microservice running behind a load balancer. Quiz, Assignment and Admin have 3 replica sets of kubernetes containers.

### Y-axis:
Y-axis axis scaling splits the application into multiple, different services. We have split our application functionality into
independent microservices.

### Z-axis:
Z-axis splits are commonly used to scale databases. Data is partitioned across a set of servers based on an attribute of 
each record. We had shared mongo cluster to partition data.  

<br/>

### Load Balancer:
Load balancer is used to distribute traffic of User microservice to scale application horizontally.

### Amazon API Gateway:
Amazon API Gateway is used to redirect request from user to different microservices. It enables user to retrieve data from mutiple microservices within single round trip.

### GoAPI: 
GOAPI works as microservies. 
  * Login/signup API- normal user can signup/login and access the application.
  * Location API- to select the available outlets of the restaurant based on zipcode.
  * Menu API- The normal user can use the API to get MENU Items while admin user can add, delete menu items as well.
  * Order API- The normal user can use this api to place any order as well as fetch it while admin user can view all the         orders and even delete a past order 
  * Payments API- 

  ## Team Contribution:

### Nivali Reddy:
* GoAPI for Login service.
* Frontend for signup and login.
* Deployed user microservice in 3 docker containers.
* Mongo Cluster with master and slave nodes.
* Schema Planning 
 
### Ramya Bandi:
* GoAPI for Location Module.
* Frontend for location page.
* Mongo Cluster with master and slave nodes.
* Deployed microservices in docker container.
* S3 bucket integration for storing and retrieving PDF.
* Schema Planning

### Panth Desai:
* GoAPI for Menu functionality.
* Developed frontend menu page and past orders design.
* Deployed microservices in docker container.
* Deployed login microservice in Amazon BeanStalk. 
* Deployed location microservice in Amazon BeanStalk
* Menu deployment in Amazon Elastic Container Service.
* Mongo Shard Cluster.
* Schema Planning.

### Udit Marolia:
* GoAPI for Order Module.
* Frontend for Order.
* Deployed microservices in docker container.
* Created Amazon API Gateway.
* Deployed Order Module in Elastic Kubernetes Service.
* Mongo Shard Cluster.
* Schema Planning.
* Deploying Front end on Heroku.

### Aayushi Gupta:
* API for Payments service
* Mongo Cluster with master and slave nodes
