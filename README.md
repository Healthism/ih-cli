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
$ curl -L https://github.com/Healthism/ih-cli/releases/download/2.0.4/ih --output /usr/local/bin/ih
$ chmod +x /usr/local/bin/ih
$ gcloud source repos clone staging-deployment /usr/local/lib/ih --project=inputhealth-chr
```
---
## Usage

#### Basic Interactive CLI
```
$ ih
```
#### Advanced Usage
- Running Command
```
$ ih run --cluster 'cluster_name' --namespace 'name_space' --release 'release_name' [commands]
```
- Configure Env
```
$ ih config --cluster 'cluster_name' --namespace 'name_space' --release 'release_name'
```
---