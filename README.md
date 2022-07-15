# polar

Inspired from the [husky-go](https://github.com/automation-co/husky)

## Docs

### Installation

```
go install github.com/plvhx/polar@latest
```

### Getting Started

You can initialize polar by typing ```polar init```

> Make sure you have git initialized.

This will make the .polar directory with the hooks folder. You can add hooks using:

```bash
$ polar add <hook> "
	<your commands for that hook>
"
```

### Example

```bash
$ polar add pre-commit "
	go build -v ./...
	go test -v ./...
"
```

### Features

> Checking commit message before committed (based on [git-karma](http://karma-runner.github.io/6.3/dev/git-commit-msg.html)). But first, you must write your own acceptable commit types in ```.polar-commitmsg-types.yaml```. You can see example configuration format below:


```yaml
types:
  - feat
  - fix
  - perf
  - docs
  - style
  - refactor
  - test
  - build
  - defect
  - foobar
  - aabbcc
```

> Then, activate the hook:

```bash
$ polar activate-commit-validator
```

And you're set. Everytime you run ```git commit``` it will validates that commit message you recently supplied.

## Blogs and Resources
- [ Get Started with Husky for go ](https://dev.to/devnull03/get-started-with-husky-for-go-31pa)
- [ Git Hooks for your Golang project ](https://dev.to/aarushgoyal/git-hooks-for-your-golang-project-1168)

## Get Familiar with Git Hooks

Learn more about git hooks from these useful resources:
- [ Customizing Git - Git Hooks ](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)
- [ Atlassian Blog on Git Hooks ](https://www.atlassian.com/git/tutorials/git-hooks)
- [ Fei's Blog | Get Started with Git Hooks ](https://medium.com/@f3igao/get-started-with-git-hooks-5a489725c639)

## License

BSD-3-Clause
