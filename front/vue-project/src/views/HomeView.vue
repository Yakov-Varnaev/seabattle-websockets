<script setup>
let ws = new WebSocket('ws://localhost:8080/ws')

ws.onmessage = function (event) {
  jsonData = JSON.parse(event.data)
  console.log(jsonData)
}

function sendMsg(x, y) {
  console.log(x, y)
  ws.send(`${x},${y}`)
}
</script>

<template>
  <main>
    <div v-for="x in [1, 2, 3]">
      <button v-for="y in [1, 2, 3]" @click="() => sendMsg(x, y)">
        [ ]
      </button>
    </div>
  </main>
</template>
