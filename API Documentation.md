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
    - 200 created
        
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