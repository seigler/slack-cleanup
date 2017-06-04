# Slack Cleanup

[![CircleCI](https://circleci.com/gh/madecomfy/slack-cleanup/tree/master.svg?style=svg&circle-token=4aee14214b029ff6bcb7664e8c403ef1288e84f4)](https://circleci.com/gh/madecomfy/slack-cleanup/tree/master)

A program for cleaning up slack file uploads. Useful if you keep hitting limits on your plan.

## Installing Slack Cleanup

    go get -u github.com/madecomfy/slack-cleanup

You will need to generate an API token that has permission to delete files. Add this token to your environment
variables. At MadeComfy we use [Direnv](https://direnv.net/) to manage local environment variables.

    export SLACK_API_TOKEN=xoxp-111111-1111111-1111111-1111111-111111

## Running Slack Cleanup

    slack-cleanup
