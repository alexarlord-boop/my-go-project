# If ```make swagger``` fails, try this:


1. Install the swagger Tool: Run the following command to install the swagger tool:

> ```GO111MODULE=on go install github.com/go-swagger/go-swagger/cmd/swagger@latest```

2. Ensure ```GOPATH/bin``` is in ```PATH```: Add the following line to your shell profile file (e.g., .bashrc, .zshrc, or .profile):

> ```export PATH=$PATH:$(go env GOPATH)/bin```

3. Then, reload your shell profile:
> ```source ~/.zshrc```

4. Verify the swagger Command: Check if the swagger command is available:
> ```which swagger```

5. Run the Makefile: Now, try running the make swagger command again:
> ```make swagger```
