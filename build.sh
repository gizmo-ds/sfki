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
docker build --rm -f "Dockerfile" -t sfki:0.0.2 . && echo -e "\033[32m Build Success \033[0m" || echo -e "\033[31m Build Fail \033[0m" && exit

cp -rf "$sfki_path/config" $blog_path
cp -rf "$sfki_path/posts" $blog_path
cp -rf "$sfki_path/blog_web" $blog_path
cp -rf "$sfki_path/docker-compose.yml" $blog_path

rm -rf $sfki_path

echo "cd ${blog_path}"
echo "docker-compose up -d --build"
