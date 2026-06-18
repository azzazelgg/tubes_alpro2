package main

import (
	"fmt"
	"strings"
)

const MAX = 1000

var daftar_ukuran = []string{
	"S",
	"M",
	"L",
	"XL",
	"XXL",
}

var daftar_kategori = []string{
	"Baju Pria",
	"Baju Wanita",
	"Bawahan Wanita",
	"Atasan Wanita",
	"Unisex",
}

type product struct {
	ID               int
	Nama             string
	Kategori         string
	Jumlah_penjualan int
	Ukuran           string
	Warna            string
	Stok             int
	Harga            float64
}

var data_product [MAX]product
var n int

func cari_kategori(kategori string) bool {
	for i := 0; i < len(daftar_kategori); i++ {
		if strings.EqualFold(daftar_kategori[i], kategori) {
			return true
		}
	}
	return false
}

func cari_ukuran(ukuran string) bool {
	for i := 0; i < len(daftar_ukuran); i++ {
		if strings.EqualFold(daftar_ukuran[i], ukuran) {
			return true
		}
	}
	return false
}

// TAMBAH DATA
func create_product() {
	if n >= MAX {
		fmt.Println("Gagal menambah produk.")
		return
	}

	fmt.Println("\n--- TAMBAH PRODUK ---")
	fmt.Print("Masukkan ID Produk     : ")
	fmt.Scan(&data_product[n].ID)

	fmt.Print("Masukkan Nama Produk   : ")
	fmt.Scan(&data_product[n].Nama)

	var pilih_kategori int
	fmt.Println("Daftar Kategori :")
	for i := 0; i < len(daftar_kategori); i++ {
		fmt.Printf("%d. %s\n", i+1, daftar_kategori[i])
	}

	fmt.Printf("Pilih Kategori (1-%d) : ", len(daftar_kategori))
	fmt.Scan(&pilih_kategori)

	if pilih_kategori < 1 || pilih_kategori > len(daftar_kategori) {
		fmt.Println("Kategori tidak valid!")
		return
	}

	data_product[n].Kategori = daftar_kategori[pilih_kategori-1]

	var pilih_ukuran int
	fmt.Println("Daftar Ukuran :")
	for i := 0; i < len(daftar_ukuran); i++ {
		fmt.Printf("%d. %s\n", i+1, daftar_ukuran[i])
	}

	fmt.Printf("Pilih Ukuran (1-%d) : ", len(daftar_ukuran))
	fmt.Scan(&pilih_ukuran)

	if pilih_ukuran < 1 || pilih_ukuran > len(daftar_ukuran) {
		fmt.Println("Ukuran tidak valid!")
		return
	}

	data_product[n].Ukuran = daftar_ukuran[pilih_ukuran-1]

	fmt.Print("Masukkan Warna         : ")
	fmt.Scan(&data_product[n].Warna)

	fmt.Print("Masukkan Jumlah Stok   : ")
	fmt.Scan(&data_product[n].Stok)

	fmt.Print("Masukkan Harga         : ")
	fmt.Scan(&data_product[n].Harga)

	data_product[n].Jumlah_penjualan = 0
	n++
	fmt.Println("Produk berhasil ditambahkan")

}

func edit_product() {
	if n == 0 {
		fmt.Println("Tidak ada produk yang bisa diubah.")
		return
	}

	var id_cari int
	fmt.Println("\n--- EDIT DATA PRODUK ---")
	fmt.Print("Masukkan ID Produk yang ingin diubah: ")
	fmt.Scan(&id_cari)

	for i := 0; i < n; i++ {
		if data_product[i].ID == id_cari {
			fmt.Println("\nData ditemukan! Silakan masukkan data baru: ")

			fmt.Printf("Nama Lama: %s -> Masukkan Nama Baru : ", data_product[i].Nama)
			fmt.Scan(&data_product[i].Nama)

			var pilih_kategori int

			fmt.Printf("Kategori Lama : %s\n", data_product[i].Kategori)

			for j := 0; j < len(daftar_kategori); j++ {
				fmt.Printf("%d. %s\n", j+1, daftar_kategori[j])
			}

			fmt.Print("Pilih Kategori Baru : ")
			fmt.Scan(&pilih_kategori)

			if pilih_kategori < 1 || pilih_kategori > len(daftar_kategori) {
				fmt.Println("Kategori tidak valid!")
				return
			}

			data_product[i].Kategori = daftar_kategori[pilih_kategori-1]

			var pilih_ukuran int
			fmt.Printf("Ukuran Lama : %s\n", data_product[i].Ukuran)

			for j := 0; j < len(daftar_ukuran); j++ {
				fmt.Printf("%d. %s\n", j+1, daftar_ukuran[j])
			}

			fmt.Print("Pilih Ukuran Baru : ")
			fmt.Scan(&pilih_ukuran)

			if pilih_ukuran < 1 || pilih_ukuran > len(daftar_ukuran) {
				fmt.Println("Ukuran tidak valid!")
				return
			}

			data_product[i].Ukuran = daftar_ukuran[pilih_ukuran-1]

			fmt.Printf("Warna Lama: %s -> Masukkan Warna Baru: ", data_product[i].Warna)
			fmt.Scan(&data_product[i].Warna)

			fmt.Printf("Stok Lama: %d -> Masukkan Stok Baru: ", data_product[i].Stok)
			fmt.Scan(&data_product[i].Stok)

			fmt.Printf("Harga Lama: Rp%.0f -> Masukkan Harga Baru: ", data_product[i].Harga)
			fmt.Scan(&data_product[i].Harga)

			fmt.Println("\n Data produk berhasil diperbarui!")
			return
		}
	}
	fmt.Println("ID produk tidak ditemukan. ")
}

