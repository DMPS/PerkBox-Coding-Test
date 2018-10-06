#Perkbox Coding Test

> A coding test in Golang written by Dalton Scott

## Requirements

### Technical Requirements

- A simple RESTful api service in Go, no frameworks should be used
- Unit tested
- Datastore of your choice
- Available on GitHub

### Product Requirements

- We would like you to build a coupons service
- You should be able to create, update, retrieve & list coupons via the API
- The list endpoint should have a variety of filters and parameters to allow clients of your API to retrieve coupons in different ways

### Example coupon

```json
{
"name": "Save £20 at Tesco",
"brand": "Tesco",
"value": 20,
"createdAt": "2018-03-01 10:15:53",
"expiry": "2019-03-01 10:15:53"
}
```