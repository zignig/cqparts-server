<template>
    <div >
        <div v-bind:model="issueItem">{{ issueItem }}</div>
    </div>
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
