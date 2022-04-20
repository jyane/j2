load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

filegroup(
    name = "config_pb_txt",
    srcs = ["config.pb.txt"],
)

go_library(
    name = "huer_lib",
    srcs = ["main.go"],
    data = [
        ":config_pb_txt",
    ],
    importpath = "j2/huer",
    visibility = ["//visibility:private"],
    deps = [
        "//golang/util",
        "//huer/proto:config",
    ],
)

go_binary(
    name = "huer",
    embed = [":huer_lib"],
    visibility = ["//visibility:public"],
)
