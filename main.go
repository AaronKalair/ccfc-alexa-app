package main

import (
	"fmt"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var Applications = map[string]interface{}{
	"/echo/": alexa.EchoApplication{
		AppID:    GetSettings().AppId,
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
}

func main() {
	alexa.Run(Applications, "3001")
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	database := InitDb(GetSettings().DatabaseLocation)
	matches := ReadItem(database)
	fmt.Println(matches[0].date)
	fmt.Println(echoReq.GetIntentName())
	echoResp.OutputSpeech("The next game is against " + matches[0].match)
}