func delete_product() {

	if n == 0 {
		fmt.Println("\n[!] Tidak ada data produk yang bisa dihapus. Silakan tambah produk terlebih dahulu.")
		return
	}

	var id_cari int
	fmt.Println("\n--- HAPUS DATA PRODUK ---")
	fmt.Print("Masukkan ID Produk yang ingin dihapus: ")
	fmt.Scan(&id_cari)

	for i := 0; i < n; i++ {
		if data_product[i].ID == id_cari {

			for j := i; j < n-1; j++ {
				data_product[j] = data_product[j+1]
			}

			n--
			fmt.Println("Produk berhasil dihapus.")
			return
		}

	}

	fmt.Println("ID Produk tidak ditemukan.")

}

func read_product() {
	// jika admin belum input data samsek
	if n == 0 {
		fmt.Println("\n[!] Data inventaris kosong")
		return
	}

	fmt.Println("\n==========================================================================================")
	fmt.Printf("%-6s | %-25s | %-10s | %-6s | %-10s | %-6s | %-12s |\n", "ID", "Nama Produk", "Kategori", "Ukuran", "Warna", "Stok", "Harga")
	fmt.Println("==========================================================================================")

	for i := 0; i < n; i++ {
		p := data_product[i]
		fmt.Printf("%-6d | %-25s | %-10s | %-6s | %-10s | %-6d | Rp%-10.0f |\n", p.ID, p.Nama, p.Kategori, p.Ukuran, p.Warna, p.Stok, p.Harga)

	}
	fmt.Println("==========================================================================================")

}

func sequential() {
	var cari_barang string

	fmt.Print("Cari Ukuran/Warna : ")
	fmt.Scan(&cari_barang)

	cari := false

	for i := 0; i < n; i++ {
		if strings.EqualFold(data_product[i].Ukuran, cari_barang) ||
			strings.EqualFold(data_product[i].Warna, cari_barang) {
			fmt.Printf("%s %s %s %s \n",
				data_product[i].Nama,
				data_product[i].Ukuran,
				data_product[i].Warna,
				data_product[i].Kategori)

			cari = true
		}
	}

	if !cari {
		fmt.Println("Data tidak ditemukan")
	}
}

func nilai_ukuran(ukuran string) int {
	switch ukuran {
	case "S":
		return 1
	case "M":
		return 2
	case "L":
		return 3
	case "XL":
		return 4
	case "XXL":
		return 5
	default:
		return 999
	}
}

func sorting_ukuran() {
	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if nilai_ukuran(data_product[j].Ukuran) < nilai_ukuran(data_product[min].Ukuran) {
				min = j
			}
		}

		data_product[i], data_product[min] = data_product[min], data_product[i]
	}
}

