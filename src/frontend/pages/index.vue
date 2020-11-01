<template>
  <div>
    <h1 :class="$style.tomorrow">What are you going to do today?</h1>
    <v-img src="01.png" max-width="160px" max-height="160px"></v-img>
    <ul>
      <li>やるべきタスクをチェック</li>
      <li>タスクを作る</li>
      <li>報酬を決める</li>
      <li>取引する</li>
      <li>実績を確認する</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component({})
export default class IndexPage extends Vue {}
</script>

<style module>
.tomorrow {
  font-family: 'Tomorrow', sans-serif;
}
</style>
