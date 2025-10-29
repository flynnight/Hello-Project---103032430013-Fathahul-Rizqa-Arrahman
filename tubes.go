package main

import "fmt"
// Ini komentar tambahan untuk tugas Git
const NMAX int = 100
type peminjam struct {
	nama string
	pinjaman float64
	tenor int
	bunga float64
	cicilan float64
	status string
	// revisi penambahan id
	id string
}

type tabPinjam [NMAX]peminjam

func pow(x float64, y int) float64 {
	hasil := 1.0
	for i:=0; i<y; i++ {
		hasil *= x
	}
	return hasil
}

func hitungCicilan(pinjaman float64, bunga float64, tenor int) float64 {
	bungaBulanan := bunga / 12 / 100
	denominator := (1 - (1 / pow(1+bungaBulanan, tenor)))
	if denominator == 0 {
		if bungaBulanan == 0 {
			if tenor == 0 { return pinjaman }
			return pinjaman / float64(tenor)
		}
		return 0
	}
	return pinjaman * bungaBulanan / denominator
}

func tambahPeminjam(A *tabPinjam, n *int) {
// revisi meminta input id unik dan mengecek apakah sudah digunakan
	var p peminjam

	if *n >= 100 {
		fmt.Println("❌ Data penuh.")
		return
	}

	fmt.Print("ID Peminjam (Unik): ")
	fmt.Scan(&p.id)
	clearInput()

	for i:=0; i<*n; i++ {
		if A[i].id == p.id {
			fmt.Print("❌ ID sudah digunakan. Harap gunakan ID lain.")
			return
		}
	}

	fmt.Print("Nama Peminjam: ")
	fmt.Scan(&p.nama)
	clearInput()

	fmt.Print("Jumlah Pinjaman: ")
	fmt.Scan(&p.pinjaman)
	clearInput()

	fmt.Print("Tenor (bulan): ")
	fmt.Scan(&p.tenor)
	clearInput()

	fmt.Print("Bunga Tahunan (%): ")
	fmt.Scan(&p.bunga)
	clearInput()

	fmt.Print("Status Pembayaran: ")
	fmt.Scan(&p.status)
	clearInput()

	p.cicilan = hitungCicilan(p.pinjaman, p.bunga, p.tenor)
	A[*n] = p
	*n++

	fmt.Println("✅ Data berhasil ditambahkan.")
}

func tampilkanPeminjam(A tabPinjam, n int) {
// revisi menambahkan tampilan ID peminjam
	if n == 0 {
		fmt.Println("ℹ️ Belum ada data peminjam.")
		return
	}

	fmt.Println("\n--- Daftar Peminjam ---")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s | ID: %s | Pinjaman: %.2f | Tenor: %d | Cicilan: %.2f | Status: %s\n",
			i+1, A[i].nama, A[i].id, A[i].pinjaman, A[i].tenor, A[i].cicilan, A[i].status)
	}
}

func ubahPeminjam(A *tabPinjam, n int) {
// revisi mengganti identifikasi dari nama menjadi id
	var id string

	if n == 0 {
		fmt.Println("ℹ️ Belum ada data untuk diubah.")
		return
	}

	fmt.Print("Masukkan ID peminjam yang ingin diubah: ")
	fmt.Scanf("%s\n", &id)
	
	found := false
	for i := 0; i < n; i++ {
		if A[i].id == id {
			fmt.Printf("Nama Peminjam: %s\n", A[i].nama)
			fmt.Print("Jumlah Pinjaman Baru: ")
			fmt.Scan(&A[i].pinjaman)
			clearInput()
			fmt.Print("Tenor Baru: ")
			fmt.Scan(&A[i].tenor)
			clearInput()
			fmt.Print("Bunga Baru: ")
			fmt.Scan(&A[i].bunga)
			clearInput()
			fmt.Print("Status Pembayaran Baru: ")
			fmt.Scan(&A[i].status)
			clearInput()
			A[i].cicilan = hitungCicilan(A[i].pinjaman, A[i].bunga, A[i].tenor)
			fmt.Println("✅ Data berhasil diperbarui.")
			found = true
			return
		}
	}

	if !found {
		fmt.Println("❌ Data tidak ditemukan.")
	}
}

func hapusPeminjam(A *tabPinjam, n *int) {
// revisi mengganti identifikasi dari nama menjadi id
	var id string

	if *n == 0 {
		fmt.Println("ℹ️ Belum ada data untuk dihapus.")
		return
	}

	fmt.Print("Masukkan ID peminjam yang ingin dihapus: ")
	fmt.Scan(&id)
	clearInput()
	
	found := false
	for i := 0; i < *n; i++ {
		if A[i].id == id {
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			fmt.Println("✅ Data berhasil dihapus.")
			found = true
			return
		}
	}
	if !found {
		fmt.Println("❌ Data tidak ditemukan.")
	}
}

