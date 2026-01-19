package main

import (
	"fmt"
	"time"
)

type Expense struct {
	Name     string
	Amount   float64
	Category string
	Date     time.Time
}

func addExpense(expenses []Expense, name string, amount float64, category string) []Expense {
	expense := Expense{
		Name:     name,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
	}
	return append(expenses, expense)
}

func displayExpenses(expenses []Expense) {
	fmt.Println("\nList of all expenses:")
	for _, expense := range expenses {
		fmt.Printf("Name: %s, Amount: %.2f, Category: %s, Date: %s\n", expense.Name, expense.Amount, expense.Category, expense.Date.Format("2006-01-02 15:04:05"))
	}
}

func displayTotalByCategory(expenses []Expense) {
	categoryTotals := make(map[string]float64)

	for _, expense := range expenses {
		categoryTotals[expense.Category] += expense.Amount
	}

	fmt.Println("\nTotal expenses by category:")
	for category, total := range categoryTotals {
		fmt.Printf("Category: %s, Total: %.2f\n", category, total)
	}
}

func main() {
	var expenses []Expense

	expenses = addExpense(expenses, "coffee", 5.50, "Food")
	expenses = addExpense(expenses, "gym membership", 50.00, "Health")
	expenses = addExpense(expenses, "groceries", 30.75, "Food")
	expenses = addExpense(expenses, "internet Bill", 60.00, "Bills")

	displayExpenses(expenses)

	displayTotalByCategory(expenses)
}
