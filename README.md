# j2
Private use

## Editor integration with bazel
Please refer this [doc](https://github.com/bazelbuild/rules_go/wiki/Editor-and-tool-integration)

Example for neovim:
``` lua
nvim_lsp.gopls.setup {
  on_attach = on_attach,
  settings = {
    gopls = {
      env = {
        GOPACKAGESDRIVER = './tools/gopackagesdriver.sh'
      },
      directoryFilters = {
        "-bazel-bin",
        "-bazel-out",
        "-bazel-testlogs",
        "-bazel-j2",
      },
    },
  },
}
```
