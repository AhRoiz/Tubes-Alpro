package main

import (
	"fmt"
)

// --- KONSTANTA DAN TIPE DATA ---

const MAX_MENU_PER_HARI = 10
const JUMLAH_HARI = 7
const JUMLAH_MENU = 5

type MenuItem struct {
	Nama  string
	Harga int
}

type Penjualan struct {
	Menu                    MenuItem
	Jumlah, TotalPendapatan int
}

type DataHarian struct {
	Hari                                     string
	PenjualanMenu                            [MAX_MENU_PER_HARI]Penjualan
	JumlahMenuTerjual, TotalPendapatanHarian int
}

type tabMingguan [JUMLAH_HARI]DataHarian
type tabMenuItem [JUMLAH_MENU]MenuItem

// --- FUNGSI UTAMA ---

func main() {
	var dataMingguan tabMingguan
	var daftarMenuTersedia tabMenuItem

	daftarMenuTersedia[0].Nama = "Bakso"
	daftarMenuTersedia[0].Harga = 13000

	daftarMenuTersedia[1].Nama = "Es Teh"
	daftarMenuTersedia[1].Harga = 3000

	daftarMenuTersedia[2].Nama = "Mie Ayam"
	daftarMenuTersedia[2].Harga = 12000

	daftarMenuTersedia[3].Nama = "Nasi Goreng"
	daftarMenuTersedia[3].Harga = 15000

	daftarMenuTersedia[4].Nama = "Soto Ayam"
	daftarMenuTersedia[4].Harga = 10000

	inisialisasiData(&dataMingguan)

	fmt.Println("=== Aplikasi Pendataan Penjualan Menu Mingguan ===")

	var pilihan string
	for pilihan != "0" {
		fmt.Println("\n============ MENU UTAMA ============")
		fmt.Println("1. Tambah Data Penjualan")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Cari Data")
		fmt.Println("0. Keluar")

		var valid bool = false
		for !valid {
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihan)
			if pilihan >= "0" && pilihan <= "3" {
				valid = true
			} else {
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
		}

		prosesMenuUtama(pilihan, &dataMingguan, daftarMenuTersedia)
	}
}

func inisialisasiData(A *tabMingguan) {
	A[0].Hari = "Senin"
	A[1].Hari = "Selasa"
	A[2].Hari = "Rabu"
	A[3].Hari = "Kamis"
	A[4].Hari = "Jumat"
	A[5].Hari = "Sabtu"
	A[6].Hari = "Minggu"
}

func prosesMenuUtama(pilihan string, A *tabMingguan, MenuTersedia tabMenuItem) {
	if pilihan == "1" {
		tambahData(A, MenuTersedia)
	} else if pilihan == "2" {
		tampilkanData(A, MenuTersedia)
	} else if pilihan == "3" {
		menuCariData(A, MenuTersedia)
	} else if pilihan == "0" {
		fmt.Println("\nTerima kasih telah menggunakan aplikasi!")
	}
}

// --- FITUR 1: TAMBAH DATA ---

