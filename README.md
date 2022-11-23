# **Go-Server-JWT-Docker-Compose**
Go server to perform CRUD operations protected by JWT authorization. 

The projects uses Docker-Compose to orchestrate both server and database containers.

# 🚀 **Run the project**


Since the project uses Docker-Compose environment

```docker-compose up --build```

> ⚠️  **If you wish to reset the environment, execute:** `docker-compose down`

# 💻 **Technologies**

Technologies used in the project

+ Docker-Compose 3.4

+ Go 1.19

+ MySQL latest

# 🌐 **REST API**

The handlers have 10 seconds of time out and the database queries 3 seconds

## Get JWT

### Request

`GET /jwt`

> ⚠️  **This process is necessary to set HTTPOnly cookie for authorization**

> ⚠️  **It is necessary to pass in the header Auth=W3UH328TF674389ynqpqipfop6ih6GYUF67Rgfryeu4745cioyg487GCWQYOY732QYDXB98U42980DJI**

### Example

`http://localhost:7000/jwt`

<hr>

## Create user

### Request

`POST /users`

> ⚠️  **user and name fields are required**

### Example

`http://localhost:7000/users`

Body

```
{
  "user": "test",
  "name": "Nametest",
  "surname": "surnameTest"
}
```

Response

```
{
  "id": 36,
  "user": "test",
  "name": "Nametest",
  "surname": "surnameTest",
  "createdAt": "2022-11-23T13:43:26.944Z",
  "updatedAt": "2022-11-23T13:43:26.944Z"
}
```

<hr>

## Get user(s)

### Request

`GET /users?`

> ⚠️  **All data will be passed as query string parameters**

> ⚠️  **If no query string parameter is provided, the API will return all data stored**

### Example

`http://localhost:7000/users?name=Nametest&surname=surnameTest`

Response

```
[
  {
    "id": 36,
    "user": "test",
    "name": "Nametest",
    "surname": "surnameTest",
    "createdAt": "2022-11-23T13:43:27Z",
    "updatedAt": "2022-11-23T13:43:27Z"
  }
]
```

<hr>

## Update user

### Request

`PUT /users`

> ⚠️  **It is necessary to provide the id**

### Example

`http://localhost:7000/users`

Body

```
{
  "id": 36,
  "name": "new",
  "surname": "new surname"
}
```

Response

```
{
  "id": 36,
  "user": "test",
  "name": "new",
  "surname": "new surname",
  "createdAt": "2022-11-23T13:43:27Z",
  "updatedAt": "2022-11-23T14:03:14Z"
}
```

<hr>

## Delete user

### Request

`DELETE /users`

> ⚠️  **The response is 204 No Content**

### Example

`http://localhost:7000/users`

Body

```
{
  "id": 36
}
```