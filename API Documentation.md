## API Documentation

* ### Menu Service:

* #### POST/CreateMenuEndpoint
        
        POST Api for adding menu items
        Accept: application/json
        
        Body:
        {
       "ItemId": "2",
       "ItemName": "CheeseBurger",
       "Price": "06",
       "Description": "A cheeseburger is a hamburger topped with cheese. Traditionally, the slice of cheese is placed on top of the meat patty.",
       "ItemType": "Main"
       }

 * #### DELETE/RemoveItemEndpoint{id}
        
        Delete Api for removing the menu item
        Accept: application/json
        
        Response:
        - Item Deleted 
        
 * #### GET/GetMenuItemsEndpoint
        
        GET Api for viewing all the Menu Items
        Accept: application/json
        
        Response:
        {
        "_id": "5dd648039a554828c30a11ae",
        "ItemId": "01",
        "ItemName": "Burgers",
        "Price": "07",
        "Description": "burger it is",
        "ItemType": "main_dish"
    }
        
 * #### GET/GetMenuItemEndpoint/{id}
        
        GET Api for viewing a particular Menu Item
        Accept: application/json
        
        Response:
        {
        "_id": "5dd648039a554828c30a11ae",
        "ItemId": "01",
        "ItemName": "Burgers",
        "Price": "07",
        "Description": "burger it is",
        "ItemType": "main_dish"
    }

* ### Location Service:

* #### Post/locations
        
        Post API for locations
        Accept: application/json
        
        Body:
        {
	    "locationId"	 : 	"34544",
	    "locationName"	:	"San Jose",		
	    "address"	:	“1 south",
	    "zipcode"	:	"12342",
	    "city"		:	"san jose",
	    "state"		:	"California",
	    "country"	:	"US",
	    "latitude"	:	"1234567",
	    "longitude"	:	"7654321",
	    "phone"		:	"1234567890",
	    "email"		:	"abc@gmail.com"
        }

        Response:
        {
        "error": "",
        "result": "Location Insertion Successful"
        }
        - 200 created
        - 400 Invalid request

 * #### Get/locations
        
        GET All Locations 
        Accept: application/json
        
        Response:
         [
         {
         "locationId": "345689",
         "locationName": "Santa Clara",
         "address": "4 south",
         "zipcode": "12342",
         "city": "san jose",
         "state": "California",
         "country": "US",
         "latitude": "1234567",
         "longitude": "7654321",
         "phone": "1234567890",
         "email": "abc@gmail.com"
         },
        
 * #### Get/locations/{locationId}
        
        GET A Location by LocationID
        Accept: application/json
        
        Response:
        {
        locationId": "345689",
        "locationName": "Santa Clara",
        "address": "4 south",
        "zipcode": "12342",
        "city": "san jose",
        "state": "California",
        "country": "US",
        "latitude": "1234567",
        "longitude": "7654321",
        "phone": "1234567890",
        "email": "abc@gmail.com"
        }

* ### Login Service:

* #### Post/signup
        
        Post API for signup
        Accept: application/json
        
        Body:
        {
	    "username" : "nivali",
        "fullname" : "Nivali",
        "emailid" : "nivali@gmail.com",
        "password" : "Abc@123"
        }

        Response:
        {
        "error": "",
        "result": "Signup Successful"
        }
        - 200 created
        - 400 Invalid request

 * #### Post/login
        
        Post API for login
        Accept: application/json
        
        Body:
        {
        “username” : “nivali”,
        “Password” : “Abc@123”
        }
        Response:
        {
        "username": "nivali",
        "fullname": "Nivali",
        "emailid": "nivali@gmail.com",
        "password": ""
        }
        or
        {
        "error": "Invalid password",
        "result": ""
        }

        - 200 created
        - 400 Invalid Request

* #### Get/users
        
        Get API for users
        Accept: application/json
        
        Response:
        [
        {
        "username": "ramyareddy",
        "fullname": "Ramya Reddy",
        "emailid": "ramyareddy@gmail.com",
    "password": "$2a$05$1c5jkWJT1qw5jc8Shnx81.i1mHKsfAEAMU2LMjqwSbRoNFRka0xVO"
        },
        {
        "username": "dhanasreeare",
        "fullname": "Dhanasree Are",
        "emailid": "dhanasreeare@gmail.com",
        "password": "$2a$05$Ht41/       sMHKhNDC6.cgYBL0eleRtYiYGD3p5UU0Rz0z6reEkGm7DCJa"
        },
        {    
        "username": "nivalireddy",
        "fullname": "Nivali Reddy",
        "emailid": "nivalireddy@gmail.com",
    "password": "$2a$05$NHkyYKCivJT2XvOJ080bPeuCxxecsnkSSW7yCS6zeryEd/1JBpy0W"
        }
        ]

        
 * #### Get API for user/{username}

        
        GET A Location by user/{username}
        Accept: application/json
        
        Response:
        {
        "username": "ramya",
        "fullname": "Ramya",
        "emailid": "ramya@gmail.com",
        "password": ""
        }

* #### Get API for user/{email}
  Accept: application/json

        Response:
        {
        "username": "ramya",
        "fullname": "Ramya",
        "emailid": "ramya@gmail.com",
        "password": ""
        }

* #### Delete API for user/{username}
  Accept: application/json
  
        Response:
        {"error":"","result":"User deleted successfully"}

* ### Order Service:
placeOrder/Post
orders/GET
getUserOrder/GET
removeOrder/DELETE

* #### POST/placeOrder
        
        POST Api for confirming the order items from cart
        Accept: application/json
        
        Body:
        {
        "UserId": "raj",
        "OrderStatus": "placed",
        "Items": [
               "CheeseBurger"
        ],
        "TotalAmount": "6"
        }
        -200 created

 * #### GET/orders
        
        Get All orders of all the users
        Accept: application/json
        
        Response:
        {
        "_id": "5dd8cab8a7e8aa8a540e0d6",
        "UserId": "raj",
        "OrderStatus": "placed",
        "Items": [
               "CheeseBurger"
        ],
        "TotalAmount": "6"
        }

        - 200 okay
        
 * #### GET/getUserOrder/{id}
        
        get orders of a particular user
        Accept: application/json
        
        Response:
        {
        "_id": "5dd8cab8a7e8aa8a540e0d6",
        "UserId": "raj",
        "OrderStatus": "placed",
        "Items": [
               "CheeseBurger"
        ],
        "TotalAmount": "6"
        }

        -200 okay
        
 * #### Delete/removeOrder/{id}
        
        delete api for according to the user id
        Accept: application/json
        
        Response:
        -200 deleted
