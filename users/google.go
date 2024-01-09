package users

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/clp-runner/runner-user/configs"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"io/ioutil"
	"net/http"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
var store = session.New()

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func Login(c fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := randToken()
	sess.Set("state", state)
	sess.Save()

	url := configs.AppConfig.GoogleLoginConfig.AuthCodeURL(state)

	c.Status(fiber.StatusSeeOther)
	err = c.Redirect().To(url)
	if err != nil {
		return err
	}
	return c.JSON(url)
}

func Callback(c fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := sess.Get("state")

	sess.Delete("state")
	sess.Save()

	if state != c.Query("state") {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	googlecon := configs.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	// postgresql 에 사용자 정보 저장
	// JWT 생성 후 쿠키 저장

	return c.SendString(string(userData))
}

func Logout(c fiber.Ctx) error {
	return nil
}
