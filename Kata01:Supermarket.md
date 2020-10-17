Kata 01 (Supermaerket Problem)



Constraint
- variable pricing
    - 3 for price of 1
  - discount can apply to one or group of items
    - buy 2 get one free
  - rouding off
    - A more complax version can be tracking the historical price of item
      - use case : user comes to return an object, on day of selling the item was as 10$ but when user come to ask for refund the price of item is 12$, so how do we ensure that refunded amount is 10$.

- generation of proper invoice with correct tax deduction and display

- having some thing to ensure to know payment status, able to retrive past orders.

- logs
    - invoice logs
  - pricing
  - have a ability to know daily profilt/loss
  - can support loyalty programs like if user spends 100$ in month then 10 off on next purchase
 
(optional, if we can include these constraint, then it is good)

- stock management (how to ensure if some object is out of stock)
-


Till now we have disc

Solution :

Objects: -
    1. Item
         - id  (PK)
           - Pricing Object
    - Metadata ID
     - expiry date
- date of buy
- date of entry
    
  2. Pricing Object
    - mrp price
    - cost price

  3. MetaData Object
    -id
          - name
- brand
    - department
    - count
4. Users
phone(PK)
User type (seller / buyer) (?)

5. Manufacturer
    -name
    - address
    - id (PK)

6. Order
Items list  [<item object>]
Order id  <uuid>
Seller id (FK to user table) 
Buyer id (FK to user table)
Booking date
Valid (bool)
7. Payment
Order id
Mode
Transaction amount
Transaction id (PK)
Status (enum) -> not done/refunded/complete/pending with bank

8.Refund

Services
UserService(phone)
CURD operation on user
Interacts with user DB
Responsibilities
Get uuid of user based on phone no
Return user object based on uuid

Item onboarding service
CurD operation on items
Interacts with item DB
Responsibilities 
Return item details based on item ID

Pricing service - 
Interact with item Table
Responsibilities-
Get sale price of item based on item id
It has to handle cases like when prices vary, example (buy2 , 1free)
Discounts based on location/time of year/sale etc


Bill generation service
Interact with order Table, rPricing service
Responsibilities
Return invoice or bill based on order id




-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


Q - How do I find which seller sold me given object id ?



Aspects to consider that might edify: KISS and YAGNI, stakeholder requirements, pricing system dev cost vs. ROI,
product return policies, tax system requirements, i18n, currency granularity and conversion, cultural and industry norms,
fraud detection, forms of payment, 3rd party coupons, forms of POS device, data storage choices,
barter and cashier discretion, and last but very important, ease of system maintenance and modification.



https://docs.google.com/document/d/1qww4v_e4KBZA95H8CrzWmXFOU_afgG4jtKFlSdcHRN4/edit?usp=sharing



Learnings

A class is closed, since it may be compiled, stored in a library, baselined, and used by client classes. But it is also open, since any new class may use it as parent, adding new features. When a descendant class is defined, there is no need to change the original or to disturb its clients.

https://stackify.com/solid-design-open-closed-principle/

https://stackoverflow.com/questions/20677249/why-we-add-underscore-before-variable-name
