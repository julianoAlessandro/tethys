package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"senha" binding:"required"`
}

var db *gorm.DB

func main() {
	// Inicializa o banco de dados SQLite
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	r := gin.Default()

	// Rota para lidar com o envio do formulário de registro
	r.POST("/register", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verifica se o nome de usuário já está em uso
		var existingUser User
		if result := db.Where("email = ?", newUser.Email).First(&existingUser); result.Error == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Nome de usuário já em uso"})
			return
		}

		// Hash da senha antes de armazenar
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash de senha"})
			return
		}

		newUser.Password = string(hashedPassword)

		// Armazena o novo usuário no banco de dados
		db.Create(&newUser)

		// Responde com sucesso
		c.JSON(http.StatusOK, gin.H{"message": "Usuário registrado com sucesso"})
	})

	// Rota para lidar com o envio do formulário de login
	r.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Username string `json:"email" binding:"required"`
			Password string `json:"senha" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Busca o usuário no banco de dados
		var user User
		if result := db.Where("username = ?", loginData.Username).First(&user); result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
			return
		}

		// Verifica a senha
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
			return
		}

		// Autenticação bem-sucedida
		c.JSON(http.StatusOK, gin.H{"message": "Login bem-sucedido", "email": user.Email})
	})

	r.Run(":8080")
}
