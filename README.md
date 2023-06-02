# vue-practice
vue 相關練習

```
### 確認有安裝nodejs，以及接下來會使用npm功能
node -v

## 安裝npm
## brew install npm
## ~/.zshrc 加入這行
## 請注意會不會跟golang的設定衝突到
export PATH=/usr/local/share/npm/bin:$PATH
## 如果有需要整合在一起，會變成這樣
## shen/go/src 是我的go path
export PATH=/Users/shen/go/src/bin:/usr/local/share/npm/bin:$PATH
```

## 裝vite (vite或vue cli選一個)
```
npm create vite@latest
```
細部設定結果如下
```
Need to install the following packages:
  create-vite@4.3.2
Ok to proceed? (y) y
✔ Project name: … vite-project
✔ Select a framework: › Vue
✔ Select a variant: › Customize with create-vue

  create-vue@3.6.4
Ok to proceed? (y) y

Vue.js - The Progressive JavaScript Framework

✔ Add TypeScript? … No 
✔ Add JSX Support? … No 
✔ Add Vue Router for Single Page Application development? …Yes
✔ Add Pinia for state management? … No 
✔ Add Vitest for Unit Testing? … No 
✔ Add an End-to-End Testing Solution? › No
✔ Add ESLint for code quality? …  Yes
✔ Add Prettier for code formatting? …  Yes

```
```
  cd vite-project
  npm install
  npm run format
  npm run dev
```

# docker設定參考
```
https://dev.to/ysmnikhil/how-to-build-with-react-or-vue-with-vite-and-docker-1a3l
```





### 這個當作備案
```
## 裝vue cli，因為有需要新建立資料夾，或許需要sudo權限
sudo npm install -g @vue/cli
## 建立專案
## vue create [專案名稱]
## 選vue3
vue create modal-project



```