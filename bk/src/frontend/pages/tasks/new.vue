<template>
  <v-container>
    <v-row>
      <v-col>
        <CreateForm />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import CreateForm from '~/components/tasks/CreateForm.vue'

@Component({
  components: {
    CreateForm
  }
})
export default class TasksNewPage extends Vue {}
</script>

<style module>

</style>
