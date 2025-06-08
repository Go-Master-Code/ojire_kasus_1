Komentar kasus 1:
•	Komentar 1:
Pada output aplikasi, terdapat hasil berikut:
x Gagal menambahkan fruit KURMA (ID: 5): ID sudah ada, tidak boleh duplikat
x Gagal menambahkan fruit Salak (ID: 5): ID sudah ada, tidak boleh duplikat
Kasus tersebut terjadi karena ID:5 sudah terlebih dahulu dipakai oleh Jeruk Bali. Untuk menjaga konsistensi data, maka 2 data berikutnya (menggunakan ID:5) ditolak dan tidak disimpan karena field ID haruslah bersifat unik atau di dalam konsep database field ID haruslah menjadi primary key.
•	Komentar 2:
Terkait dengan komentar 1, seharusnya diberikan sistem penomoran auto increment pada field ID agar penomoran ID dilakukan secara otomatis oleh sistem tanpa adanya resiko duplikasi data ID (yang seharusnya merupakan primary key)
•	Komentar 3:
Diperlukan konsistensi standar penulisan nama buah (fruitName) misalnya, data teks harus tersimpan dalam mode All Caps ataupun Small Caps. Pada contoh kasus ini misalnya terdapat data Apel (ID:1) dan apel (ID:3) yang merupakan contoh inkonsistensi penulisan data. Bisa saja kedua item tersebut merupakan item yang sama namun terinput 2x atau keduanya adalah item berbeda varian namun tidak dituliskan secara detil, misalnya Apel Washington dan Apel Fuji RRC. Inilah pentingnya untuk membuat standarisasi penulisan data untuk menghindari kesalahpahaman.
•	Komentar 4:
Masih terkait konsistensi data, pada contoh kasus ini terdapat data Kurma (ID:2) dan KURMA (ID:5). Jika diteliti lebih dalam keduanya memiliki stock yang sama yaitu 20. Ada kemungkinan bahwa data yang sama terinput sebanyak dua kali. Kasus ini bisa kembali menjadi kendala berkaitan dengan duplikasi data.
