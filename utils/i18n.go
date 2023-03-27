package utils

import (
	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var I18nBundle goi18n.Bundle

func Load() {
	// Create a new language bundle with the default language (English).
	bundle := goi18n.NewBundle(language.English)
	// Register toml unmarshal function.
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// Load translations for other languages.
	bundle.MustLoadMessageFile("./resources/langs/active.en.toml")
	bundle.MustLoadMessageFile("./resources/langs/active.fr.toml")

	I18nBundle = *bundle
}

func GetLocale(ctx *fiber.Ctx) string {
	return ctx.Get("Accept-Language")
}

func GetLocalizer(ctx *fiber.Ctx) *goi18n.Localizer {
	return goi18n.NewLocalizer(&I18nBundle, GetLocale(ctx))
}

func Localize(ctx *fiber.Ctx, messageId string, template map[string]string) string {
	localizer := GetLocalizer(ctx)
	return localizer.MustLocalize(&goi18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: template,
	})
}
