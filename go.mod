module github.com/infobloxopen/atlas-app-toolkit

go 1.14

require (
	contrib.go.opencensus.io/exporter/ocagent v0.7.0
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/dgrijalva/jwt-go v3.2.1-0.20200107013213-dc14462fd587+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/inflection v1.0.0
	github.com/lib/pq v1.3.1-0.20200116171513-9eb3fc897d6f
	github.com/sirupsen/logrus v1.8.0
	github.com/speps/go-hashids/v2 v2.0.1
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.22.4
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106
	google.golang.org/grpc v1.45.0
	google.golang.org/grpc/examples v0.0.0-20210715165331-ce7bdf50abb1 // indirect
	google.golang.org/protobuf v1.27.1
)
