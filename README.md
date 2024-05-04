# CampusTrade
This is a simple managing system like project built by GoLang
### Project Structure
![image](./assets/project_structure.png)
## Environment
First you need to make sure that you have installed the `go` environment.

If you have already installed `go`,Execute the following command in powershell.Otherwise you should download `go` from [GoLang](https://golang.google.cn/)


this command makes sure that you have successfully download the `go`.

```pwsh
# Use Bash/Zsh/Fish/PowerShell
go env
```
## Run project locally
First you should the packages which this project depend on.We provided a simple way to install these  packages.
### Windows
Open your PowerShell
```pwsh
cd .\scripts\
# Set the execution policy, if it has not been set
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope Process

# run the script
./init.ps1

cd ..
```

### Linux
Open your Bash
```sh
cd ./scripts/
chmod +x init.sh
./init.sh
cd ..
```
