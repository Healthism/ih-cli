# IH-CLI

## Installation

### Requirement

- Google Cloud SDK

### Setup

```
$ curl %%%%%
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