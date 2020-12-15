#使うDockerイメージ
FROM mysql:5.7

#ポートを開ける
EXPOSE 3306

#MySQL設定ファイルをイメージ内にコピー
ADD docker/mysql/my.cnf /etc/mysql/conf.d/my.cnf

RUN touch /var/log/mysql/mysqld.log

#docker runに実行される
CMD ["mysqld"]