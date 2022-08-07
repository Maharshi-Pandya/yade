# Project Structure

The root module is named `yade` and it contains packages for encryption and decryption algorithms.

```bash
yade
+--- cli.go (Driver code file)
+--- symclassic
|   +--- caesar
|   +--- hill
|   +--- playfair
|   ...
+--- symmodern
|   +--- aes
|   +--- des
|   +--- tripleDes
|   ...
+--- utils
|   +--- utils.go
```

Each go file, in the packages, contains an `Encrypt` and a `Decrypt` function which performs the required operation.