### 1.Install Delve Debugger

```sh
go install github.com/go-delve/delve/cmd/dlv@latest
```

### 2. Verify Installation

```sh
dlv version

```

### 3. Explore Delve Help Commands

```sh
dlv
```

### 4. Update & Install Delve via VS Code

```sh
    1.Open Command Palette (Cmd + Shift + P / Ctrl + Shift + P)
    2.Search for Go: Install/Update Tools
    3.Select dlv latest
```

### 5. Configure VS Code Debugger

```sh
    1.Go to Run and Debug in VS Code (left sidebar).
    2.Click Create a launch.json file → choose Go Launch.
    3.Add the following snippet to .vscode/launch.json:
```

```sh
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type":"go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/dsa/main.go",
        }
    ]
}
```

### 6. finally you can see `.vscode file`

```sh
A .vscode folder will be created with your debugging configuration.
Now, hit Run and Debug ▶️ to start debugging with Delve
```
