package main

import "fmt"

func main() {
    var pilihan string
    var angka1 float64
    var angka2 float64
    var operator string
    var hasil float64
    var error string
    var status bool

    status = false
    
    for !status {
        error = ""
        
        fmt.Print("\nMasukkan angka pertama: ")
        fmt.Scanln(&angka1)
        
        fmt.Print("Pilih operator (+, -, *, /): ")
        fmt.Scanln(&operator)
        
        fmt.Print("Masukkan angka kedua: ")
        fmt.Scanln(&angka2)
        
        switch operator {
        case "+":
            hasil = angka1 + angka2
        case "-":
            hasil = angka1 - angka2
        case "*":
            hasil = angka1 * angka2
        case "/":
            if angka2 != 0 {
                hasil = angka1 / angka2
            } else {
                error = "Tidak bisa dibagi dengan nol!"
            }
        default:
            error = "Operator tidak valid!"
        }
        
        if error != "" {
            fmt.Println("ERROR:", error)
        } else {
            fmt.Printf("Hasil: %.2f %s %.2f = %.2f\n", angka1, operator, angka2, hasil)
        }
        
        fmt.Print("\nHitung lagi? (y/n): ")
        fmt.Scanln(&pilihan)
        
        if pilihan != "y" && pilihan != "Y" {
            fmt.Println("Terima kasih!")
            status = true
        }
    }
}