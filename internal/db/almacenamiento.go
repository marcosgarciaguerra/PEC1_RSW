package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"pec2/internal/models"
	_ "modernc.org/sqlite"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func InitDB() {
	var err error
	os.MkdirAll(filepath.Join("internal", "db"), 0755)
	dbPath := filepath.Join("internal", "db", "platinum.db")

	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("Error abriendo base de datos SQLite:", err)
	}

	createTables()
}

func createTables() {
	userTableInfo := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL,
		apellidos TEXT,
		fecha_nacimiento TEXT,
		direccion TEXT,
		telefono TEXT,
		correo TEXT UNIQUE,
		documento TEXT,
		metodo_pago TEXT,
		numero_pago TEXT,
		password TEXT NOT NULL,
		suscripcion_activa BOOLEAN DEFAULT 1
	);`

	_, err := DB.Exec(userTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla usuarios: ", err)
	}

	resenaTableInfo := `
	CREATE TABLE IF NOT EXISTS resenas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		autor TEXT NOT NULL,
		puntuacion INTEGER NOT NULL,
		texto TEXT NOT NULL
	);`

	_, err = DB.Exec(resenaTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla resenas: ", err)
	}

	sesionTableInfo := `
	CREATE TABLE IF NOT EXISTS sesiones (
		id TEXT PRIMARY KEY,
		correo TEXT NOT NULL
	);`

	_, err = DB.Exec(sesionTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla sesiones: ", err)
	}

	var count int
	DB.QueryRow("SELECT COUNT(*) FROM resenas").Scan(&count)
	if count == 0 {
		_, _ = DB.Exec(`INSERT INTO resenas (autor, puntuacion, texto) VALUES 
		('Carlos G.', 5, 'El mejor gimnasio de la ciudad. Las máquinas son increíbles.'),
		('María P.', 4, 'Muy buen ambiente, aunque a veces hay mucha gente en hora punta.'),
		('Luis R.', 5, 'Los fisioterapeutas me curaron una lesión de hombro que llevaba meses arrastrando.')`)
	}

	CrearTablaArticulos()
}

func GuardarUsuario(u models.Usuario) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare(`INSERT INTO usuarios (nombre, apellidos, fecha_nacimiento, direccion, telefono, correo, documento, metodo_pago, numero_pago, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Nombre, u.Apellidos, u.FechaNacimiento, u.Direccion, u.Telefono, u.Correo, u.Documento, u.MetodoPago, u.NumeroPago, string(hash))
	return err
}

func ObtenerResenas() []models.Resena {
	var resenas []models.Resena
	rows, err := DB.Query("SELECT autor, puntuacion, texto FROM resenas ORDER BY id DESC")
	if err != nil {
		log.Println("Error obteniendo reseñas:", err)
		return resenas
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Resena
		if err := rows.Scan(&r.Autor, &r.Puntuacion, &r.Texto); err == nil {
			resenas = append(resenas, r)
		}
	}
	return resenas
}

func GuardarResena(r models.Resena) error {
	stmt, err := DB.Prepare("INSERT INTO resenas (autor, puntuacion, texto) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.Autor, r.Puntuacion, r.Texto)
	return err
}

func ObtenerSocioPorCorreo(correo string) *models.Socio {
	var s models.Socio
	err := DB.QueryRow("SELECT id, nombre, password, suscripcion_activa FROM usuarios WHERE correo = ?", correo).
		Scan(&s.ID, &s.Nombre, &s.Contrasena, &s.SuscripcionActiva)
	if err != nil {
		return nil
	}
	return &s
}

func ObtenerSocioPorNombre(nombre string) *models.Socio {
	var s models.Socio
	err := DB.QueryRow("SELECT id, nombre, password, suscripcion_activa FROM usuarios WHERE nombre = ?", nombre).
		Scan(&s.ID, &s.Nombre, &s.Contrasena, &s.SuscripcionActiva)
	if err != nil {
		return nil
	}
	return &s
}

func ObtenerUsuarioPorCorreo(correo string) *models.Usuario {
	var u models.Usuario
	err := DB.QueryRow("SELECT id, nombre, apellidos, correo, direccion, telefono, metodo_pago FROM usuarios WHERE correo = ?", correo).
		Scan(&u.ID, &u.Nombre, &u.Apellidos, &u.Correo, &u.Direccion, &u.Telefono, &u.MetodoPago)
	if err != nil {
		return nil
	}
	return &u
}

func CrearSesion(idSesion, correo string) error {
	_, err := DB.Exec("INSERT INTO sesiones (id, correo) VALUES (?, ?)", idSesion, correo)
	return err
}

func ObtenerCorreoPorSesion(idSesion string) string {
	var correo string
	err := DB.QueryRow("SELECT correo FROM sesiones WHERE id = ?", idSesion).Scan(&correo)
	if err != nil {
		return ""
	}
	return correo
}

func EliminarSesion(idSesion string) error {
	_, err := DB.Exec("DELETE FROM sesiones WHERE id = ?", idSesion)
	return err
}
