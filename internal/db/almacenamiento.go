package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"pec2/internal/models"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
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
	migrarDatosSensiblesLegacy()
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
		documento TEXT UNIQUE,
		plan TEXT DEFAULT 'basico',
		metodo_pago TEXT,
		numero_pago TEXT,
		password TEXT NOT NULL,
		suscripcion_activa BOOLEAN DEFAULT 1
	);`

	_, err := DB.Exec(userTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla usuarios: ", err)
	}

	// Asegurar compatibilidad si la tabla ya existía sin la restricción UNIQUE
	_, _ = DB.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_usuarios_documento ON usuarios (documento);")
	// Migración ligera para bases de datos creadas antes de añadir el plan.
	_, _ = DB.Exec("ALTER TABLE usuarios ADD COLUMN plan TEXT DEFAULT 'basico';")

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

	clasesTableInfo := `
	CREATE TABLE IF NOT EXISTS clases (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL,
		entrenador TEXT NOT NULL,
		aforo INTEGER NOT NULL,
		horario TEXT NOT NULL,
		descripcion TEXT NOT NULL,
		lugar TEXT NOT NULL
	);`

	_, err = DB.Exec(clasesTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla clases: ", err)
	}

	reservasTableInfo := `
	CREATE TABLE IF NOT EXISTS reservas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		socio_id INTEGER NOT NULL,
		actividad_id INTEGER NOT NULL,
		fecha_asist TEXT NOT NULL,
		FOREIGN KEY(socio_id) REFERENCES usuarios(id),
		FOREIGN KEY(actividad_id) REFERENCES clases(id)
	);`

	_, err = DB.Exec(reservasTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla reservas: ", err)
	}

	pedidosTableInfo := `
	CREATE TABLE IF NOT EXISTS pedidos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		usuario_id INTEGER NOT NULL,
		fecha TEXT NOT NULL,
		direccion_envio TEXT NOT NULL,
		metodo_pago TEXT,
		total REAL NOT NULL,
		estado TEXT NOT NULL,
		FOREIGN KEY(usuario_id) REFERENCES usuarios(id)
	);`

	_, err = DB.Exec(pedidosTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla pedidos: ", err)
	}

	pedidoItemsTableInfo := `
	CREATE TABLE IF NOT EXISTS pedido_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		pedido_id INTEGER NOT NULL,
		articulo_id INTEGER NOT NULL,
		nombre TEXT NOT NULL,
		precio_unitario REAL NOT NULL,
		cantidad INTEGER NOT NULL,
		subtotal REAL NOT NULL,
		FOREIGN KEY(pedido_id) REFERENCES pedidos(id)
	);`

	_, err = DB.Exec(pedidoItemsTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla pedido_items: ", err)
	}

	sessionsTableInfo := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		usuario_id INTEGER NOT NULL,
		token_hash TEXT UNIQUE NOT NULL,
		expires_at DATETIME NOT NULL,
		revoked_at DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(usuario_id) REFERENCES usuarios(id)
	);`
	_, err = DB.Exec(sessionsTableInfo)
	if err != nil {
		log.Fatal("Error creando tabla sessions: ", err)
	}

	_, err = DB.Exec("CREATE INDEX IF NOT EXISTS idx_sessions_token_hash ON sessions(token_hash);")
	if err != nil {
		log.Fatal("Error creando indice de sessions: ", err)
	}

	var count int
	DB.QueryRow("SELECT COUNT(*) FROM resenas").Scan(&count)
	if count == 0 {
		_, _ = DB.Exec(`INSERT INTO resenas (autor, puntuacion, texto) VALUES 
		('Carlos G.', 5, 'El mejor gimnasio de la ciudad. Las máquinas son increíbles.'),
		('María P.', 4, 'Muy buen ambiente, aunque a veces hay mucha gente en hora punta.'),
		('Luis R.', 5, 'Los fisioterapeutas me curaron una lesión de hombro que llevaba meses arrastrando.')`)
	}

	var countClases int
	DB.QueryRow("SELECT COUNT(*) FROM clases").Scan(&countClases)
	if countClases == 0 {
		_, _ = DB.Exec(`INSERT INTO clases (nombre, entrenador, aforo, horario, descripcion, lugar) VALUES 
		('Ciclo Indoor (Spinning)', 'Carlos Méndez', 20, 'Lunes, 13 de abril de 2026 a las 18:00h', 'Entrenamiento cardiovascular en bicicleta estática.', 'Sala de Spinning (Planta Baja)'),
		('Yoga Vinyasa', 'Elena Rostova', 15, 'Martes, 14 de abril de 2026 a las 09:00h', 'Práctica fluida y dinámica para mejorar flexibilidad.', 'Sala Zen (Planta Alta)'),
		('HIIT', 'Marcos Silva', 12, 'Miércoles, 15 de abril de 2026 a las 19:30h', 'Entrenamiento de intervalos de alta intensidad.', 'Zona Funcional'),
		('BodyPump', 'Laura Gómez', 20, 'Jueves, 16 de abril de 2026 a las 14:00h', 'Tonificación muscular con barras y discos.', 'Sala de Actividades Dirigidas 1'),
		('Zumba', 'Sofía Valdés', 25, 'Viernes, 17 de abril de 2026 a las 18:30h', 'Baile aeróbico con ritmos latinos.', 'Sala de Actividades Dirigidas 2'),
		('Pilates Mat', 'Javier Ruiz', 15, 'Sábado, 18 de abril de 2026 a las 10:00h', 'Ejercicios para fortalecer core y postura.', 'Sala Zen (Planta Alta)'),
		('Fitboxing', 'David Castro', 12, 'Domingo, 19 de abril de 2026 a las 11:00h', 'Mix de boxeo al saco sin contacto con cardio.', 'Zona de Combate'),
		('AquaGym', 'Ana Morales', 20, 'Lunes, 20 de abril de 2026 a las 08:30h', 'Gimnasia aeróbica de bajo impacto en el agua.', 'Piscina Climatizada')`)
	}
}

