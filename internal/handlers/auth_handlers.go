package handlers

import (
	"encoding/json"
	"fmt"
	"google-auth-go-v2/internal/infra"
	"google-auth-go-v2/internal/models"
	"google-auth-go-v2/internal/repository"
	"google-auth-go-v2/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type googleAuthResponse struct {
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Id            string `json:"id"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	LocalId       string `json:"locale"`
	VerifiedEmail bool   `json:"verified_email"`
}

func LoginHandler(c *fiber.Ctx) error {
	state, err := utils.GenerateRandomString()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate random string",
		})
	}

	session, err := infra.GetSessionStore().Get(c)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get session",
		})
	}

	session.Set("state", state)
	if err := session.Save(); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save session",
		})
	}

	url := infra.GetAuthProvider().AuthCodeURL(state)

	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func CallbackHandler(c *fiber.Ctx) error {
	state := c.Query("state")
	session, err := infra.GetSessionStore().Get(c)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get session",
		})
	}

	if state != session.Get("state") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid state",
		})
	}

	code := c.Query("code")
	token, err := infra.GetAuthProvider().Exchange(c.Context(), code)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to exchange token",
		})
	}

	client := infra.GetAuthProvider().Client(c.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user info",
		})
	}
	defer resp.Body.Close()

	authResp := googleAuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode user info",
		})
	}

	user, err := repository.GetById(authResp.Id)

	if err != nil {
		user = &models.User{
			Id: authResp.Id,
		}
		updateUserFields(user, authResp)

		err = repository.Create(user)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}
	} else {
		if user.Email != authResp.Email ||
			user.FirstName != authResp.GivenName ||
			user.LastName != authResp.FamilyName ||
			user.Picture != authResp.Picture {
			updateUserFields(user, authResp)
			err = repository.Update(user)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to update user",
				})
			}
		}
	}

	claims := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error signing token",
		})
	}

	return c.JSON(fiber.Map{"token": t})
}

func updateUserFields(user *models.User, authResp googleAuthResponse) {
	user.Email = authResp.Email
	user.FirstName = authResp.GivenName
	user.LastName = authResp.FamilyName
	user.Picture = authResp.Picture
}
