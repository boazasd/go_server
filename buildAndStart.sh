export PATH=$PATH:/usr/local/go/bin
export GOPATH="$HOME/go"
PATH="$GOPATH/bin:$PATH"
echo "==== Building and starting agora_server ===="
templ generate
go build -o /var/www/agora_server .
rsync -a assets/ /var/www/agora_server/assets
sudo service agora_server restart
