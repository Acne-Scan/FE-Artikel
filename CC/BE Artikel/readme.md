# Migration
```
migrate -database "mysql://root:@tcp(localhost:3306)/acne_scan" -path internal/infrastructure/database/migrations up
```
```
migrate create -ext sql -dir internal/infrastructure/database/migrations (namafile).sql
```
