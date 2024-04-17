<script setup>
import { ref } from "vue"
const field = ref([
  [null, null, null],
  [null, null, null],
  [null, null, null],
])

let ws = new WebSocket('ws://localhost:8080/ws')

ws.onmessage = function (event) {
  let jsonData = JSON.parse(event.data)
  console.log(jsonData)
  if (jsonData.type === "cmd") {
    const { x, y, value } = jsonData
    console.log(x, y, value)
    console.log(field)
    field.value[x][y] = value
  }
}

function sendMsg(x, y) {
  console.log(x, y)
  ws.send(JSON.stringify({ x, y }))
}
</script>

<template>
  <main>
    <div v-for="[x, row] in field.entries()">
      <button v-for="[y, cell] in row.entries()" @click="() => sendMsg(x, y)">
        {{ cell ? cell : "[]" }}
      </button>
    </div>
  </main>
</template>
