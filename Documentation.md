# CMPE - 281 Team Hackathon Project
# QuizzBox
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
BurgerPoint is an interactive platform for ordering burgers. It has a login user who will create a order and then he can review the order by 
looking at the cart and then place the order. The admin can add and delete the items from the backend. No user iterface has been provided for it.
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

## Team Contribution:
