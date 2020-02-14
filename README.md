## *URL SHORTENER*


Nuestro API va a permitir generar una nueva url corta . Esta url generada será un código alfanumérico de 6 dígitos 

Funcionalidad

* Un endpoint HTTP que reciba un URL como parámetro y la guarde asociada a un código de 6 caracteres 

* Un endpoint HTTP que reciba un código (el generado en el paso anterior) y redirija al
URL asociado

## RUN 

 Local enviroment:
 
 - Clone this project 
 - Exec `go run main.go`

## USAGE

**Create a short URL**

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



**Redirect with a Short URL**

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

