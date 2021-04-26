# go-kit

[![Go](https://github.com/evenlab/go-kit/actions/workflows/go.yml/badge.svg)](https://github.com/evenlab/go-kit/actions/workflows/go.yml)

### Golang KIT of helpful lightweight and tested functionality for reuse in other projects

# `Makefile` Usage on Windows

## `choco` Setup Chocolatey CLI

> For latest setup info see [Chocolatey CLI Documentation](https://docs.chocolatey.org/en-us/choco/setup)

### Requirements:
* Windows 7+ / Windows Server 2003+
* PowerShell v2+ (Not PowerShell Core yet though)(minimum is v3 for install from this website due to TLS 1.2 requirement)
* .NET Framework 4+ (the installation will attempt to install .NET 4.0 if you do not have it installed)(minimum is 4.5 for install from this website due to TLS 1.2 requirement)
> That's it! All you need is choco.exe (that you get from the installation scripts) and you are good to go! No Visual Studio required.

### Installing Chocolatey
Chocolatey installs in seconds. You are just a few steps from running choco right now!

* Ensure that you are using an administrative shell.
* Copy the text specific to your command shell: cmd.exe or powershell.exe.
* Copy/Paste the text into your shell and press Enter.
* Wait a few seconds for the command to complete.

### `PowerShell:` run the following command

```shell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

### `cmd:` run the following command

```shell
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
```

> If you don't see any errors, you are ready to use Chocolatey! Type choco or choco -? now, or see Getting Started for usage instructions.

## `make` Installing

```shell
choco install make
```
### Now you will be able to use `make` on Windows.
