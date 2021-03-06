#### Bash completion: add this line to ~/.bashrc

```sh
source <(ugc completion bash)
```

#### Zsh completion: add this line to ~/.zshrc

```sh
source <(ugc completion zsh)
```

#### Powershell completion: add this line to profile.ps1

This requires powershell > 5.0, which comes with Win10 but can be downloaded for Win7/Win8

```powershell
ugc.exe completion powershell | Out-String | Invoke-Expression
```

#### Inspired by

* [grc](https://github.com/garabik/grc)
* [ohmyzsh](https://github.com/ohmyzsh/ohmyzsh)
* [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting)

#### License: Public domain