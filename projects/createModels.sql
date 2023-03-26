CREATE TABLE customers (
    "id" UUID ,
    "name" VARCHAR , 
    "phone" VARCHAR 
)

CREATE TABLE users (
    "id" UUID ,
    "name" VARCHAR , 
    "phone" VARCHAR 
);

CREATE TABLE couriers (
    "id" UUID ,
    "name" VARCHAR , 
    "phone_number" VARCHAR 
);

CREATE TABLE products (
    "id" UUID ,
    "name" VARCHAR , 
    "price" NUMERIC 
);

CREATE TABLE categories (
    "id" UUID ,
    "name" VARCHAR ,
    "product_id" UUID references product(id)
);

CREATE TABLE orders (
    "id" UUID ,
    "name" VARCHAR ,
    "price" NUMERIC ,
    "phone_number" VARCHAR ,
    "user_id" UUID references user(id)
    "customer_id" UUID references customer(id)
    "courier_id" UUID  references courier(id)
    "product_id" UUID references product(id)
    "quantity" NUMERIC
);