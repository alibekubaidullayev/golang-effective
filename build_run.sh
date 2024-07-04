go work init ./src
go build -o app ./src

if [ $? -eq 0 ]; then
  ./app
else
  echo "Build failed"
fi
