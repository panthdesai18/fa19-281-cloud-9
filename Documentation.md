# CMPE - 281 Team Hackathon Project
# The Counter Custom Burger
## Team Members:
* ### Panth Desai
* ### Udit Marolia
* ### Nivali Reddy
* ### Ramya Bandi
* ### Ayushi Gupta
<br/>

## Architecture Diagram:

## Description:

### The Counter Custom Burger:  
This is an interactive platform for ordering burgers. It has a login user who will first select the location of the outlet using zipcode, create an order, review the order by looking at the cart and then place the order. The admin can add and delete the items from the backend. No user iterface has been provided for the admin.
This platform will be very helpful to users to place orders.


### Frontend - User/Admin:  
User will interact with the appliction from frontend. Request from here will be directed to Amazon API Gateway. API gateway will interact with microservices.
**Technology**: React, HTML, CSS. 

<br/>

## AKF Scale Cube:

#### Figure:

### X-axis:
X-axis scaling consists of running multiple copies of an application behind a load balancer. We have 3 docker images for User
microservice running behind a load balancer.

### Y-axis:
Y-axis axis scaling splits the application into multiple, different services. We split our application functionality into
independent microservices.

### Z-axis:
Z-axis splits are commonly used to scale databases. Data is partitioned across a set of servers based on an attribute of 
each record. We had sharded mongo cluster to partition data.  

<br/>

### Load Balancer:
Load balancer is used to distribute traffic of User microservice to scale application horizontally.

### Amazon API Gateway:
Amazon API Gateway is used to redirect request from user to different microservices. It enables user to retrieve data from mutiple microservices within single round trip.
