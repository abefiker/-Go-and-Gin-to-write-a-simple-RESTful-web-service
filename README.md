# -Go-and-Gin-to-write-a-simple-RESTful-web-service
go run . to start a server
and the at seconde terminal Invoke-RestMethod -Uri "http://localhost:4000/albums" -Method GET -Headers @{"Content-Type"="application/json"}
Invoke-RestMethod -Uri "http://localhost:4000/albums/id" -Method GET -Headers @{"Content-Type"="application/json"}
Invoke-RestMethod -Uri "http://localhost:4000/albums" -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
