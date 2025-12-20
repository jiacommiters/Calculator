package main
import "fmt"
func main() {
var hasil, angkasebelum, angka, error float64
var op, oplama, ulang string
var lanjut, lagi bool

lagi = true
fmt.Print("==Selamat Datang Di Kalkulator Digital== \n")
for lagi {
	fmt.Print("\nMasukkan angka: \n")
	fmt.Scan(&angka)
	
	hasil, angkasebelum, error = 0, 0, 0
	// oplama = "+"
	lanjut = true
	for lanjut {
		fmt.Scan(&op)

		if oplama == "+" {
			hasil = hasil + angkasebelum
			angkasebelum = angka
		} else if oplama == "-" {
			hasil = hasil + angkasebelum
			angkasebelum = -angka
		} else if oplama == "*" {
			angkasebelum = angkasebelum * angka
		} else if oplama == "/" {
			angkasebelum = angkasebelum / angka
		}

		if op == "=" {
			lanjut = false
		} else {
			fmt.Scan(&angka)
			oplama = op
		}
		
		if op == "/" && angka == 0 {
			error = error+1
		}
	}
	if error == 1 {
		fmt.Print("\nError, Tidak Bisa Membagi 0!\n")
	} else {
	fmt.Println(hasil + angkasebelum, "\n")
	}
	fmt.Println("Hitung lagi? y/n")
	fmt.Scan(&ulang)
	switch ulang {
		case "y": lagi = true
		case "n": lagi = false
	}
}
fmt.Print("\n ===Terima Kasih=== \noFARREN DIMAZ FAUZI \noMUHAMMAD DZIA ILHAQI \noMUHAMMAD FAJAR RAMADHAN")
}