func sequentialSearch(A tabPinjam, n int) {
// revisi mengganti identifikasi dari nama menjadi id
	var id string

	if n == 0 {
		fmt.Println("ℹ️ Belum ada data untuk dicari.")
		return
	}

	fmt.Print("Masukkan ID yang dicari: ")
	fmt.Scan(&id)
	clearInput()

	for i := 0; i < n; i++ {
		if A[i].id == id {
			fmt.Printf("✅ Data ditemukan:\nNama: %s\nID: %s\nPinjaman: %.2f\nTenor: %d\nBunga: %.2f%%\nCicilan: %.2f\nStatus: %s\n", A[i].nama, A[i].id, A[i].pinjaman, A[i].tenor, A[i].bunga, A[i].cicilan, A[i].status)
			return
		}
	}
	fmt.Println("❌ Data tidak ditemukan.")
}

func binarySearch(A tabPinjam, n int) {
// revisi menambahkan pemanggilan fungsi InsertionSortNama agar selalu mengurutkan nama terlebih dahulu ketika memanggil fungsi
	var nama string
	var left, right, mid int
	
	insertionSortNama(&A, n)

	if n == 0 {
		fmt.Println("ℹ️ Belum ada data untuk dicari.")
		return
	}

	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	clearInput()

	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if A[mid].nama == nama {
			fmt.Printf("✅ Data ditemukan:\nNama: %s\nID: %s\nPinjaman: %.2f\nTenor: %d\nBunga: %.2f%%\nCicilan: %.2f\nStatus: %s\n", A[mid].nama, A[mid].id, A[mid].pinjaman, A[mid].tenor, A[mid].bunga, A[mid].cicilan, A[mid].status)
			return
		} else if A[mid].nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("❌ Data tidak ditemukan.")
}

func insertionSortNama(A *tabPinjam, n int) {
	var i, pass int
	var temp peminjam

	if n < 2 {
		fmt.Println("ℹ️ Tidak cukup data untuk diurutkan.")
		return
	}

	pass = 1
	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.nama < A[i-1].nama {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
	fmt.Println("✅ Data diurutkan berdasarkan nama.")
}

func selectionSortPinjaman(A *tabPinjam, n int) {
	var i, idxMin, pass int
	var temp peminjam

	if n < 2 {
		fmt.Println("ℹ️ Tidak cukup data untuk diurutkan.")
		return
	}

	for pass = 0; pass < n-1; pass++ {
		idxMin = pass
		for i = pass + 1; i < n; i++ {
			if A[i].pinjaman < A[idxMin].pinjaman {
				idxMin = i
			}
		}
		if idxMin != pass {
			temp = A[pass]
			A[pass] = A[idxMin]
			A[idxMin] = temp
		}
	}
	fmt.Println("✅ Data diurutkan berdasarkan jumlah pinjaman.")
}

func laporanTotalPinjaman(A tabPinjam, n int) {
	var total float64
	var i int

	if n == 0 {
		fmt.Println("ℹ️ Belum ada data pinjaman.")
		return
	}

	total = 0.0
	for i = 0; i < n; i++ {
		total += A[i].pinjaman
	}
	fmt.Printf("\n📊 Total Seluruh Pinjaman: %.2f\n", total)
}

func menu() {

	var pilih, nData int
	var data tabPinjam

	for {
		fmt.Println()
		fmt.Println("=== MENU PINJAMAN ===")
		fmt.Println(" 1. Tambah Pinjaman")
		fmt.Println(" 2. Tampilkan Semua Pinjaman")
		fmt.Println(" 3. Ubah Data Peminjam")
		fmt.Println(" 4. Hapus Data Peminjam")
		fmt.Println(" 5. Urutkan Berdasarkan Nama (Ascending)")
		fmt.Println(" 6. Urutkan Berdasarkan Jumlah Pinjaman (Ascending)")
		fmt.Println(" 7. Cari Peminjam Berdasarkan ID (Sequential Search)")
		fmt.Println(" 8. Cari Peminjam (Binary Search)")
		fmt.Println(" 9. Laporan Total Pinjaman")
		fmt.Println(" 0. Keluar")
		fmt.Print(" Pilih menu: ")
		fmt.Scan(&pilih)
		clearInput()
		switch pilih {
		case 1: 
			tambahPeminjam(&data, &nData)
		case 2:
			tampilkanPeminjam(data, nData)
		case 3:
			ubahPeminjam(&data, nData)
		case 4:
			hapusPeminjam(&data, &nData)
		case 5:
			insertionSortNama(&data, nData)
		case 6:
			selectionSortPinjaman(&data, nData)
		case 7:
			sequentialSearch(data, nData)
		case 8:
			binarySearch(data, nData)
		case 9:
			laporanTotalPinjaman(data, nData)
		case 0:
			fmt.Println("👋 Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid.")
		}
	}
}

func main() {
	register()
	menu()
}

func register() {

	var user, pass string
	fmt.Println("---Register User---")
	fmt.Print("Masukkan username: ")
	fmt.Scanln(&user)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&pass)

	fmt.Printf("✅ Registrasi berhasil. Selamat datang, %s!\n", user)
}

func clearInput() {
	var s string
	fmt.Scanln(&s)
}