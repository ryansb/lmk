LMK.go
======
A dead-simple piece of machinery to let you know when something happens via
SNS.

You build it by moving `conf.go.empty` to `conf.go` then filling in your AWS
keys and the topic ARN you want to publish to.

Once you've build LMK with your keys, you can copy it to any machine with no
configuration and use it. Great to have on all your machines, and so easy!

Usage
=====
If you have a long-running task (say, a big file copy) and you don't want to
watch it but need to make sure it's ok, here's how you'd do it.

Your shell will interpret "$?" to the exit code of long-running-task, so you'll
know if it failed.

`long-running-task; lmk long-running-task exited code "$?"`

Whatever args you pass will be truncated to 20 characters for the message
subject, and the body will contain the rest.
