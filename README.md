# IH-CLI

## Installation

### Requirement

- Google Cloud SDK
- kubectl

### Setup
> Be sure to remove old 'ih' cli if you have installed previously. 

```shell
# initialize your GCP session
$ gcloud init

# initialize default kubernetes cluster
$ gcloud container clusters get-credentials staging --zone northamerica-northeast1-a

# linux users use : https://github.com/Healthism/ih-cli/releases/download/1.0.3/ih-linux
$ curl -L https://github.com/Healthism/ih-cli/releases/download/1.0.3/ih --output /usr/local/bin/ih
$ chmod +x /usr/local/bin/ih
$ gcloud source repos clone staging-deployment /usr/local/lib/ih --project=inputhealth-chr
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
--release [string]
```
- [REQUIRED] a helm release in staging cluster (ex. chr-qa-backend-1, chr-qa-desktop, chr-qa-socket-1 ...)

```
--verbose
```
- [OPTIONAL] [DEFAULT: false] verbose output
```
--help
```
- [OPTIONAL] show help

### Loacl Flags
#### cmd : update
```
--manual
```
- [OPTIONAL] [DEFAULT: false] open text editor instead to update configration

#### cmd : run
```
--command [string]
```
- [OPTIONAL] [DEFAULT: rails console] run specific command
---
## Example Usage

### Launching Rails console for 'qa backend 1'
```
$ ih -r chr-qa-backend-1 run
```

### Executing rake task for 'qa backend 1'
```
$ ih -r chr-qa-backend-1 run -c "rake db:migrate"
```

### Changing Configuration
```
$ ih -r chr-qa-backend-1 update $key1=$value1 $key2=$value2 ....
```

### Editiing Configuration Using Text Editor
```
$ ih --release chr-qa-socket-1 update --manual
```
---
