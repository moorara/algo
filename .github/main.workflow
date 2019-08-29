workflow "Main" {
  on = "push"
  resolves = [ "Lint", "Test" ]
}

action "Lint" {
  uses = "docker://golangci/golangci-lint:latest"
  runs = [ "golangci-lint", "run" ]
  args = [ "--new", "--deadline", "2m" ]
}

action "Test" {
  uses = "docker://golang:1.12"
  runs = [ "go", "test" ]
  args = [ "-race", "./..." ]
}
