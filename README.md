API Pembuatan Rekening Pinjaman
API ini menyediakan endpoint untuk membuat rekening pinjaman dan menghasilkan tabel angsurannya.

Fitur :
Membuat rekening pinjaman dengan parameter kustom :
- Plafond
- Lama Pinjaman (bulan)
- Suku Bunga (tahunan)
- Tanggal Mulai Angsuran

Menghasilkan tabel angsuran yang termasuk :
- Nomor bulan
- Angsuran pokok
- Angsuran bunga
- Total angsuran 
- Sisa pinjaman

Prasyarat :
- Golang (versi [1.22.1])
- PostgreSQL (versi 13)