package main

import "testing"

func initEmployee() *EmployeeNode {
	employees := []EmployeeNode{
		{"Steve", 400, 150, nil},
		{"David", 190, 400, nil},
		{"Alan", 100, 150, nil},
		{"Martin", 220, 100, nil},
		{"Jamie", 150, NullID, nil},
		{"Alex", 275, 100, nil},
	}
	root := buildEmployeeTree(employees)
	return root
}
func TestBuildEmployeeTree(t *testing.T) {
	root := initEmployee()
	root.printEmployee()
}

func TestChangeEmployeeOrder(t *testing.T) {
	employees := []EmployeeNode{
		{"Steve", 400, 150, nil},
		{"David", 190, 400, nil},
		{"Jamie", 150, NullID, nil},
		{"Alex", 275, 100, nil},
		{"Alan", 100, 150, nil},
		{"Martin", 220, 100, nil},
	}
	printTabularData(employees)
	root := buildEmployeeTree(employees)
	root.printEmployee()
}

func TestInsertEmployee(t *testing.T) {
	root := initEmployee()
	t.Log("Before -------------------")
	root.printEmployee()

	root = addEmployee(root, &EmployeeNode{"Kevin", 200, 100, nil})
	t.Log("After -------------------")
	root.printEmployee()

}

func TestFindEmployeeWithManagerID(t *testing.T) {
	root := initEmployee()
	employeeID := 100
	e := root.findEmpolyeeWithManagerID(employeeID)
	if e == nil {
		t.Fatal("Failed to find employee ID ", employeeID)
	} else {
		e.printEmployee()
	}
}

func TestAddInvalidEmployee(t *testing.T) {
	root := initEmployee()
	e := &EmployeeNode{"James", 290, 19, nil}
	root = addEmployee(root, e)
	if root == nil {
		t.Log("Failed to add", e.name)
	}
}
