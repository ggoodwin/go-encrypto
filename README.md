<div align="center">
	<h1><img alt="go-encrypto logo" src="https://github.com/ggoodwin/go-encrypto/blob/master/logo.png" height="200" /><br />
		Encrypto Go Package
	</h1>

  [![GMan#0001](https://dcbadge.vercel.app/api/shield/179795086543028224)](https://discord.id/?prefill=179795086543028224)

  [![Go Reference](https://pkg.go.dev/badge/ggoodwin/go-encrypto.svg)](https://pkg.go.dev/github.com/ggoodwin/go-encrypto) [![Go Version](https://img.shields.io/github/go-mod/go-version/ggoodwin/go-encrypto)](https://go.dev/) ![Size](https://img.shields.io/github/languages/code-size/ggoodwin/go-encrypto) [![Last Commit](https://img.shields.io/github/last-commit/ggoodwin/go-encrypto)](https://github.com/ggoodwin/go-encrypto/commits/master) [![License](https://img.shields.io/github/license/ggoodwin/go-encrypto)](https://github.com/ggoodwin/go-encrypto/blob/master/LICENSE.md)

  [![GoReportCard](https://goreportcard.com/badge/github.com/ggoodwin/go-encrypto)](https://goreportcard.com/report/github.com/ggoodwin/go-encrypto) [![CodeFactor](https://www.codefactor.io/repository/github/ggoodwin/go-encrypto/badge)](https://www.codefactor.io/repository/github/ggoodwin/go-encrypto) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/17f51d3e54264211b19220ce470783ae)](https://app.codacy.com/gh/ggoodwin/go-encrypto/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![codecov](https://codecov.io/gh/ggoodwin/go-encrypto/branch/master/graph/badge.svg?token=YNDB8EF3ZN)](https://codecov.io/gh/ggoodwin/go-encrypto)

  [![Build](https://github.com/ggoodwin/go-encrypto/actions/workflows/go.yml/badge.svg)](https://github.com/ggoodwin/go-encrypto/actions/workflows/go.yml) [![lint](https://github.com/ggoodwin/go-encrypto/actions/workflows/lint.yml/badge.svg)](https://github.com/ggoodwin/go-encrypto/actions/workflows/lint.yml) [![CodeQL](https://github.com/ggoodwin/go-encrypto/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/ggoodwin/go-encrypto/actions/workflows/github-code-scanning/codeql)

</div>

<hr/>

## 🌟 How it works

This package encrypts and decrypts a file using a password provided.

## 📦 Installation and Usage

### Go

Make sure you have `Go` installed on your machine.

You can check by running the following command in the `console`

```powershell
go version
```

If you don't have `Go` installed, you can download it from [here](https://go.dev/dl/).

### Add to your project

Run the following command in the `console`, in the `project directory`, to install the library with `go get`

```plain
go get github.com/ggoodwin/go-encrypto
```

### Importing

Add the import to your `.go` file

```go
import encrypto "github.com/ggoodwin/go-encrypto"
```

## 💰 Usage

### Check if the Market is Open at Current Time

```go
// Encrypt File
// Parameters: key byte[], inputFile string, outputFile string
error := encrypto.encryptFile(key, inputFile, outputFile)
if error != nil {
    panic(error)
}

// Decrypt File
// Parameters: key byte[], inputFile string, outputFile string
error := encrypto.decryptFile(key, inputFile, outputFile)
if error != nil {
    panic(error)
}
```

## 💻 Dependencies

- [`Go`](https://go.dev/)

## 🙇‍♂️ Issues and Contributing

If you find an issue with this library, please report the issue using our [GITHUB-ISSUES] or check out the [SECURITY] details if it is security related.

If you'd like, I welcome any contributions. Please read the [CONTRIBUTING] document then [FORK] this library and submit a [PULL-REQUEST]. Make sure to click `compare across forks` to see your fork.

## ⚖️ License

This project is under the MIT License. See the [LICENSE] file for the full license text.

## 📜 Changes

Check out our [CHANGELOG]

## 👍🏻 Code of Conduct

Please read my [CODE-OF-CONDUCT] before contributing or engaging in discussions.

<!-- Links -->
[LICENSE]: https://github.com/ggoodwin/go-encrypto/blob/master/LICENSE.md
[CHANGELOG]: https://github.com/ggoodwin/go-encrypto/blob/master/CHANGELOG.md
[SECURITY]: https://github.com/ggoodwin/go-encrypto/blob/master/SECURITY.md
[FORK]: https://github.com/ggoodwin/go-encrypto/fork
[PULL-REQUEST]: https://github.com/ggoodwin/go-encrypto/compare
[CODE-OF-CONDUCT]: https://github.com/ggoodwin/go-encrypto/blob/master/CODE_OF_CONDUCT.md
[CONTRIBUTING]: https://github.com/ggoodwin/go-encrypto/blob/master/CONTRIBUTING.md
[GITHUB-ISSUES]: https://github.com/ggoodwin/go-encrypto/issues
