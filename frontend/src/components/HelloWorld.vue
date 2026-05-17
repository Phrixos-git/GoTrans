<script setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'
import {ProcessMessage} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  //resultText: "Please enter your name below 👇",
　// 1. メッセージ用の変数を追加
  myMessage: "",
})

/**function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}
*/

// 2. 送信用の関数を追加
function sendMessage() {
  ProcessMessage(data.myMessage).then(result => {
    data.resultText = result
    data.myMessage = "" // 送信後に枠を空にする
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <!--<div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
    </div>-->
     <!-- 3. 新しく追加するメッセージボックス -->
    <div class="input-box" style="margin-top: 20px;">
      <input v-model="data.myMessage" class="input" type="text" placeholder="メッセージを入力"/>
      <button class="btn" @click="sendMessage">送信</button>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
