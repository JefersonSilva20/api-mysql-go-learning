# Exemplo de API-REST com Go!
 Olá! Este é apenas um exemplo de uma API-REST utilizando MYSQL como database para realizar operações CRUD no Model de pessoas, construída na linguagem Go.
 ## Endpoints e Responses
 - **CREATE:** localhost:8090/create 
	 -  **Response Status Code:** 201
 - **UPDATE:** localhost:8090/update/{ID}
	 -  **Response Status Code:** 200
 - **GET-ALL:** localhost:8090/get
	 -   **Response Status Code:** 200
 - **DELETE:** localhost:8090/delete/{ID}
	 -  **Response Status Code:** 204
 - **GET-{ID}:** localhost:8090/get/{ID}
	 -  **Response Status Code:** 200

## Documents

Para facilitar existe uma Collection que você pode usar no Postman na pasta `documents\Postman_collection.json`

## Tecnologias utilizadas

 - GORM
 - net/http
 - gorilla/mux