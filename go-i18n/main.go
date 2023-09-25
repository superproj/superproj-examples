// Command example runs a sample webserver that uses go-i18n/v2/i18n.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func main() {
	// 1. 创建一个 Bundle，用于应用程序的生命周期
	bundle := i18n.NewBundle(language.English)

	// 2. 将翻译配置加载到 Bundle 实例中
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	// No need to load en.yaml since we are providing default translations.
	// bundle.MustLoadMessageFile("en.yaml")
	bundle.MustLoadMessageFile("zh.yaml")

	// 定义一个HTTP请求处理函数。
	// 在这个函数中，首先获取请求中的语言参数和 Accept-Language 头部信息，用于确定用户的语言首选项
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")

		// 3. 创建一个 Localizer 以用于一组语言首选项
		localizer := i18n.NewLocalizer(bundle, lang, accept)

		// 获取请求中的参数
		name := r.FormValue("name")
		if name == "" {
			name = "Bob"
		}

		unreadEmailCount, _ := strconv.ParseInt(r.FormValue("unreadEmailCount"), 10, 64)

		// 4. 使用 Localizer 查找消息
		helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloPerson",
				Other: "Hello {{.Name}}",
			},
			TemplateData: map[string]string{
				"Name": name,
			},
		})

		myUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
			// 设置默认的 Message
			DefaultMessage: &i18n.Message{
				ID:          "MyUnreadEmails",
				Description: "The number of unread emails I have",
				One:         "I have {{.PluralCount}} unread email.",
				Other:       "I have {{.PluralCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
		})

		personUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "PersonUnreadEmails",
				Description: "The number of unread emails a person has",
				One:         "{{.Name}} has {{.UnreadEmailCount}} unread email.",
				Other:       "{{.Name}} has {{.UnreadEmailCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
			TemplateData: map[string]interface{}{
				"Name":             name,
				"UnreadEmailCount": unreadEmailCount,
			},
		})

		// 构造响应消息
		message := map[string]interface{}{
			"title":              helloPerson,
			"myUnreadEmails":     myUnreadEmails,
			"personUnreadEmails": personUnreadEmails,
		}

		// 将响应消息转换为 JSON 格式并写入响应
		resp, _ := json.Marshal(message)
		fmt.Fprintln(w, string(resp))
	})

	// 启动 HTTP 服务器
	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
