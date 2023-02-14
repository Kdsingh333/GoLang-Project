package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4812272197328-4791074613748-IvipU6Go5PrOrqydiyybYZKt")
	os.Setenv("CHANNEL_ID", "C04NZUW0D62")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"basic_golang_notes.pdf"}
	
	for i:=0; i<len(fileArr); i++{
		parmas := slack.FileUploadParameters{
			Channels : channelArr,
			File: fileArr[i],
		}
		file,err := api.UploadFile(parmas)
		if err!= nil{
			fmt.Printf("%s\n",err)
			return 
		}
		fmt.Printf("Name : %s,URL:%s\n",file.Name,file.URL)
	}
	

}
