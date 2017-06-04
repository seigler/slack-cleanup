# Slack Cleanup

[![CircleCI](https://circleci.com/gh/madecomfy/slack-cleanup/tree/master.svg?style=svg&circle-token=4aee14214b029ff6bcb7664e8c403ef1288e84f4)](https://circleci.com/gh/madecomfy/slack-cleanup/tree/master)

A program for cleaning up slack files to free up space.

## Setup

Install Slack Cleanup

    go get github.com/madecomfy/slack-cleanup

You will need to generate an API token that has permission to delete files. Add this token to your environment
variables (we use Direnv to manage env vars).

    export SLACK_API_TOKEN=xoxp-111111-1111111-1111111-1111111-111111
