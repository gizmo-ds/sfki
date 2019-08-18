#!/bin/bash
work_path=$(pwd)
blog_path="${work_path}/sfki_blog"
mkdir $blog_path >/dev/null 2>&1

sfki_path="${work_path}/sfki"
git clone https://github.com/NectarFish/sfki.git $sfki_path
cd $sfki_path
docker stop sfki >/dev/null 2>&1
docker rm sfki >/dev/null 2>&1
docker rmi sfki:0.0.2 >/dev/null 2>&1
if docker build --rm -f "Dockerfile" -t sfki:0.0.2 .; then
  cp -rf "$sfki_path/config" "$blog_path/config"
  cp -rf "$sfki_path/posts" "$blog_path/posts"
  cp -rf "$sfki_path/blog_web" "$blog_path/blog_web"
  cp -rf "$sfki_path/docker-compose.yml" "$blog_path/docker-compose.yml"

  rm -rf $sfki_path
  echo -e "\033[32m Build Success \033[0m"
  echo "================================="
  echo "cd ${blog_path}"
  echo "docker-compose up -d --build"
  echo "================================="
else
  echo -e "\033[31m Build Fail \033[0m"
  exit
fi
