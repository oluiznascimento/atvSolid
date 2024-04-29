package main

import "fmt"

// Printer é uma interface para impressoras.
type Printer interface {
    Print()
}

// Scanner é uma interface para scanners.
type Scanner interface {
    Scan()
}

// Photocopier é uma interface para fotocopiadoras.
type Photocopier interface {
    Printer
    Scanner
}

// CanonPrinter é uma estrutura que representa uma impressora Canon.
type CanonPrinter struct{}

// Print implementa o método Print para CanonPrinter.
func (cp CanonPrinter) Print() {
    fmt.Println("Printing document")
}

// CanonScanner é uma estrutura que representa um scanner Canon.
type CanonScanner struct{}

// Scan implementa o método Scan para CanonScanner.
func (cs CanonScanner) Scan() {
    fmt.Println("Scanning document")
}

// CanonPhotocopier é uma estrutura que representa uma fotocopiadora Canon.
type CanonPhotocopier struct {
    printer Printer
    scanner Scanner
}

// NewCanonPhotocopier cria uma nova instância de CanonPhotocopier.
func NewCanonPhotocopier(printer Printer, scanner Scanner) *CanonPhotocopier {
    return &CanonPhotocopier{
        printer: printer,
        scanner: scanner,
    }
}

// Print faz a fotocopiadora Canon imprimir.
func (cp *CanonPhotocopier) Print() {
    cp.printer.Print()
}

// Scan faz a fotocopiadora Canon digitalizar.
func (cp *CanonPhotocopier) Scan() {
    cp.scanner.Scan()
}

func main() {
    canonPrinter := CanonPrinter{}
    canonScanner := CanonScanner{}

    canonPhotocopier := NewCanonPhotocopier(canonPrinter, canonScanner)
    canonPhotocopier.Print()
    canonPhotocopier.Scan()
}
