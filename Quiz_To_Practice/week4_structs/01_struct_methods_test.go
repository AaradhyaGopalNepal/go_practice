package week4_structs

import (
	"testing"
	"time"
	"strings"
	"math"
	"github.com/stretchr/testify/assert"
)

func TestBankAccount(t *testing.T) {
	t.Skip("Remove this line after implementing BankAccount")

	// Create accounts
	// account1 := &BankAccount{
	// 	accountNumber: "ACC001",
	// 	owner: "John Doe",
	// 	balance: 1000.0,
	// }
	// account2 := &BankAccount{
	// 	accountNumber: "ACC002",
	// 	owner: "Jane Smith",
	// 	balance: 500.0,
	// }

	// Test deposit
	// err := account1.Deposit(500)
	// assert.NoError(t, err)
	// assert.Equal(t, 1500.0, account1.GetBalance())

	// Test withdrawal
	// err = account1.Withdraw(200)
	// assert.NoError(t, err)
	// assert.Equal(t, 1300.0, account1.GetBalance())

	// Test insufficient funds
	// err = account2.Withdraw(1000)
	// assert.Error(t, err)

	// Test transfer
	// err = account1.Transfer(account2, 300)
	// assert.NoError(t, err)
	// assert.Equal(t, 1000.0, account1.GetBalance())
	// assert.Equal(t, 800.0, account2.GetBalance())

	// Test transaction history
	// history := account1.GetTransactionHistory(3)
	// assert.Len(t, history, 3)
}

func TestLibrary(t *testing.T) {
	t.Skip("Remove this line after implementing Library")

	// lib := &Library{
	// 	books: []Book{},
	// 	members: []Member{},
	// 	loans: make(map[string]Loan),
	// }

	// Add books
	// book1 := Book{
	// 	ISBN: "123",
	// 	title: "Go Programming",
	// 	author: "John Doe",
	// 	totalCopies: 2,
	// 	availableCopies: 2,
	// }
	// lib.AddBook(book1)

	// Register member
	// member := Member{
	// 	ID: "M001",
	// 	name: "Alice",
	// 	email: "alice@example.com",
	// }
	// lib.RegisterMember(member)

	// Borrow book
	// err := lib.BorrowBook("M001", "123")
	// assert.NoError(t, err)

	// Check availability decreased
	// books := lib.SearchBooks("Go Programming")
	// assert.Equal(t, 1, books[0].availableCopies)

	// Return book
	// err = lib.ReturnBook("M001", "123")
	// assert.NoError(t, err)
}

func TestShapes(t *testing.T) {
	t.Skip("Remove this line after implementing Shapes")

	// Create shapes
	// circle := &Circle{
	// 	Shape: Shape{Name: "Circle1", Color: "red"},
	// 	Center: Point{X: 0, Y: 0},
	// 	Radius: 5,
	// }

	// rectangle := &Rectangle{
	// 	Shape: Shape{Name: "Rect1", Color: "blue"},
	// 	TopLeft: Point{X: 0, Y: 0},
	// 	Width: 10,
	// 	Height: 5,
	// }

	// Test area
	// assert.Equal(t, 25*math.Pi, circle.Area())
	// assert.Equal(t, 50.0, rectangle.Area())

	// Test perimeter
	// assert.Equal(t, 10*math.Pi, circle.Perimeter())
	// assert.Equal(t, 30.0, rectangle.Perimeter())

	// Test contains
	// assert.True(t, circle.Contains(Point{X: 3, Y: 3}))
	// assert.False(t, circle.Contains(Point{X: 10, Y: 10}))

	// Test move
	// circle.Move(5, 5)
	// assert.Equal(t, Point{X: 5, Y: 5}, circle.Center)

	// Test scale
	// circle.Scale(2)
	// assert.Equal(t, 10.0, circle.Radius)
}

func TestOrderedMap(t *testing.T) {
	t.Skip("Remove this line after implementing OrderedMap")

	// om := NewOrderedMap()

	// Test set and get
	// om.Set("first", 1)
	// om.Set("second", 2)
	// om.Set("third", 3)

	// val, ok := om.Get("second")
	// assert.True(t, ok)
	// assert.Equal(t, 2, val)

	// Test order preservation
	// keys := om.Keys()
	// assert.Equal(t, []string{"first", "second", "third"}, keys)

	// Test delete
	// deleted := om.Delete("second")
	// assert.True(t, deleted)
	// assert.Equal(t, 2, om.Len())

	// keys = om.Keys()
	// assert.Equal(t, []string{"first", "third"}, keys)

	// Test ForEach
	// var result []string
	// om.ForEach(func(key string, value interface{}) {
	// 	result = append(result, fmt.Sprintf("%s:%v", key, value))
	// })
	// assert.Equal(t, []string{"first:1", "third:3"}, result)
}

func TestEmployeeHierarchy(t *testing.T) {
	t.Skip("Remove this line after implementing Employee hierarchy")

	// Create employees
	// emp1 := &Employee{
	// 	Person: Person{Name: "John", Age: 30, Email: "john@example.com"},
	// 	EmployeeID: "E001",
	// 	Department: "Engineering",
	// 	Salary: 75000,
	// 	HireDate: time.Now().AddDate(-2, 0, 0),
	// }

	// emp2 := &Employee{
	// 	Person: Person{Name: "Jane", Age: 28, Email: "jane@example.com"},
	// 	EmployeeID: "E002",
	// 	Department: "Engineering",
	// 	Salary: 70000,
	// 	HireDate: time.Now().AddDate(-1, 0, 0),
	// }

	// Create manager
	// mgr := &Manager{
	// 	Employee: Employee{
	// 		Person: Person{Name: "Bob", Age: 40, Email: "bob@example.com"},
	// 		EmployeeID: "M001",
	// 		Department: "Engineering",
	// 		Salary: 100000,
	// 		HireDate: time.Now().AddDate(-5, 0, 0),
	// 	},
	// 	Budget: 500000,
	// }

	// Test manager methods
	// mgr.AddTeamMember(emp1)
	// mgr.AddTeamMember(emp2)
	// assert.Equal(t, 2, mgr.GetTeamSize())
	// assert.Equal(t, 145000.0, mgr.GetTotalTeamSalary())

	// Test employee methods
	// emp1.GiveRaise(10)
	// assert.Equal(t, 82500.0, emp1.Salary)

	// Create executive
	// exec := &Executive{
	// 	Manager: *mgr,
	// 	StockOptions: 10000,
	// 	Departments: []string{"Engineering", "Product"},
	// }

	// Test executive methods
	// exec.ManageDepartment("Sales")
	// assert.Contains(t, exec.Departments, "Sales")
}