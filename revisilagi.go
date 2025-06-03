package main

import "fmt"

const NMAX = 5

type Item struct {
	Nama     string
	Jenis    string
	Kotor    int
	Bersih   int
	Jumlah   int
	TotalKot int
	TotalBrs int
}

type DaftarMenu [NMAX]Item
type DataTransaksi [NMAX]Item

func main() {
	menu := initMenu()
	var transaksi DataTransaksi
	var jumlahData int
	var sudahDipilih [NMAX]bool

	jumlahInput := inputJumlahData()
	for i := 0; i < jumlahInput; i++ {
		fmt.Printf("\nData ke-%d\n", i+1)
		tampilkanMenuTersisa(menu, sudahDipilih)

		var pilihan int
		fmt.Print("Pilih nomor menu: ")
		fmt.Scan(&pilihan)

		// Validasi input
		if pilihan < 1 || pilihan > NMAX {
			fmt.Println("Pilihan tidak valid.")
			i--
			continue
		}
		if sudahDipilih[pilihan-1] {
			fmt.Println("Menu tersebut sudah dipilih sebelumnya. Silakan pilih menu lain.")
			i--
			continue
		}

		var jumlah int
		fmt.Printf("Masukkan jumlah terjual untuk %s: ", menu[pilihan-1].Nama)
		fmt.Scan(&jumlah)

		tambahTransaksi(menu[pilihan-1], jumlah, &transaksi, &jumlahData)
		sudahDipilih[pilihan-1] = true
	}

	urutkanBersihDesc(&transaksi, jumlahData)
	tampilkanHasil(transaksi, jumlahData)
	tampilkanKategori(transaksi, jumlahData)
}

// FUNGSI 1: Inisialisasi daftar menu
func initMenu() DaftarMenu {
	return DaftarMenu{
		{"ayam geprek", "makanan", 17000, 2000, 0, 0, 0},
		{"ayam bakar", "makanan", 18000, 3000, 0, 0, 0},
		{"ayam serundeng", "makanan", 18000, 4000, 0, 0, 0},
		{"es cendol", "minuman", 12000, 4000, 0, 0, 0},
		{"es teh", "minuman", 4000, 2000, 0, 0, 0},
	}
}

// FUNGSI 2: Input jumlah jenis data
func inputJumlahData() int {
	var n int
	fmt.Printf("Masukkan jumlah jenis menu yang terjual (maksimal %d): ", NMAX)
	fmt.Scan(&n)
	if n < 1 || n > NMAX {
		fmt.Println("Jumlah tidak valid.")
		return inputJumlahData()
	}
	return n
}

// FUNGSI 3: Tampilkan menu yang belum dipilih
func tampilkanMenuTersisa(menu DaftarMenu, sudahDipilih [NMAX]bool) {
	fmt.Println("\n=== Menu yang Tersedia ===")
	for i := 0; i < NMAX; i++ {
		if !sudahDipilih[i] {
			fmt.Printf("%d. %s (%s) - Kotor: Rp%d | Bersih: Rp%d\n",
				i+1, menu[i].Nama, menu[i].Jenis, menu[i].Kotor, menu[i].Bersih)
		}
	}
}

// FUNGSI 4: Tambah transaksi
func tambahTransaksi(item Item, jumlah int, transaksi *DataTransaksi, jumlahData *int) {
	item.Jumlah = jumlah
	item.TotalKot = jumlah * item.Kotor
	item.TotalBrs = jumlah * item.Bersih
	transaksi[*jumlahData] = item
	*jumlahData++
}

// FUNGSI 5: Urutkan transaksi berdasarkan bersih (descending)
func urutkanBersihDesc(transaksi *DataTransaksi, jumlahData int) {
	for i := 0; i < jumlahData-1; i++ {
		for j := i + 1; j < jumlahData; j++ {
			if transaksi[i].TotalBrs < transaksi[j].TotalBrs {
				transaksi[i], transaksi[j] = transaksi[j], transaksi[i]
			}
		}
	}
}

// FUNGSI 6: Tampilkan hasil akhir transaksi
func tampilkanHasil(transaksi DataTransaksi, jumlahData int) {
	fmt.Println("\n=== Hasil Penjualan ===")
	var totalKotor, totalBersih int
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("- %s (%s): %d terjual | Kotor: Rp%d | Bersih: Rp%d\n",
			transaksi[i].Nama, transaksi[i].Jenis, transaksi[i].Jumlah, transaksi[i].TotalKot, transaksi[i].TotalBrs)
		totalKotor += transaksi[i].TotalKot
		totalBersih += transaksi[i].TotalBrs
	}
	fmt.Println("--------------------------------------")
	fmt.Printf("Total Pendapatan Kotor: Rp%d\n", totalKotor)
	fmt.Printf("Total Pendapatan Bersih: Rp%d\n", totalBersih)
}

// FUNGSI 7: Tampilkan total per kategori
func tampilkanKategori(transaksi DataTransaksi, jumlahData int) {
	var totalMakanan, totalMinuman int
	for i := 0; i < jumlahData; i++ {
		if transaksi[i].Jenis == "makanan" {
			totalMakanan += transaksi[i].TotalBrs
		} else if transaksi[i].Jenis == "minuman" {
			totalMinuman += transaksi[i].TotalBrs
		}
	}
	fmt.Println("\n=== Total Bersih per Kategori ===")
	fmt.Printf("Makanan: Rp%d\n", totalMakanan)
	fmt.Printf("Minuman: Rp%d\n", totalMinuman)
}
