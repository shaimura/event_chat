FROM node:lts-alpine

WORKDIR /app


COPY docker/vue/package*.json ./

# RUN apk update && \
#     npm install && \
#     npm install -g npm && \
#     npm install -g vue-cli && \
#     npm install -g http-server

RUN npm install
# 静的コンテンツを配信するシンプルな http サーバをインストールする


# `package.json` と `package-lock.json` （あれば）を両方コピーする

# プロジェクトの依存ライブラリをインストールする
# RUN npm install

# カレントワーキングディレクトリ(つまり 'app' フォルダ)にプロジェクトのファイルやフォルダをコピーする
COPY docker/vue/ .


RUN npm install axios && \
    npm install vue-router && \
    npm install vuex && \
    npm install vuex-persistedstate && \
    npm install vue2-dropzone



# 本番向けに圧縮しながらアプリケーションをビルドする
RUN npm run build

EXPOSE 8080
CMD [ "http-server", "dist" ]

