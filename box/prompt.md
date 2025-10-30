# INGAT, INI CUMAN FILE, GAK ADA FOLDER
# SEBAGIAN COMMAND YANG KAMU KASIH, AKU HAPUS KARENA KURASA KURANG PENTING
# AKU CUMAN MINTA HANDLER, GAK MINTA COBRA COMMAND

ini adalah salah satu contoh handler, yang akan menangani encrypt si cli
func Add(ctx *fiber.Ctx) (err error) {
	file, err := ctx.FormFile("")
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "", err)
	}

	hash, err := storage.Box.Upload(file)
	if err != nil {
		return
	}

	if encrypt := ctx.FormValue("encrypt", "false"); encrypt == "true" {
		return db.DB.Create(&db.File{
			Name:    file.Filename,
			Hash:    hash,
			Encrypt: true,
			UserID:  auth.UserId(ctx),
		}).Error
	}

	return db.DB.Create(&db.File{
		Name:   file.Filename,
		Hash:   hash,
		UserID: auth.UserId(ctx),
	}).Error
}

aku punya model :
package db

import "gorm.io/gorm"

type File struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `gorm:"not null;index"`
	Name    string `gorm:"size:255;not null"`
	Hash    string
	Encrypt bool
	gorm.Model
}

ini db :
package db

import "gorm.io/gorm"

var DB *gorm.DB

interface untuk storage, ini udah jadi, gak oerlu dibuay ulang,bisa dikases global pakai obejvt Box:
type Storage interface {
	Upload(file *multipart.FileHeader) (string, error)
	Get(hash string) (io.ReadCloser, error)
	Delete(hash string) error
}

kita akan support command yang ini aja:
# FILE OPERATIONS
add         // Upload file dengan optional encryption
get         // Download file dari storage, auto-decrypt jika perlu  
list        // List semua files dengan metadata dasar
delete      // Hapus file permanen dari system
copy        // Duplikat file
info        // Menampilkan detail lengkap sebuah file
size        // Check size file
count       // Hitung total files milik user

# ENCRYPTION & SECURITY
encrypt     // Encrypt file yang sudah ada di system
decrypt     // Decrypt file untuk processing
lock        // Kunci akses ke file dengan password
unlock      // Buka kunci file dengan password
share       // Generate temporary share link dengan expiry
revoke      // Batalkan share link sebelum waktunya

# ORGANIZATION & METADATA
tag         // Tambah atau remove tags dari file
meta        // Edit custom metadata file
pin         // Pin important files untuk quick access
unpin       // Unpin files

# BATCH OPERATIONS  
import      // Bulk import dari local folder
export      // Export files ke local directory
cleanup     // Hapus files berdasarkan rules, yang ini gak usah dulu, bisa kubuat sendiri

# PRIVACY MANAGEMENT
hide        // Sembunyikan file dari public listing
unhide      // Tampilkan hidden files
trash // pindahkan file ke trash, ini biaa kubuat sendiri
recover     // Emergency recovery options

# SYSTEM & MAINTENANCE
duplicates  // Find duplicate files by content

# CLI PRODUCTIVITY
recent      // List recently accessed files
select      // Interactive file picker
pipe        // Output file ke stdout untuk piping, ini akan kubuat sendiri
quick-add   // Auto-generate filename & tags, ini bisa nanti dulu
history     // Command execution history

# DEVELOPER TOOLS
export-json // Export data sebagai JSON

# BACKUP & RECOVERY, ini aku masih bingung jadi nanti aja
snapshot    // Create system snapshot
export-all  // Export semua data + metadata
