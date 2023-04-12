<template>
  <v-container>
    <v-row>
      <v-col>
        <h1 :class="$style.tomorrow">What are you going to do today?</h1>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-card color="#BBDEFB">
          <div class="d-flex flex-no-wrap justify-space-between">
            <div>
              <v-card-title class="headline">check my tasks</v-card-title>
              <v-card-subtitle>if you want to ...</v-card-subtitle>
              <v-card-actions>
                <v-btn text to="/tasks/list">Do it now</v-btn>
              </v-card-actions>
            </div>
            <v-avatar class="ma-3" size="125" tile>
              <v-img src="top/checkMyTasks.png"></v-img>
            </v-avatar>
          </div>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-card color="#DCEDC8">
          <div class="d-flex flex-no-wrap justify-space-between">
            <div>
              <v-card-title class="headline">create task</v-card-title>
              <v-card-subtitle>if you want to ...</v-card-subtitle>
              <v-card-actions>
                <v-btn text to="/tasks/new">Do it now</v-btn>
              </v-card-actions>
            </div>
            <v-avatar class="ma-3" size="125" tile>
              <v-img src="top/createTask.png"></v-img>
            </v-avatar>
          </div>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-card color="#FFE0B2">
          <div class="d-flex flex-no-wrap justify-space-between">
            <div>
              <v-card-title class="headline">decide your rewards</v-card-title>
              <v-card-subtitle>if you want to ...</v-card-subtitle>
              <v-card-actions>
                <v-btn text to="/rewards/new">Do it now</v-btn>
              </v-card-actions>
            </div>
            <v-avatar class="ma-3" size="125" tile>
              <v-img src="top/decideYourRewards.png"></v-img>
            </v-avatar>
          </div>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-card color="#D7CCC8">
          <div class="d-flex flex-no-wrap justify-space-between">
            <div>
              <v-card-title class="headline">trade</v-card-title>
              <v-card-subtitle>if you want to ...</v-card-subtitle>
              <v-card-actions>
                <v-btn text to="/trades/list">Do it now</v-btn>
              </v-card-actions>
            </div>
            <v-avatar class="ma-3" size="125" tile>
              <v-img src="top/trade.png"></v-img>
            </v-avatar>
          </div>
        </v-card>
      </v-col>
      <v-col cols="12">
        <v-card color="#D1C4E9">
          <div class="d-flex flex-no-wrap justify-space-between">
            <div>
              <v-card-title class="headline">check your results</v-card-title>
              <v-card-subtitle>if you want to ...</v-card-subtitle>
              <v-card-actions>
                <v-btn text to="/results/list">Do it now</v-btn>
              </v-card-actions>
            </div>
            <v-avatar class="ma-3" size="125" tile>
              <v-img src="top/checkYourResults.png"></v-img>
            </v-avatar>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
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
