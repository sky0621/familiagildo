<template>
  <v-app>
    <!--
      ヘッダー
    -->
    <v-app-bar fixed app color="primary">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title><h3 :class="$style.tomorrow">kaubandus</h3></v-toolbar-title>
    </v-app-bar>

    <!--
      サイドメニュー
    -->
    <v-navigation-drawer v-model="drawer" fixed app>
      <v-list>
        <!--
          Home
        -->
        <v-list-item to="/">
          <v-list-item-icon><v-icon>mdi-home</v-icon></v-list-item-icon>
          <v-list-item-title>Home</v-list-item-title>
        </v-list-item>
        <!--
          Tasks
        -->
        <v-list-group prepend-icon="mdi-format-list-text" no-action>
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title>Tasks</v-list-item-title>
            </v-list-item-content>
          </template>

          <v-list-item to="/tasks/new">
            <v-list-item-content>
              <v-list-item-title>you created</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item to="/tasks/list">
            <v-list-item-title>you should do</v-list-item-title>
          </v-list-item>
        </v-list-group>
        <!--
          Rewards
        -->
        <v-list-group prepend-icon="mdi-treasure-chest" no-action>
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title>Rewards</v-list-item-title>
            </v-list-item-content>
          </template>

          <v-list-item to="/rewards/set">
            <v-list-item-title>you set</v-list-item-title>
          </v-list-item>
        </v-list-group>
        <!--
          Trades
        -->
        <v-list-group prepend-icon="mdi-handshake" no-action>
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title>Trades</v-list-item-title>
            </v-list-item-content>
          </template>

          <v-list-item to="/trades/new">
            <v-list-item-title>create new</v-list-item-title>
          </v-list-item>
          <v-list-item to="/trades/list">
            <v-list-item-title>list</v-list-item-title>
          </v-list-item>
        </v-list-group>
        <!--
          Results
        -->
        <v-list-group prepend-icon="mdi-file-chart" no-action>
          <template v-slot:activator>
            <v-list-item-title>Results</v-list-item-title>
          </template>
            <v-list-item to="/results/rewards">
              <v-list-item-title>The reward you got</v-list-item-title>
            </v-list-item>
        </v-list-group>
      </v-list>
    </v-navigation-drawer>

    <!--
      コンテンツ
    -->
    <v-main>
      <nuxt />
    </v-main>

    <!--
      フッター
    -->
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

<style module>
.tomorrow {
  font-family: 'Tomorrow', sans-serif;
}
</style>
