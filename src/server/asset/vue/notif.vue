<template>
    <b-notification>
        <div v-bind:model="issueItem">{{ issueItem }}</div>
    </b-notification>
</template>

<script>
export default {
    props: {
        issueItem: {
            type: String,
        },
    },
    methods:{
    }
}
</script>

<style>
</style>
