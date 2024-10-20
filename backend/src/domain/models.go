package domain

type CardType struct {
	ID          int64  `json:"id"`          // ID del tipo de tarjeta
	Name        string `json:"name"`        // Nombre del tipo de tarjeta
	Description string `json:"description"` // Descripción del tipo de tarjeta
}

type User struct {
	ID        int64  `json:"id"`         // ID del usuario
	Name      string `json:"name"`       // Nombre del usuario
	Password  string `json:"password"`   // Contraseña del usuario
	BirthDate string `json:"birth_date"` // Fecha de nacimiento
	Gender    string `json:"gender"`     // Género del usuario
	Cards     []Card `json:"cards"`      // Relación uno a muchos con Cards
}

type Card struct {
	ID        int64    `json:"id"`         // ID de la tarjeta
	Title     string   `json:"title"`      // Título de la tarjeta
	Content   string   `json:"content"`    // Contenido de la tarjeta
	IsDefault bool     `json:"is_default"` // Si es tarjeta predeterminada
	UserID    int64    `json:"user_id"`    // ID del usuario que posee la tarjeta (FK a User)
	TypeID    int64    `json:"type_id"`    // ID del tipo de tarjeta (FK a CardType)
	Type      CardType `json:"type"`       // Relación con el tipo de tarjeta
}

type GenderUser int

const (
	Male   GenderUser = iota + 1
	Female GenderUser = iota
	Other  GenderUser = iota
)

func (g GenderUser) String() string {
	return [...]string{"Male", "Female", "Other"}[g-1]
}

func (g GenderUser) EnumIndex() int {
	return int(g)
}

