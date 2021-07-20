package main

type sensor func() kelvin


// Var1
// func measureTemperature(samples int, s func() sensor)

//  Var2
// func measureTemperature(samples int, s sensor)



func drawTable(rows int, getRow func(row int) (string, string))

type getRow func() (string, string)
func drawTable(rows int, gr getRow)