# TicketingApi

## Routes

#### Authentication

`POST` [/auth/signup](#sign-up)  
`GET` [/auth/login](#log-in)  

#### User management

`POST` [/users/create/support](#create-a-support-user)  
`POST` [/users/create/client](#create-a-client-user)  
`PATCH` [/users/edit/{id_user}](#edit-a-user)  

#### Tickets
`POST` [/tickets/create](#create-a-ticket)  
`GET` [/tickets/{id}](#get-a-ticket)  
`GET` [/tickets](#get-a-list-of-tickets)  
`PATCH` [/tickets/{id}/edit](#edit-a-ticket)  
`POST` [/tickets/{id}/close](#close-a-ticket)  
`POST` [/tickets/{id}/archive](#archive-a-ticket)  

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
| **title** | Title |
| **author** | [User struct](#user-struct) |
| **status** | (Open / Pending reply / Closed) |
| *(optional)* **messages** | Array of [Message struct](#message-struct) |

#### Create a ticket

Method:   `POST`  
Endpoint: `/tickets/create`  

| Request | Response |
|---------|----------|
| **title :** Ticket name | **ticket :** [Ticket struct](#ticket-struct) |
| **message :** Ticket Description | |

#### Get a ticket

Method:   `GET`  
Endpoint: `/tickets/{id}`  

| Request | Response |
|---------|----------|
| | **ticket :** [Ticket struct](#ticket-struct) |

#### Get a list of tickets

Method:   `GET`  
Endpoint: `/tickets` 

URI Parameters:  
*(optional)* **user :** Id of the author

| Request | Response |
|---------|----------|
| | **tickets :** Array of [Ticket struct](#ticket-struct) |

#### Edit a ticket

Method:   `PATCH`  
Endpoint: `/tickets/{id}/edit`  

*A client can only modify a ticket created by himself*

| Request | Response |
|---------|----------|
| **title :** Title | **ticket :** [Ticket struct](#ticket-struct) |

### Close a ticket

Method:   `POST`  
Endpoint: `/tickets/{id}/close`  

*Only for support*

| Request | Response |
|---------|----------|
| | **ticket :** [Ticket struct](#ticket-struct) |

#### Archive a ticket

Method:   `POST`  
Endpoint: `/tickets/{id}/archive`  

| Request | Response |
|---------|----------|
| | **ticket :** [Ticket struct](#ticket-struct) |

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

Method:   `.`  
Endpoint: `.`  

| Request | Response |
|---------|----------|
| | |
