## GO Fast Project

A GO minimal project initializer to test ideas.

```shell
git clone git@github.com:mindwingx/fast-project.git 
```

- build the binary
```shell
go build -o prj main.go
```

- move to `~/bin` (linux/mac) to be accessible globally
```shell
mv prj ~/bin
```

- make the `~/bin` directory if not exists. then add `export PATH="$HOME/bin:$PATH"
` to the `~/.bashrc` or `~/.zshrc`

- then, apply the changes:

```shell
source ~/.bashrc
```

or

```shell
source ~/.zshrc
```

- now, each time by execute the `prj` command, you can have the minimal project in the desired path