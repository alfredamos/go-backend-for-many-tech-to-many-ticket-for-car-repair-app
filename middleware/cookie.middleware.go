package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx, tokenPath, tokenName, tokenValue string, expiresIn time.Time) {
	c.Cookie(&fiber.Cookie{
		Name:     tokenName,
		Value:    tokenValue,
		Path:     tokenPath,
		Domain:   "localhost",
		Expires:  expiresIn, //time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
	})

}

func GetCookie(c *fiber.Ctx, tokenName string) string {
	return c.Cookies(tokenName)
}

func DeleteCookie(c *fiber.Ctx, tokenName, tokenPath string) {

	c.Cookie(&fiber.Cookie{
		Name:     tokenName,
		Value:    "",                               // Clear the value
		Expires:  time.Now().Add(-3 * time.Second), // Set an expired time for older browsers (optional but recommended)
		MaxAge:   -1,                               // Set MaxAge to a negative value to delete the cookie immediately
		HTTPOnly: true,                             // Should match the original cookie's HttpOnly setting
		Path:     tokenPath,                        // Should match the original cookie's Path setting
		Domain:   "localhost",                      // Should match the original cookie's Domain setting (or leave empty for host cookie)
		Secure:   false,
	})
}
