# Project Structure

The root module is named `yade` and it contains packages for encryption and decryption algorithms.

```bash
yade
+--- cli.go (Driver code file)
+--- symmetricClassic
|   +--- caesar.go
|   +--- hill.go
|   +--- playfair.go
|   ...
+--- symmetricModern
|   +--- aes.go
|   +--- des.go
|   +--- tripleDes.go
|   ...
+--- utils
|   +--- matrix.go
```

Each go file, in the packages, contains an `encrypt` and a `decrypt` function which performs the required operation.