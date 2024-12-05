package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"utils"

	"github.com/gin-gonic/gin"
)

// RegisterUser registers a new user in the database using RegistrationInfo
func RegisterUser(c *gin.Context) {
	var registrationInfo RegistrationInfo

	// Bind the incoming JSON to the registrationInfo struct
	if err := c.ShouldBindJSON(&registrationInfo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Verify password format
	err := utils.ValidatePasswordFormat(registrationInfo.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Password must be between 8 and 20 characters"})
		return
	}

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(registrationInfo.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// Insert the user into the database
	query := `INSERT INTO users (username, email, hashed_password) VALUES (?, ?, ?)`
	_, err = DB.Exec(query, registrationInfo.Name, registrationInfo.Email, string(hashedPassword))
	if err != nil {
		fmt.Println("Error inserting user:", err)
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

func GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid user ID:", err)
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	// Query the database for the user by ID
	query := `SELECT id, username, email, created_at FROM users WHERE id = ?`
	row := DB.QueryRow(query, id)

	var user User
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": fmt.Sprintf("User with ID %d not found", id)})
		} else {
			log.Println("Error fetching user:", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(200, user)
}

// LoginUser handles user login by verifying the password and issuing a JWT
func LoginUser(c *gin.Context) {
	var loginInfo LoginInfo

	// Bind the incoming JSON to the login info struct
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Retrieve user from the database
	var user User
	var hashedPassword string
	query := `SELECT id, username, email, hashed_password FROM users WHERE email = ?`
	row := DB.QueryRow(query, loginInfo.Email)

	// Scan the retrieved data into the User struct
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			// Return an error if no user is found
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		// Return any other errors
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// Compare the provided password with the stored hashed password
	passwordMatch := utils.CheckPasswordHash(loginInfo.Password, hashedPassword)
	if !passwordMatch {
		// If the password doesn't match, return an error
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create the JWT token
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the JWT token to the client
	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

// ChangePassword allows a user to change their password
func ChangePassword(c *gin.Context) {
	var input struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	// Get the user ID from the context (set by middleware after authentication)
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	fmt.Println("userId", userId)

	// Bind the incoming JSON to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Verify password format
	if err := utils.ValidatePasswordFormat(input.NewPassword); err != nil {
		c.JSON(400, gin.H{"error": "Password must be between 8 and 20 characters"})
		return
	}

	// Retrieve the user's current hashed password from the database
	var hashedPassword string
	query := `SELECT hashed_password FROM users WHERE id = ?`
	err := DB.QueryRow(query, userId).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			log.Println("Error querying user:", err)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Verify the old password
	passwordMatch := utils.CheckPasswordHash(input.OldPassword, hashedPassword)
	if !passwordMatch {
		c.JSON(400, gin.H{"error": "Old password is incorrect"})
		return
	}

	// Hash the new password
	newHashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		log.Println("Error hashing new password:", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// Update the password in the database
	updateQuery := `UPDATE users SET hashed_password = ? WHERE id = ?`
	_, err = DB.Exec(updateQuery, newHashedPassword, userId)
	if err != nil {
		log.Println("Error updating password:", err)
		c.JSON(500, gin.H{"error": "Failed to change password"})
		return
	}

	c.JSON(200, gin.H{"message": "Password changed successfully"})
}
