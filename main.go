package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// define struct fruit with json tag
type Fruit struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Type  FruitType `json:"type"`
	Stock int       `json:"stock"`
}

// Custom type string
type FruitType string

// Constants untuk FruitType
const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

// untuk menghitung jumlah const
var AllFruitTypes = []FruitType{
	FruitTypeImport,
	FruitTypeLocal,
}

// Struct untuk menampung data fruit yang valid
type FruitList struct {
	fruits []Fruit
}

// Cek apakah ID fruit sudah ada
func (fl *FruitList) Exists(id int) bool {
	for _, f := range fl.fruits {
		if f.ID == id {
			return true
		}
	}
	return false
}

// Tambah fruit dengan cek ID unik
func (fl *FruitList) Add(fruit Fruit) error {
	if fl.Exists(fruit.ID) { //Jika ID fruit sudah ada pada struct data fruit valid, data ditolak
		return errors.New("ID sudah ada, tidak boleh duplikat")
	}
	fl.fruits = append(fl.fruits, fruit) //Masukkan data ke dalam struct data fruit valid
	return nil
}

func main() {
	// JSON input dengan 7 data fruit
	jsonData := `[
        {"id": 1, "name": "Apel", "type": "IMPORT", "stock":10},
        {"id": 2, "name": "Kurma", "type": "IMPORT", "stock":20},
		{"id": 3, "name": "apel", "type": "IMPORT", "stock":50},
		{"id": 4, "name": "Manggis", "type": "LOCAL", "stock":100},
		{"id": 5, "name": "Jeruk Bali", "type": "LOCAL", "stock":10},
		{"id": 5, "name": "KURMA", "type": "IMPORT", "stock":20},
		{"id": 5, "name": "Salak", "type": "LOCAL", "stock":150}
    ]`

	var fruitsFromJson []Fruit

	// Parsing JSON
	if err := json.Unmarshal([]byte(jsonData), &fruitsFromJson); err != nil {
		fmt.Println("Gagal parse JSON:", err)
		return
	}

	fruitList := FruitList{}

	// Tambah tiap fruit ke fruitList dengan cek duplikat ID
	for _, f := range fruitsFromJson {
		if err := fruitList.Add(f); err != nil {
			fmt.Printf("x Gagal menambahkan fruit %s (ID: %d): %v\n", f.Name, f.ID, err)
		} else {
			fmt.Printf("- Buah %s (ID: %d) berhasil ditambahkan\n", f.Name, f.ID)
		}
	}

	// Tampilkan daftar fruit yang berhasil ditambahkan
	fmt.Println("\nFruit list:")
	fmt.Println("========================================")
	fmt.Printf("%-5s %-15s %-10s %-4s\n", "ID", "Fruit", "Type", "Stock")
	fmt.Println("========================================")
	for _, f := range fruitList.fruits {
		//WITHOUT FORMAT
		//fmt.Printf("- ID: %d, Buah: %s, Type: %s, Stock: %d\n", f.ID, f.Name, f.Type, f.Stock)
		fmt.Printf("%-5d %-15s %-10s %-4d\n", f.ID, f.Name, f.Type, f.Stock)
	}
	fmt.Println("========================================")

	fmt.Println("\nSoal 1. Buah apa saja yang dimiliki Andi? (fruitName)")
	//hanya nama buah
	for _, f := range fruitList.fruits {
		fmt.Printf("- ID: %d, Buah: %s\n", f.ID, f.Name)
	}

	fmt.Println("\nSoal 2. Anda memisahkan buahnya menjadi beberapa wadah berdasarkan tipe buah (fruitType)\nBerapa jumlah wadah yang dibutuhkan?\nDan ada buah apa saja di masing-masing wadah?")

	fmt.Println("\nJumlah wadah yang dibutuhkan:", len(AllFruitTypes))

	fmt.Println("Daftar buah di wadah LOCAL:")

	fruitListLocal := FruitList{}
	for _, f := range fruitList.fruits {
		if f.Type == "LOCAL" {
			fruitListLocal.fruits = append(fruitListLocal.fruits, f) //Masukkan data ke dalam struct data fruit valid
		}
	}
	for _, l := range fruitListLocal.fruits {
		fmt.Printf("- ID: %d, Buah: %s, Stock: %d\n", l.ID, l.Name, l.Stock)
	}

	fmt.Println("Daftar buah di wadah IMPORT:")
	fruitListImport := FruitList{}
	for _, f := range fruitList.fruits {
		if f.Type == "IMPORT" {
			fruitListImport.fruits = append(fruitListImport.fruits, f) //Masukkan data ke dalam struct data fruit valid
		}
	}
	for _, l := range fruitListImport.fruits {
		fmt.Printf("- ID: %d, Buah: %s, Stock: %d\n", l.ID, l.Name, l.Stock)
	}

	fmt.Println("\nSoal 3. Berapa total stock buah yang ada di masing-masing wadah?")
	//LOCAL
	totalBuahLokal := 0
	for _, l := range fruitListLocal.fruits {
		totalBuahLokal = totalBuahLokal + l.Stock
	}
	fmt.Printf("- Total buah LOCAL: %d\n", totalBuahLokal)

	//IMPORT
	totalBuahImport := 0
	for _, i := range fruitListImport.fruits {
		totalBuahImport = totalBuahImport + i.Stock
	}
	fmt.Printf("- Total buah IMPORT: %d\n", totalBuahImport)

	fmt.Println("\nSoal 4. Apakah ada komentar terkait kasus di atas?")
	fmt.Println("- Ya, jawaban saya lampirkan pada file Komentar kasus1.docx")
}
