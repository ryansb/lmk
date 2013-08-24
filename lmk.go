package main

import (
	"flag"
	"fmt"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/exp/sns"
	"os"
	"strings"
)

var onlyFailure int
var specificSubject string
var printVersion bool

const nonFail int = -999

func PublishMessage(msgOpts *sns.PublishOpt) {
	var conn *sns.SNS = sns.New(aws.Auth{AccessKeyId, AccessSecret}, aws.USEast)
	_, err := conn.Publish(msgOpts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	flag.IntVar(&onlyFailure, "only-failure", nonFail, "Only send "+
		"notification if the prior command exited non-zero. Need to pass $? "+
		" to this command.")
	flag.StringVar(&specificSubject, "subject", "", "Specify the subject "+
		"(defaults to beginning of message)")
	flag.BoolVar(&printVersion, "v", false, "Print version number")
}

func main() {
	flag.Parse()

	if printVersion {
		fmt.Println("lmk version 0.1")
		return
	}

	var msg string = strings.Join(flag.Args(), " ")

	var msgOpts sns.PublishOpt = sns.PublishOpt{
		Message:  msg,
		Subject:  msg,
		TopicArn: SnsTopicArn,
	}

	if specificSubject != "" {
		msgOpts.Subject = specificSubject
	} else if len(msg) > 20 {
		msgOpts.Subject = msg[:20] + "..."
	}

	if onlyFailure != nonFail {
		if onlyFailure != 0 {
			PublishMessage(&msgOpts)
		}
		return
	} else {
		PublishMessage(&msgOpts)
	}
}
