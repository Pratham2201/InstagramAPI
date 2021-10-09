# InstagramAPI
Created from Scratch using Mongodb and go Standard library

Features:

Reusable -> The http requests are divided into functions to provide resusability

Naming Standards ->
  User Attributes Prefix (Struct) : U_
  Post Attributes Prefix (Struct) : P_
  Global variables : G_
  Other variables are to be considered as local variables
  
Server Thread Safe -> Mutex Locks are used to create server thread safe

Unit Test -> Unit Tests are created for each function
  
MongoDb Collection (Create the MongoDb collection named 'Test' with collection 'User' and 'Post' on localhost/27017) ->
  Users :
    Id
    Name
    Email
    Password!

  Posts:
    Id
    Caption
    Image URL
    Posted Timestamp
* All the attributes are included in the collection

# All the Endpoints ->

* POST -- /users
![Screenshot (55)](https://user-images.githubusercontent.com/55658008/136670287-2c2009af-247f-479b-bcd4-2fd91c3b6c22.png)

* GET -- /user/<user_id>
![Screenshot (56)](https://user-images.githubusercontent.com/55658008/136670328-6ff63e10-1b94-4f2c-ab4c-80747a0f3071.png)

* POST -- /posts
![Screenshot (57)](https://user-images.githubusercontent.com/55658008/136670347-61bcf87b-9a25-4884-baee-e47365ad6aa0.png)

* GET -- /posts/<post_id>
![Screenshot (58)](https://user-images.githubusercontent.com/55658008/136670383-5a485898-0f4a-4c30-b5e7-adeeb06139ab.png)

* GET -- /posts/users/<user_id>
![Screenshot (55)](https://user-images.githubusercontent.com/55658008/136670391-f15a7981-d042-4a53-a033-e482de5a2941.png)
