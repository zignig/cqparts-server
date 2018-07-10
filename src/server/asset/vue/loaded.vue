<template lang="html">
  <div>

        <sui-progress v-bind:model="loading"
          progress
          :percent="loading"
        >
        </sui-progress>
  </div>
</template>

<script>
export default {
  props: ['loading','active'],
  methods: {
  },
  computed: {
  }
}
</script>

<style>

</style>