func migrarDatosSensiblesLegacy() {
	rows, err := DB.Query("SELECT id, password, metodo_pago, numero_pago FROM usuarios")
	if err != nil {
		log.Println("No se pudo revisar migración de datos sensibles:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var password, metodoPago, numeroPago string
		if err := rows.Scan(&id, &password, &metodoPago, &numeroPago); err != nil {
			continue
		}

		if !strings.HasPrefix(password, "$2") && strings.TrimSpace(password) != "" {
			hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if hashErr == nil {
				_, _ = DB.Exec("UPDATE usuarios SET password = ? WHERE id = ?", string(hash), id)
			}
		}

		if strings.TrimSpace(numeroPago) == "" {
			continue
		}
		if strings.HasPrefix(numeroPago, "****") || strings.HasPrefix(numeroPago, "IBAN-****") {
			continue
		}

		clean := strings.ReplaceAll(strings.TrimSpace(numeroPago), " ", "")
		if len(clean) >= 4 {
			last4 := clean[len(clean)-4:]
			masked := "****" + last4
			if metodoPago == "domiciliacion" {
				masked = "IBAN-****" + strings.ToUpper(last4)
			}
			_, _ = DB.Exec("UPDATE usuarios SET numero_pago = ? WHERE id = ?", masked, id)
		} else {
			_, _ = DB.Exec("UPDATE usuarios SET numero_pago = ? WHERE id = ?", "REDACTED", id)
		}
	}
}

func GuardarUsuario(u models.Usuario) error {
	stmt, err := DB.Prepare(`INSERT INTO usuarios (nombre, apellidos, fecha_nacimiento, direccion, telefono, correo, documento, plan, metodo_pago, numero_pago, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Nombre, u.Apellidos, u.FechaNacimiento, u.Direccion, u.Telefono, u.Correo, u.Documento, u.Plan, u.MetodoPago, u.NumeroPago, u.Password)
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
		Scan(&s.ID, &s.Nombre, &s.ContrasenaHash, &s.SuscripcionActiva)
	if err != nil {
		return nil
	}
	return &s
}

func ObtenerSocioPorNombre(nombre string) *models.Socio {
	var s models.Socio
	err := DB.QueryRow("SELECT id, nombre, password, suscripcion_activa FROM usuarios WHERE nombre = ?", nombre).
		Scan(&s.ID, &s.Nombre, &s.ContrasenaHash, &s.SuscripcionActiva)
	if err != nil {
		return nil
	}
	return &s
}

func ObtenerUsuarioPorCorreo(correo string) *models.Usuario {
	var u models.Usuario
	err := DB.QueryRow("SELECT id, nombre, apellidos, correo, direccion, telefono, plan, metodo_pago, numero_pago FROM usuarios WHERE correo = ?", correo).
		Scan(&u.ID, &u.Nombre, &u.Apellidos, &u.Correo, &u.Direccion, &u.Telefono, &u.Plan, &u.MetodoPago, &u.NumeroPago)
	if err != nil {
		return nil
	}
	return &u
}

func ObtenerUsuarioPorID(id int) *models.Usuario {
	var u models.Usuario
	err := DB.QueryRow("SELECT id, nombre, apellidos, correo, direccion, telefono, plan, metodo_pago, numero_pago FROM usuarios WHERE id = ?", id).
		Scan(&u.ID, &u.Nombre, &u.Apellidos, &u.Correo, &u.Direccion, &u.Telefono, &u.Plan, &u.MetodoPago, &u.NumeroPago)
	if err != nil {
		return nil
	}
	return &u
}

func ObtenerSocioPorID(id int) *models.Socio {
	var s models.Socio
	err := DB.QueryRow("SELECT id, nombre, password, suscripcion_activa FROM usuarios WHERE id = ?", id).
		Scan(&s.ID, &s.Nombre, &s.ContrasenaHash, &s.SuscripcionActiva)
	if err != nil {
		return nil
	}
	return &s
}

func ActualizarPasswordHash(usuarioID int, passwordHash string) error {
	_, err := DB.Exec("UPDATE usuarios SET password = ? WHERE id = ?", passwordHash, usuarioID)
	return err
}

func CrearSesion(usuarioID int, tokenHash string, expiresAt time.Time) error {
	_, err := DB.Exec(
		"INSERT INTO sessions (usuario_id, token_hash, expires_at) VALUES (?, ?, ?)",
		usuarioID, tokenHash, expiresAt.UTC().Format(time.RFC3339),
	)
	return err
}

func RevocarSesion(tokenHash string) error {
	_, err := DB.Exec("UPDATE sessions SET revoked_at = ? WHERE token_hash = ? AND revoked_at IS NULL", time.Now().UTC().Format(time.RFC3339), tokenHash)
	return err
}

func ObtenerUsuarioPorTokenSesion(tokenHash string) *models.Usuario {
	var usuarioID int
	var expiresAtStr string
	var revokedAt sql.NullString
	err := DB.QueryRow(
		"SELECT usuario_id, expires_at, revoked_at FROM sessions WHERE token_hash = ? ORDER BY id DESC LIMIT 1",
		tokenHash,
	).Scan(&usuarioID, &expiresAtStr, &revokedAt)
	if err != nil || revokedAt.Valid {
		return nil
	}
	expiresAt, err := time.Parse(time.RFC3339, expiresAtStr)
	if err != nil || time.Now().UTC().After(expiresAt) {
		return nil
	}
	return ObtenerUsuarioPorID(usuarioID)
}

func ObtenerSocioPorTokenSesion(tokenHash string) *models.Socio {
	var usuarioID int
	var expiresAtStr string
	var revokedAt sql.NullString
	err := DB.QueryRow(
		"SELECT usuario_id, expires_at, revoked_at FROM sessions WHERE token_hash = ? ORDER BY id DESC LIMIT 1",
		tokenHash,
	).Scan(&usuarioID, &expiresAtStr, &revokedAt)
	if err != nil || revokedAt.Valid {
		return nil
	}
	expiresAt, err := time.Parse(time.RFC3339, expiresAtStr)
	if err != nil || time.Now().UTC().After(expiresAt) {
		return nil
	}
	return ObtenerSocioPorID(usuarioID)
}

func GuardarPedido(p *models.Pedido) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.Exec(
		`INSERT INTO pedidos (usuario_id, fecha, direccion_envio, metodo_pago, total, estado) VALUES (?, ?, ?, ?, ?, ?)`,
		p.UsuarioID, p.Fecha, p.DireccionEnvio, p.MetodoPago, p.Total, p.Estado,
	)
	if err != nil {
		return err
	}

	pedidoID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(pedidoID)

	stmt, err := tx.Prepare(`INSERT INTO pedido_items (pedido_id, articulo_id, nombre, precio_unitario, cantidad, subtotal) VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range p.Items {
		_, err = stmt.Exec(p.ID, item.ArticuloID, item.Nombre, item.PrecioUnitario, item.Cantidad, item.Subtotal)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
