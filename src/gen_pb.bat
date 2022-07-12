if NOT EXIST %~dp0common\pb mkdir %~dp0common\pb
protoc --go_out=%~dp0common\pb  --proto_path=%~dp0proto\ *.proto --plugin=protoc-gen-go --experimental_allow_proto3_optional