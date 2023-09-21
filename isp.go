package main

import "fmt"

// Interface Segregation Principle

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {
	fmt.Println("Printing document")
}

func (m MultiFunctionPrinter) Fax(d Document) {
	fmt.Println("Faxing document")
}

func (m MultiFunctionPrinter) Scan(d Document) {
	fmt.Println("Scanning document")
}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {
	fmt.Println("Printing document")
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("Operation not supported")
}

func (o OldFashionedPrinter) Scan(d Document) {
	panic("Operation not supported")
}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {
	fmt.Println("Printing document")
}

type Photocopier struct {
}

func (p Photocopier) Print(d Document) {
	fmt.Println("Printing document")
}

func (p Photocopier) Scan(d Document) {
	fmt.Println("Scanning document")
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	ofp := OldFashionedPrinter{}
	ofp.Scan(Document{})
}
