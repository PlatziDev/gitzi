# Gitzi

v0.99

Slack notifier project for Github issues.

## Requirements

- Go > 1.5 (We tested Gitzi with Golang 1.7.4)
- (Glide)[https://github.com/Masterminds/glide]

## Installation

- Download from release file:
  - Windows
  - Linux
  - OS x
- Create slack_users.yaml
  - Example:
  ```yaml
  users:
    <GithubUsername>: <SlackUsername>
  ```
- Run

## Build

- `$ go get github.com/PlatziDev/gitzi`
- `$ cd $GOPATH/src/github.com/PlatziDev/gitzi`
- `$ glide install`
- `$ go build`

## Note

We on Platzi using Gitzi with two big repositories, right now we don't have problems managing the same secret with both
repos.