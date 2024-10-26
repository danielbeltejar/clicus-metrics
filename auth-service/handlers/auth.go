package handlers

import (
	"auth-service/models"
	"auth-service/utils"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var (
	userCollection    *mongo.Collection
	jwtSecret         string
	allowUserCreation bool
)

// InitAuthService initializes the auth service with MongoDB URI, JWT secret, and user creation flag.
func InitAuthService(mongoURI, secret string, allowCreation bool) {
	jwtSecret = secret
	allowUserCreation = allowCreation
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	userCollection = client.Database("clicusmetrics").Collection("users")
}

// @Summary Register a new user
// @Description Create a new user account with a username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} map[string]string "message": "User created successfully"
// @Failure 403 {object} map[string]string "error": "User registration disabled"
// @Failure 500 {object} map[string]string "error": "Error creating user"
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	if !allowUserCreation {
		http.Error(w, "User registration disabled", http.StatusForbidden)
		return
	}

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body map[string]string true "Login request"
// @Success 200 {object} map[string]string "token": "JWT token"
// @Failure 401 {object} map[string]string "error": "Invalid username or password"
// @Failure 500 {object} map[string]string "error": "Error generating token"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&loginRequest)

	var user models.User
	err := userCollection.FindOne(context.TODO(), map[string]string{"username": loginRequest.Username}).Decode(&user)
	if err != nil || !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username, jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
