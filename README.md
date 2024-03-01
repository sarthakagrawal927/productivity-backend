# Backend

[Endpoints](https://galactic-escape-804413.postman.co/workspace/Stumble~8b18c535-4c33-4445-9f41-3fd645691d7d/collection/20124508-30a6e9dc-c460-473e-a0c3-e2fc2c5e066d?action=share&creator=20124508&active-environment=26183107-3fa6a988-04b0-403c-8575-7ec046deff9c)

## Notes

- Gorm adds deletedAt column in each table, and it always soft deletes. That's why there are no delete status in the models. Though it is to promote soft delete I feel having extra column can be avoided if I already have a status column in most tables which can be purposed into delete. Will do that after the project is in reasonable shape.

## URLS

https://github.com/golang-jwt/jwt

https://developers.google.com/calendar/api/quickstart/go

https://github.com/search?q=google.golang.org%2Fapi%2Fcalendar%2Fv3+language%3AGo&type=code

## command to generate secretKey in ubuntu

openssl rand -base64 32 > secret_key.txt
