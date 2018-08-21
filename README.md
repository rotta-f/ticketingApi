# TicketingApi

## Routes

#### Authentication

`POST` [/auth/signup](#sign-up)

## Authentication

#### Sign up

Method: `POST`  
Endpoint: `/auth/signup`  

|Request | Response|
|--------|---------|
| **firstname :** User first name | **token :** Access token |
| **lastname :**  User last name | |
| **email :** User email | |
| **password :**  User password | |

#### Log in

Method:   `GET`  
Endpoint: `/auth/login`  

| Request | Response |
|---------|----------|
| **email :** User email | **token :** Access token |
| **password :** User password | |

  
  
## User management

#### Create a support user

Method:   `POST`  
Endpoint: `/users/create/support`

*Only for support*

| Request | Response |
|---------|----------|
| **email :** User email | **firstname :** User first name |
| | **lastname :** User last name |
| | **email :** User email |
| | **password :** User password |

#### Create a client user

Method:   `POST`  
Endpoint: `/users/create/client`

*Only for support*

| Request | Response |
|---------|----------|
| **email :** User email | **firstname :** User first name |
| | **lastname :** User last name |
| | **email :** User email |
| | **type :** User type |
| | **password :** User password |

#### Edit a user

Method:   `PATCH`  
Endpoint: `/users/edit/{id_user}` 

*A client can only modify itself*

| Request | Response |
|---------|----------|
| *(optional)* **firstname :** User first name | **firstname :** User first name |
| *(optional)* **lastname :**  User last name | **lastname :**  User last name |
| *(optional)* **email :** User email | **email :** User email |
| *(optional)* **type :** User type, *only for support* | **type :** User type |
| *(optional)* **password :**  User password | |

Method:   `.`  
Endpoint: `.`  

| Request | Response |
|---------|----------|
| | |
