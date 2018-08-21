# TicketingApi

## Routes

#### Authentication

`POST` [/auth/signup](#sign-up)  
`GET` [/auth/login](#log-in)  

#### User management

`POST` [/users/create/support](#create-a-support-user)  
`POST` [/users/create/client](#create-a-client-user)  
`PATCH` [/users/edit/{id_user}](#edit-a-user)  

## Authentication

#### Sign up

Method: `POST`  
Endpoint: `/auth/signup`  

|Request | Response|
|--------|---------|
| **firstname :** User first name | **token :** Access token |
| **lastname :**  User last name | **user :** [User struct](#user-struct) |
| **email :** User email | |
| **password :**  User password | |

#### Log in

Method:   `GET`  
Endpoint: `/auth/login`  

| Request | Response |
|---------|----------|
| **email :** User email | **token :** Access token |
| **password :**  User password | **user :** [User struct](#user-struct) |  
  
## User management

#### User struct

| Fields | Description |
|--------|-------------|
| **id** | Unique id |
| **firstname** | First name |
| **lastname** | Last name |
| **type** | (Support / Client) |
| **email** | Email |

#### Create a support user

Method:   `POST`  
Endpoint: `/users/create/support`

*Only for support*

| Request | Response |
|---------|----------|
| **email :** User email | **user :** [User struct](#user-struct) |
| | **password :** User password |

#### Create a client user

Method:   `POST`  
Endpoint: `/users/create/client`

*Only for support*

| Request | Response |
|---------|----------|
| **email :** User email | **user :** [User struct](#user-struct) |  
| | **password :** User password |

#### Edit a user

Method:   `PATCH`  
Endpoint: `/users/edit/{id_user}` 

*A client can only modify itself*

| Request | Response |
|---------|----------|
| *(optional)* **firstname :** User first name | **user :** [User struct](#user-struct) |
| *(optional)* **lastname :**  User last name | |
| *(optional)* **email :** User email | |
| *(optional)* **type :** User type, *only for support* | |
| *(optional)* **password :**  User password | |

## Tickets

#### Ticket struct

| Fields | Description |
|--------|-------------|
| **id** | Unique id |
| **author** | [User struct](#user-struct) |
| **status** | (Open / Pending reply / Closed) |
| **messages** | Array of [Message struct](#message-struct) |

#### Create a ticket

Method:   `POST`  
Endpoint: `/tickets/create`  

| Request | Response |
|---------|----------|
| | |

Method:   `.`  
Endpoint: `.`  

| Request | Response |
|---------|----------|
| | |

Method:   `.`  
Endpoint: `.`  

| Request | Response |
|---------|----------|
| | |
