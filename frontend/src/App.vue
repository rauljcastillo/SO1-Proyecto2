<template>
  <nav>
    <div id="title">
      <h1>SO1</h1>
    </div>
    <div id="descrip">
      <h1>Votaciones</h1>
    </div>
  </nav>
  <div class="main">
    <div id="botones">
      <button @click="getLogs">Obtener logs</button>
    </div>
    <div id="section-logs">
      <p v-for="elem in data">{{elem.date}} -- {{ elem.name }} -- {{ elem.album }}</p>
    </div>
  </div>
</template>

<script setup>

import {ref} from 'vue' 
let data=ref([])
async function getLogs() {
  let datito=await fetch("https://logs-api-tvp5bmf6nq-uc.a.run.app/getLogs")
  data.value= await datito.json()
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  height: 80px;
  text-align: center;
  background-color: goldenrod;
  
}

nav #title {
  width: 200px;
}

nav #descrip {
  width: 200px;
  
}

.main {
  width: 100vw;
  height: calc(100vh - 80px);
  display: flex;
  flex-direction: column;;
}

#botones {
  width: 100%;
  height: 10%;
  display: flex;
  justify-content: center;
  padding: 10px;
}
#botones > button {
	width: 200px;
  height: 40px;
	border-radius: 20px;
  background-color: greenyellow;
}

#section-logs {
  width: 100vw;
  height: 100vh;
  background-color: rgb(44, 43, 43);
  padding: 40px;
  color: white;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>