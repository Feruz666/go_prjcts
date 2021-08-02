package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: lenght %v, capasity %v %v\n", label, len(slice), cap(slice), slice)
}


func main() {

	// dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}

	// dump("dwarfs", dwarfs)
	// dump("dwarfs [1:2]", dwarfs[1:2])
	// dwarfs = append(dwarfs, "Orcus", "Kvavar", "Sedna")
	dwarfs1 := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}  // Длина 5, вместимость 5
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacius", "Kvavar", "Sedna")
	fmt.Println(len(dwarfs3), cap(dwarfs3))
}	