func binary_search() {
	if n == 0 {
		fmt.Println("Data Tidak Ditemukan")
	}

	sorting_ukuran()

	var cari_ukuran string
	fmt.Print("Masukkan Data yang ingin di cari : ")
	fmt.Scan(&cari_ukuran)

	target := nilai_ukuran(strings.ToUpper(cari_ukuran))

	kiri := 0
	kanan := n - 1
	ditemukan := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		nilai_tengah := nilai_ukuran(data_product[tengah].Ukuran)

		if nilai_tengah == target {
			ditemukan = tengah
			break
		} else if target < nilai_tengah {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	// jika data tidak ditemukan
	if ditemukan == -1 {
		fmt.Println("Produk dengan ukuran tersebut tidak ditemukan")
		return
	}

	fmt.Println("\n Produk telah ditemukan")

	awal := ditemukan
	for awal > 0 &&
		nilai_ukuran(data_product[awal-1].Ukuran) == target {
		awal--
	}

	// menampilkan semua produk dengan ukuran yang sama
	for i := awal; i < n && nilai_ukuran(data_product[i].Ukuran) == target; i++ {
		fmt.Printf(
			"ID: %d | Nama: %s | Kategori: %s | Ukuran: %s | Warna: %s | Stok: %d | Harga: Rp%.0f\n",
			data_product[i].ID,
			data_product[i].Nama,
			data_product[i].Kategori,
			data_product[i].Ukuran,
			data_product[i].Warna,
			data_product[i].Stok,
			data_product[i].Harga,
		)
	}

}

func menu_selection_sort() {
	if n == 0 {
		fmt.Println("Data produk kosong")
		return
	}

	var pilihan int

	fmt.Println("\n=== SELECTION SORT ===")
	fmt.Println("1. Harga Termurah -> Termahal")
	fmt.Println("2. Harga Termahal -> Termurah")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		selection_sort_asc()
		fmt.Println("\nData berhasil diurutkan berdasarkan harga (ASC)")
		read_product()

	case 2:
		selection_sort_harga_desc()
		fmt.Println("\nData berhasil diurutkan berdasarkan harga (DESC)")
		read_product()

	default:
		fmt.Println("Pilihan tidak valid")
	}
}

// sorting berdasarkan harga
func selection_sort_asc() {
	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if data_product[j].Harga < data_product[min].Harga {
				min = j
			}
		}

		data_product[i], data_product[min] =
			data_product[min], data_product[i]
	}
}

func selection_sort_harga_desc() {
	for i := 0; i < n-1; i++ {
		max := i

		for j := i + 1; j < n; j++ {
			if data_product[j].Harga > data_product[max].Harga {
				max = j
			}
		}

		data_product[i], data_product[max] =
			data_product[max], data_product[i]
	}
}

func menu_insertion_sort() {
	if n == 0 {
		fmt.Println("Data produk kosong")
		return
	}

	var pilihan int

	fmt.Println("\n=== INSERTION SORT ===")
	fmt.Println("1. Stok Terkecil -> Terbesar")
	fmt.Println("2. Stok Terbesar -> Terkecil")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		insertion_sort_asc()
		fmt.Println("\nData berhasil diurutkan berdasarkan stok (ASC)")
		read_product()

	case 2:
		insertion_sort_desc()
		fmt.Println("\nData berhasil diurutkan berdasarkan stok (DESC)")
		read_product()

	default:
		fmt.Println("Pilihan tidak valid")
	}
}

// sorting berdasarkan stock
func insertion_sort_asc() {
	for i := 1; i < n; i++ {
		temp := data_product[i]
		j := i - 1

		for j >= 0 && data_product[j].Stok > temp.Stok {
			data_product[j+1] = data_product[j]
			j--
		}

		data_product[j+1] = temp
	}
}

func insertion_sort_desc() {
	for i := 1; i < n; i++ {
		temp := data_product[i]
		j := i - 1

		for j >= 0 && data_product[j].Stok < temp.Stok {
			data_product[j+1] = data_product[j]
			j--
		}

		data_product[j+1] = temp
	}
}

func statistik_produk_terlaris() {
	if n == 0 {
		fmt.Println("Data Masih Kosong")
		return
	}

	maksimal_stock := data_product[0].Stok
	indeks := 0

	for i := 0; i < n; i++ {
		if data_product[i].Stok > maksimal_stock {
			maksimal_stock = data_product[i].Stok
			indeks = i
		}
	}

	fmt.Println("\n===== STATISTIK PRODUK PALING LARIS =====")
	fmt.Printf("ID Produk      : %d\n", data_product[indeks].ID)
	fmt.Printf("Nama Produk    : %s\n", data_product[indeks].Nama)
	fmt.Printf("Kategori       : %s\n", data_product[indeks].Kategori)
	fmt.Printf("Ukuran         : %s\n", data_product[indeks].Ukuran)
	fmt.Printf("Warna          : %s\n", data_product[indeks].Warna)
	fmt.Printf("Jumlah Stok    : %d\n", data_product[indeks].Stok)
	fmt.Printf("Harga          : Rp%.0f\n", data_product[indeks].Harga)
	fmt.Println("=========================================")
}

// func total_stok_per_kategori() {
// 	if n == 0 {
// 		fmt.Println("Data Produk Masih Kosong")
// 	}

