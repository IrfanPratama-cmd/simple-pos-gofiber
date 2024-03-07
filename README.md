## Developer
Muhammad Irfan Pratama

## Summary
This is a simple POS (Point Of Sale) project REST API integrated with xendit payment gateway (https://www.xendit.co/id/). The features available in this project are login, register, email verification using Gmail SMTP, and master data, features for users include displaying product lists, transaction, and payment gateway

## Database 

### ERD
https://drive.google.com/file/d/1ijdO7np8shaq7wQgEEWxZQ2lQPBP4s6-/view?usp=sharing

### Master Data

```sql
CREATE TABLE [category](
[id] [varchar](36) NOT NULL,
[category_code] [nvarchar](255) NOT NULL,
[category_name] [nvarchar](255) NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
[deleted_at] [datetime] NULL,
)
```

```sql
CREATE TABLE [brand](
[id] [varchar](36) NOT NULL,
[brand_code] [nvarchar](255) NOT NULL,
[brand_name] [nvarchar](255) NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
[deleted_at] [datetime] NULL,
)
```

```sql
CREATE TABLE [product](
[id] [varchar](36) NOT NULL,
[brand_id] [varchar](36)NOT NULL,
[category_id] [varchar](36)NOT NULL,
[product_code] [nvarchar](255) NOT NULL,
[product_name] [nvarchar](255)NOT NULL,
[quantity] [integer](11) NOT NULL,
[price] [decimal](22) NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
[deleted_at] [datetime] NULL,
)
```

```sql
CREATE TABLE [product_asset](
[id] [varchar](36) NOT NULL,
[product_id] [varchar](36)NOT NULL,
[file_name] [nvarchar](255) NOT NULL,
[file_path] [nvarchar](255) NOT NULL,
[is_primary] [bool] NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
[deleted_at] [datetime] NULL,
)
```

## User Authentication

```sql
CREATE TABLE [users](
[id] [varchar](36) NOT NULL,
[role_id] [varchar](36)NOT NULL,
[username] [nvarchar](255) NOT NULL,
[email] [nvarchar](256) NOT NULL,
[password] [nvarchar](256) NOT NULL,
[activated_at] [datetime] NULL,
[verification_code] [nvarchar](256) NOT NULL,
[verification_expiration] [datetime] NULL,
[is_activated] [bool] NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
[deleted_at] [datetime] NULL,
)
```

```sql
CREATE TABLE [customers](
[id] [varchar](36) NOT NULL,
[user_id] [varchar](36)NOT NULL,
[customer_name] [nvarchar](255)  NULL,
[mobile] [nvarchar](256)  NULL,
[alternate_number] [nvarchar](256)  NULL,
[email] [nvarchar](256)  NULL,
[website] [nvarchar](256)  NULL,
[address] [nvarchar](256)  NULL,
[zip_code] [nvarchar](256) NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
[deleted_at] [datetime] NULL,
)
```
## Transaction

```sql
CREATE TABLE [carts](
[id] [varchar](36) NOT NULL,
[customer_id] [varchar](36)NOT NULL,
[product_id] [varchar](36)NOT NULL,
[qty] [integer](11) NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
[deleted_at] [datetime] NULL,
)
```

```sql
CREATE TABLE [transactions](
[id] [varchar](36) NOT NULL,
[user_id] [varchar](36)NOT NULL,
[customer_id] [varchar](36)NOT NULL,
[transaction_date] [datetime] NOT NULL,
[invoice_no] [nvarchar](256)  NULL,
[transaction_status] [enum('cancel','pending','paid')]  NULL ,
[transaction_type] [enum('cash','cashless')]  NULL ,
[total_amount] [decimal](22) NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
[deleted_at] [datetime] NULL,
)
```

```sql
CREATE TABLE [transaction_details](
[id] [varchar](36) NOT NULL,
[customer_id] [varchar](36)NOT NULL,
[product_id] [varchar](36)NOT NULL,
[transaction_id] [varchar](36)NOT NULL,
[qty] [integer](11) NOT NULL,
[amount] [decimal](22) NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
)
```

```sql
CREATE TABLE [payment](
[id] [varchar](36) NOT NULL,
[checkout_link] [nvarchar](256)  NULL,
[external_id] [varchar](36)NOT NULL,
[payment_status] [enum('cancel','pending','settled')]  NULL ,
[amount] [decimal](22) NOT NULL,
[created_at] [datetime] NOT NULL,
[updated_at] [datetime] NULL,
)
```

## G-Mail SMTP Setting env
```
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=app-password
SMTP_HOST=smtp.gmail.com
SMT_PORT=587
```


## Install and run the application on a local development environment
This project uses docker to run it

```
docker-compose up -d
```

```
docker-compose exec go go run .
```

Basepath : /api/v1

## Documentation API

### Postman
Postman link : https://drive.google.com/file/d/1o25Yfgi_5TZEk9KapbHHloWNvuJBIpR8/view?usp=sharing

### Swagger
```
swag init
```

Access local swagger ui documentation after running docker : http://localhost:8080/

### Filter, Sorting, and Pagination Library

Click this github link to see details and examples https://github.com/morkid/paginate

## Run Unit Test in Specific Folder

```
docker-compose exec go go test -v -coverprofile cover.txt ./app/controller/brand
```