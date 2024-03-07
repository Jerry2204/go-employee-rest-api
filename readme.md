# Employee Management API 

This document provides instructions on how to set up and run this Golang application using Docker. It also includes information on setting up a PostgreSQL container for your application.

## Prerequisites
- Git
- Docker

## Cloning the Repository
1. Run the following command to clone this project to your local machine:
   ```bash
   git clone https://github.com/Jerry2204/go-employee-rest-api
   ```
2. Navigate to the project folder:
   ```bash
   cd go-employee-rest-api
   ```
   


## Setting Up PostgreSQL Container

1. Copy the `.env.example` file to `.env` in the root of your project directory:

   ```bash
   cp .env.example .env
2. Open the `.env` file and set up the configuration for your PostgreSQL database. Replace the placeholders with your actual database credentials:

   ```bash
   DB_USER=your_postgres_user
   DB_PASSWORD=your_postgres_password
   DB_NAME=your_database_name

## Running This Golang Application

1. Run the following command to start the PostgreSQL container and this Golang application:

   ```bash
   docker-compose up --build

1. This Golang application should now be running inside a Docker container. You can access it at http://localhost:8080 in your web browser or you can access the API using your API Management Tools.

   ```bash
   docker-compose up --build

## Stopping This Golang Application
To stop and remove the PostgreSQL and Golang containers, use the following command:
```
docker compose down
``` 
#

# API for Management Employee Specification

## Create Employee
+ Endpoint : ``/api/employees``
+ HTTP Method : ``POST``
+ Request Header :
    + Accept : ``application/json``
+ Request Body :
```json
{
    "firstName"       : "Jerry Andrianto",  
    "lastName"        : "Pangaribuan",  
    "email"           : "jerryandrianto@example.com"
}
```

+ Response Body (Success) :
+ Response Code ``201``

```json
{
    "status"   : "OK",
    "message"  : "Data Created Successfully",
    "data"     : [
        {
            "firstName": "Jerry Andrianto",
            "lastName" : "Pangaribuan",
            "email"    : "jerryandrianto@example.com"
        }
    ]
}
```

+ Response Body (Fail) :
+ Response Code ``400``

```json
{
    "status"   : "error",
    "error"    : "Bad Request!"
}
```

+ Response Code ``500``

```json
{
    "status"   : "error",
    "error"    : "Internal Server Error!"
}
```

## Get All Employee

+ Endpoint : ``/api/employees``
+ HTTP Method : ``GET``
+ Request Header :
    + Accept : ``application/json``

+ Response Body (Success) :
+ Status Code ``200``

```json
{
    "status"   : "OK",
    "message"  : "Get Data Sucessfully!",
    "data"     : [
        {
            "ID"        : 28,
            "CreatedAt" : "2024-03-05T17:59:34.092438Z",
            "UpdatedAt" : "2024-03-06T14:21:06.891792Z",
            "DeletedAt" : null,
            "firstName" : "Irin Chandra",
            "lastName"  : "Pangaribuans",
            "email"     : "irin@example.com",
            "hireDate"  : "0001-01-01T00:00:00Z"
        },
        {
            "ID"        : 37,
            "CreatedAt" : "2024-03-06T09:16:11.37196Z",
            "UpdatedAt" : "2024-03-06T09:16:11.37196Z",
            "DeletedAt" : null,
            "firstName" : "Alendra Yudistira",
            "lastName"  : "Hutahaean",
            "email"     : "alen@example.com",
            "hireDate"  : "0001-01-01T00:00:00Z"
        },
        {
            "ID"        : 41,
            "CreatedAt" : "2024-03-06T10:04:44.27142Z",
            "UpdatedAt" : "2024-03-06T10:04:44.27142Z",
            "DeletedAt" : null,
            "firstName" : "Revalino Adhistana",
            "lastName"  : "Hutahaean",
            "email"     : "alen@example.com",
            "hireDate"  : "2024-03-06T10:04:44.270006Z"
        }
    ]
}
```

## Get Employee By Id

+ Endpoint : ``/api/employees/{employeeId}``
+ HTTP Method : ``GET``
+ Request Header :
    + Accept : ``application/json``

+ Response Body (Success) :
+ Status Code ``200``

```json
{
    "message" : "Get Data Sucessfully!",
    "status"  : "OK",
    "data"    : {
      "ID": 28,
      "CreatedAt": "2024-03-05T17:59:34.092438Z",
      "UpdatedAt": "2024-03-06T14:21:06.891792Z",
      "DeletedAt": null,
      "firstName": "Irin Chandra",
      "lastName": "Pangaribuans",
      "email": "irin@example.com",
      "hireDate": "0001-01-01T00:00:00Z"
    }
}
```

+ Response Body (Fail) :
+ Status Code ``404``

```json
{
    "status"   : "error",
    "message"  : "Data Not Found!"
}
```

## Update Employee By Id
+ Endpoint : ``/api/employees/{employeeId}``
+ HTTP Method : ``PUT``
+ Request Header :
    + Accept : ``application/json``
+ Request Body :
```json
{
    "firstName"       : "Jerry Andrianto",  
    "lastName"        : "Pangaribuan",  
    "email"           : "jerryandrianto@example.com"
}
```

+ Response Body (Success) :
+ Status Code ``200``

```json
{
  "status"  : "OK",
  "message" : "Data Updated Successfully!",
  "data"    : {
    "id": 28,
    "firstName": "Jerry Andrianto",
    "lastName": "Pangaribuan",
    "email": "jerryandrianto@example.com",
    "hireDate": "2024-03-06T10:04:44.270006Z"
  }
}
```

+ Response Body (Fail) :
+ Status Code ``404``

```json
{
    "status"   : "error",
    "message"  : "Data Not Found!"
}
```

+ Response Code ``400``

```json
{
    "status"   : "error",
    "error"    : "Bad Request!"
}
```

+ Response Code ``500``

```json
{
    "status"   : "error",
    "error"    : "Internal Server Error!"
}
```

## Delete Employee By Id
+ Endpoint : ``/api/employees/{employeeId}``
+ HTTP Method : ``DELETE``
+ Request Header :
    + Accept : ``application/json``
+ Response Body (Success) :
+ Status Code ``200``

```json
{
  "status"  : "OK",
  "message" : "Data Deleted Successfully!",
  "data"    : {
    "id": 28,
    "firstName": "Jerry Andrianto",
    "lastName": "Pangaribuan",
    "email": "jerryandrianto@example.com",
    "hireDate": "2024-03-06T10:04:44.270006Z"
  }
}
```

+ Response Body (Fail) :
+ Status Code ``404``

```json
{
    "status"   : "error",
    "message"  : "Data Not Found!"
}
```
+ Response Code ``500``

```json
{
    "status"   : "error",
    "error"    : "Internal Server Error!"
}
```

