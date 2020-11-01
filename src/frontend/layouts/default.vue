<template>
  <v-app>
    <v-app-bar fixed app color="primary">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title>kaubandus</v-toolbar-title>
    </v-app-bar>
    <v-navigation-drawer v-model="drawer" fixed app>
      <!-- Home -->
      <v-list>
        <v-list-item to="/">
          <v-list-item-icon><v-icon>mdi-home</v-icon></v-list-item-icon>
          <v-list-item-title>Home</v-list-item-title>
        </v-list-item>
      </v-list>

      <!-- Tasks -->
      <v-list-group prepend-icon="mdi-format-list-text">
        <template v-slot:activator>
          <v-list-item-title>Tasks</v-list-item-title>
        </template>
        <v-list-item-group>
          <v-list-item>
            <v-list-item-title>you created</v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>you should do</v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list-group>

      <!-- Rewards -->
      <v-list-group prepend-icon="mdi-treasure-chest">
        <template v-slot:activator>
          <v-list-item-title>Rewards</v-list-item-title>
        </template>
        <v-list-item-group>
          <v-list-item>
            <v-list-item-title>you set</v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list-group>

      <!-- Trades -->
      <v-list-group prepend-icon="mdi-handshake">
        <template v-slot:activator>
          <v-list-item-title>Trades</v-list-item-title>
        </template>
        <v-list-item-group>
          <v-list-item>
            <v-list-item-title>create new</v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>list</v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list-group>

      <!-- Results -->
      <v-list-group prepend-icon="mdi-file-chart">
        <template v-slot:activator>
          <v-list-item-title>Results</v-list-item-title>
        </template>
        <v-list-item-group>
          <v-list-item>
            <v-list-item-title>The reward you got</v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list-group>
    </v-navigation-drawer>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
    <v-footer app>
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { Component, Provide, Vue } from 'nuxt-property-decorator'

@Component({})
export default class DefaultLayout extends Vue {
  @Provide()
  drawer: boolean = false
}
</script>
