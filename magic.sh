TOKEN=$1
rm go.sum
touch go.sum
export GOPRIVATE=github.com/thesixnetwork/sixnft ##! To make it able to download private repo
git config --global url."https://${TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/" ##! To make it able to download private repo
go mod tidy -e -go=1.17 && go mod tidy -e -go=1.18
git config --global --unset url."https://${TOKEN}:x-oauth-basic@github.com/".insteadOf  ##! To prevent it from cannot access to other repo