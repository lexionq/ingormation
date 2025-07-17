# ingormation
Detailed system analysis tool for pentesters and everyday use. With developed Go. 🦫

<p align="center">
  <img src="https://img.shields.io/github/license/lexionq/ingormation?style=for-the-badge&color=blue">
  <img src="https://img.shields.io/github/languages/top/lexionq/ingormation?style=for-the-badge&color=cyan">
  <img src="https://img.shields.io/github/v/release/lexionq/ingormation?style=for-the-badge&color=purple">  
</p>
<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/lexionq/ingormation?style=for-the-badge&color=darkblue">
  <img src="https://goreportcard.com/badge/github.com/lexionq/ingormation?style=for-the-badge&color=purple">
  <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/lexionq/ingormation/go.yml?style=for-the-badge&color=darkgreen">
</p>

## Disclaimer
>[!CAUTION]
> ***This tool is created solely to perform system analysis in ethical environments. The developer, lexionq, expressly disclaims any responsibility for any technical or legal issues arising from the use of this tool on non-ethical systems.***

## What system information can you access with this tool?
- Hostname
- Host ID
- OS
- Kernel Version
- Kernel Arch
- Platform
- Platform Version
- Platform Family
- Virtualization Role
- Virtualization System
- Boot Time & Uptime
- CPU Cores
- Core ID
- Microcode
- Family
- Model
- Model NAme
- Mhz
- Physical ID & Vendor ID
- CacheSize
- Flags
- CPU Count
- CPU Usage per Second
- Total Memory
- Used Memory ( MB & Percent )
- Free Memory
- Total Swap
- Used Swap ( MB )
- Free Swap
- Disks
- Mountpoints
- Filesystem
- Total Disk (GB)
- Used Disk
- Free Disk

## Usage on Linux (Kali, Parrot, Mint and others)
The program has already been compiled for you. Therefore, you can download the compiled version directly from this repository:

### 1. Go to "Releases"

### 2. And download "ingormation" file

**OR**

```bash
git clone https://github.com/lexionq/ingormation.git

cd ingormation

chmod +x ingormation(if require)

./ingormation
```

Or, you can run the Go file downloaded from the repository:

```bash
git clone https://github.com/lexionq/ingormation.git

cd ingormation

go run ingormation.go

```

## Build on Your Computer
You can compile it yourself if you wish (there is no need to do so, as the program has already been compiled for you). All you need to do to compile it on your own computer is the following:

`go build`

## Go Report Card Results
<img width="492" height="751" alt="image" src="https://github.com/user-attachments/assets/30f24f50-be8d-4f2e-9bc2-f562a60f0e4e" />
