# E-COMMERCE
It is an dynamic website,where user can buy and add the product.This site mainly consists of the server,client and database where,
- Server-side : Golang
- Client-side : html, css,bootstrap 
- Database: Sqlite

## Application Requirement
For the developement of this website we need certain package which should be be installed.For that we should get those package and run the command go get (package url), and then go mod init and finally go mod tidy. The packages we need are listed below-

1. For  Golang https://go.dev/dl/
2. For  Gorilla/mux package for router ```go get github.com/gorilla/mux```
3. Gorm package for connecting to database ```go get github.com/jinzhu/gorm```
 4. For sqlite ```go get github.com/jinzhu/gorm/dialects/sqlite```
4. For token jwt ```github.com/dgrijalva/jwt-go```
5. For tests ```go get github.com/golang/mock/gomock```
6. For session ```github.com/gin-contrib/sessions```


We have also created docker image of this project.Docker is a platform for managing ,deploying and building the containerized applications.It comes with the command line interface using which we can do all the operations which the website provides.

Steps for dockerizing the project to the docker images are-

1. Create the docker file outside the controllers and models.
2. Inside that dockerfile write code and save that file
3. In the terminal run some  docker command .

        In the terminal run the following commands to create the docker image .
        step 1-  docker build -t (image name) .
        step 2-docker image ls
        step 3- docker run -p (port name here ....:....) -it (image name)

    
 we can run this website in terminal through the command [go run main.go] or we can also do [ docker run -p (port) -it (image name)]
 
 ## Features and  Api
 ### User Api
| Api  |Description 
|---|---
|Register   |To register new user   
|Loginuser|To generate the token right after login
|Login   |To login existing user   
|Logout   |for logout   


### Product Api

| Api   |Description   |   
|---|---
|Create   | To create product  | 
|Delete   | It is for deleteing product  |   
|Update   | To update  |   
|Checkout|To check out from cart
|Product|To view the product details
|Addproduct|To add product from the user side





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