package main

// +gen slice:"Where,GroupBy[int],Any,Count,Distinct,First,Select[Product],SortBy"
type Product struct {
	ID int64
	Name string
	Stock int64
	Price float64
}
