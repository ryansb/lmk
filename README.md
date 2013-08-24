LMK.go
======
A dead-simple piece of machinery to let you know when something happens via
SNS.

You build it by moving `conf.go.empty` to `conf.go` then filling in your AWS
keys and the topic ARN you want to publish to.

Once you've build LMK with your keys, you can copy it to any machine with no
configuration and use it. Great to have on all your machines, and so easy!

To build on your system:

```
$ vim conf.go conf.go.example # see Usage
$ go get launchpad.net/goamz/aws launchpad.net/goamz/exp/sns
$ go build
# if you need a build for a different architecture
$ cp lmk lmk.x64
$ GOARCH=386 go build -a
$ cp lmk lmk.i386
```

I keep a `lmk.x64` and `lmk.i386` handy so I don't have to rebuild every time I
want it on a different architecture.

Usage
-----
If you have a long-running task (say, a big file copy) and you don't want to
watch it but need to make sure it's ok, here's how you'd do it.

Your shell will interpret "$?" to the exit code of long-running-task, so you'll
know if it failed.

`long-running-task; lmk long-running-task exited code "$?"`

Whatever args you pass will be truncated to 20 characters for the message
subject, and the body will contain the rest.

If you only want a notification if the command fails, pass `$?` to the
`--only-failure` option.

`long-running-task-that-might-fail; lmk --only-failure $? long-running-task
failed. Bad news dude.`

You can also specify the subject.

`task; lmk --subject "RE: Your Task" the task completed`

Permissions
-----------
Since it's possible to extract the keys directly from the binary, I recommend
having a separate IAM role for this tool. If the binary happens to get
somewhere you don't want it (and someone starts spamming you) it makes it easy
to proactively prevent them from doing other things (like getting into your S3)
and to cut them off if they over-publish.

To do this, make a new IAM user and give them these permissions:

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "sns:Publish"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:sns:THE-ARN-FOR-YOUR-TOPIC"
    }
  ]
}
```
