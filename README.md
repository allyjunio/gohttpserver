#REST API in Golang with mySql Database

# Install go lang

# Configuration

        dataSourceName: root:root@/teamwork

# Installation

        go get -u github.com/go-sql-driver/mysql
        go get -u github.com/elgs/gosqljson
        go get -u github.com/gorilla/mux

        git clone https://github.com/allyjunio/gohttpserver.git
        cd gohttpserver
        go run server.go

And open http://localhost:8080

# Endpoint 1: GET '/depts/{deptNo}/bonuses' ex: http://localhost:8080/depts/1/bonuses
# Endpoint 2: GET '/employees/lowearners'
# Endpoint 3: GET '/employees'




