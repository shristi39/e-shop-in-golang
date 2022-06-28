### E-COMMERCE
E-commerce is also known as the electronic commerce.It is an  complete dynamic website,where user can buy and add the product after they logined .User should register first before they login.They should create the new account.We can add the product from the site and from the api as well.Customers access an online store to browse through and place orders for products or services via their own choices.It is business to customer.
### Why this site?
 People can get to connect direct  and one can find any necessary things within the site with a click.They can do shopping from the browser from anywhere where the internet is accesible.They doesnt have to interact with the people and this kind of  site  can be suitable in pandemic situations.
## Technology stacks
There is client and server side in our website where 
 
 1 .Client-side:Html,css,bootstrap

2 .Server-side:Golang



## Setup Instruction
We have develop this website using the golang language.you can get the go package in(https://go.dev/dl/). We have done it in the version go 1.16, but  haven't tested in any other version.

 We have  use the gorilla/mux packages for the router. Gorilla/mux basically implements a request router and dispatcher for matching incoming requests to their respective handler.you can get the package in (https://github.com/gorilla/mux).
 
 We have also used the token that is generated during the login,we used the jwt token ,we can get the package in (https://github.com/dgrijalva/jwt-go).

 We have also used the session in this site where the login info is saved after the registration for this we used the session packages that is available for golang the used package is (https://github.com/gin-contrib/sessions).


 We also have created the mock file for the testing the coverage of the packages and functions the package we used for creating the mock file is (https://github.com/golang/mock)

The GORM is fantastic ORM library for Golang, aims to be developer friendly. It is an ORM library for dealing with relational databases. This gorm library is developed on the top of database/sql package. We have used the gorm for the query and to  manipulate data from a database using an object-oriented paradigm(https://github.com/jinzhu/gorm)

We have also created docker image of this project.Docker is a platform for managing ,deploying and building the containerized applications.It comes with the command line interface using which we can do all the operations which the website provides.We can get the information on building the image on(https://docs.docker.com/)

Steps for dockerizing the project to the docker images are-

1. Create the docker file outside the controllers and models.
2. Inside that dockerfile write code(where we include 

   FROM golang:1.16 

    WORKDIR /app 

    ADD . .

   RUN go build -o main .

   CMD /app/main
) and save that file

3. In the terminal run some  docker command .

In the terminal run the following commands to create the docker image .
 
 step 1-  docker build -t (image name) .
 
 step 2-docker image ls
 
 step 3- docker run -p (port name here ....:....) -it (image name)


## Running Instruction

When we finish setting the required package and instruction after that we can run the code by (go run main.go)or 
we can run this website in terminal through the command [ docker run -p (port) -it (image name)]
 
 ## Features and  Api
 ### User Api
| Api  |Description 
|---|---
|r.GET"/register"  |When there comes the new user  at first they register during this api wil call and registration form will appear 
|r.POST"/register"|After the registration there comes the data in the POST form and this api will call which checks either the name,email,password match the given security conditions if it matches then it will redirect to the home page or it will show the unsuccesful message.
|r.POST"/user/login"|This api will be called after login and by calling this api there will generate the user token. 
|r.GET"/login" |When the existing and registered user signin this api will called and user get logined if the login info is incorrect then their appear the unsuccesful message.
|r.GET"/logout"  |This api will call when the existing user will logout from  the site  


### Product Api

| Api   |Description   |   
|---|---
|r.POST"/create"   | This api will called to create the product from the postman | 
|r.DELETE"/Delete"   | When we have to delete the product we call this api in order to d delete the product.we delete the product from the postman by providing the product details. |   
|r.Get"/checkout| This api is called when the user checksout from the cart 
|r.GET"/poduct/:id"|To view the product details from the product id.when the user click the view product then this api will respond according and show the product details.
|r.GET"/addproduct"|When the user logedin then there comes the add sign from which user can add the product accordingly,when the user submit the this api will called and the product will be addedand shown in the home page





## Preview

Home page
![home page](/ecommerce_image/home.png)
mycart page
![cart page](/ecommerce_image/mycart.png)
register page
![register page](/ecommerce_image/register.png)
login page
![login page](/ecommerce_image/login.png)


## Test coverage

Models  test 
![models test](/ecommerce_image/modelstest.png)
Controllers test
![controllers test](/ecommerce_image/contrtest.png)