func tambahData(A *tabMingguan, MenuTersedia tabMenuItem) {
	var pilihanHari int
	var idxHari int = -1

	fmt.Println("\nPilih hari untuk menambah data (1-7), atau 0 untuk kembali:")
	var i int = 0
	for i < JUMLAH_HARI {
		fmt.Printf("%d. %s\n", i+1, A[i].Hari)
		i++
	}
	fmt.Println("0. Kembali")

	var valid bool = false
	for !valid {
		fmt.Print("\nPilihan Anda: ")
		fmt.Scan(&pilihanHari)
		if pilihanHari >= 0 && pilihanHari <= JUMLAH_HARI {
			valid = true
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}

	if pilihanHari > 0 {
		idxHari = pilihanHari - 1

		var lanjutTambahMenu bool = true
		for lanjutTambahMenu {
			fmt.Printf("\n--- Menambah Data untuk Hari %s ---\n", A[idxHari].Hari)

			var pilihanMenu int
			fmt.Println("Pilih menu yang ingin ditambahkan datanya:")
			i = 0
			for i < JUMLAH_MENU {
				var menu MenuItem = MenuTersedia[i]
				fmt.Printf("%d. %s (Rp %d)\n", i+1, menu.Nama, menu.Harga)
				i++
			}
			fmt.Println("0. Selesai menambah menu untuk hari ini")

			var validMenu bool = false
			for !validMenu {
				fmt.Print("\nPilihan Anda: ")
				fmt.Scan(&pilihanMenu)
				if pilihanMenu >= 0 && pilihanMenu <= JUMLAH_MENU {
					validMenu = true
				} else {
					fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
				}
			}

			if pilihanMenu > 0 {
				var idxMenu int = pilihanMenu - 1
				var menuTerpilih MenuItem = MenuTersedia[idxMenu]
				var jumlah int
				fmt.Printf("Masukkan jumlah penjualan untuk %s: ", menuTerpilih.Nama)
				fmt.Scan(&jumlah)

				var indeksData int = -1
				var menuDitemukan bool = false
				i = 0
				for i < A[idxHari].JumlahMenuTerjual && !menuDitemukan {
					if A[idxHari].PenjualanMenu[i].Menu.Nama == menuTerpilih.Nama {
						indeksData = i
						menuDitemukan = true
					}
					i++
				}

				if indeksData != -1 {
					var dataLama *Penjualan = &A[idxHari].PenjualanMenu[indeksData]
					A[idxHari].TotalPendapatanHarian -= dataLama.TotalPendapatan
					dataLama.Jumlah += jumlah
					dataLama.TotalPendapatan = dataLama.Jumlah * dataLama.Menu.Harga
					A[idxHari].TotalPendapatanHarian += dataLama.TotalPendapatan
					fmt.Println("Data berhasil diperbarui.")
				} else {
					if A[idxHari].JumlahMenuTerjual < MAX_MENU_PER_HARI {
						var n int = A[idxHari].JumlahMenuTerjual
						var dataBaru *Penjualan = &A[idxHari].PenjualanMenu[n]
						dataBaru.Menu = menuTerpilih
						dataBaru.Jumlah = jumlah
						dataBaru.TotalPendapatan = jumlah * menuTerpilih.Harga
						A[idxHari].TotalPendapatanHarian += dataBaru.TotalPendapatan
						A[idxHari].JumlahMenuTerjual++
						fmt.Println("Data berhasil ditambahkan.")
					} else {
						fmt.Println("Data menu untuk hari ini sudah penuh.")
					}
				}
			} else {
				lanjutTambahMenu = false
			}
		}
	}
}

// --- FITUR 2: TAMPILKAN DATA ---

func tampilkanData(A *tabMingguan, MenuTersedia tabMenuItem) {
	var pilihan string
	for pilihan != "0" {
		fmt.Println("\n--- Menu Tampilkan Data ---")
		fmt.Println("1. Tampilkan Rincian Penjualan per Hari")
		fmt.Println("2. Tampilkan Total Pendapatan per Hari")
		fmt.Println("3. Urutkan & Tampilkan Total Tertinggi per Hari (Insertion Sort)")
		fmt.Println("4. Tampilkan Menu Terlaris Mingguan (Selection Sort)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilihan)

		if pilihan == "1" {
			tampilkanRincianPenjualan(A, MenuTersedia)
		} else if pilihan == "2" {
			tampilkanTotalPendapatanPerHari(A)
		} else if pilihan == "3" {
			urutkanDanTampilkanTotalTertinggi(A)
		} else if pilihan == "4" {
			tampilkanMenuTerlarisMingguan(A, MenuTersedia)
		} else if pilihan != "0" {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanRincianPenjualan(A *tabMingguan, MenuTersedia tabMenuItem) {
	fmt.Println("\n--- Rincian Penjualan per Hari ---")
	var i int = 0
	for i < JUMLAH_HARI {
		var hari DataHarian = A[i]
		fmt.Printf("\n====================== %s ======================\n", hari.Hari)

		var j int = 0
		for j < JUMLAH_MENU {
			var menuTersedia MenuItem = MenuTersedia[j]
			var jumlahTerjual int = 0
			var pendapatanMenu int = 0
			var ditemukan bool = false

			var k int = 0
			for k < hari.JumlahMenuTerjual && !ditemukan {
				var penjualanHarian Penjualan = hari.PenjualanMenu[k]
				if penjualanHarian.Menu.Nama == menuTersedia.Nama {
					jumlahTerjual = penjualanHarian.Jumlah
					pendapatanMenu = penjualanHarian.TotalPendapatan
					ditemukan = true
				}
				k++
			}

			fmt.Printf("%-2d. %-12s jumlah terjual %-4d pendapatan Rp %d\n", j+1, menuTersedia.Nama, jumlahTerjual, pendapatanMenu)
			j++
		}

		fmt.Printf("--------------------------------------------------\n")
		fmt.Printf("Total Pendapatan Hari Ini: Rp %d\n", hari.TotalPendapatanHarian)
		fmt.Println()
		i++
	}
}

func tampilkanTotalPendapatanPerHari(A *tabMingguan) {
	fmt.Println("\n=== Total Pendapatan per Hari (diurutkan berdasarkan hari) ===")
	var i int = 0
	for i < JUMLAH_HARI {
		fmt.Printf("%d. %-7s: Rp %d\n", i+1, A[i].Hari, A[i].TotalPendapatanHarian)
		i++
	}
	fmt.Println()
}

func urutkanDanTampilkanTotalTertinggi(A *tabMingguan) {
	fmt.Println("\n=== Total Penjualan Tertinggi per Hari (Diurutkan dengan Insertion Sort) ===")
	var dataUrut tabMingguan = *A
	var i int = 1
	for i < JUMLAH_HARI {
		var temp DataHarian = dataUrut[i]
		var j int = i - 1
		for j >= 0 && dataUrut[j].TotalPendapatanHarian < temp.TotalPendapatanHarian {
			dataUrut[j+1] = dataUrut[j]
			j--
		}
		dataUrut[j+1] = temp
		i++
	}
	i = 0
	for i < JUMLAH_HARI {
		fmt.Printf("%d. %-7s: Rp %d\n", i+1, dataUrut[i].Hari, dataUrut[i].TotalPendapatanHarian)
		i++
	}
	fmt.Println()
}

func tampilkanMenuTerlarisMingguan(A *tabMingguan, MenuTersedia tabMenuItem) {
	fmt.Println("\n=== Menu dengan Penjualan Terbanyak per Minggu (Selection Sort) ===")

	var totalPenjualanMenu [JUMLAH_MENU]Penjualan
	var i int = 0
	for i < JUMLAH_MENU {
		totalPenjualanMenu[i] = Penjualan{
			Menu:            MenuTersedia[i],
			Jumlah:          0,
			TotalPendapatan: 0,
		}
		i++
	}

	i = 0
	for i < JUMLAH_HARI {
		var j int = 0
		for j < A[i].JumlahMenuTerjual {
			var penjualanHarian Penjualan = A[i].PenjualanMenu[j]
			var k int = 0
			var ditemukan bool = false
			for k < JUMLAH_MENU {
				// Kondisi hanya akan berjalan jika menu cocok DAN belum ditemukan sebelumnya
				if totalPenjualanMenu[k].Menu.Nama == penjualanHarian.Menu.Nama && !ditemukan {
					totalPenjualanMenu[k].Jumlah += penjualanHarian.Jumlah
					totalPenjualanMenu[k].TotalPendapatan += penjualanHarian.TotalPendapatan
					ditemukan = true
				}
				k++
			}
			j++
		}
		i++
	}

	var n int = JUMLAH_MENU
	i = 0
	for i < n-1 {
		var maxIdx int = i
		var j int = i + 1
		for j < n {
			if totalPenjualanMenu[j].Jumlah > totalPenjualanMenu[maxIdx].Jumlah {
				maxIdx = j
			}
			j++
		}
		// Fix: Corrected typo from totalPenjualan to totalPenjualanMenu
		totalPenjualanMenu[i], totalPenjualanMenu[maxIdx] = totalPenjualanMenu[maxIdx], totalPenjualanMenu[i]
		i++
	}

	i = 0
	for i < JUMLAH_MENU {
		var p Penjualan = totalPenjualanMenu[i]
		fmt.Printf("%d. %-12s: Total terjual %-4d (Total Pendapatan Rp %d)\n", i+1, p.Menu.Nama, p.Jumlah, p.TotalPendapatan)
		i++
	}
	fmt.Println()
}

// --- FITUR 3: CARI DATA ---

func menuCariData(A *tabMingguan, MenuTersedia tabMenuItem) {
	var pilihan string
	for pilihan != "0" {
		fmt.Println("\n=== Menu Pencarian Data ===")
		fmt.Println("1. Cari Data per Hari (Sequential Search)")
		fmt.Println("2. Lacak Penjualan per Menu (Binary Search)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilihan)

		if pilihan == "1" {
			cariDataPerHariSequential(A)
		} else if pilihan == "2" {
			lacakPenjualanPerMenuBinary(A, MenuTersedia)
		} else if pilihan != "0" {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func cariDataPerHariSequential(A *tabMingguan) {
	fmt.Println("\nPilih hari yang ingin dicari datanya:")
	var i int = 0
	for i < JUMLAH_HARI {
		fmt.Printf("%d. %s\n", i+1, A[i].Hari)
		i++
	}
	fmt.Println("0. Kembali")

	var pilihan int
	var valid bool = false
	for !valid {
		fmt.Print("\nPilihan Anda: ")
		fmt.Scan(&pilihan)
		if pilihan >= 0 && pilihan <= JUMLAH_HARI {
			valid = true
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}

	if pilihan > 0 {
		var idxHari int = -1
		var ditemukan bool = false
		i = 0
		for i < JUMLAH_HARI && !ditemukan {
			if (i + 1) == pilihan {
				idxHari = i
				ditemukan = true
			}
			i++
		}

		if ditemukan {
			var hari DataHarian = A[idxHari]
			fmt.Printf("\n=== Detail Penjualan Hari %s ===\n", hari.Hari)
			fmt.Printf("Total Pendapatan: Rp %d\n", hari.TotalPendapatanHarian)
			fmt.Println("Rincian per Menu:")
			if hari.JumlahMenuTerjual == 0 {
				fmt.Println("  - Tidak ada penjualan tercatat.")
			} else {
				var j int = 0
				for j < hari.JumlahMenuTerjual {
					var p Penjualan = hari.PenjualanMenu[j]
					fmt.Printf("  - %-12s: %d terjual (Total Rp %d)\n", p.Menu.Nama, p.Jumlah, p.TotalPendapatan)
					j++
				}
			}
			fmt.Println()
		} else {
			fmt.Println("Hari tidak ditemukan.")
		}
	}
}

func lacakPenjualanPerMenuBinary(A *tabMingguan, MenuTersedia tabMenuItem) {
	fmt.Println("\nPilih menu yang ingin dilacak:")
	var i int = 0
	for i < JUMLAH_MENU {
		fmt.Printf("%d. %s\n", i+1, MenuTersedia[i].Nama)
		i++
	}
	fmt.Println("0. Kembali")

	var pilihan int
	var valid bool = false
	for !valid {
		fmt.Print("\nPilihan Anda: ")
		fmt.Scan(&pilihan)
		if pilihan >= 0 && pilihan <= JUMLAH_MENU {
			valid = true
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}

	if pilihan > 0 {
		var namaMenu string = MenuTersedia[pilihan-1].Nama
		var idxMenu int = -1
		var low int = 0
		var high int = JUMLAH_MENU - 1
		for low <= high {
			var mid int = low + (high-low)/2
			if MenuTersedia[mid].Nama == namaMenu {
				idxMenu = mid
				low = high + 1
			} else if MenuTersedia[mid].Nama < namaMenu {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		if idxMenu != -1 {
			var menuTerpilih MenuItem = MenuTersedia[idxMenu]
			fmt.Printf("\n=== Laporan Penjualan untuk: %s ===\n", menuTerpilih.Nama)
			var penjualanDitemukan bool = false
			var totalTerjual int = 0
			var totalPendapatanMenu int = 0
			i = 0
			for i < JUMLAH_HARI {
				var hari DataHarian = A[i]
				var j int = 0
				for j < hari.JumlahMenuTerjual {
					var penjualan Penjualan = hari.PenjualanMenu[j]
					if penjualan.Menu.Nama == menuTerpilih.Nama {
						fmt.Printf("- %s: %d terjual (Rp %d)\n", hari.Hari, penjualan.Jumlah, penjualan.TotalPendapatan)
						penjualanDitemukan = true
						totalTerjual += penjualan.Jumlah
						totalPendapatanMenu += penjualan.TotalPendapatan
					}
					j++
				}
				i++
			}
			if !penjualanDitemukan {
				fmt.Println("Tidak ada data penjualan yang ditemukan untuk menu ini.")
			} else {
				fmt.Println()
				fmt.Printf("Total Terjual Seminggu: %d\n", totalTerjual)
				fmt.Printf("Total Pendapatan Seminggu: Rp %d\n", totalPendapatanMenu)
			}
			fmt.Println()
		} else {
			fmt.Println("Terjadi kesalahan: Menu tidak ditemukan.")
		}
	}
}