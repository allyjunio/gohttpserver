// server.go
//
// REST APIs with Go and MySql.
//
// Usage:
//
//   # run go server in the background
//   $ go run server.go

package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/elgs/gosqljson"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"fmt"
	"time"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error
var theCase string

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TeamWork!!!!\n"))
}

func getEmployeeLowEarners(w http.ResponseWriter, r *http.Request) {
	data, _ := gosqljson.QueryDbToMapJSON(db, theCase,
		"select employe.emp_id as employeeId, " +
				    	    "employe.emp_name as employeeName, " +
							"DATE_FORMAT(employe.hire_date, '%d-%b-%Y') as hireDate, " +
							"employe.salary as salary " +
					"from employees employe " +
					"where salary < (select salary from employees employee where employee.hire_date > employe.hire_date order by hire_date limit 1) " +
					"order by hire_date ")
	w.Write([]byte(data))
}

func getBonusDepartment(w http.ResponseWriter, r *http.Request) {
	deptNo := mux.Vars(r)["deptNo"]
	data, _ := gosqljson.QueryDbToMapJSON(db, theCase,
		"select  employe.dept_no as deptId, " +
						"sum(employe.salary) as totalSalary," +
						"sum((employe.salary * (select sum(bonus.type) * 10 as bonus  from bonuses bonus where bonus.emp_id = employe.emp_id)/100)) as totalBonuses " +
					"from employees employe " +
					"where employe.dept_no = " + deptNo +
					" group by dept_no")
	w.Write([]byte(data))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	data, _ := gosqljson.QueryDbToMapJSON(db, theCase, "SELECT emp_id as employeeId, emp_name as employeeName, dept_no as deptId, salary, DATE_FORMAT(hire_date, '%d-%b-%Y') as hireDate FROM employees")
	w.Write([]byte(data))
}

func main() {
	db, err = sql.Open("mysql", "root:root@/teamwork")
	theCase = "lower" // "lower", "upper", "camel" or the orignal case if this is anything other than these three

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/depts/{deptNo:[0-9]+}/bonuses", getBonusDepartment).Methods("GET")
	r.HandleFunc("/employees/lowearners", getEmployeeLowEarners).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}