// 	var total_stok [5]int

// 	for i := 0; i < n; i++ {
// 		for j := 0; i < len(daftar_kategori); i++ {
// 			total_stok[j] += data_product[i].Stok
// 		}
// 	}

// 	fmt.Println("\n===== TOTAL STOK PER KATEGORI =====")

// 	for i := 0; i < len(daftar_kategori); i++ {
// 		fmt.Printf(
// 			"%-20s : %d stok\n",
// 			daftar_kategori[i],
// 			total_stok[i],
// 		)
// 	}

// 	fmt.Println("===================================")
// }

func filter_harga() {
	var min, max float64

	fmt.Println("\n========================================================")
	fmt.Println("         FILTER PRODUK BERDASARKAN HARGA")
	fmt.Println("========================================================")

	fmt.Print("Masukkan Nilai Minimal harga : ")
	fmt.Scan(&min)

	fmt.Print("Masukkan Maksimal harga : ")
	fmt.Scan(&max)

	ditemukan := false
	total := 0

	fmt.Println("\n=====================================================================================")
	fmt.Printf(
		"%-6s %-24s %-16s %-8s %-10s\n",
		"ID",
		"Nama Produk",
		"Kategori",
		"Stok",
		"Harga",
	)
	fmt.Println("=====================================================================================")

	for i := 0; i < n; i++ {
		if data_product[i].Harga >= min && data_product[i].Harga <= max {
			fmt.Printf(
				"%-6d %-24s %-16s %-8d Rp%-10.0f\n",
				data_product[i].ID,
				data_product[i].Nama,
				data_product[i].Kategori,
				data_product[i].Stok,
				data_product[i].Harga,
			)

			ditemukan = true
			total++
		}
	}

	fmt.Println("=====================================================================================")

	if ditemukan {

		fmt.Printf(
			"\nTotal Produk Ditemukan : %d\n",
			total,
		)

	} else {

		fmt.Println("\nProduk tidak ditemukan")
	}
}

func tambah_kategori() {
	var kategori_baru string

	fmt.Print("Masukkan kategori baru : ")
	fmt.Scan(&kategori_baru)

	if cari_kategori(kategori_baru) {
		fmt.Println("Kategori sudah ada")
		return
	}

	daftar_kategori = append(daftar_kategori, kategori_baru)
	fmt.Println("Kategori berhasil ditambahkan")
}

func tambah_ukuran() {
	var ukuran_baru string

	fmt.Print("Masukkan ukuran baru : ")
	fmt.Scan(&ukuran_baru)

	ukuran_baru = strings.ToUpper(ukuran_baru)

	if cari_ukuran(ukuran_baru) {
		fmt.Println("Ukuran sudah ada")
		return
	}

	daftar_ukuran = append(daftar_ukuran, ukuran_baru)
	fmt.Println("Ukuran berhasil ditambahkan")
}

func header() {
	fmt.Println("+==========================================+")
	fmt.Println("|            APLIKASI SiFASHION             |")
	fmt.Println("| SISTEM MANAJEMEN INVENTARIS PRODUK FASHION|")
	fmt.Println("+==========================================+")
}

func menu() {
	header()
	fmt.Println("+======================================+")
	fmt.Println("|          SILAHKAN PILIH MENU         |")
	fmt.Println("+======================================+")
	fmt.Println("| 1. Tambah Produk                     |")
	fmt.Println("| 2. Tampil Produk                     |")
	fmt.Println("| 3. Edit Produk                       |")
	fmt.Println("| 4. Hapus Produk                      |")
	fmt.Println("| 5. Sequential Search                 |")
	fmt.Println("| 6. Binary Search                     |")
	fmt.Println("| 7. Selection Sort                    |")
	fmt.Println("| 8. Insertion Sort                    |")
	fmt.Println("| 9. Statistik                         |")
	fmt.Println("| 10. Filter Harga          		    |")
	fmt.Println("| 11. Tambah Kategori                  |")
	fmt.Println("| 12. Tambah Ukuran                   	|")
	fmt.Println("| 0. Keluar                            |")
	fmt.Println("+======================================+")
}

func main() {
	var pilih int

	for {
		menu()
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			create_product()
		case 2:
			read_product()
		case 3:
			edit_product()
		case 4:
			delete_product()
		case 5:
			sequential()
		case 6:
			binary_search()
		case 7:
			menu_selection_sort()
		case 8:
			menu_insertion_sort()
		case 9:
			statistik_produk_terlaris()
		case 10:
			filter_harga()
		case 11:
			tambah_kategori()
		case 12:
			tambah_ukuran()
		case 0:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}
