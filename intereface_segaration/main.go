package main

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

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {

}

func (o OldFashionedPrinter) Fax(d Document) {

}

func (o OldFashionedPrinter) Scan(d Document) {

}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {

}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {

}

func (p Photocopier) Scanner(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

//decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {}

func main() {
	//ofp := OldFashionedPrinter{}
}

/*
You should put too much in an interface. You shouldnt try to throw everything
into one interface. Sometimes it makes sense to break up an interface into
several interfaces.

Break up interfaces into separate parts that people will definitely need.
This is no guarentee that someone with print will need fax.

you can combine interfaces that represents a multifunction device
*/
