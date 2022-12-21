read -p "Enter Github Token: " TOKEN
if [ -z "$TOKEN"]; then
    TOKEN=${GIT_TOKEN}
fi
rm go.sum
touch go.sum
git config --global url."https://${TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
go mod tidy -e -go=1.17 && go mod tidy -e -go=1.18