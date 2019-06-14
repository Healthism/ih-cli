# IH-CLI

## Installation

### Requirement

- Google Cloud SDK

### Setup
> Be sure to remove old 'ih' cli if you have installed previously. 

```
$ curl -L https://github.com/Healthism/ih-cli/releases/download/1.0.0/ih --output /usr/local/bin/ih
$ chmod +x /usr/local/bin/ih
$ git clone ssh://rlee@inputhealth.com@source.developers.google.com:2022/p/inputhealth-chr/r/staging-deployment /usr/local/lib/ih
```
---
## Usage

```
$ ih [GLOBAL FLAGS] [COMMAND] [LOCAL FLAGS] [ARGUMENTS]
```
---
## Flags
### Global Flags

```
--release=[string]
```
- [REQUIRED] a helm release in staging cluster (ex. chr-qa-backend-1, chr-qa-desktop, chr-qa-socket-1 ...)

```
--verbose=[boolean]
```
- [OPTIONAL] [DEFAULT: false] verbose output
```
--help
```
- [OPTIONAL] show help

### Loacl Flags
#### cmd : update
```
--manual=[boolean]
```
- [OPTIONAL] [DEFAULT: false] open text editor instead to update configration

#### cmd : run
```
--command=[string]
```
- [OPTIONAL] [DEFAULT: rails console] run specific command
---
## Example Usage

### Laucnhing console for 'qa backend 1'
```
$ ih --release=chr-qa-backend-1 run
```

### Changing Configuration
```
$ ih --release=chr-qa-backend-1 update $key1=$value1 $key2=$value2 ....
```

### Editiing Configuration Using Text Editor
```
$ ih --release=chr-qa-socket-1 update --manual
```
---