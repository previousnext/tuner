Tuner
=====

Configures system based on 3 variables:

* Max - The maximum amount of memory available
* Proc - The amount of memory per process
* Multiplier - A method of boosting how many procs can be run

## Usage

**Apache**

```bash
$ tuner --conf=apache
<IfModule mpm_prefork_module>
	StartServers		8
	MinSpareServers		4
	MaxSpareServers		4
	MaxRequestWorkers	8
	MaxConnectionsPerChild  8
</IfModule>
```

**PHP**

```bash
$ tuner --conf=php
memory_limit = 128M
```

## Advanced usage

```bash
$ export TUNER_CONF=apache
$ export TUNER_MAX=1024
$ export TUNER_PROC=64
$ export TUNER_MULTIPLIER=3
$ tuner --conf=apache
<IfModule mpm_prefork_module>
	StartServers		48
	MinSpareServers		24
	MaxSpareServers		24
	MaxRequestWorkers	48
	MaxConnectionsPerChild  48
</IfModule>
```
