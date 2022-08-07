# Motivation

The idea behind **YADE** is that a user can provide a `file` or a `plain-text string` along with an algorithm, that he wants to encrypt/decrypt the file with. YADE will perform the required operation and the user can either display it or save it in a file.

```shell
$ yade -i=inputfile --algo=aes -o=outputfile --encrypt
```

The user can choose an algorithm from the list of algorithms YADE implements.


The list of encryption/decryption algorithms are:

- Symmetric (classic)
    - Caesar / Rot - 13
    - Vigenere
    - Hill
    - Substitution / Atbash
    - Playfair
    - ADFGVX
    - Byte Addition
    - XOR
    - Vernam / OTP
    - Homophone
    - Permutation / Transposition
    - Solitaire
    - Scytale / Rail Fence

- Symmetric (modern)
    - AES (Advanced Encryption Standard)
    - DES (Data Encryption Standard)
    - Triple DES
    - RSA Security
    - Blowfish
    - TwoFish