# Prime numbers generator

This is just a sample project to test Go language parallel execution capabilities (channels and wait groups).

## Requirements
- The make tool
- Go sdk

## Build

Use included [Makefile]:
```sh
make build
```

## Running

You can use included [Makefile] to understand how it works:

```sh
make run NUMBERS=100 JOBS=3
```

![demo](docs/demo.gif?raw=true "Demo")

## Diagram
This diagram can help you to understand the flow used for this little program.
![diagram](docs/go-parallel-execution-flux.png?raw=true "Diagram")

[Makefile]: Makefile "Makefile"