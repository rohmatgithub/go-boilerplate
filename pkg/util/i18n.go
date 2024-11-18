package util

import (
	"boilerplate/pkg/configs"
	"encoding/json"
	"fmt"
	"path"
	"reflect"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func InitializeI18n() {
	bundle = i18n.NewBundle(language.English)
	// bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(path.Join(configs.App.I18nPath, "error/id.json"))
	bundle.MustLoadMessageFile(path.Join(configs.App.I18nPath, "error/en.json"))
	bundle.MustLoadMessageFile(path.Join(configs.App.I18nPath, "constanta/id.json"))
	bundle.MustLoadMessageFile(path.Join(configs.App.I18nPath, "constanta/en.json"))
	bundle.MustLoadMessageFile(path.Join(configs.App.I18nPath, "common/id.json"))
	bundle.MustLoadMessageFile(path.Join(configs.App.I18nPath, "common/en.json"))
}

func GetI18nErrorMessage(locale string, messageID string, templateData map[string]interface{}) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("err : ", r)
		}
	}()
	localizer := i18n.NewLocalizer(bundle, locale)

	for key, value := range templateData {
		dtType := reflect.TypeOf(value)
		if dtType.Kind() == reflect.String {
			templateData[key] = localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: value.(string)})
		}
	}

	translation := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
	return translation
}
