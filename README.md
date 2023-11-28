# Train Ticketing Golang App

# Requirements
  All API referenced are gRPC APIs, not REST ones.
- I want to board a train from London to France. The train ticket will cost $20. 1. Create API where you can submit a purchase for a ticket. Details included in the receipt are:
``` From, To, User, price paid. User should include first and last name, email address ```
- The user is allocated a seat in the train. Assume the train has only 2 sections, section A and section B.
- An API that shows the details of the receipt for the user
- An API that lets you view the users and seat they are allocated by the requested section
- An API to remove a user from the train
- An API to modify a user's seat
- An API to make a discount (i.e. TBD123, $1 off)

# Best Code Structure  
```ticketing_app
.
├── README.md
├── api
│   ├── get_receipt_details.go
│   ├── get_user_seats.go
│   ├── modify_user_seat.go
│   ├── module.go
│   ├── submit_purchase.go
│   ├── remove_user.go
│   ├── apply_discount.go
│   └── utils.go
├── go.mod
├── go.sum
├── main.go
├── proto
│   └── ticket
│       └── ticket.proto
└── proto-gen
    └── ticket
        ├── ticket.pb.go
        └── ticket_grpc.pb.go
```

# GRPC APIS
- ```PurchaseTicket```: Submit a purchase for a ticket.
- ```GetReceiptDetails```: Get details of the receipt for a user.
- ```GetUserSeats```: Get users and their allocated seats by section.
- ```RemoveUser```: Remove a user from the train.
- ```ModifyUserSeat```: Modify a user's seat.
- ```ApplyDiscount```: Apply a Discount.

# Run Server
```
$ go build .
$ ./ticket_app
```
