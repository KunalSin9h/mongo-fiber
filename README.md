![banner](https://tiddi.kunalsin9h.com/CdUx188)

## Project Documentation

Important Points

1. The main executable is `./cmd/main.go`

So, to run the project, run the following command:

```bash
go run ./cmd/*
```

This will start the server at port ":5000"

so the url will be `http://localhost:5000`

2. We can set some environment variables to change the behaviour of the server

The default mongo db connection is `mongodb://localhost:27017`

To change this we can set the environment variable `MONGODB_URI`

#### In Linux

```bash
export MONGODB_URI="mongodb://localhost:27017"
```

```bash
go run ./cmd/*
```

#### In Windows

```bash
set MONGODB_URI="mongodb://localhost:27017"
```

```bash
go run ./cmd/*
```

# API Documentation

## 1. Login

#### POST /api/login

with body

```json
{
  "email": "admin@stationery.shop",
  "password": "admin"
}
```

This will return `access_token` which is an `JWT` token

return body

```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ"
  }
}
```

## 2. Register New User

Only an owner of the shop can register new user

Need to pass the `access_token` in the query

#### POST /api/register?access_token={access_token}

with body

```json
{
  "image": "https://www.gravatar.com/av",
  "name": "New Staff",
  "email": "staff@stationery.shop",
  "phone": "111111111",
  "password": "staff",
  "role": "staff"
}
```

this will return the newly created user's id

```json
{
  "data": {
    "user_id": "647c9a59806696b8adc4e266"
  }
}
```

## 3. Add Inventory

Need to pass the `access_token` in the query

#### POST /api/inventory?access_token={access_token}

with body

```json
{
  "image": "https://img.prodcut.com",
  "name": "Drawing Book",
  "description": "Colored drawing book",
  "category": "drawing",
  "subCategory": "book",
  "quantity": 10,
  "price": 150.0,
  "supplier": "classmate",
  "min_stock": 100
}
```

this will return the newly created inventory's id

```json
{
  "data": {
    "inventory_id": "647c9a59806696b8adc4e266"
  }
}
```

## 4. Get Inventory

Need to pass the `access_token` in the query

#### GET /api/inventory?access_token={access_token}

This can get the entire inventory or you can specify multiple query
to specify the inventory you want to get

**Query**

1. `category` - to get inventory of a specific category

   - GET /api/inventory?access_token={access_token}&category={stationery}

2. `sub_category` - to get inventory of a specific sub category

   - GET /api/inventory?access_token={access_token}&sub_category={book}

3. `name` - to get inventory of a specific name

   - GET /api/inventory?access_token={access_token}&name={Drawing Book}

4. `supplier` - to get inventory of a specific supplier
   - GET /api/inventory?access_token={access_token}&supplier={classmate}

## 5. Add Sales

Need to pass the `access_token` in the query

#### POST /api/sales?access_token={access_token}

with body

```json
{
  "customer_name": "John Doe",
  "customer_email": "john@doe.com",
  "customer_phone": "1234567890",
  "product": "Pentonic Blue Dot Pen",
  "quantity": 5,
  "price": 10.5
}
```

this will return the newly created sales's id

```json
{
  "data": {
    "sales_id": "647c9a59806696b8adc4e266"
  }
}
```

This will update the inventory quantity

## 6. Get Sales

Need to pass the `access_token` in the query

#### GET /api/sales?access_token={access_token}

this will return list of all sales with date and time

# Testing

For testing an dummy Owner is created with email `admin@stationery.shop` and password `admin`

You can open `postman` and start testing by login with this user

The **exported** `Postman` API Collection is in the `postman` folder

You can import this file in `postman` and start testing the api's

---
