package main

import (
	"fmt"
	"log"
	"strconv"
)

//NullID is invalid employee ID
const NullID = -1

//EmployeeNode is the tree node
type EmployeeNode struct {
	name      string
	id        int
	managerID int
	children  []*EmployeeNode
}

func (e *EmployeeNode) findEmpolyeeWithManagerID(managerID int) *EmployeeNode {
	var result *EmployeeNode = nil
	if e.id == managerID {
		return e
	}

	for _, child := range e.children {
		result = child.findEmpolyeeWithManagerID(managerID)
		if result != nil {
			break
		}
	}
	return result
}

func (e *EmployeeNode) addChildren(newEmploye *EmployeeNode) {
	e.children = append(e.children, newEmploye)
}

func (e EmployeeNode) builderEmployeeString(tabLevel int) string {
	builder := ""
	for i := 0; i < tabLevel; i++ {
		builder += "\t"
	}
	builder += e.name
	builder += "\n"
	for _, child := range e.children {
		builder += child.builderEmployeeString(tabLevel + 1)
	}
	return builder
}

func (e EmployeeNode) printEmployee() {
	builder := e.builderEmployeeString(0)
	fmt.Println(builder)
}

func addEmployee(root *EmployeeNode, newEmployee *EmployeeNode) *EmployeeNode {
	if root == nil {
		return nil
	}
	if newEmployee.managerID == NullID {
		//newEmployee is the new root
		if root.managerID == newEmployee.id {
			newEmployee.addChildren(root)
			return newEmployee
		}
		return nil
	}
	manager := root.findEmpolyeeWithManagerID(newEmployee.managerID)
	if manager != nil {
		manager.addChildren(newEmployee)
		return root
	}
	return nil
}

func buildEmployeeTree(nodes []EmployeeNode) *EmployeeNode {
	var root *EmployeeNode = nil

	for i := 0; i < len(nodes); i++ {
		newEmployee := &nodes[i]
		if newEmployee.managerID == NullID && root == nil {
			//CEO of the company doesn't have a manager
			root = newEmployee
		} else {
			//try to add newEmployee to root tree
			result := addEmployee(root, newEmployee)
			//result is nil if newEmployee's manager isn't in the root tree, find its manager in nodes[i+1:]
			if result == nil {
				for j := i + 1; j < len(nodes); j++ {
					result = addEmployee(&nodes[j], newEmployee)
					if result != nil {
						break
					}
				}
				//the employee's manager does not exist
				if result == nil {
					log.Fatal("Can not find employee ", newEmployee.name, "'s manager ", newEmployee.managerID)
				}
			}
		}
	}
	if root == nil {
		panic("The company does not have an CEO!!!")
	}
	return root
}

func printTabularData(nodes []EmployeeNode) {
	for _, node := range nodes {
		str := node.name + "\t" + strconv.Itoa(node.id) + "\t" + strconv.Itoa(node.managerID)
		fmt.Println(str)
	}
	fmt.Println("------------------------------------")
}

func main() {
	nodes := []EmployeeNode{
		{"Alan", 100, 150, nil},
		{"Martin", 220, 100, nil},
		{"Jamie", 150, NullID, nil},
		{"Alex", 275, 100, nil},
		{"Steve", 400, 150, nil},
		{"David", 190, 400, nil},
	}
	printTabularData(nodes)
	root := buildEmployeeTree(nodes)
	root.printEmployee()
}
