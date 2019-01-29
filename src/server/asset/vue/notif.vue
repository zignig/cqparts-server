<template>
    <b-notification>
        <div v-bind:model="issueItem">{{ issueItem }}</div>
    </b-notification>
</template>

<script>
export default {
    data() {
        return {
            "active": false
        }
    },
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
