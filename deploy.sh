# scp -i .ssh/old/id_rsa_2 -r ../go_server ubuntu@129.159.137.246:~
rsync -r -v --exclude .git -e ssh -i .ssh/old/id_rsa_2 ./ ubuntu@129.159.137.246:~/agora_server
ssh -i .ssh/old/id_rsa_2 -t ubuntu@129.159.137.246 'cd ~/agora_server && sh ./buildAndStart.sh'