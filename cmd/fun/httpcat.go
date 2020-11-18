package fun

import (
	"math/rand"

	"github.com/zorbyte/whiskey/arg"
	"github.com/zorbyte/whiskey/cmd"
)

var statusCodes = []string{
	"100", "101", "102",
	"200", "201", "202", "203", "204", "205", "206", "207",
	"300", "301", "302", "303", "304", "305", "306", "307",
	"400", "401", "402", "403", "404", "405", "406", "408", "409",
	"410", "411", "412", "413", "414", "415", "416", "417", "418",
	"420", "421", "422", "423", "424", "425", "426", "428", "429",
	"431", "444", "450", "451", "499",
	"500", "501", "502", "503", "504", "505", "506", "507", "508", "509",
	"510", "511", "599",
}

func init() {
	cmd := cmd.New()
	cmd.Name("httpcat")
	cmd.Aliases("cat", "http")
	cmd.Description("Grab a httpcat")
	cmd.Use(httpcat)
	cmd.Arg(&arg.Argument{
		Name: "code",
	})

	Category.AddCommand(cmd.Command())
}

func httpcat(ctx *cmd.Context, next cmd.NextFunc) {
	if len(ctx.RawArgs) > 0 {
		arg := ctx.Args["code"].(string)
		if stringInSlice(arg, statusCodes) {
			ctx.Send(httpCat(arg))
		} else if code := itemFromAinB(ctx.RawArgs, statusCodes); code != "" {
			ctx.Send(httpCat(code))
		} else {
			ctx.Send(httpCat("404"))
		}
	} else {
		code := randomFromSlice(statusCodes)
		ctx.Send(httpCat(code))
	}
}

func httpCat(code string) string {
	return "https://http.cat/" + code
}

func randomFromSlice(slice []string) string {
	i := rand.Intn(len(slice))
	return slice[i]
}

func stringInSlice(str string, slice []string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func itemFromAinB(a []string, b []string) string {
	for _, item := range b {
		if stringInSlice(item, a) {
			return item
		}
	}
	return ""
}
