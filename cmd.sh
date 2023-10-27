#!/bin/bash
source utils.sh
echo -------- "$1" "$2" "$3" "$4" --------

### git
# ./cmd.sh git push
# ./cmd.sh git clear
if [ "$1" == "git" ]; then
  commandExists git
  if [ "$2" == "push" ]; then
    branch=$(git symbolic-ref --short -q HEAD)
    log "------------ git push ${branch} ------------"
    time=$(date '+%F %T')
    msg='auto push at '${time}
    git add .
    git commit -m "${msg}"
    git push origin "${branch}"
    echo "git push origin ${branch} success"
  elif [ "$2" == "tag" ]; then
    log "------------ git add tag ------------"
    latest_tag=$(git describe --abbrev=0 --tags)
    major_version=$(echo "$latest_tag" | cut -d. -f1)
    minor_version=$(echo "$latest_tag" | cut -d. -f2)
    patch_version=$(echo "$latest_tag" | cut -d. -f3)
    patch_version=$((patch_version + 1))
    new_version="$major_version.$minor_version.$patch_version"
    git tag "$new_version"
    git push origin --tags
    echo "git add tag ${new_version} success "
  elif [ "$2" == "clear" ]; then
    log "------------ git clear commit ------------"
    remoteUrl=$(git config --get remote.origin.url)
    log "$remoteUrl"
    rm -rf .git
    git init
    git add .
    git commit -am "init"
    git remote add origin "${remoteUrl}"
    git push origin master --force
    echo "clear commit success"
  fi
fi
