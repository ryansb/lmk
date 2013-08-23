package main

import (
	"fmt"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/exp/sns"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	msg := strings.Join(os.Args[1:], " ")
	var conn *sns.SNS = sns.New(aws.Auth{AccessKeyId, AccessSecret}, aws.USEast)

	msgOpts := sns.PublishOpt{
		Message:  msg,
		Subject:  msg,
		TopicArn: SnsTopicArn,
	}

	if len(msg) > 20 {
		msgOpts.Subject = msg[:20] + "..."
	}

	_, err := conn.Publish(&msgOpts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}
