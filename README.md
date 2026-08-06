# Config Transpiler

A go transpiler that converts cloud-config YAML into Butane YAML 

The project was built as a part of my preparation for the Linux foundation
mentorship project, "CNCF - Flatcar Container Linux: Cloud-Init to Butane YAML config transpiler".

The goal was to understand design decisions, architecture and the pipeline for the transpiler while
following production-oriented go practices.

### Features

- CLI
- User conversion for name and gecos

### Project structure
- config/ - contains struct definitions for cloud-config and butane YAML formats.
- parser/ - converts input cloud-config YAML to cloud-config structs
- converter/ - converts cloud-init structs to Butane structs
- writer/ - writes Butane structs into the output YAML file
- examples/ - contains examples (input.yaml and output.yaml)
