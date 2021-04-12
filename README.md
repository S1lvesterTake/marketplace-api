# Condition
1. Customers were able to put items in their cart, check out, and then pay. After several days, many of our customers received calls from
our Customer Service department stating that their orders have been canceled due to stock unavailability.
2. These bad reviews generally come within a week after our 12.12 event, in which we held a large flash sale and set up other major
discounts to promote our store.

# Describe
1. This incident could happen because data inconsistencies in the tables in the database related to product stock and the system is less informative and preventive in notifying the stock of goods directly to consumers when putting goods into carts or when making transactions

2. add guard in the following processes:
  - When you put the products into the cart, make sure backend always check the stock of the products.
  - When doing checkout it will reduce the stock of products.
  - When consumen do not pay off the payment after the payment deadline, it will return the stock of products.
  - send notification via email or push notification when the stock of items put in runs out



## Requirement
1. Install GO version 1.15.4 or newer
2. MySQL
## How to run
1. run `go run main.go`

## postman
`https://www.getpostman.com/collections/1e46ec23ca6117c49a32`