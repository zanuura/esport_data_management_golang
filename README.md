## **Aplikasi Pengelolaan Data E-Sports Tournament**

---

### **Pendahuluan**

Tugas Besar ini bertujuan untuk membuat sebuah aplikasi konsol sederhana yang berfungsi untuk mengelola data turnamen e-sports. Aplikasi ini dirancang untuk digunakan oleh panitia turnamen maupun pemain yang ingin memantau klasemen. Fungsionalitas utama mencakup manajemen data tim, pencatatan hasil pertandingan, serta penyajian statistik performa tim dan pencarian data menggunakan algoritma dasar struktur data seperti pencarian dan pengurutan.

---

### **Deskripsi TuBes**

Aplikasi ini merupakan program berbasis terminal menggunakan bahasa pemrograman Go (`Golang`) yang memiliki fitur-fitur berikut:

#### **Fitur Utama:**

1. **Manajemen Tim:**

   * Tambah, ubah, dan hapus data tim.
   * Setiap tim menyimpan data: nama, jumlah pertandingan, kemenangan, kekalahan, skor yang dicetak, dan skor yang diterima.

2. **Pencatatan Hasil Pertandingan:**

   * Pengguna dapat memasukkan hasil pertandingan antara dua tim.
   * Sistem akan secara otomatis memperbarui data tim yang terlibat.

3. **Klasemen Otomatis:**

   * Menampilkan urutan tim berdasarkan jumlah kemenangan dan selisih skor (skor masuk dikurangi skor kebobolan).

4. **Pencarian Tim:**

   * Mencari tim berdasarkan nama menggunakan metode **Sequential Search** dan **Binary Search** (binary disesuaikan jika diperlukan berdasarkan urutan).

5. **Pengurutan Tim:**

   * Pengguna dapat memilih antara **Selection Sort** dan **Insertion Sort**.
   * Kriteria pengurutan: jumlah kemenangan, kekalahan, jumlah pertandingan, dan selisih skor.

6. **Statistik Tim Terbaik:**

   * Menampilkan 3 tim dengan **win rate tertinggi**.

7. **Riwayat Pertandingan dan Tabel Tim:**

   * Melihat semua data pertandingan dan daftar tim yang tercatat.

---

### **Tantangan dan Solusi**

| Tantangan                                                 | Solusi                                                                                                                           |
| --------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| Penyimpanan data secara permanen belum tersedia           | Saat ini data disimpan sementara di variabel global slice. Solusi ke depan bisa menggunakan file `.json` atau integrasi database |
| Binary Search hanya bekerja jika data telah terurut       | Implementasi pencarian tetap menggunakan sequential search agar fleksibel.                                                       |
| Validasi input pengguna masih manual                      | Penanganan kesalahan dilakukan dengan `strconv.Atoi` dan pengecekan input kosong atau tidak valid                                |
| Struktur pengurutan cukup kompleks karena banyak kriteria | Dibuat fungsi `getCompareFunc` untuk membangun logika perbandingan berdasarkan input pengguna                                    |

---

### **Kesimpulan dan Rekomendasi**

Aplikasi ini berhasil memenuhi seluruh spesifikasi fungsional:

* Pengelolaan tim dan hasil pertandingan,
* Klasemen dan statistik performa,
* Pencarian dan pengurutan dengan algoritma klasik.

Namun, masih terdapat peluang pengembangan lebih lanjut seperti:

* **Penyimpanan data permanen** (menggunakan file atau database).
* **Antarmuka pengguna grafis (GUI)** menggunakan web (React/HTML) atau desktop (TUI/GTK).
* **Login multi-user dan peran (admin, penonton)**.
* **Penghitungan statistik pemain individu** jika data pemain ditambahkan.



---



