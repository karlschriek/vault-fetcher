export VAULT_ADDR=http://127.0.0.1:8200
export VAULT_TOKEN=root

echo ">>>> Input"
cat input.yaml

echo ""
echo ""
echo ">>>> Output"
cat input.yaml | go run main.go
