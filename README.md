## Config Transpiler

A config transpiler written in Go that converts cloud-config YAML files
into Butane YAML files.

## Motivation

This repository is a prototype for the project named ```CNCF - Flatcar Container Linux:
Cloud-Init to Butane YAML config transpiler (2026 Term 3)``` which is under LFX 2026 fall term 
program.

## Architecture

![Config Transpiler Architecture](./doc/images/config-transpiler-architecture.png)

After the work of the config transpiler, butane compiler can convert the output butane yaml file
into an ignition file which can then be used for provisioning of a Flatcar Linux machine.

## QEMU Validation

1) config-transpiler converts cloud-config YAML file from examples into a butane format YAML file
named ```output.yaml```

![step-1](./doc/images/output-1.png)

2) The output butane format YAML file is then used by the Butane compiler in another directory.

![step-2](./doc/images/output-2.png)

3) The ```config.ign``` is then used by QEMU for a Flatcar Image.

![step-3](./doc/images/output-3.png)

## Features supported

Since this is a prototype, only a small subset of the features needed are supported.
For more details : [Supported Features](./doc/supported-features.md)

## Example usage

Assume a cloud-config YAML file, that has the following contents:

```yaml
#cloud-config

runcmd:
  - echo hello
  - echo bye
groups:
  - devs
  - bosses
users:
  - name: foobar
    gecos: Foo Bar. H
    groups: devs, bosses
    shell: /bin/dash
    uid: 1004
  - name: barfoo
    gecos: Bar Foo. I
    groups: devs
    shell: /bin/bash
    uid: 1768
```

It will be converted to:

```yaml
variant: flatcar
version: 1.0.0
passwd:
    users:
        - name: foobar
          gecos: Foo Bar. H
          groups:
            - devs
            - bosses
          shell: /bin/dash
          uid: 1004
        - name: barfoo
          gecos: Bar Foo. I
          groups:
            - devs
          shell: /bin/bash
          uid: 1768
    groups:
        - name: devs
        - name: bosses
systemd:
    units:
        - name: cloud-init-runcmd-0.service
          enabled: true
          contents: |-
            [Unit]
            Description=cloud-init runcmd unit

            [Service]
            Type=oneshot
            ExecStart=/bin/sh -c 'echo hello'

            [Install]
            WantedBy=multi-user.target

        - name: cloud-init-runcmd-1.service
          enabled: true
          contents: |-
            [Unit]
            Description=cloud-init runcmd unit

            [Service]
            Type=oneshot
            ExecStart=/bin/sh -c 'echo bye'

            [Install]
            WantedBy=multi-user.target

```


## Using the CLI

```bash
./config-transpiler <input_file_name.yaml> <output_file_name.yaml>
```

## Next Steps

Checkout out the [Supported Features](./doc/supported-features.md) section for future plans.
To quickly run the CLI, use the pre-written config files in examples section by running:

```
./config-transpiler examples/input.yaml <output_filename>
```

