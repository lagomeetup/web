#!/bin/bash

st=$(git status --short)
if [ -n "$st" ]; then
    echo "error: uncommited changes"
    exit 1
fi

~/go_appengine/appcfg.py update .
git tag -f deployed
git push
git push --tags -f
