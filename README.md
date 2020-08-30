# Introduction
This Go program  converts the employee tabular data into the hierarchical format.  For example, the input data
```
Employee Name 	id      Manager id 
Alan 	        100 	150 
Martin 	        220 	100 
Jamie 	        150 	 
Alex 	        275 	100 
Steve 	        400 	150 
David 	        190 	400 
```
The hierarchical format looks like: 
```
Jamie 	 	 
 	Alan 	 
 	 	Matin 
 	 	Alex 
 	Steve 	 
 	 	David 
```

# Setup
You'll need to install Golang, as well as the git and make commands. Open terminal window, clone this repository and cd into it.
```
git clone https://github.com/mullumaus/mnt-golang.git
cd mnt-golang/src
```
# Source code
The src directory has following files:
1. main.go : the main go program
2. main_test.go: the unit test suite
3. Makefile: build, run, test and clean up this program

The code builds tabular employee data into a hierarchy using recursion, and use a tree data structure to represent the  hierarchy. Each employee is a node that has a list of children representing the directly report employees.
```
type EmployeeNode struct {
	name      string
	id        int
	managerID int
	children  []*EmployeeNode
}
```
The program defined a node slice (nodes) to store the input data. buildEmployeeTree() function builds the tree from the begining of the slice. Each item in the slice is a single node tree. The algorithm is to merge these single node trees together according to manager relationship.
1. if nodes[i]'s manager ID is null, this node must be CEO, it's the root of the tree
2. nodes[i]'s manager must be in root's tree or nodes[i+1:], call addEmployee() function to find its manager and add into correct tree
3. Print the hierarchical data by printEmployee() function.  This function calls builderEmployeeString() to build the output string using recursion. 
   
# Build
Run 'make build' command to build the code
```
make build
```

# Run
Use 'make run' command to run the program, output is
```
go run main.go
Alan    100     150
Martin  220     100
Jamie   150     -1
Alex    275     100
Steve   400     150
David   190     400
------------------------------------
Jamie
        Alan
                Martin
                Alex
        Steve
                David
```
# Run unit test
Invoke command 'make test' to run the test cases defined in main_test.go. The test suites include five test cases
1. TestBuildEmployeeTree: input flat employee data, and print the hirarchy format
2. TestChangeEmployeeOrder: change the input data order to verify the output is correct
3. TestInsertEmployee: insert an employee to the existing tree
4. TestFindEmployeeWithManagerID: print a manager's employees (directly&non-directly report)
5. TestAddInvalidEmployee: add an invalid employee should fail
   
# Clean up
Run 'make clean' to delete the exectuable file

# Time complexity 
The searching in this algorithm includes two parts:
1. Searching in the root tree. Search tree has worst case complexity of O(i) where i is the number of nodes in the tree. In general, time complexity is O(h) where h is height.
2. Searching in the remaining nodes in slice. Search data slice has complexity of O(j) where i+j=n




