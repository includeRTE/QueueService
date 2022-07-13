@echo off
set CURRENT_PATH = %~dp0
if NOT EXIST %CURRENT_PATH%common\pb mkdir %CURRENT_PATH%common\pb

for %%i in (%CURRENT_PATH%proto\*.proto) do (
    protoc .\%%i --go_out=%CURRENT_PATH%common\pb --plugin=%GOPATH%\bin\protoc-gen-go --experimental_allow_proto3_optional
)
