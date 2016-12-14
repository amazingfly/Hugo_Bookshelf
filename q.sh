#!/bin/bash

hugo --theme=hugo_theme_robust
cd public
git commit -am "quick update"
git push
