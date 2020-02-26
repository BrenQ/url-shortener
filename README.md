# *URL SHORTENER*


Nuestro API va a permitir generar una nueva url corta . Esta url generada será un código alfanumérico de 6 dígitos 

Funcionalidad

* Un endpoint HTTP que reciba un URL como parámetro y la guarde asociada a un código de 6 caracteres 

* Un endpoint HTTP que reciba un código (el generado en el paso anterior) y redirija al
URL asociado

## RUN 

 ###Local environment:
 
** Requirements**
  - [Go 1.12](https://github.com/golang/go "Go 1.12")
  - [Mongo DB 4.2](https://github.com/mongodb/mongo "Mongo DB Last version")
  
  
 - Clone this project with
	`git clone https://github.com/BrenQ/url-shortener.git`
	
 - Add a .env.local file with the following variables:

	
	# env.local
    DB_HOST=127.0.0.1
    DB_USER=user
    DB_PASSWORD=password
    DB_PORT=27017
    DB_TIMEOUT=5
    DB_NAME=url
    
    # Server
    SERVER_ADDR=127.0.0.1
    SERVER_PORT=8080
    
    HOST=localhost

 - In mongo, you must create credentials as described in the .env.local file ( DB_USER and DB_PASSWORD)
 - Finally , exec `go run main.go` and test in `127.0.0.1:8080`

### Run with Docker (docker-compose)

If you have Docker and *docker-compose* you can run the following command:

`docker-compose up --build`

After that, in a local environment the application will run in the port
`:8000` in the`0.0.0.0`

## USAGE

###Create a short URL

  **URL**
    
    /links
    
  **Method:**
    
    `POST`
    
-   **URL Params**
    
    **Required:**
    
    `url=[alphanumeric]`
    
-   **Header:**
    
    `Content-Type : application/json`
    
	
-	**Body**

>     	
>     {
>     	"url": "https://www.google.com"
>     }

	
    
-   **Success Response:**
    
    -   **Code:**  200
			
		> {
		>    "link": "example.com/p7R0yE",
		>    "orig": "https://www.google.com",
		>    "short": "p7R0yE"
		> }

    OR
    
    -   **Code:**  404 Bad Request  
        **Content:**  `{ "message": "Invalid request" }`
		
		   **Code:**  404 Bad Request  
           **Content:**  `{ "message": "Url invalid" }`

		 **Code:**  422 Unprocessable entity  
        **Content:**  `{ "message": "Unable to process request. Try Again" }`


###Redirect with a Short URL

-   **URL**
    
    /:code
    
-   **Method:**
    
    `GET`
    
-   **URL Params**
    
    **Required:**
    
    `code=[integer]`
    
-   **Success Response:**
    
    -   **Code:**  301  
    
    OR
    
    -   **Code:**  404 Bad Request  
        **Content:**  `{ "message": "We could not process url info" }`

		 **Code:**  404 NOT FOUND  
        **Content:**  `{ "message": "Code not registered" }`
-

