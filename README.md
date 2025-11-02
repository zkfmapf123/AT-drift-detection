# AT-Atlantis-Drift-Detection

## Bot Token Required

![1](./public/1.png)
![2](./public/2.png)
![3](./public/3.png)
![4](./public/4.png)
![5](./public/5.png)
![6](./public/6.png)
![7](./public/7.png)
![8](./public/8.png)

## Parameters

| Flag | Short | Description | Example |
|------|-------|-------------|---------|
| `--github-token` | `-g` | The Github token | `ghp_xxx` |
| `--github-repo-ref` | `-f` | The Github repository reference | `main`, `master` |
| `--atlantis-url` | `-u` | The Atlantis URL | `https://atlantis.example.com` |
| `--atlantis-token` | `-t` | The Atlantis token | `your-api-secret` |
| `--atlantis-repository` | `-r` | Atlantis Repository | `owner/repo-name` |
| `--atlantis-config` | `-c` | Atlantis Config File | `atlantis.yaml` |
| `--slack-bot-token` | `-s` | Slack Bot Token | `xoxb-xxx` |
| `--slack-channel` | `-l` | Slack Channel | `C024BE91L` |

## Execute

```sh
    make build
    make run
```