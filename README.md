Tuner
=====

Configures system based on available memory. Ideal for containers.

## Usage

```bash
$ tuner
Tuning system to: Memory = 1024, PHP = 128, Procs = 8
```

## Advanced usage

```bash
$ export TUNER_MEMORY=2048
$ export TUNER_PHP_MEMORY=192
$ tuner
Tuning system to: Memory = 2048, PHP = 192, Procs = 10
```
