package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         uint   `db:"id,primary_key"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     uint   `db:"salary"`
}

func main() {
	dsn := "root:12345677@tcp(127.0.0.1:3306)/grom?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	//if err := createTableEmployees(db); err != nil {
	//	panic(err)
	//}
	//println("员工表创建成功")

	//employees := []Employee{
	//	{Name: "张三", Department: "技术部", Salary: 15000},
	//	{Name: "李四", Department: "技术部", Salary: 12000},
	//	{Name: "王五", Department: "营销部", Salary: 8000},
	//}
	//if err := createEmployees(db, employees); err != nil {
	//	panic(err)
	//}
	//println("员工数据插入成功")

	//查询技术部所有员工
	//techEmployees, err := getEmployeesByDepartment(db, "技术部")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("技术部员工:", techEmployees)

	//查询工资最高的员工信息
	employee, err := getHeightSalaryEmployees(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("最高工资员工信息:", employee)
}

// 创建表
func createTableEmployees(db *sqlx.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS employees (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        department VARCHAR(50) NOT NULL,
        salary INT NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    `
	_, err := db.Exec(createTableSQL)
	return err
}

// 插入数据
func createEmployees(db *sqlx.DB, employees []Employee) error {
	//单个创建
	//_, err := db.Exec("insert into employees (name,department,salary) values (?,?,?)", "张三", "技术部", 10000)
	//if err != nil {
	//	return err
	//}
	//return nil

	for _, employee := range employees {
		_, err := db.NamedExec("insert into employees (name,department,salary) values (:name,:department,:salary)", employee)
		if err != nil {
			return err
		}
	}

	return nil

}

// 查询指定部门的所有员工
func getEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee

	err := db.Select(&employees, "SELECT id, name, department, salary FROM employees WHERE department = ?", department)
	if err != nil {
		return nil, fmt.Errorf("查询员工失败: %v", err)
	}

	return employees, nil
}

// 查询工资最高的员工信息
func getHeightSalaryEmployees(db *sqlx.DB) (Employee, error) {
	var employee Employee
	err := db.Get(&employee, "select id, name, department, salary from employees order by salary desc limit 1")
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}
